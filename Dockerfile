FROM golang:1.19 AS build

ENV CGO_ENABLED 0
ENV GOOS linux
ENV GOPROXY https://goproxy.cn,direct

WORKDIR /root
COPY go.mod .
COPY go.sum .
RUN go mod download

COPY cmd cmd
COPY kitex_gen kitex_gen
COPY pkg pkg
COPY internal internal
COPY scripts/build_all.sh build_all.sh

RUN bash ./build_all.sh

FROM ubuntu AS script
WORKDIR /root
COPY scripts/wait.sh wait.sh
COPY configs configs

FROM golang:1.19 AS runtime

ENV CGO_ENABLED 0
ENV GOOS linux
ENV GOPROXY https://goproxy.cn,direct

WORKDIR /root
COPY --from=build /root .
COPY --from=script /root/wait.sh ./wait.sh
COPY --from=script /root/configs ./configs
COPY scripts/wait.sh wait.sh
RUN sed -i "s@http://\(deb\|security\).debian.org@https://mirrors.tuna.tsinghua.edu.cn@g" /etc/apt/sources.list \
    && apt update \
    && apt install -y netcat \
    && apt install -y ffmpeg \
    && apt install -y wget \
    && wget https://ghproxy.com/https://raw.githubusercontent.com/eficode/wait-for/v2.2.3/wait-for
