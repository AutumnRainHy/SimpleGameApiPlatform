package game

import (
	"github.com/gorilla/websocket"
)

// 单牌结构；
type Card struct {
	Number int32 // 点数；
	Flower int32 // 花色；
}

// 玩家结构；
type Player struct {
	ID       int32           // 玩家标识；
	Purse    int32           // 玩家钱包；
	Status   int32           // 玩家状态；
	HandCard []Card          // 玩家手牌；
	Conn     *websocket.Conn // 玩家会话；
}
