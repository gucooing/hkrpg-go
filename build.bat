@echo off
setlocal enabledelayedexpansion

go mod download
go mod verify

set CGO_ENABLED=1
set GOARCH=amd64
set GOOS=windows

for /f "delims=" %%i in ('go env GOHOSTOS') do set GOENV=%%i
for /f "delims=" %%i in ('go env GOHOSTARCH') do set GOENV=!GOENV!/%%i

echo GOENV !GOENV!

set "OUT_DIR=./build"

set "PE_PATH=.\cmd\hkrpg-go-pe\hkrpg-go.go"
set "DISPATCH_PATH=.\cmd\dispatch\dispatch.go"
set "GAMESERVER_PATH=.\cmd\gameserver\gameserver.go"
set "GATE_PATH=.\cmd\gateserver\gateserver.go"
set "MUIP_PATH=.\cmd\muipserver\muipserver.go"
set "MULTI_PATH=.\cmd\multiserver\multiserver.go"
set "NODE_PATH=.\cmd\nodeserver\nodeserver.go"
set "ROBOT_PATH=.\cmd\robot\robot.go"

for %%f in ("%PE_PATH%" "%DISPATCH_PATH%" "%GAMESERVER_PATH%" "%GATE_PATH%" "%MUIP_PATH%" "%MULTI_PATH%" "%NODE_PATH%" "%ROBOT_PATH%") do (
    set "FILENAME=%%~nxf"
    set "OUTPUT_NAME=%%~nf"
    echo Building !FILENAME!...
    go build -ldflags="-s -w" -o "!OUT_DIR!/!OUTPUT_NAME!.exe" %%f
)

endlocal