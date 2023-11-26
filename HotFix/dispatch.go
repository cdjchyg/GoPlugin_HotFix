package main

import (
	"GoHotFix/Struct"
	"fmt"
)

var pm Struct.IPlayerMgr

func DispatchMsg(param interface {}) interface{} {
	switch t := param.(type) {
	case *Struct.SetPlayerMgrReq:
		req := param.(*Struct.SetPlayerMgrReq)
		pm = req.PM
		return nil
	case *Struct.CreatePlayerReq:
		req := param.(*Struct.CreatePlayerReq)
		pm.AddPlayer(req.PID)
		// fmt.Println(req.PID)
		return &Struct.CreatePlayerAck{}
	case *Struct.UpdateNameReq:
		req := param.(*Struct.UpdateNameReq)
		p := pm.GetPlayer(req.PID)
		p.Name = req.Name
		// fmt.Println(p.Name)
		return &Struct.UpdateNameAck{}
	case *Struct.GetPlayerNameReq:
		req := param.(*Struct.GetPlayerNameReq)
		p := pm.GetPlayer(req.PID)
		return &Struct.GetPlayerNameAck{
			Name: p.Name,
			PID: p.PID,
		}
	default:
		fmt.Println("unknown Msg", t)
	}
	return nil
}