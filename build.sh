#!/bin/bash
if command -v go &> /dev/null; then
    echo "Go 环境已安装"
    go version
else
    echo "Go 环境未安装"
    exit 1
fi

set CGO_ENABLED=1
go mod download
go mod verify

GOENV=$(go env GOHOSTOS)
GOENV="$GOENV/$(go env GOHOSTARCH)"
echo "GOENV $GOENV"

OUT_DIR=./build

PE_PATH=./cmd/hkrpg-go/hkrpg-go.go
DISPATCH_PATH=./cmd/dispatch/dispatch.go
GAMESERVER_PATH=./cmd/gameserver/gameserver.go
GATE_PATH=./cmd/gateserver/gateserver.go
MUIP_PATH=./cmd/muipserver/muipserver.go
MULTI_PATH=./cmd/multiserver/multiserver.go
NODE_PATH=./cmd/nodeserver/nodeserver.go
ROBOT_PATH=./cmd/robot/robot.go

for file in $PE_PATH $DISPATCH_PATH $GAMESERVER_PATH $GATE_PATH $MUIP_PATH $MULTI_PATH $NODE_PATH $ROBOT_PATH; do
    FILENAME=$(basename $file)
    OUTPUT_NAME=${FILENAME%.*}
    echo "Building $FILENAME..."
    go build -ldflags="-s -w" -o $OUT_DIR/$OUTPUT_NAME $file
done