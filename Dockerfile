FROM golang:1.23-alpine as builder
LABEL authors="gucooing"

RUN sed -i 's/dl-cdn.alpinelinux.org/mirrors.tuna.tsinghua.edu.cn/g' /etc/apk/repositories
RUN apk add --no-cache gcc musl-dev linux-headers git bash
WORKDIR /usr/hkrpg
ADD go.mod .
ADD go.sum .
ENV CGO_ENABLED=1
RUN go mod download && go mod verify
COPY . .
RUN go build -tags netgo -o /usr/hkrpg/hkrpg-go ./cmd/hkrpg-go-pe/hkrpg-go.go

FROM alpine:latest
WORKDIR /usr/hkrpg
COPY --from=builder /usr/hkrpg/hkrpg-go /usr/hkrpg/hkrpg-go
COPY --from=builder /usr/hkrpg/data/ /usr/hkrpg/data/
COPY --from=builder /usr/hkrpg/start.sh /usr/hkrpg/start.sh
RUN chmod +x /usr/hkrpg/start.sh
EXPOSE 8080/tcp 20041/udp
ENTRYPOINT ["bash","./start.sh"]