package redpacketmodel

import (
	"encoding/json"
	"fmt"
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
	RedPacketId   uint64 `xorm:"'red_packet_id' bigint pk"`
	RedPacketName string `xorm:"'red_packet_name' varchar(255) notnull"`
	RedPacketType uint32 `xorm:"'red_packet_type' int notnull"`
	Quantity      uint32 `xorm:"'quantity' int notnull"`
	Count         uint32 `xorm:"'count' int notnull"`
	Ctime         uint64 `xorm:"'ctime' int notnull"`
	Mtime         uint64 `xorm:"'mtime' int notnull"`
	ExtraData     []byte `xorm:"'extra_data' text"`
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
