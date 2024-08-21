FROM golang:1.23-alpine as builder
LABEL authors="gucooing"

RUN sed -i 's/dl-cdn.alpinelinux.org/mirrors.tuna.tsinghua.edu.cn/g' /etc/apk/repositories
RUN apk add --no-cache gcc musl-dev linux-headers
WORKDIR /usr/hkrpg
ADD go.mod .
ADD go.sum .
RUN go env -w GO111MODULE=on
RUN go env -w GOPROXY=https://goproxy.cn,direct
ENV CGO_ENABLED=1
RUN go mod download && go mod verify
COPY . .
RUN go build -tags netgo -o /usr/hkrpg/hkrpg-go hkrpg-go-pe.go

FROM alpine:latest
WORKDIR /usr/hkrpg
COPY --from=builder /usr/hkrpg/hkrpg-go /usr/hkrpg/hkrpg-go
COPY --from=builder /usr/hkrpg/data/ /usr/hkrpg/data/
EXPOSE 8080/tcp 20041/udp
ENTRYPOINT ["./hkrpg-go"]