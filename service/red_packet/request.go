package redpacketservice

import "github.com/shengchaohua/red-packet-backend/data/enum"

// CreateRedPacketRequest defines the request
// RedPacketName - optional
// Quantity - the max quantity of people that can open the red packet
// Amount - the money mount in the red packet
type CreateRedPacketRequest struct {
	RequestId         string                 `json:"request_id,omitempty"`
	UserId            uint64                 `json:"user_id,omitempty"`
	RedPacketCategory enum.RedPacketCategory `json:"red_packet_category,omitempty"`
	RedPacketType     enum.RedPacketType     `json:"red_packet_type,omitempty"`
	RedPacketName     string                 `json:"red_packet_name,omitempty"`
	Quantity          uint32                 `json:"quantity,omitempty"`
	Amount            uint32                 `json:"amount,omitempty"`
}

// CreateRedPacketResponse defines the reponse
type CreateRedPacketResponse struct {
	RequestId   string `json:"request_id,omitempty"`
	RedPacketId uint64 `json:"red_packet_id,omitempty"`
}
