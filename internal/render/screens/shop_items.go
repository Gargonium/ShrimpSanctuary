package screens

import (
	"ShrimpSanctuary/internal/config"
	"ShrimpSanctuary/internal/input"
)

type ShrimpItem struct {
	BuyButton  *input.Button
	Cost       int
	ShrimpType config.ShrimpType
}

func NewShrimpItem(btn *input.Button, st int) *ShrimpItem {
	si := new(ShrimpItem)
	si.BuyButton = btn
	si.ShrimpType = config.ShrimpsTypesInShop[st]
	si.Cost = config.ShrimpCost[si.ShrimpType]
	return si
}

type WallpaperItem struct {
	BuyButton *input.Button
	Cost      int
	Type      config.WallpaperState
	IsBought  bool
	IsActive  bool
}

func NewWallpaperItem(btn *input.Button, wt int) *WallpaperItem {
	wi := new(WallpaperItem)
	wi.BuyButton = btn
	wi.Type = config.WallpaperTypesInShop[wt]
	wi.Cost = config.WallpaperCost[wi.Type]
	wi.IsBought = false
	wi.IsActive = false
	return wi
}
