package redpacketmodel

import "github.com/shengchaohua/red-packet-backend/common/enum"

const (
	TableRedPacket = "red_packet_tab"

	ColumnRedPacketId   = "red_packet_id"
	ColumnRedPacketName = "red_packet_name"
)

// RedPacketTab defines the red packet table in DB
type RedPacketTab struct {
	RedPacketId   uint64 `xorm:"'red_packet_id' bigint pk"`
	RedPacketName string `xorm:"'red_packet_name' varchar(255) not null"`
	RedPacketType uint32 `xorm:"'red_packet_type' int not null"`
	Quantity      uint32
	Count         uint32
	Ctime         uint32
	Mtime         uint32
	ExtraData     []byte
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
