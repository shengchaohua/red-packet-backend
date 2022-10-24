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
		return "RedPacketCategory(peer_to_peer)"
	case RedPacketCategoryGroup:
		return "RedPacketCategory(group)"
	case RedPacketCategoryGroupExclusive:
		return "RedPacketCategory(group_exclusive)"
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
		return "RedPacketType(random_amount)"
	case RedPacketTypeIdenticalAmount:
		return "RedPacketType(identical_amount)"
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
		return "TransactionType(create_red_packet)"
	case OpenRedPacket:
		return "TransactionType(open_red_packet)"
	case RefundRedPacket:
		return "TransactionType(refund_red_packet)"
	}
	return ""
}

type TransactionStatus uint32

const (
	TransactionStatusSuccess = 1
)

func (transactionStatus TransactionStatus) String() string {
	switch transactionStatus {
	case TransactionStatusSuccess:
		return "TransactionStatus(success)"
	}
	return ""
}
