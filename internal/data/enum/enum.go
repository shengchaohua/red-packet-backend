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
