# Build stage
FROM golang:1.17.3 as build

LABEL maintainer="Khiops Team <khiops.team@orange.com>"

RUN apt-get update && \
    apt-get install --no-install-recommends -y upx-ucl && \
    # apt-get install --no-install-recommends -y unzip && \
    apt-get clean && rm -rf /var/lib/apt/lists/*

#RUN curl -LO https://github.com/protocolbuffers/protobuf/releases/download/v3.14.0/protoc-3.14.0-linux-x86_64.zip
#RUN unzip protoc-3.14.0-linux-x86_64.zip
COPY . /src
ENV CGO_ENABLED 0
WORKDIR /src
#RUN make install 
#RUN make generate
RUN go build -ldflags="-s -w" -o /service && ls -lh /service

RUN upx-ucl --brute /service && ls -lh /service
RUN useradd -rm -d /home/ubuntu -s /bin/bash ubuntu
USER ubuntu

# Runtime stage
FROM busybox:glibc as runtime
COPY --from=build /service /service

ENTRYPOINT ["/service"]
