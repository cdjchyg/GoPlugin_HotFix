package Struct


type SetPlayerMgrReq struct {
	PM IPlayerMgr
}

type CreatePlayerReq struct {
	PID int64
}

type CreatePlayerAck struct {
}

type UpdateNameReq struct {
	PID int64
	Name string
}

type UpdateNameAck struct {

}

type GetPlayerNameReq struct {
	PID int64
}

type GetPlayerNameAck struct {
	PID int64
	Name string
} 
