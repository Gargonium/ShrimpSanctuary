package sound_bar

import (
	"ShrimpSanctuary/assets"
	rl "github.com/gen2brain/raylib-go/raylib"
)

type SoundBar struct {
	isMuted       bool
	musicVolume   float32
	effectsVolume float32
	ts            *assets.AssetStorage
}

func NewSoundBar(ts *assets.AssetStorage) *SoundBar {
	sb := new(SoundBar)
	sb.ts = ts
	sb.isMuted = false

	if !rl.IsAudioDeviceReady() {
		rl.TraceLog(rl.LogError, "Аудиоустройство не готово! Проверьте драйверы.")
	}

	sb.musicVolume = 0.1
	sb.effectsVolume = 0.1

	return sb
}

func (sb *SoundBar) Update() {
	rl.UpdateMusicStream(sb.ts.BackgroundMusic)
	rl.UpdateMusicStream(sb.ts.AquariumSound)
	rl.UpdateMusicStream(sb.ts.CleanSound)
	rl.UpdateMusicStream(sb.ts.FoodDropSound)
}

func (sb *SoundBar) ChangeMusicVolume(volume float32) {
	rl.SetMusicVolume(sb.ts.BackgroundMusic, volume)
	if !sb.isMuted {
		sb.musicVolume = volume
	}
}

func (sb *SoundBar) ChangeEffectsVolume(volume float32) {
	rl.SetMusicVolume(sb.ts.AquariumSound, volume)
	rl.SetMusicVolume(sb.ts.CleanSound, volume)
	rl.SetMusicVolume(sb.ts.FoodDropSound, volume)
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
	sb.ts.BackgroundMusic.Looping = true
	rl.SetMusicVolume(sb.ts.BackgroundMusic, sb.musicVolume)
	rl.PlayMusicStream(sb.ts.BackgroundMusic)
}

func (sb *SoundBar) StopBgMusic() {
	rl.StopMusicStream(sb.ts.BackgroundMusic)
}

func (sb *SoundBar) PlayAquariumSound() {
	sb.ts.AquariumSound.Looping = true
	rl.SetMusicVolume(sb.ts.AquariumSound, sb.effectsVolume)
	rl.PlayMusicStream(sb.ts.AquariumSound)
}

func (sb *SoundBar) StopAquariumSound() {
	rl.StopMusicStream(sb.ts.AquariumSound)
}

func (sb *SoundBar) PlayCleanSound() {
	sb.ts.CleanSound.Looping = false
	rl.SetMusicVolume(sb.ts.CleanSound, sb.effectsVolume)
	rl.PlayMusicStream(sb.ts.CleanSound)
}

func (sb *SoundBar) StopCleanSound() {
	rl.StopMusicStream(sb.ts.CleanSound)
}

func (sb *SoundBar) PlayFoodDropSound() {
	sb.ts.FoodDropSound.Looping = false
	rl.SetMusicVolume(sb.ts.FoodDropSound, sb.effectsVolume)
	rl.PlayMusicStream(sb.ts.FoodDropSound)
}

func (sb *SoundBar) StopFoodDropSound() {
	rl.StopMusicStream(sb.ts.FoodDropSound)
}

func (sb *SoundBar) UnloadAll() {
	rl.UnloadMusicStream(sb.ts.BackgroundMusic)
	rl.UnloadMusicStream(sb.ts.AquariumSound)
	rl.UnloadMusicStream(sb.ts.CleanSound)
	rl.UnloadMusicStream(sb.ts.FoodDropSound)
}

func (sb *SoundBar) IsMuted() bool {
	return sb.isMuted
}

func (sb *SoundBar) GetMusicVolume() float32 {
	return sb.musicVolume
}

func (sb *SoundBar) GetEffectsVolume() float32 {
	return sb.effectsVolume
}
