package userdm

import (
	"context"
	"fmt"

	usermodel "github.com/shengchaohua/red-packet-backend/internal/data/model/user"
	"github.com/shengchaohua/red-packet-backend/internal/pkg/database"
	"github.com/shengchaohua/red-packet-backend/internal/utils"
	"xorm.io/xorm"
)

type defaultDM struct {
	database.EngineManager
	tableName string
}

func NewDefaultDM(engineManager database.EngineManager) DM {
	return &defaultDM{
		EngineManager: engineManager,
		tableName:     usermodel.UserTableName,
	}
}

func (dm *defaultDM) InsertWithSession(
	ctx context.Context,
	session *xorm.Session,
	user *usermodel.User,
) error {
	if session == nil {
		return ErrParam.WithMsg("session cannot be nil")
	}

	now := utils.GetCurrentTime()
	user.Ctime = now
	user.Mtime = now

	userWalletTxnTab, err := user.ModelToTab()
	if err != nil {
		return ErrData.Wrap(err)
	}

	affected, err := session.Table(dm.tableName).Insert(userWalletTxnTab)
	if err != nil {
		return ErrInsert.WrapWithMsg(err, fmt.Sprintf("insert_db_error|user_name=%s",
			user.Username,
		))
	}
	if affected == 0 {
		return ErrInsert.WrapWithMsg(err, fmt.Sprintf("insert_db_failed|user_name=%s",
			user.Username,
		))
	}

	return nil
}
