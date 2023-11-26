package player

import (
	"GoHotFix/Struct"
)

type PlayerMgr struct {
	players map[int64]*Struct.Player
}

func NewPlayerMgr() *PlayerMgr {
	return &PlayerMgr{
		players: make(map[int64]*Struct.Player),
	}
}

func (pm *PlayerMgr) AddPlayer(id int64) {
	pm.players[id] = &Struct.Player{
		PID: id,
	}
}

func (pm *PlayerMgr) GetPlayer(id int64) *Struct.Player {
	p, ok := pm.players[id]
	if !ok {
		return nil
	}

	return p
}
