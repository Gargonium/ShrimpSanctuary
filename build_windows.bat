@echo off
setlocal enabledelayedexpansion



REM Сборка для выбранной платформы
echo Start build for windows amd64...
set GOOS=windows
set GOARCH=amd64

go build -ldflags="-H windowsgui -s -w" -o "ShrimpSanctuary_windows.exe" ./cmd