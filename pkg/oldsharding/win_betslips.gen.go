// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.

package oldsharding

import (
	"context"

	"gorm.io/gorm"
	"gorm.io/gorm/clause"
	"gorm.io/gorm/schema"

	"gorm.io/gen"
	"gorm.io/gen/field"

	"gorm.io/plugin/dbresolver"

	"migration/pkg/model"
)

func newWinBetslips(db *gorm.DB, opts ...gen.DOOption) winBetslips {
	_winBetslips := winBetslips{}

	_winBetslips.winBetslipsDo.UseDB(db, opts...)
	_winBetslips.winBetslipsDo.UseModel(&model.WinBetslips{})

	tableName := _winBetslips.winBetslipsDo.TableName()
	_winBetslips.ALL = field.NewAsterisk(tableName)
	_winBetslips.ID = field.NewInt64(tableName, "id")
	_winBetslips.RoundID = field.NewString(tableName, "round_id")
	_winBetslips.TransactionID = field.NewString(tableName, "transaction_id")
	_winBetslips.XbStatus = field.NewInt32(tableName, "xb_status")
	_winBetslips.XbUID = field.NewInt32(tableName, "xb_uid")
	_winBetslips.XbUsername = field.NewString(tableName, "xb_username")
	_winBetslips.MerchantID = field.NewInt32(tableName, "merchant_id")
	_winBetslips.XbProfit = field.NewFloat64(tableName, "xb_profit")
	_winBetslips.Stake = field.NewFloat64(tableName, "stake")
	_winBetslips.ValidStake = field.NewFloat64(tableName, "valid_stake")
	_winBetslips.Payout = field.NewFloat64(tableName, "payout")
	_winBetslips.CoinRefund = field.NewFloat64(tableName, "coin_refund")
	_winBetslips.CoinBefore = field.NewFloat64(tableName, "coin_before")
	_winBetslips.GameProviderSubtypeID = field.NewInt32(tableName, "game_provider_subtype_id")
	_winBetslips.GameListID = field.NewInt32(tableName, "game_list_id")
	_winBetslips.GamePagcorID = field.NewInt32(tableName, "game_pagcor_id")
	_winBetslips.GameProviderID = field.NewInt32(tableName, "game_provider_id")
	_winBetslips.AmountType = field.NewInt32(tableName, "amount_type")
	_winBetslips.DtStarted = field.NewInt64(tableName, "dt_started")
	_winBetslips.DtCompleted = field.NewInt64(tableName, "dt_completed")
	_winBetslips.WinTransactionID = field.NewString(tableName, "win_transaction_id")
	_winBetslips.CreateTimeStr = field.NewString(tableName, "create_time_str")
	_winBetslips.CreatedAt = field.NewInt32(tableName, "created_at")
	_winBetslips.UpdatedAt = field.NewInt32(tableName, "updated_at")
	_winBetslips.GameTypeID = field.NewInt32(tableName, "game_type_id")

	_winBetslips.fillFieldMap()

	return _winBetslips
}

type winBetslips struct {
	winBetslipsDo

	ALL                   field.Asterisk
	ID                    field.Int64   // 主键
	RoundID               field.String  // 回合id
	TransactionID         field.String  // 注单号 对应三方拉单transaction_id
	XbStatus              field.Int32   // 注单状态 1:待开彩  2:完成  3: 退款
	XbUID                 field.Int32   // 对应user表id
	XbUsername            field.String  // 对应user表username
	MerchantID            field.Int32   // 商户id
	XbProfit              field.Float64 // 盈亏金额
	Stake                 field.Float64 // 投注
	ValidStake            field.Float64 // 有效投注金额
	Payout                field.Float64 // 派彩
	CoinRefund            field.Float64 // 退款金额
	CoinBefore            field.Float64 // 投注前金额
	GameProviderSubtypeID field.Int32   // 游戏id对应game_provider_subtype表id
	GameListID            field.Int32   // 游戏id对应game_list表id
	GamePagcorID          field.Int32   // pagcor分组id
	GameProviderID        field.Int32   // 游戏供应商id
	AmountType            field.Int32   // 投注方式 1:现金，2:奖金 3:免费旋转 4:活动免费旋转
	DtStarted             field.Int64   // 游戏开始时间
	DtCompleted           field.Int64   // 游戏结束时间
	WinTransactionID      field.String  // 开奖交易单号
	CreateTimeStr         field.String  // 投注时间
	CreatedAt             field.Int32
	UpdatedAt             field.Int32
	GameTypeID            field.Int32 // 游戏分组id

	fieldMap map[string]field.Expr
}

func (w winBetslips) Table(newTableName string) *winBetslips {
	w.winBetslipsDo.UseTable(newTableName)
	return w.updateTableName(newTableName)
}

func (w winBetslips) As(alias string) *winBetslips {
	w.winBetslipsDo.DO = *(w.winBetslipsDo.As(alias).(*gen.DO))
	return w.updateTableName(alias)
}

func (w *winBetslips) updateTableName(table string) *winBetslips {
	w.ALL = field.NewAsterisk(table)
	w.ID = field.NewInt64(table, "id")
	w.RoundID = field.NewString(table, "round_id")
	w.TransactionID = field.NewString(table, "transaction_id")
	w.XbStatus = field.NewInt32(table, "xb_status")
	w.XbUID = field.NewInt32(table, "xb_uid")
	w.XbUsername = field.NewString(table, "xb_username")
	w.MerchantID = field.NewInt32(table, "merchant_id")
	w.XbProfit = field.NewFloat64(table, "xb_profit")
	w.Stake = field.NewFloat64(table, "stake")
	w.ValidStake = field.NewFloat64(table, "valid_stake")
	w.Payout = field.NewFloat64(table, "payout")
	w.CoinRefund = field.NewFloat64(table, "coin_refund")
	w.CoinBefore = field.NewFloat64(table, "coin_before")
	w.GameProviderSubtypeID = field.NewInt32(table, "game_provider_subtype_id")
	w.GameListID = field.NewInt32(table, "game_list_id")
	w.GamePagcorID = field.NewInt32(table, "game_pagcor_id")
	w.GameProviderID = field.NewInt32(table, "game_provider_id")
	w.AmountType = field.NewInt32(table, "amount_type")
	w.DtStarted = field.NewInt64(table, "dt_started")
	w.DtCompleted = field.NewInt64(table, "dt_completed")
	w.WinTransactionID = field.NewString(table, "win_transaction_id")
	w.CreateTimeStr = field.NewString(table, "create_time_str")
	w.CreatedAt = field.NewInt32(table, "created_at")
	w.UpdatedAt = field.NewInt32(table, "updated_at")
	w.GameTypeID = field.NewInt32(table, "game_type_id")

	w.fillFieldMap()

	return w
}

func (w *winBetslips) GetFieldByName(fieldName string) (field.OrderExpr, bool) {
	_f, ok := w.fieldMap[fieldName]
	if !ok || _f == nil {
		return nil, false
	}
	_oe, ok := _f.(field.OrderExpr)
	return _oe, ok
}

func (w *winBetslips) fillFieldMap() {
	w.fieldMap = make(map[string]field.Expr, 25)
	w.fieldMap["id"] = w.ID
	w.fieldMap["round_id"] = w.RoundID
	w.fieldMap["transaction_id"] = w.TransactionID
	w.fieldMap["xb_status"] = w.XbStatus
	w.fieldMap["xb_uid"] = w.XbUID
	w.fieldMap["xb_username"] = w.XbUsername
	w.fieldMap["merchant_id"] = w.MerchantID
	w.fieldMap["xb_profit"] = w.XbProfit
	w.fieldMap["stake"] = w.Stake
	w.fieldMap["valid_stake"] = w.ValidStake
	w.fieldMap["payout"] = w.Payout
	w.fieldMap["coin_refund"] = w.CoinRefund
	w.fieldMap["coin_before"] = w.CoinBefore
	w.fieldMap["game_provider_subtype_id"] = w.GameProviderSubtypeID
	w.fieldMap["game_list_id"] = w.GameListID
	w.fieldMap["game_pagcor_id"] = w.GamePagcorID
	w.fieldMap["game_provider_id"] = w.GameProviderID
	w.fieldMap["amount_type"] = w.AmountType
	w.fieldMap["dt_started"] = w.DtStarted
	w.fieldMap["dt_completed"] = w.DtCompleted
	w.fieldMap["win_transaction_id"] = w.WinTransactionID
	w.fieldMap["create_time_str"] = w.CreateTimeStr
	w.fieldMap["created_at"] = w.CreatedAt
	w.fieldMap["updated_at"] = w.UpdatedAt
	w.fieldMap["game_type_id"] = w.GameTypeID
}

func (w winBetslips) clone(db *gorm.DB) winBetslips {
	w.winBetslipsDo.ReplaceConnPool(db.Statement.ConnPool)
	return w
}

func (w winBetslips) replaceDB(db *gorm.DB) winBetslips {
	w.winBetslipsDo.ReplaceDB(db)
	return w
}

type winBetslipsDo struct{ gen.DO }

func (w winBetslipsDo) Debug() *winBetslipsDo {
	return w.withDO(w.DO.Debug())
}

func (w winBetslipsDo) WithContext(ctx context.Context) *winBetslipsDo {
	return w.withDO(w.DO.WithContext(ctx))
}

func (w winBetslipsDo) ReadDB() *winBetslipsDo {
	return w.Clauses(dbresolver.Read)
}

func (w winBetslipsDo) WriteDB() *winBetslipsDo {
	return w.Clauses(dbresolver.Write)
}

func (w winBetslipsDo) Session(config *gorm.Session) *winBetslipsDo {
	return w.withDO(w.DO.Session(config))
}

func (w winBetslipsDo) Clauses(conds ...clause.Expression) *winBetslipsDo {
	return w.withDO(w.DO.Clauses(conds...))
}

func (w winBetslipsDo) Returning(value interface{}, columns ...string) *winBetslipsDo {
	return w.withDO(w.DO.Returning(value, columns...))
}

func (w winBetslipsDo) Not(conds ...gen.Condition) *winBetslipsDo {
	return w.withDO(w.DO.Not(conds...))
}

func (w winBetslipsDo) Or(conds ...gen.Condition) *winBetslipsDo {
	return w.withDO(w.DO.Or(conds...))
}

func (w winBetslipsDo) Select(conds ...field.Expr) *winBetslipsDo {
	return w.withDO(w.DO.Select(conds...))
}

func (w winBetslipsDo) Where(conds ...gen.Condition) *winBetslipsDo {
	return w.withDO(w.DO.Where(conds...))
}

func (w winBetslipsDo) Order(conds ...field.Expr) *winBetslipsDo {
	return w.withDO(w.DO.Order(conds...))
}

func (w winBetslipsDo) Distinct(cols ...field.Expr) *winBetslipsDo {
	return w.withDO(w.DO.Distinct(cols...))
}

func (w winBetslipsDo) Omit(cols ...field.Expr) *winBetslipsDo {
	return w.withDO(w.DO.Omit(cols...))
}

func (w winBetslipsDo) Join(table schema.Tabler, on ...field.Expr) *winBetslipsDo {
	return w.withDO(w.DO.Join(table, on...))
}

func (w winBetslipsDo) LeftJoin(table schema.Tabler, on ...field.Expr) *winBetslipsDo {
	return w.withDO(w.DO.LeftJoin(table, on...))
}

func (w winBetslipsDo) RightJoin(table schema.Tabler, on ...field.Expr) *winBetslipsDo {
	return w.withDO(w.DO.RightJoin(table, on...))
}

func (w winBetslipsDo) Group(cols ...field.Expr) *winBetslipsDo {
	return w.withDO(w.DO.Group(cols...))
}

func (w winBetslipsDo) Having(conds ...gen.Condition) *winBetslipsDo {
	return w.withDO(w.DO.Having(conds...))
}

func (w winBetslipsDo) Limit(limit int) *winBetslipsDo {
	return w.withDO(w.DO.Limit(limit))
}

func (w winBetslipsDo) Offset(offset int) *winBetslipsDo {
	return w.withDO(w.DO.Offset(offset))
}

func (w winBetslipsDo) Scopes(funcs ...func(gen.Dao) gen.Dao) *winBetslipsDo {
	return w.withDO(w.DO.Scopes(funcs...))
}

func (w winBetslipsDo) Unscoped() *winBetslipsDo {
	return w.withDO(w.DO.Unscoped())
}

func (w winBetslipsDo) Create(values ...*model.WinBetslips) error {
	if len(values) == 0 {
		return nil
	}
	return w.DO.Create(values)
}

func (w winBetslipsDo) CreateInBatches(values []*model.WinBetslips, batchSize int) error {
	return w.DO.CreateInBatches(values, batchSize)
}

// Save : !!! underlying implementation is different with GORM
// The method is equivalent to executing the statement: db.Clauses(clause.OnConflict{UpdateAll: true}).Create(values)
func (w winBetslipsDo) Save(values ...*model.WinBetslips) error {
	if len(values) == 0 {
		return nil
	}
	return w.DO.Save(values)
}

func (w winBetslipsDo) First() (*model.WinBetslips, error) {
	if result, err := w.DO.First(); err != nil {
		return nil, err
	} else {
		return result.(*model.WinBetslips), nil
	}
}

func (w winBetslipsDo) Take() (*model.WinBetslips, error) {
	if result, err := w.DO.Take(); err != nil {
		return nil, err
	} else {
		return result.(*model.WinBetslips), nil
	}
}

func (w winBetslipsDo) Last() (*model.WinBetslips, error) {
	if result, err := w.DO.Last(); err != nil {
		return nil, err
	} else {
		return result.(*model.WinBetslips), nil
	}
}

func (w winBetslipsDo) Find() ([]*model.WinBetslips, error) {
	result, err := w.DO.Find()
	return result.([]*model.WinBetslips), err
}

func (w winBetslipsDo) FindInBatch(batchSize int, fc func(tx gen.Dao, batch int) error) (results []*model.WinBetslips, err error) {
	buf := make([]*model.WinBetslips, 0, batchSize)
	err = w.DO.FindInBatches(&buf, batchSize, func(tx gen.Dao, batch int) error {
		defer func() { results = append(results, buf...) }()
		return fc(tx, batch)
	})
	return results, err
}

func (w winBetslipsDo) FindInBatches(result *[]*model.WinBetslips, batchSize int, fc func(tx gen.Dao, batch int) error) error {
	return w.DO.FindInBatches(result, batchSize, fc)
}

func (w winBetslipsDo) Attrs(attrs ...field.AssignExpr) *winBetslipsDo {
	return w.withDO(w.DO.Attrs(attrs...))
}

func (w winBetslipsDo) Assign(attrs ...field.AssignExpr) *winBetslipsDo {
	return w.withDO(w.DO.Assign(attrs...))
}

func (w winBetslipsDo) Joins(fields ...field.RelationField) *winBetslipsDo {
	for _, _f := range fields {
		w = *w.withDO(w.DO.Joins(_f))
	}
	return &w
}

func (w winBetslipsDo) Preload(fields ...field.RelationField) *winBetslipsDo {
	for _, _f := range fields {
		w = *w.withDO(w.DO.Preload(_f))
	}
	return &w
}

func (w winBetslipsDo) FirstOrInit() (*model.WinBetslips, error) {
	if result, err := w.DO.FirstOrInit(); err != nil {
		return nil, err
	} else {
		return result.(*model.WinBetslips), nil
	}
}

func (w winBetslipsDo) FirstOrCreate() (*model.WinBetslips, error) {
	if result, err := w.DO.FirstOrCreate(); err != nil {
		return nil, err
	} else {
		return result.(*model.WinBetslips), nil
	}
}

func (w winBetslipsDo) FindByPage(offset int, limit int) (result []*model.WinBetslips, count int64, err error) {
	result, err = w.Offset(offset).Limit(limit).Find()
	if err != nil {
		return
	}

	if size := len(result); 0 < limit && 0 < size && size < limit {
		count = int64(size + offset)
		return
	}

	count, err = w.Offset(-1).Limit(-1).Count()
	return
}

func (w winBetslipsDo) ScanByPage(result interface{}, offset int, limit int) (count int64, err error) {
	count, err = w.Count()
	if err != nil {
		return
	}

	err = w.Offset(offset).Limit(limit).Scan(result)
	return
}

func (w winBetslipsDo) Scan(result interface{}) (err error) {
	return w.DO.Scan(result)
}

func (w winBetslipsDo) Delete(models ...*model.WinBetslips) (result gen.ResultInfo, err error) {
	return w.DO.Delete(models)
}

func (w *winBetslipsDo) withDO(do gen.Dao) *winBetslipsDo {
	w.DO = *do.(*gen.DO)
	return w
}
