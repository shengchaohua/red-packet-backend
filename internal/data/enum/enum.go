package enum

type RedPacketCategory uint32

const (
	RedPacketCategoryP2P            RedPacketCategory = 1 // send red packet from peer to peer
	RedPacketCategoryGroup          RedPacketCategory = 2 // send red packet in group
	RedPacketCategoryGroupExclusive RedPacketCategory = 3 // send red packet in group to a specific person
)

func (redPacketCategory RedPacketCategory) String() string {
	switch redPacketCategory {
	case RedPacketCategoryP2P:
		return "peer_to_peer"
	case RedPacketCategoryGroup:
		return "group"
	case RedPacketCategoryGroupExclusive:
		return "group_exclusive"
	}
	return ""
}

type RedPacketType uint32

const (
	RedPacketTypeRandomAmount    RedPacketType = 1 // red packet with random amount
	RedPacketTypeIdenticalAmount RedPacketType = 2 // red packet with identical amount
)

func (redPacketType RedPacketType) String() string {
	switch redPacketType {
	case RedPacketTypeRandomAmount:
		return "random_amount"
	case RedPacketTypeIdenticalAmount:
		return "identical_amount"
	}
	return ""
}

type TransactionType uint32

const (
	CreateRedPacket TransactionType = 1
	OpenRedPacket   TransactionType = 2
	RefundRedPacket TransactionType = 3 // red packet has remaining money
)

func (transactionType TransactionType) String() string {
	switch transactionType {
	case CreateRedPacket:
		return "create_red_packet"
	case OpenRedPacket:
		return "open_red_packet"
	case RefundRedPacket:
		return "refund_red_packet"
	}
	return ""
}
