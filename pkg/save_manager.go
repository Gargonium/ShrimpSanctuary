package pkg

// В начале файла с Game или в отдельном файле save.go
import (
	"ShrimpSanctuary/internal/config"
	"ShrimpSanctuary/internal/game"
	"ShrimpSanctuary/internal/game/entities"
	"encoding/json"
	"fmt"
	rl "github.com/gen2brain/raylib-go/raylib"
	"os"
	"path/filepath"
	"time"
)

// SaveData - структура данных для сохранения
type SaveData struct {
	Version   string          `json:"version"`   // Версия игры для миграций
	Timestamp time.Time       `json:"timestamp"` // Когда сохранено
	Aquarium  AquariumData    `json:"aquarium"`  // Состояние аквариума
	Shrimps   []ShrimpData    `json:"shrimps"`   // Состояние креветок
	Pollution []PollutionData `json:"poluttion"`
}

// Упрощенные структуры для сериализации (без каналов, функций и т.д.)
type AquariumData struct {
	Money             int                     `json:"money"`
	Wallpaper         config.WallpaperState   `json:"wallpaper"`
	UnlockedWallpaper []config.WallpaperState `json:"unlockedWallpaper"`
}

type ShrimpData struct {
	Type config.ShrimpType `json:"type"`
}

type PollutionData struct {
	PositionX  float32
	PositionY  float32
	Durability int32
}

// SaveManager - менеджер сохранений
type SaveManager struct {
	SaveDir string
}

// NewSaveManager создает менеджер сохранений
func NewSaveManager() *SaveManager {
	// Определяем путь к папке сохранений

	cwd, _ := os.Getwd()
	saveDir := filepath.Join(cwd, "saves")
	println("0.", saveDir)

	// Создаем папку если не существует
	err := os.MkdirAll(saveDir, 0755)
	if err != nil {
		println(err)
	}

	println("3.", saveDir)
	return &SaveManager{
		SaveDir: saveDir,
	}
}

// getSavePath возвращает путь к файлу сохранения
func (sm *SaveManager) getSavePath() string {
	return filepath.Join(sm.SaveDir, fmt.Sprintf("save.json"))
}

// getBackupPath возвращает путь к бэкапу
func (sm *SaveManager) getBackupPath() string {
	return filepath.Join(sm.SaveDir, fmt.Sprintf("save.backup"))
}

// SaveGame сохраняет игру в указанный слот
func (sm *SaveManager) SaveGame(game *game.Game) error {
	// Сначала сохраняем бэкап
	if sm.SaveExists() {
		backupPath := sm.getBackupPath()
		savePath := sm.getSavePath()
		err := os.Rename(savePath, backupPath)
		if err != nil {
			println(err)
		}
	}

	// Подготавливаем данные для сохранения
	saveData := SaveData{
		Version:   "1.0",
		Timestamp: time.Now(),
		Aquarium: AquariumData{
			Money:             game.Money,
			UnlockedWallpaper: game.UnlockedWallpaper,
			Wallpaper:         game.WallpaperState,
		},
	}

	// Сохраняем креветок
	for _, shrimp := range game.Shrimps {
		shrimpData := ShrimpData{
			Type: shrimp.Type,
		}
		saveData.Shrimps = append(saveData.Shrimps, shrimpData)
	}
	for _, pol := range game.Pollution {
		polData := PollutionData{
			PositionX:  pol.Position.X,
			PositionY:  pol.Position.Y,
			Durability: pol.Durability,
		}
		saveData.Pollution = append(saveData.Pollution, polData)
	}

	// Сериализуем в JSON
	data, err := json.MarshalIndent(saveData, "", "  ")
	if err != nil {
		return err
	}

	// Сохраняем в файл
	savePath := sm.getSavePath()
	return os.WriteFile(savePath, data, 0644)
}

// LoadGame загружает игру из слота
func (sm *SaveManager) LoadGame(game *game.Game) error {
	savePath := sm.getSavePath()
	data, err := os.ReadFile(savePath)
	if err != nil {
		// Пробуем загрузить из бэкапа
		backupPath := sm.getBackupPath()
		data, err = os.ReadFile(backupPath)
		if err != nil {
			return err
		}
	}

	var saveData SaveData
	if err := json.Unmarshal(data, &saveData); err != nil {
		return err
	}

	game.Money = saveData.Aquarium.Money

	// Восстанавливаем креветок
	game.Shrimps = nil
	for _, shrimpData := range saveData.Shrimps {
		shrimp := entities.NewShrimp(shrimpData.Type)
		game.Shrimps = append(game.Shrimps, shrimp)
	}
	game.Pollution = nil
	for _, polData := range saveData.Pollution {
		pol := new(entities.Pollute)
		pol.Position = rl.NewVector2(polData.PositionX, polData.PositionY)
		pol.Durability = polData.Durability
		game.Pollution = append(game.Pollution, pol)
	}

	game.WallpaperState = saveData.Aquarium.Wallpaper
	game.UnlockedWallpaper = saveData.Aquarium.UnlockedWallpaper

	return nil
}

// AutoSave выполняет автосохранение
func (sm *SaveManager) AutoSave(game *game.Game) error {
	return sm.SaveGame(game)
}

func (sm *SaveManager) SaveExists() bool {
	path := sm.getSavePath()
	_, err := os.Stat(path)
	return err == nil
}
