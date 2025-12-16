package sound_bar

import (
	"ShrimpSanctuary/internal/config"
	rl "github.com/gen2brain/raylib-go/raylib"
)

type SoundBar struct {
	isMuted       bool
	musicVolume   float32
	effectsVolume float32
	background    rl.Music
}

func NewSoundBar() *SoundBar {
	sb := new(SoundBar)
	sb.background = loadSound(config.BgMusicPath)
	sb.isMuted = false

	if !rl.IsAudioDeviceReady() {
		rl.TraceLog(rl.LogError, "Аудиоустройство не готово! Проверьте драйверы.")
	}

	return sb
}

func (sb *SoundBar) Update() {
	rl.UpdateMusicStream(sb.background)
}

func (sb *SoundBar) ChangeMusicVolume(volume float32) {
	rl.SetMusicVolume(sb.background, volume)
	if !sb.isMuted {
		sb.musicVolume = volume
	}
}

func (sb *SoundBar) ChangeEffectsVolume(volume float32) {
	if !sb.isMuted {
		sb.effectsVolume = volume
	}
}

func (sb *SoundBar) Mute() {
	sb.isMuted = !sb.isMuted
	if sb.isMuted {
		sb.ChangeMusicVolume(0)
		sb.ChangeEffectsVolume(0)
	} else {
		sb.ChangeMusicVolume(sb.musicVolume)
		sb.ChangeEffectsVolume(sb.effectsVolume)
	}
}

func (sb *SoundBar) PlayBgMusic() {
	sb.background.Looping = true
	sb.musicVolume = 0.1
	sb.effectsVolume = 0.1
	rl.SetMusicVolume(sb.background, sb.musicVolume)
	rl.PlayMusicStream(sb.background)
}

func (sb *SoundBar) StopBgMusic() {
	rl.StopMusicStream(sb.background)
	rl.UnloadMusicStream(sb.background)
}

func (sb *SoundBar) IsMuted() bool {
	return sb.isMuted
}

func loadSound(musicPath string) rl.Music {
	music := rl.LoadMusicStream(musicPath)

	if music.Stream.Buffer == nil {
		rl.TraceLog(rl.LogError, "Не удалось загрузить музыку! Проверьте путь к файлу.")
	}
	return music
}

func (sb *SoundBar) GetMusicVolume() float32 {
	return sb.musicVolume
}

func (sb *SoundBar) GetEffectsVolume() float32 {
	return sb.effectsVolume
}
