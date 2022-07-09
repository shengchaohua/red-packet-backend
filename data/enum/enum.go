package enum

type RedPacketCategory uint32

const (
	RedPacketCategoryP2P   RedPacketCategory = 1 // send red packet from peer to peer
	RedPacketCategoryGroup RedPacketCategory = 2 // send red packet in group
)

type RedPacketType uint32

const (
	RedPacketTypeRandomAmount    RedPacketType = 1 // red packet with random amount
	RedPacketTypeIdenticalAmount RedPacketType = 2 // red packet with identical amount
)
