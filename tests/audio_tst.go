package main

import (
	"fmt"
	rl "github.com/gen2brain/raylib-go/raylib"
	"runtime"
)

func main() {
	fmt.Printf("OS: %s, Arch: %s\n", runtime.GOOS, runtime.GOARCH)

	// 1. Включаем аудио через флаги (корректный путь)
	// Флаг для аудио называется FlagAudioDevice, а не FlagAudioDeviceEnabled
	rl.SetConfigFlags(rl.FlagWindowResizable)

	// 2. Инициализация окна
	rl.InitWindow(800, 450, "Raylib Audio Debug")
	defer rl.CloseWindow()

	// 3. Инициализация аудио
	fmt.Println("Инициализирую аудиоустройство...")
	rl.InitAudioDevice()
	defer rl.CloseAudioDevice()

	// 4. Проверка готовности
	isReady := rl.IsAudioDeviceReady()
	fmt.Printf("Аудиоустройство готово: %v\n", isReady)

	if isReady {
		fmt.Println("✅ SUCCESS: Аудио инициализировано!")

		// 5. Пробуем загрузить и воспроизвести очень простой WAV-файл
		// Создайте тестовый файл test.wav в той же папке что и программа
		sound := rl.LoadSound("../assets/sounds/Background.wav")

		if sound.Stream.Buffer != nil {
			fmt.Println("Звук загружен, воспроизвожу...")
			rl.PlaySound(sound)

			// Ждем 2 секунды
			rl.SetTargetFPS(60)
			frames := 0
			for frames < 120 {
				rl.BeginDrawing()
				rl.ClearBackground(rl.Black)
				rl.DrawText("Звук играет...", 50, 200, 30, rl.Green)
				rl.EndDrawing()
				frames++
			}

			rl.UnloadSound(sound)
			fmt.Println("Тест завершен успешно!")
		} else {
			fmt.Println("❌ Не удалось загрузить тестовый звук")
		}

	} else {
		fmt.Println("❌ FAIL: Аудиоустройство не инициализировано.")
		fmt.Println("\nВозможные причины в порядке вероятности:")
		fmt.Println("1. Конфликт с другими аудиопрограммами (закройте Discord, браузер и т.д.)")
		fmt.Println("2. Устаревшая версия raylib-go")
		fmt.Println("3. Отсутствуют необходимые системные библиотеки")
		fmt.Println("4. Антивирус блокирует доступ")

		// Просто показываем окно с ошибкой
		rl.SetTargetFPS(60)
		for !rl.WindowShouldClose() {
			rl.BeginDrawing()
			rl.ClearBackground(rl.Red)
			rl.DrawText("АУДИО НЕ РАБОТАЕТ!", 50, 150, 40, rl.White)
			rl.DrawText("1. Закройте все программы использующие звук", 50, 220, 20, rl.White)
			rl.DrawText("2. Проверьте настройки звука Windows", 50, 250, 20, rl.White)
			rl.DrawText("3. Нажмите ESC для выхода", 50, 300, 20, rl.White)
			rl.EndDrawing()
		}
	}
}
