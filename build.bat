@echo off
setlocal enabledelayedexpansion

REM Выбор архитектуры
echo Choose architecture:
echo 1. Windows 64-bit (default)
echo 2. Windows 32-bit
echo 3. Linux 64-bit
echo 4. macOS 64-bit
set /p choice="Enter the number [1]: "

if "%choice%"=="" set choice=1



if "%choice%"=="2" (
    set ARCH=386
    set SUFFIX=_32bit.exe
    set OS=windows
) else if "%choice%"=="3" (
    set ARCH=amd64
    set SUFFIX=_linux
    set OS=linux
) else if "%choice%"=="4" (
    set ARCH=amd64
    set SUFFIX=_macos
    set OS=darwin
) else (
    set ARCH=amd64
    set SUFFIX=.exe
    set OS=windows
)

REM Создаем папку для релиза с указанием архитектуры
set RELEASE_FOLDER=release_%OS%_%ARCH%
if not exist "%RELEASE_FOLDER%" mkdir "%RELEASE_FOLDER%"

REM Сборка для выбранной платформы
echo Start build for %OS% %ARCH%...
set GOOS=%OS%
set GOARCH=%ARCH%
if "%OS%"=="windows" (
    go build -ldflags="-H windowsgui -s -w" -o "%RELEASE_FOLDER%\ShrimpSanctuary%SUFFIX%" ./cmd
)  else (
    go build -ldflags="-s -w" -o "%RELEASE_FOLDER%\ShrimpSanctuary%SUFFIX%" ./cmd
)


cd "%RELEASE_FOLDER%"
powershell "Compress-Archive -Path * -DestinationPath 'ShrimpSanctuary_%OS%_%ARCH%.zip' -Force"
cd ..

echo.
echo Archive created: %RELEASE_FOLDER%\ShrimpSanctuary_%OS%_%ARCH%.zip
pause