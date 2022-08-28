package redpacketmodel

import (
	"encoding/json"
	"fmt"

	"github.com/shengchaohua/red-packet-backend/data/enum"
)

const (
	TableRedPacket = "red_packet_tab"

	ColumnRedPacketId   = "red_packet_id"
	ColumnRedPacketName = "red_packet_name"
)

// RedPacket defines the red packet class
type RedPacket struct {
	*RedPacketTab
	ExtraData *RedPacketExtraData
}

// RedPacketTab defines the red packet table in DB
type RedPacketTab struct {
	RedPacketId       uint64                 `xorm:"'red_packet_id' bigint pk autoincr"`
	RedPacketName     string                 `xorm:"'red_packet_name' varchar(255) notnull"`
	RedPacketCategory enum.RedPacketCategory `xorm:"'red_packet_category' int notnull"`
	RedPacketType     enum.RedPacketType     `xorm:"'red_packet_type' int notnull"`
	Quantity          uint32                 `xorm:"'quantity' int notnull"`
	RemainingQuantity uint32                 `xorm:"'remaining_quantity' int notnull"`
	Amount            uint32                 `xorm:"'count' int notnull"`
	Ctime             uint32                 `xorm:"'ctime' int notnull"`
	Mtime             uint32                 `xorm:"'mtime' int notnull"`
	ExtraData         []byte                 `xorm:"'extra_data' text"`
}

// RedPacketExtraData defines the extra data in red packet
type RedPacketExtraData struct {
}

func (model *RedPacket) ModelToTab() (*RedPacketTab, error) {
	if model == nil {
		return nil, fmt.Errorf("red packet model is nil")
	}

	tab := model.RedPacketTab

	extraDataBytes, err := json.Marshal(model.ExtraData)
	if err != nil {
		return nil, fmt.Errorf("fail to marshal red packet extra data: %s", err.Error())
	}
	tab.ExtraData = extraDataBytes

	return tab, nil
}

func (tab *RedPacketTab) TabToModel() (*RedPacket, error) {
	if tab == nil {
		return nil, fmt.Errorf("red packet tab is nil")
	}

	model := &RedPacket{
		RedPacketTab: tab,
	}

	extraData := &RedPacketExtraData{}
	if err := json.Unmarshal(tab.ExtraData, extraData); err != nil {
		return nil, fmt.Errorf("fail to unmarshal red packet extra data: %s", err.Error())
	}
	model.ExtraData = extraData

	return model, nil
}
