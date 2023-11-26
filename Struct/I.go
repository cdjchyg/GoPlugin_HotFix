package Struct

type IPlayerMgr interface {
	AddPlayer(id int64)
	GetPlayer(id int64) *Player
}