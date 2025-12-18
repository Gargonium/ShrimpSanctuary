@echo off
setlocal enabledelayedexpansion

REM Создаем папку для релиза с указанием архитектуры
set RELEASE_FOLDER=release_windows_amd64
if not exist "%RELEASE_FOLDER%" mkdir "%RELEASE_FOLDER%"

REM Сборка для выбранной платформы
echo Start build for windows amd64...
set GOOS=windows
set GOARCH=amd64

go build -ldflags="-H windowsgui -s -w" -o "%RELEASE_FOLDER%\ShrimpSanctuary.exe" ./cmd

pause