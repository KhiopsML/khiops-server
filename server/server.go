package server

import (
	"bytes"
	"context"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"os"
	"os/exec"
	"sync"
	"time"

	longrunning "cloud.google.com/go/longrunning/autogen/longrunningpb"
	"github.com/gofrs/uuid"
	"github.com/golang/protobuf/ptypes/empty"
	khiops "gitlab.tech.orange/khiops-tools/docker-khiops/go/gen/proto"
	"golang.org/x/net/html/charset"
	"golang.org/x/text/transform"
	code "google.golang.org/genproto/googleapis/rpc/code"
	status "google.golang.org/genproto/googleapis/rpc/status"
	"google.golang.org/protobuf/types/known/anypb"
)

var khiopsCmd, khiopsCoclusteringCmd string

// Job holds info about a particular job
type Job struct {
	cmd                  *exec.Cmd
	stdoutBuf, stderrBuf bytes.Buffer
	error                *status.Status
	cancelled            bool
	ch                   chan error
}

// Backend implements the protobuf interface
type Backend struct {
	ctx   context.Context
	mu    *sync.RWMutex
	items map[string]*Job
	// Embed the unimplemented servers
	khiops.UnimplementedKhiopsServer
	khiops.UnimplementedOperationsServer
}

// New initializes a new Backend struct.
func New(ctx context.Context) *Backend {
	return &Backend{
		ctx:   ctx,
		mu:    &sync.RWMutex{},
		items: make(map[string]*Job),
	}
}

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func init() {
	khiopsCmd = os.Getenv("KHIOPS_CMD")
	if khiopsCmd == "" {
		khiopsCmd = "/usr/bin/khiops"
		fmt.Println("KHIOPS_CMD not set, default to " + khiopsCmd)
	} else {
		fmt.Println("KHIOPS_CMD set to " + khiopsCmd)
	}
	khiopsCoclusteringCmd = os.Getenv("KHIOPS_COCLUSTERING_CMD")
	if khiopsCoclusteringCmd == "" {
		khiopsCoclusteringCmd = "/usr/bin/khiops_coclustering"
		fmt.Println("KHIOPS_COCLUSTERING_CMD not set, default to " + khiopsCoclusteringCmd)
	} else {
		fmt.Println("KHIOPS_COCLUSTERING_CMD set to " + khiopsCoclusteringCmd)
	}
}

type decodeWriter struct {
	ow io.Writer
}

func newDecodeWriter(w io.Writer) *decodeWriter {
	return &decodeWriter{ow: w}
}

func (w *decodeWriter) Write(p []byte) (int, error) {
	dp := _Decode(p)
	//fmt.Printf("In: %d out: %d\n", len(p), len(dp))
	_, err := w.ow.Write(dp)
	return len(p), err
}

// RunBatch runs a khiops scenario.
func (b *Backend) RunBatch(ctx context.Context, in *khiops.BatchRequest) (*longrunning.Operation, error) {
	b.mu.Lock()
	defer b.mu.Unlock()

	f, err := ioutil.TempFile("/tmp", "scenario")
	check(err)
	scenario_path := f.Name()
	if s := in.GetScenario(); s != "" {
		_, err = f.WriteString(s)
		check(err)
	}
	if b := in.GetBinaryScenario(); b != nil {
		_, err = f.Write(b)
		check(err)
	}
	if p := in.GetScenarioPath(); p != "" {
		scenario_path = p
	}

	binary := khiopsCmd
	if in.Tool == khiops.BatchRequest_KHIOPS_COCLUSTERING {
		binary = khiopsCoclusteringCmd
	}
	cmd := exec.CommandContext(b.ctx, binary, "-b", "-i", scenario_path)

	job := &Job{
		cmd: cmd,
	}
	cmd.Stdout = io.MultiWriter(newDecodeWriter(os.Stdout), &job.stdoutBuf)
	cmd.Stderr = io.MultiWriter(newDecodeWriter(os.Stderr), &job.stderrBuf)

	job.ch = make(chan error)

	id := uuid.Must(uuid.NewV4()).String()
	b.items[id] = job
	fmt.Println("Items len=", len(b.items))

	op := &longrunning.Operation{
		Name: id,
	}

	err = cmd.Start()
	if err != nil {
		log.Printf("cmd.Start() failed with '%s'\n", err)
		job.error = &status.Status{
			Code:    int32(code.Code_INVALID_ARGUMENT),
			Message: fmt.Sprint(err),
		}
		op.Done = true
		op.Result = &longrunning.Operation_Error{
			Error: job.error,
		}
	} else {
		// wait for the program to end in a goroutine
		go func() {
			job.ch <- cmd.Wait()
			if err != nil {
				log.Printf("cmd.Wait() failed with '%s'\n", err)
			}
			// logic to run once process finished. Send err in channel if necessary
			defer os.Remove(f.Name())
		}()
	}

	return op, nil
}

func _FetchOperationStatus(name string, job *Job) *longrunning.Operation {
	out := new(longrunning.Operation)
	out.Name = name
	// Set other fields
	out.Done = job.error != nil || (job.cmd.ProcessState != nil && job.cmd.ProcessState.Exited())
	if job.error != nil {
		out.Result = &longrunning.Operation_Error{
			Error: job.error,
		}
	} else if job.cancelled {
		out.Result = &longrunning.Operation_Error{
			Error: &status.Status{
				Code:    int32(code.Code_CANCELLED),
				Message: "operation cancelled before completion",
			},
		}
	} else {
		if out.Done {
			resp := &khiops.BatchResponse{
				Status: int32(job.cmd.ProcessState.ExitCode()),
				Output: string(_Decode(job.stdoutBuf.Bytes())),
				Error:  string(_Decode(job.stderrBuf.Bytes())),
			}
			a, err := anypb.New(resp)
			if err != nil {
				fmt.Println(err)
			}
			out.Result = &longrunning.Operation_Response{
				Response: a,
			}
		}
	}
	return out
}

func _Decode(b []byte) []byte {
	e, _, _ := charset.DetermineEncoding(b, "")
	ub, _ := ioutil.ReadAll(transform.NewReader(bytes.NewReader(b), e.NewDecoder()))
	return ub
}

// ListOperations lists all current operations
func (b *Backend) ListOperations(ctx context.Context, in *longrunning.ListOperationsRequest) (*longrunning.ListOperationsResponse, error) {
	b.mu.RLock()
	defer b.mu.RUnlock()

	fmt.Println("List operations, CurLen:", len(b.items))

	out := &longrunning.ListOperationsResponse{}
	for name, job := range b.items {
		fmt.Println("Name:", name)
		out.Operations = append(out.Operations, _FetchOperationStatus(name, job))
	}

	return out, nil
}

// GetOperation retrieves the status of an ongoing operation.
func (b *Backend) GetOperation(ctx context.Context, in *longrunning.GetOperationRequest) (*longrunning.Operation, error) {
	b.mu.RLock()
	defer b.mu.RUnlock()

	fmt.Println("Get Name:", in.Name, "CurLen:", len(b.items))

	job, ok := b.items[in.Name]
	if ok {
		out := _FetchOperationStatus(in.Name, job)
		return out, nil
	} else {
		// Non-existent operation, return an error
		fmt.Println("Key:", in.Name, "not found, return error")
		out := new(longrunning.Operation)
		out.Name = in.Name
		out.Done = true
		out.Result = &longrunning.Operation_Error{
			Error: &status.Status{
				Code:    int32(code.Code_INVALID_ARGUMENT),
				Message: "Requested item does not exist.",
			},
		}
		return out, nil //status.Errorf(codes.InvalidArgument, "requested item does not exist")
	}
}

// DeleteOperation deletes an ongoing operation.
func (b *Backend) DeleteOperation(ctx context.Context, in *longrunning.DeleteOperationRequest) (*empty.Empty, error) {
	b.mu.RLock()
	defer b.mu.RUnlock()

	out := new(empty.Empty)
	fmt.Println("Delete Name:", in.Name, "CurLen:", len(b.items))

	job, ok := b.items[in.Name]
	if ok {
		// Kill process too since we do not want to hang the CPU
		err := job.cmd.Process.Kill()
		if err != nil {
			fmt.Println("Kill process failed: ", err)
		}
		err = job.cmd.Process.Release()
		if err != nil {
			fmt.Println("Release process failed: ", err)
		}
		delete(b.items, in.Name)
		return out, nil
	} else {
		// Non-existent operation, return an error
		fmt.Println("Key:", in.Name, "not found, return error")
		// FIXME: return error apparently does not work well for gRPC-REST conversion
		return out, nil //status1.Errorf(codes.NotFound, "not found")
	}
}

// CancelOperation cancels an ongoing operation.
func (b *Backend) CancelOperation(ctx context.Context, in *longrunning.CancelOperationRequest) (*empty.Empty, error) {
	b.mu.RLock()
	defer b.mu.RUnlock()

	out := new(empty.Empty)
	fmt.Println("Cancel Name:", in.Name, "CurLen:", len(b.items))

	job, ok := b.items[in.Name]
	if ok {
		err := job.cmd.Process.Kill()
		if err != nil {
			fmt.Println("Kill process failed: ", err)
		}
		job.cancelled = true
		return out, nil
	} else {
		// Non-existent operation, return an error
		fmt.Println("Key:", in.Name, "not found, return error")
		// FIXME: return error apparently does not work well for gRPC-REST conversion
		return out, nil //status1.Errorf(codes.NotFound, "not found")
	}
}

// WaitOperation holds caller until operation completes.
func (b *Backend) WaitOperation(ctx context.Context, in *longrunning.WaitOperationRequest) (*longrunning.Operation, error) {
	// Warning! Manual lock/unlock of mutex!
	fmt.Println("Wait Name:", in.Name, "CurLen:", len(b.items))

	b.mu.RLock()
	job, ok := b.items[in.Name]
	if ok {
		out := _FetchOperationStatus(in.Name, job)
		ch := job.ch
		b.mu.RUnlock()

		if out.Done {
			return out, nil
		} else {
			if in.Timeout != nil {
				fmt.Println("Start waiting for cmd, timeout=" + in.Timeout.String())
			} else {
				fmt.Println("Start waiting for cmd")
			}
			timeout := make(chan bool, 1)
			go func() {
				if in.Timeout != nil {
					time.Sleep(time.Duration(in.Timeout.Seconds * int64(time.Second)))
					timeout <- true
				}
			}()

			select {
			case err := <-ch:
				// done! check error
				if err != nil {
					log.Printf("cmd.Wait() failed with '%s'\n", err)
				}
				fmt.Println("Finished, return updated job status")
			case <-timeout:
				// the read from ch has timed out
				fmt.Println("Wait timed out, return updated job status")
			}
			// FIXME if job has been deleted, we should return error
			out = _FetchOperationStatus(in.Name, job)
			return out, nil
		}

	} else {
		// Non-existent operation, return an error
		b.mu.RUnlock()
		fmt.Println("Key:", in.Name, "not found, return error")
		out := new(longrunning.Operation)
		out.Name = in.Name
		out.Done = true
		out.Result = &longrunning.Operation_Error{
			Error: &status.Status{
				Code:    int32(code.Code_INVALID_ARGUMENT),
				Message: "Requested item does not exist.",
			},
		}
		return out, nil //status.Errorf(codes.InvalidArgument, "requested item does not exist")
	}
}
