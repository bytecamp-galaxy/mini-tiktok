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
COPY scripts/wait.sh wait.sh
COPY scripts/build_all.sh build_all.sh

RUN bash ./build_all.sh \
    && sed -i "s@http://\(deb\|security\).debian.org@https://mirrors.tuna.tsinghua.edu.cn@g" /etc/apt/sources.list  \
    && apt update \
    && apt install -y netcat