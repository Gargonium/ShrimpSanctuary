#!/bin/bash
# build_linux.sh - скрипт сборки для Linux

echo "========================================"
echo "  Building Shrimp Sanctuary for Linux"
echo "========================================"

# Переходим в папку проекта (если скрипт запускается не оттуда)
cd "$(dirname "$0")" || cd /mnt/c/Projects/Go/ShrimpSanctuary


echo "📁 Current directory: $(pwd)"

# 1. Проверяем зависимости
echo "🔍 Checking dependencies..."
if ! command -v go &> /dev/null; then
    echo "❌ Go not installed!"
    echo "Install: sudo apt install golang"
    exit 1
fi

# 2. Устанавливаем системные зависимости
echo "📦 Installing system libraries..."
sudo apt update > /dev/null 2>&1
sudo apt install -y gcc libgl1-mesa-dev libxi-dev xorg-dev > /dev/null 2>&1

# 3. Чистим go.mod
echo "🧹 Cleaning go.mod..."
if grep -q "toolchain" go.mod; then
    echo "Removing toolchain directive..."
    sed -i '/toolchain/d' go.mod
fi

# 4. Обновляем модули
echo "🔄 Updating Go modules..."
go mod tidy

# 5. Компилируем
echo "⚡ Compiling..."
CGO_ENABLED=1 GOOS=linux GOARCH=amd64 \
go build -ldflags="-s -w" \
-o ShrimpSanctuary_linux \
../cmd

# 6. Проверяем результат
if [ -f "ShrimpSanctuary_linux" ]; then
    echo "✅ SUCCESS: ShrimpSanctuary_linux created!"
    echo ""
    echo "📊 File information:"
    echo "  Size: $(du -h ShrimpSanctuary_linux | cut -f1)"
    echo "  Type: $(file ShrimpSanctuary_linux | cut -d: -f2)"
    echo ""
    echo "🚀 To run: ./ShrimpSanctuary_linux"
    echo "📁 File is in: $(pwd)/ShrimpSanctuary_linux"

    # Копируем в Windows для удобства (опционально)
    if [ -d "/mnt/c" ]; then
        cp ShrimpSanctuary_linux /mnt/c/Users/$USER/Desktop/ShrimpSanctuary_linux 2>/dev/null
        echo "📋 Also copied to Windows Desktop"
    fi
else
    echo "❌ FAILED: Build unsuccessful"
    echo "Check errors above"
    exit 1
fi

echo "========================================"