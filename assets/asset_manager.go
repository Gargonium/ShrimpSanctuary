package assets

import (
	"embed"
	rl "github.com/gen2brain/raylib-go/raylib"
	"os"
	"path/filepath"
)

//go:embed assets/fonts/*.ttf
var fontAssets embed.FS

//go:embed assets/sounds/*.mp3
var soundAssets embed.FS

//go:embed assets/sprites/Other/*.png
var otherTextureAssets embed.FS

//go:embed assets/sprites/Screens/*.png
var screensTextureAssets embed.FS

//go:embed assets/sprites/Shrimps/*.png
var shrimpsTextureAssets embed.FS

//go:embed assets/sprites/Wallpapers/*.png
var wallpapersTextureAssets embed.FS

type AssetManager struct {
	tempFiles []string // Для очистки
}

func NewAssetManager() *AssetManager {
	return new(AssetManager)
}

func (am *AssetManager) LoadMusic(path string) rl.Music {
	data, err := soundAssets.ReadFile(path)
	if err != nil {
		// Возвращаем пустую музыку
		return rl.Music{}
	}

	// Создаем временный файл
	tmpFile, err := os.CreateTemp("", "music_*"+filepath.Ext(path))
	if err != nil {
		panic(err)
	}
	defer func(tmpFile *os.File) {
		err := tmpFile.Close()
		if err != nil {
			println(err)
		}
	}(tmpFile)

	_, _ = tmpFile.Write(data)
	am.tempFiles = append(am.tempFiles, tmpFile.Name())

	return rl.LoadMusicStream(tmpFile.Name())
}

// LoadFont загружает шрифт
func (am *AssetManager) LoadFont(path string) rl.Font {
	data, err := fontAssets.ReadFile(path)
	if err != nil {
		panic("Шрифт не найден: " + path)
	}

	// Некоторые версии raylib-go поддерживают загрузку шрифтов из памяти
	// Если нет - используем временный файл
	tmpFile, err := os.CreateTemp("", "font_*"+filepath.Ext(path))
	if err != nil {
		panic(err)
	}
	_, _ = tmpFile.Write(data)
	_ = tmpFile.Close()

	am.tempFiles = append(am.tempFiles, tmpFile.Name())

	return rl.LoadFont(tmpFile.Name())
}

// LoadTexture загружает текстуру
func (am *AssetManager) LoadOtherTexture(path string) rl.Texture2D {
	data, err := otherTextureAssets.ReadFile(path)
	if err != nil {
		panic("Текстура не найдена: " + path)
	}

	// Raylib пока не поддерживает загрузку текстур из памяти напрямую
	// Поэтому используем временный файл
	tmpFile, err := os.CreateTemp("", "texture_*"+filepath.Ext(path))
	if err != nil {
		panic(err)
	}
	_, _ = tmpFile.Write(data)
	_ = tmpFile.Close()

	am.tempFiles = append(am.tempFiles, tmpFile.Name())

	texture := rl.LoadTexture(tmpFile.Name())
	return texture
}

func (am *AssetManager) LoadScreensTexture(path string) rl.Texture2D {
	data, err := screensTextureAssets.ReadFile(path)
	if err != nil {
		panic("Текстура не найдена: " + path)
	}

	// Raylib пока не поддерживает загрузку текстур из памяти напрямую
	// Поэтому используем временный файл
	tmpFile, err := os.CreateTemp("", "texture_*"+filepath.Ext(path))
	if err != nil {
		panic(err)
	}
	_, _ = tmpFile.Write(data)
	_ = tmpFile.Close()

	am.tempFiles = append(am.tempFiles, tmpFile.Name())

	texture := rl.LoadTexture(tmpFile.Name())
	return texture
}

func (am *AssetManager) LoadShrimpsTexture(path string) rl.Texture2D {
	data, err := shrimpsTextureAssets.ReadFile(path)
	if err != nil {
		panic("Текстура не найдена: " + path)
	}

	// Raylib пока не поддерживает загрузку текстур из памяти напрямую
	// Поэтому используем временный файл
	tmpFile, err := os.CreateTemp("", "texture_*"+filepath.Ext(path))
	if err != nil {
		panic(err)
	}
	_, _ = tmpFile.Write(data)
	_ = tmpFile.Close()

	am.tempFiles = append(am.tempFiles, tmpFile.Name())

	texture := rl.LoadTexture(tmpFile.Name())
	return texture
}

func (am *AssetManager) LoadWallpapersTexture(path string) rl.Texture2D {
	data, err := wallpapersTextureAssets.ReadFile(path)
	if err != nil {
		panic("Текстура не найдена: " + path)
	}

	// Raylib пока не поддерживает загрузку текстур из памяти напрямую
	// Поэтому используем временный файл
	tmpFile, err := os.CreateTemp("", "texture_*"+filepath.Ext(path))
	if err != nil {
		panic(err)
	}
	_, _ = tmpFile.Write(data)
	_ = tmpFile.Close()

	am.tempFiles = append(am.tempFiles, tmpFile.Name())

	texture := rl.LoadTexture(tmpFile.Name())
	return texture
}

// Cleanup удаляет временные файлы
func (am *AssetManager) Cleanup() {
	for _, file := range am.tempFiles {
		_ = os.Remove(file)
	}
	am.tempFiles = nil
}
