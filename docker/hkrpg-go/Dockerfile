FROM golang:1.23-alpine as builder
LABEL authors="gucooing"

WORKDIR /usr/hkrpg
ADD go.mod .
ADD go.sum .
RUN go mod download && go mod verify
COPY . .
RUN go build -tags netgo -o /usr/hkrpg/hkrpg-go ./cmd/hkrpg-go/hkrpg-go.go

FROM alpine:latest
RUN apk add --no-cache git bash
WORKDIR /usr/hkrpg
COPY --from=builder /usr/hkrpg/hkrpg-go /usr/hkrpg/hkrpg-go
COPY --from=builder /usr/hkrpg/data/ /usr/hkrpg/data/
COPY --from=builder /usr/hkrpg/start.sh /usr/hkrpg/start.sh
RUN chmod +x /usr/hkrpg/start.sh
EXPOSE 8080/tcp 20041/udp
ENTRYPOINT ["bash","start.sh"]