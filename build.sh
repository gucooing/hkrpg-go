#!/bin/bash
if command -v go &> /dev/null; then
    echo "Go 环境已安装"
    go version
else
    echo "Go 环境未安装"
    exit 1
fi

go mod tidy

OUT_DIR=./build
PROJECT_NAME=hkrpg-go
TARGET_PLATFORMS="linux/amd64 linux/arm64 windows/amd64 windows/arm64"

DISPATCH_PATH=./cmd/dispatch/dispatch.go
GAMESERVER_PATH=./cmd/gameserver/gameserver.go
GATE_PATH=./cmd/gateserver/gateserver.go
MUIP_PATH=./cmd/muipserver/muipserver.go
MULTI_PATH=./cmd/multiserver/multiserver.go
NODE_PATH=./cmd/nodeserver/nodeserver.go
ROBOT_PATH=./cmd/robot/robot.go

for file in $MAIN_PATH $DISPATCH_PATH $GAMESERVER_PATH $GATE_PATH $MUIP_PATH $MULTI_PATH $NODE_PATH $ROBOT_PATH; do
  for platform in $TARGET_PLATFORMS; do
    GOOS=$(echo $platform | cut -d'/' -f1)
    GOARCH=$(echo $platform | cut -d'/' -f2)
    FILENAME=$(basename $file)
    OUTPUT_NAME=${FILENAME%.*}_$GOOS"_"$GOARCH
    if [ $GOOS = "windows" ]; then
      OUTPUT_NAME="$OUTPUT_NAME.exe"
    fi
    echo "Building $OUTPUT_NAME..."
    env GOOS=$GOOS GOARCH=$GOARCH go build -o $OUT_DIR/$OUTPUT_NAME $file
  done
done