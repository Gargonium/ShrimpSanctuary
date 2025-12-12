package sound_bar

import (
	"ShrimpSanctuary/internal/config"
	rl "github.com/gen2brain/raylib-go/raylib"
)

type SoundBar struct {
	background rl.Music
}

func NewSoundBar() SoundBar {
	sb := SoundBar{}
	sb.background = loadSound(config.BgMusicPath)

	if !rl.IsAudioDeviceReady() {
		rl.TraceLog(rl.LogError, "Аудиоустройство не готово! Проверьте драйверы.")
	}

	return sb
}

func (sb *SoundBar) Update() {
	rl.UpdateMusicStream(sb.background)
}

func (sb *SoundBar) PlayBgMusic() {
	sb.background.Looping = true
	rl.SetMusicVolume(sb.background, 0.5)
	rl.PlayMusicStream(sb.background)
}

func (sb *SoundBar) StopBgMusic() {
	rl.StopMusicStream(sb.background)
	rl.UnloadMusicStream(sb.background)
}

func loadSound(musicPath string) rl.Music {
	music := rl.LoadMusicStream(musicPath)

	if music.Stream.Buffer == nil {
		rl.TraceLog(rl.LogError, "Не удалось загрузить музыку! Проверьте путь к файлу.")
	}
	return music
}
