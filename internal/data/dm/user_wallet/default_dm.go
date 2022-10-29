package userwalletdm

import (
	"context"
	"fmt"

	"xorm.io/xorm"

	userwalletmodel "github.com/shengchaohua/red-packet-backend/internal/data/model/user_wallet"
	"github.com/shengchaohua/red-packet-backend/internal/pkg/database"
	"github.com/shengchaohua/red-packet-backend/internal/pkg/logger"
)

type defaultDM struct {
	database.EngineManager
	tableName string
}

func NewDefaultDM(engineManager database.EngineManager) DM {
	return &defaultDM{
		EngineManager: engineManager,
		tableName:     userwalletmodel.UserWalletTableName,
	}
}

func (dm *defaultDM) InsertWithSession(
	ctx context.Context,
	session *xorm.Session,
	userWallet *userwalletmodel.UserWallet,
) error {
	if session != nil {
		return ErrParam.WithMsg("[InsertWithSession]session_cannot_be_nil")
	}

	return dm.insert(ctx, session, userWallet)
}

func (dm *defaultDM) insert(
	ctx context.Context,
	session *xorm.Session,
	userWallet *userwalletmodel.UserWallet,
) error {
	if session == nil {
		logger.Logger(ctx).Info("[defaultDM.insert]session_is_nil,use_new_session")
		session = dm.GetMasterEngine().Table(dm.tableName)
	}

	redPacketTab, err := userWallet.ModelToTab()
	if err != nil {
		return ErrData.Wrap(err)
	}

	affected, err := session.Table(dm.tableName).InsertOne(redPacketTab)
	if err != nil {
		return ErrInsert.WrapWithMsg(err, fmt.Sprintf("insert db error|user_id=%d", userWallet.Id))
	}
	if affected == 0 {
		return ErrInsert.WrapWithMsg(err, fmt.Sprintf("insert db failed|user_id=%d", userWallet.Id))
	}

	return nil
}

func (dm *defaultDM) UpdateWithSession(
	ctx context.Context,
	session *xorm.Session,
	userWallet *userwalletmodel.UserWallet,
) error {
	if session == nil {
		return ErrParam.WithMsg("[UpdateByIdWithSession]session cannot be nil")
	}

	return dm.update(ctx, session, userWallet)
}

func (dm *defaultDM) update(
	ctx context.Context,
	session *xorm.Session,
	userWallet *userwalletmodel.UserWallet,
) error {
	if session == nil {
		logger.Logger(ctx).Info("[defaultDM.update]session_is_nil,use_new_session")
		session = dm.GetMasterEngine().Table(dm.tableName)
	}

	userWalletTab, err := userWallet.ModelToTab()
	if err != nil {
		return ErrData.Wrap(err)
	}

	affected, err := session.Where("id = ?", userWallet.Id).
		AllCols().
		Update(userWalletTab)
	if err != nil {
		return ErrUpdate.WrapWithMsg(err, fmt.Sprintf("update db error|id=%d,user_id=%d",
			userWallet.Id, userWallet.UserId,
		))
	}
	if affected == 0 {
		logger.Logger(ctx).Info("[defaultDM.update]update_nothing")
		return nil
	}

	return nil
}

func (dm *defaultDM) LoadByUserIdWithSessionForUpdate(
	ctx context.Context,
	session *xorm.Session,
	userId uint64,
) (*userwalletmodel.UserWallet, error) {
	var userWalletTab = &userwalletmodel.UserWalletTab{}

	has, err := session.
		Where("user_id = ?", userId).
		ForUpdate().
		Get(userWalletTab)
	if err != nil {
		return nil, ErrQuery.WrapWithMsg(err, fmt.Sprintf("query db error|user_id=%d", userId))
	}
	if !has {
		return nil, nil
	}

	return userWalletTab.TabToModel()
}

func (dm *defaultDM) LoadByUserId(
	ctx context.Context,
	userId uint64,
	querySlave bool,
	queryMaster bool,
) (*userwalletmodel.UserWallet, error) {
	if !querySlave && !queryMaster {
		return nil, ErrParam.WithMsg("both querySlave and queryMaster are false")
	}

	var (
		userWallet = &userwalletmodel.UserWallet{}
		err        error
	)

	if querySlave {
		userWallet, err = dm.loadById(ctx, nil, userId, false)
		if err != nil {
			return nil, err
		}
		if userWallet != nil {
			return userWallet, nil
		}
	}

	if queryMaster {
		userWallet, err = dm.loadById(ctx, nil, userId, true)
		if err != nil {
			return nil, err
		}
		if userWallet == nil {
			return nil, nil
		}
	}

	return userWallet, nil
}

func (dm *defaultDM) loadById(
	ctx context.Context,
	session *xorm.Session,
	userId uint64,
	queryMaster bool,
) (*userwalletmodel.UserWallet, error) {
	var (
		userWalletTab = &userwalletmodel.UserWalletTab{}
		engine        = dm.GetSlaveEngine()
	)
	if queryMaster {
		engine = dm.GetMasterEngine()
	}
	if session == nil {
		logger.Logger(ctx).Info("session_is_nil,use_default_session")
		session = engine.Table(dm.tableName)
	}

	has, err := engine.Table(dm.tableName).
		Where("user_id = ?", userId).
		Get(userWalletTab)
	if err != nil {
		return nil, ErrQuery.WrapWithMsg(err, fmt.Sprintf("query db error|user_id=%d", userId))
	}
	if !has {
		return nil, nil
	}

	return userWalletTab.TabToModel()
}
