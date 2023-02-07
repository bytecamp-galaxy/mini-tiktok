#FROM ubuntu AS build
#RUN apt-get update && \
#    apt-get install -y ffmpeg
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

WORKDIR /root/cmd/publish
RUN chmod +x build.sh \
    && ./build.sh \
    && sed -i "s@http://\(deb\|security\).debian.org@https://mirrors.tuna.tsinghua.edu.cn@g" /etc/apt/sources.list  \
    && apt update \
    && apt install -y netcat \
    && apt install -y ffmpeg

WORKDIR /root
ENTRYPOINT ["/root/cmd/publish/output/publish"]