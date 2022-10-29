package redpacketmodel

import (
	"encoding/json"
	"fmt"

	"github.com/shengchaohua/red-packet-backend/internal/data/enum"
)

const (
	RedPacketTableName = "red_packet_tab"
)

// RedPacket defines the red packet class
type RedPacket struct {
	*RedPacketTab
	ExtraData *RedPacketExtraData
}

// RedPacketTab defines the red packet table in DB
type RedPacketTab struct {
	Id                uint64                 `xorm:"'id' bigint pk autoincr"`
	RedPacketName     string                 `xorm:"'red_packet_name' varchar(255) notnull"`
	RedPacketCategory enum.RedPacketCategory `xorm:"'red_packet_category' int notnull"`
	RedPacketType     enum.RedPacketType     `xorm:"'red_packet_type' int notnull"`
	Quantity          uint32                 `xorm:"'quantity' int notnull"`
	Amount            uint32                 `xorm:"'amount' int notnull"`
	RemainingQuantity uint32                 `xorm:"'remaining_quantity' int notnull"`
	Ctime             uint32                 `xorm:"'ctime' int notnull"`
	Mtime             uint32                 `xorm:"'mtime' int notnull"`
	ExtraData         []byte                 `xorm:"'extra_data' blob"`
}

// RedPacketExtraData defines the extra data in red packet
type RedPacketExtraData struct{}

// ModelToTab convert red packet model to tab
func (model *RedPacket) ModelToTab() (*RedPacketTab, error) {
	if model == nil {
		return nil, fmt.Errorf("red packet model is nil")
	}

	tab := model.RedPacketTab

	extraDataBytes, err := json.Marshal(model.ExtraData)
	if err != nil {
		return nil, fmt.Errorf("fail to marshal red packet extra data: %w", err)
	}
	tab.ExtraData = extraDataBytes

	return tab, nil
}

// // ModelToTab convert red packet tab to model
func (tab *RedPacketTab) TabToModel() (*RedPacket, error) {
	if tab == nil {
		return nil, fmt.Errorf("red packet tab is nil")
	}

	model := &RedPacket{
		RedPacketTab: tab,
	}

	extraData := &RedPacketExtraData{}
	if err := json.Unmarshal(tab.ExtraData, extraData); err != nil {
		return nil, fmt.Errorf("fail to unmarshal red packet extra data: %w", err)
	}
	model.ExtraData = extraData

	return model, nil
}
