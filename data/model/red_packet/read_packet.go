package redpacketmodel

import "github.com/shengchaohua/red-packet-backend/data/enum"

// RedPacketTab defines the red packet table in DB
type RedPacketTab struct {
	RedPacketId   uint64
	RedPacketName string
	RedPacketType uint32
	Quantity      uint32
	Count         uint32
}

// RedPacketExtraData defines the extra data in red packet
type RedPacketExtraData struct {
}

// RedPacket defines the red packet class
type RedPacket struct {
	*RedPacketTab

	RedPacketCategory enum.RedPacketCategory
	RedPacketType     enum.RedPacketType

	ExtraData *RedPacketExtraData
}
