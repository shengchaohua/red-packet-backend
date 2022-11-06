package redpacketmodel

import (
	"encoding/json"
	"fmt"

	"github.com/shengchaohua/red-packet-backend/internal/data/enum"
)

const (
	RedPacketTableName = "red_packet_tab"
)

type RedPacket struct {
	*RedPacketTab
	ExtraData *RedPacketExtraData
}

type RedPacketTab struct {
	Id                  uint64                   `xorm:"'id' bigint unsigned pk autoincr"`
	RedPacketName       string                   `xorm:"'red_packet_name' varchar(255) notnull"`
	RedPacketCategory   enum.RedPacketCategory   `xorm:"'red_packet_category' int unsigned notnull"`
	RedPacketResultType enum.RedPacketResultType `xorm:"'red_packet_result_type' int unsigned notnull"`
	Quantity            uint32                   `xorm:"'quantity' int unsigned notnull"`
	Amount              uint32                   `xorm:"'amount' int unsigned notnull"`
	RemainingQuantity   uint32                   `xorm:"'remaining_quantity' int unsigned notnull"`
	Ctime               uint32                   `xorm:"'ctime' int unsigned notnull"`
	Mtime               uint32                   `xorm:"'mtime' int unsigned notnull"`
	ExtraData           []byte                   `xorm:"'extra_data' blob"`
}

type RedPacketExtraData struct {
	ReceiverUserId uint64 `json:"receiver_user_id,omitempty"` // valid for P2P red packet
	GroupId        uint64 `json:"group_id,omitempty"`         // valid for Group red packet
}

func (model *RedPacket) ModelToTab() (*RedPacketTab, error) {
	if model == nil {
		return nil, fmt.Errorf("red packet model is nil")
	}

	tab := model.RedPacketTab

	extraDataBytes, err := json.Marshal(model.ExtraData)
	if err != nil {
		return nil, fmt.Errorf("marshal red packet extra data error: %w", err)
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
		return nil, fmt.Errorf("unmarshal red packet extra data error: %w", err)
	}
	model.ExtraData = extraData

	return model, nil
}
