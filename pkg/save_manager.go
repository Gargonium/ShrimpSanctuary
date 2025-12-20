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
	Version        string          `json:"version"`   // Версия игры для миграций
	Timestamp      time.Time       `json:"timestamp"` // Когда сохранено
	Aquarium       AquariumData    `json:"aquarium"`  // Состояние аквариума
	Shrimps        []ShrimpData    `json:"shrimps"`   // Состояние креветок
	Pollution      []PollutionData `json:"poluttion"`
	SettingsData   SettingsData    `json:"settings"`
	StatisticsData StatisticsData  `json:"statistics"`
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
	PositionX  float32 `json:"position-x"`
	PositionY  float32 `json:"position-y"`
	Durability int32   `json:"durability"`
}

type SettingsData struct {
	MusicVolume   float32 `json:"music-volume"`
	EffectsVolume float32 `json:"effects-volume"`
}

type StatisticsData struct {
	Achievements    []bool                    `json:"achievements"`
	ShrimpsFed      int                       `json:"shrimpsFed"`
	AquariumCleaned int                       `json:"aquariumCleaned"`
	ShrimpsCount    map[config.ShrimpType]int `json:"shrimpsCount"`
	WallpapersCount int                       `json:"wallpapersCount"`
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

	// Создаем папку если не существует
	err := os.MkdirAll(saveDir, 0755)
	if err != nil {
		println(err)
	}

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
		SettingsData: SettingsData{
			MusicVolume:   game.SoundBar.GetMusicVolume(),
			EffectsVolume: game.SoundBar.GetEffectsVolume(),
		},
		StatisticsData: StatisticsData{
			Achievements:    game.Statistics.Achievements,
			ShrimpsFed:      game.Statistics.ShrimpsFed,
			AquariumCleaned: game.Statistics.AquariumCleaned,
			ShrimpsCount:    game.Statistics.ShrimpsCount,
			WallpapersCount: game.Statistics.WallpapersCount,
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

	game.SoundBar.ChangeMusicVolume(saveData.SettingsData.MusicVolume)
	game.SoundBar.ChangeEffectsVolume(saveData.SettingsData.EffectsVolume)

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

	game.Statistics.Achievements = saveData.StatisticsData.Achievements
	game.Statistics.ShrimpsFed = saveData.StatisticsData.ShrimpsFed
	game.Statistics.AquariumCleaned = saveData.StatisticsData.AquariumCleaned
	game.Statistics.ShrimpsCount = saveData.StatisticsData.ShrimpsCount
	game.Statistics.WallpapersCount = saveData.StatisticsData.WallpapersCount

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
