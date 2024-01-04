package Game

import (
	spb "github.com/gucooing/hkrpg-go/protocol/server"
)

func (g *Game) GetScene() *spb.Scene {
	if g.PlayerPb.Scene == nil {
		g.PlayerPb.Scene = &spb.Scene{EntryId: 1010101}
	}
	return g.PlayerPb.Scene
}

func (g *Game) GetPos() *spb.VectorBin {
	if g.PlayerPb.Pos == nil {
		g.PlayerPb.Pos = &spb.VectorBin{
			X: -43300,
			Y: 6,
			Z: -37960,
		}
	}
	return g.PlayerPb.Pos
}

func (g *Game) GetRot() *spb.VectorBin {
	if g.PlayerPb.Rot == nil {
		g.PlayerPb.Rot = &spb.VectorBin{
			X: 0,
			Y: 90000,
			Z: 0,
		}
	}
	return g.PlayerPb.Rot
}
