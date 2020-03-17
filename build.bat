@echo off

echo Generating app.json
go build github.com/Tnze/CoolQ-Golang-SDK/v2/tools/cqcfg
go generate
IF ERRORLEVEL 1 pause

echo Setting env vars
SET CGO_LDFLAGS=-Wl,--kill-at
SET CGO_ENABLED=1
SET GOOS=windows
SET GOARCH=386
SET GOPROXY=https://goproxy.cn

echo Building app.dll
go build -ldflags "-s -w" -buildmode=c-shared -o app.dll
IF ERRORLEVEL 1 pause

:: Copy app.dll amd app.json
SET DevDir=D:\酷Q Air\dev\com.poicraft.bot
if defined DevDir (
    echo Coping files
    for %%f in (app.dll,app.json) do move %%f "%DevDir%\%%f" > nul
    IF ERRORLEVEL 1 pause
)
