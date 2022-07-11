package redpacketservice

import "github.com/shengchaohua/red-packet-backend/data/enum"

// CreateRedPacketRequest defines the request to create red packet
type CreateRedPacketRequest struct {
	RedPacketCategory enum.RedPacketCategory
	RedPacketType     enum.RedPacketType
	RedPacketName     string // optional
	Quantity          uint32 // the max quantity of people that can redeem the red packet
	Amount            uint32 // the money mount in the red packet
}

type CreateRedPacketResponse struct {
	RedPacketId uint64
}
