package room

import (
	"SimpleGameApiPlatform/server/game"

	"github.com/gorilla/websocket"
)

var (
	FGFGameRoomQueue []FGFRoom // 炸金花游戏房间队列；
)

var (
	FGFWaitQueue []*websocket.Conn            // 炸金花等待队列；
	FGFConnChan  = make(chan *websocket.Conn) // 炸金花会话通道；
)

// 炸金花房间结构；
type FGFRoom struct {
	Status     int32           // 房间状态；
	Number     int32           // 房间号码；
	Banker     int32           // 当前庄家；
	Prices     int32           // 当前底注；
	Rounder    int32           // 回合玩家；
	RoundCount int32           // 回合记数；
	PrizePool  int32           // 房间奖池；
	RoomInfo   []game.Player   // 房间信息；
	ConnQueue  *websocket.Conn // 会话队列；
	//MsgPkgChan chan Message    // 消息通道；
}

//
