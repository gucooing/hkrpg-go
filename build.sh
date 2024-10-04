#!/bin/bash
if command -v go &> /dev/null; then
    echo "Go 环境已安装"
    go version
else
    echo "Go 环境未安装"
    exit 1
fi

go mod download
go mod verify
export CGO_ENABLED=0

PLATFORMS="linux/amd64 linux/arm64 windows/amd64 windows/arm64"

OUT_DIR=./build

PE_PATH=./cmd/hkrpg-go/hkrpg-go.go
DISPATCH_PATH=./cmd/dispatch/dispatch.go
GAMESERVER_PATH=./cmd/gameserver/gameserver.go
GATE_PATH=./cmd/gateserver/gateserver.go
MUIP_PATH=./cmd/muipserver/muipserver.go
NODE_PATH=./cmd/nodeserver/nodeserver.go
ROBOT_PATH=./cmd/robot/robot.go

for file in $PE_PATH $DISPATCH_PATH $GAMESERVER_PATH $GATE_PATH $MUIP_PATH $NODE_PATH $ROBOT_PATH; do
  for platform in $PLATFORMS; do
      export GOOS=$(echo $platform | cut -d'/' -f1)
      export GOARCH=$(echo $platform | cut -d'/' -f2)
      CURRENT_OUT_DIR=$OUT_DIR/$GOOS-$GOARCH
      FILENAME=$(basename $file)
      OUTPUT_NAME=${FILENAME%.*}_$GOOS"_"$GOARCH
      if [ $GOOS = "windows" ]; then
        OUTPUT_NAME="$OUTPUT_NAME.exe"
      fi
      echo "Building $OUTPUT_NAME..."
      go build -ldflags="-s -w" -o $CURRENT_OUT_DIR/$OUTPUT_NAME $file
    done
done