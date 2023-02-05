FROM golang:1.19

ENV CGO_ENABLED 0
ENV GOOS linux
ENV GOPROXY https://goproxy.cn,direct

WORKDIR /root
COPY go.mod .
COPY go.sum .
RUN go mod download

COPY cmd cmd
COPY kitex_gen kitex_gen
COPY configs configs
COPY pkg pkg
COPY internal internal

RUN cd configs \
    && rm global.yaml \
    && mv docker.yaml global.yaml

WORKDIR /root/cmd/user
RUN chmod +x build.sh
RUN ./build.sh

WORKDIR /root
ENTRYPOINT ["/root/cmd/user/output/user"]