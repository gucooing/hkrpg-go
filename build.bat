@echo off
setlocal enabledelayedexpansion

go mod download
go mod verify

set "PLATFORMS=windows/amd64 windows/arm64 linux/amd64 linux/arm64"

set "OUT_DIR=./build"

set "PE_PATH=.\cmd\hkrpg-go\hkrpg-go.go"
set "DISPATCH_PATH=.\cmd\dispatch\dispatch.go"
set "GAMESERVER_PATH=.\cmd\gameserver\gameserver.go"
set "GATE_PATH=.\cmd\gateserver\gateserver.go"
set "MUIP_PATH=.\cmd\muipserver\muipserver.go"
set "MULTI_PATH=.\cmd\multiserver\multiserver.go"
set "NODE_PATH=.\cmd\nodeserver\nodeserver.go"
set "ROBOT_PATH=.\cmd\robot\robot.go"

for %%p in (%PLATFORMS%) do (
    for /f "tokens=1,2 delims=/" %%a in ("%%p") do (
        set "GOOS=%%a"
        set "GOARCH=%%b"

        echo Compiling for GOOS=!GOOS! GOARCH=!GOARCH!...

        set "CURRENT_OUT_DIR=!OUT_DIR!/!GOOS!-!GOARCH!"
        if not exist "!CURRENT_OUT_DIR!" mkdir "!CURRENT_OUT_DIR!"


        for %%f in ("%PE_PATH%" "%DISPATCH_PATH%" "%GAMESERVER_PATH%" "%GATE_PATH%" "%MUIP_PATH%" "%MULTI_PATH%" "%NODE_PATH%" "%ROBOT_PATH%") do (
            set "FILENAME=%%~nxf"
            set "OUTPUT_NAME=%%~nf"

            if "!GOOS!"=="windows" (
                go build -ldflags="-s -w" -o "!CURRENT_OUT_DIR!/!OUTPUT_NAME!_!GOOS!_!GOARCH!.exe" %%f
            ) else (
                go build -ldflags="-s -w" -o "!CURRENT_OUT_DIR!/!OUTPUT_NAME!_!GOOS!_!GOARCH!" %%f
            )
        )
    )
)

endlocal
