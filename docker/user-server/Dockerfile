FROM golang:1.19

ENV CGO_ENABLED 0
ENV GOOS linux
ENV GOPROXY https://goproxy.cn,direct

WORKDIR /root
COPY go.mod .
COPY go.sum .
RUN go mod download

COPY user-server user-server
COPY configs configs
COPY pkg pkg
WORKDIR /root/user-server
RUN chmod +x build.sh
RUN ./build.sh

WORKDIR /root
ENTRYPOINT ["/root/user-server/output/bin/user"]