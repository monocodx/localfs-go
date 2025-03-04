@echo off

@rem remove previous build output
rmdir "bin" /s /q

@rem remove build cache
go clean -cache

@rem executable binary
go build -C src  -o ..\bin\localfs.exe  -ldflags="-s -w" -x

pause