package enum

type RedPacketCategory uint32

const (
	RedPacketCategoryP2P   RedPacketCategory = 1 // send red packet from peer to peer
	RedPacketCategoryGroup RedPacketCategory = 2 // send red packet in group
)

func (redPacketCategory RedPacketCategory) String() string {
	switch redPacketCategory {
	case RedPacketCategoryP2P:
		return "peer_to_peer"
	case RedPacketCategoryGroup:
		return "group"
	}
	return ""
}

type RedPacketResultType uint32

const (
	RedPacketResultTypeRandomAmount    RedPacketResultType = 1
	RedPacketResultTypeIdenticalAmount RedPacketResultType = 2
)

func (redPacketResultType RedPacketResultType) String() string {
	switch redPacketResultType {
	case RedPacketResultTypeRandomAmount:
		return "random_amount"
	case RedPacketResultTypeIdenticalAmount:
		return "identical_amount"
	}
	return ""
}

type TransactionType uint32

const (
	TopupUserWallet TransactionType = 1
	CreateRedPacket TransactionType = 2
	OpenRedPacket   TransactionType = 3
	ReturnRedPacket TransactionType = 4 // red packet has remaining money
)

func (transactionType TransactionType) String() string {
	switch transactionType {
	case TopupUserWallet:
		return "topup_user_wallet"
	case CreateRedPacket:
		return "create_red_packet"
	case OpenRedPacket:
		return "open_red_packet"
	case ReturnRedPacket:
		return "refund_red_packet"
	}
	return ""
}
