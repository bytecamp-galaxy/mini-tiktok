FROM golang:1.19

ENV CGO_ENABLED 0
ENV GOOS linux
ENV GOPROXY https://goproxy.cn,direct

WORKDIR /root
COPY go.mod .
COPY go.sum .
RUN go mod download

COPY api-server api-server
COPY configs configs
COPY pkg pkg
# TODO(vgalaxy): archive kitex_gen
COPY user-server user-server
COPY comment-server comment-server
COPY favorite-server favorite-server
COPY feed-server feed-server
WORKDIR /root/api-server
RUN chmod +x build.sh
RUN ./build.sh

WORKDIR /root
ENTRYPOINT ["/root/api-server/output/bin/api"]