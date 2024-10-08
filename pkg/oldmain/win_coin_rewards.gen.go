// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.

package oldmain

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

func newWinCoinReward(db *gorm.DB, opts ...gen.DOOption) winCoinReward {
	_winCoinReward := winCoinReward{}

	_winCoinReward.winCoinRewardDo.UseDB(db, opts...)
	_winCoinReward.winCoinRewardDo.UseModel(&model.WinCoinReward{})

	tableName := _winCoinReward.winCoinRewardDo.TableName()
	_winCoinReward.ALL = field.NewAsterisk(tableName)
	_winCoinReward.ID = field.NewInt64(tableName, "id")
	_winCoinReward.UID = field.NewInt32(tableName, "uid")
	_winCoinReward.Username = field.NewString(tableName, "username")
	_winCoinReward.Coin = field.NewFloat64(tableName, "coin")
	_winCoinReward.CoinBefore = field.NewFloat64(tableName, "coin_before")
	_winCoinReward.ReferID = field.NewInt32(tableName, "refer_id")
	_winCoinReward.LadderName = field.NewString(tableName, "ladder_name")
	_winCoinReward.ReferCode = field.NewString(tableName, "refer_code")
	_winCoinReward.FlowClaim = field.NewInt32(tableName, "flow_claim")
	_winCoinReward.Codes = field.NewFloat64(tableName, "codes")
	_winCoinReward.EndedAt = field.NewInt32(tableName, "ended_at")
	_winCoinReward.Info = field.NewString(tableName, "info")
	_winCoinReward.Status = field.NewInt32(tableName, "status")
	_winCoinReward.CreatedAt = field.NewInt32(tableName, "created_at")
	_winCoinReward.UpdatedAt = field.NewInt32(tableName, "updated_at")

	_winCoinReward.fillFieldMap()

	return _winCoinReward
}

// winCoinReward 奖金表
type winCoinReward struct {
	winCoinRewardDo

	ALL        field.Asterisk
	ID         field.Int64
	UID        field.Int32   // UID
	Username   field.String  // 用户名
	Coin       field.Float64 // 金额
	CoinBefore field.Float64 // 即时金额
	ReferID    field.Int32   // 关联ID(活动表ID)
	LadderName field.String  // 新活动，等级code
	ReferCode  field.String  // 关联Code(活动表Code)
	FlowClaim  field.Int32   // 流水倍数
	Codes      field.Float64 // 所需打码量
	EndedAt    field.Int32   // 活动结束时间
	Info       field.String  // 备注
	Status     field.Int32   // 状态:0-申请中 1-已满足 2-已派发3-已结束
	CreatedAt  field.Int32
	UpdatedAt  field.Int32

	fieldMap map[string]field.Expr
}

func (w winCoinReward) Table(newTableName string) *winCoinReward {
	w.winCoinRewardDo.UseTable(newTableName)
	return w.updateTableName(newTableName)
}

func (w winCoinReward) As(alias string) *winCoinReward {
	w.winCoinRewardDo.DO = *(w.winCoinRewardDo.As(alias).(*gen.DO))
	return w.updateTableName(alias)
}

func (w *winCoinReward) updateTableName(table string) *winCoinReward {
	w.ALL = field.NewAsterisk(table)
	w.ID = field.NewInt64(table, "id")
	w.UID = field.NewInt32(table, "uid")
	w.Username = field.NewString(table, "username")
	w.Coin = field.NewFloat64(table, "coin")
	w.CoinBefore = field.NewFloat64(table, "coin_before")
	w.ReferID = field.NewInt32(table, "refer_id")
	w.LadderName = field.NewString(table, "ladder_name")
	w.ReferCode = field.NewString(table, "refer_code")
	w.FlowClaim = field.NewInt32(table, "flow_claim")
	w.Codes = field.NewFloat64(table, "codes")
	w.EndedAt = field.NewInt32(table, "ended_at")
	w.Info = field.NewString(table, "info")
	w.Status = field.NewInt32(table, "status")
	w.CreatedAt = field.NewInt32(table, "created_at")
	w.UpdatedAt = field.NewInt32(table, "updated_at")

	w.fillFieldMap()

	return w
}

func (w *winCoinReward) GetFieldByName(fieldName string) (field.OrderExpr, bool) {
	_f, ok := w.fieldMap[fieldName]
	if !ok || _f == nil {
		return nil, false
	}
	_oe, ok := _f.(field.OrderExpr)
	return _oe, ok
}

func (w *winCoinReward) fillFieldMap() {
	w.fieldMap = make(map[string]field.Expr, 15)
	w.fieldMap["id"] = w.ID
	w.fieldMap["uid"] = w.UID
	w.fieldMap["username"] = w.Username
	w.fieldMap["coin"] = w.Coin
	w.fieldMap["coin_before"] = w.CoinBefore
	w.fieldMap["refer_id"] = w.ReferID
	w.fieldMap["ladder_name"] = w.LadderName
	w.fieldMap["refer_code"] = w.ReferCode
	w.fieldMap["flow_claim"] = w.FlowClaim
	w.fieldMap["codes"] = w.Codes
	w.fieldMap["ended_at"] = w.EndedAt
	w.fieldMap["info"] = w.Info
	w.fieldMap["status"] = w.Status
	w.fieldMap["created_at"] = w.CreatedAt
	w.fieldMap["updated_at"] = w.UpdatedAt
}

func (w winCoinReward) clone(db *gorm.DB) winCoinReward {
	w.winCoinRewardDo.ReplaceConnPool(db.Statement.ConnPool)
	return w
}

func (w winCoinReward) replaceDB(db *gorm.DB) winCoinReward {
	w.winCoinRewardDo.ReplaceDB(db)
	return w
}

type winCoinRewardDo struct{ gen.DO }

func (w winCoinRewardDo) Debug() *winCoinRewardDo {
	return w.withDO(w.DO.Debug())
}

func (w winCoinRewardDo) WithContext(ctx context.Context) *winCoinRewardDo {
	return w.withDO(w.DO.WithContext(ctx))
}

func (w winCoinRewardDo) ReadDB() *winCoinRewardDo {
	return w.Clauses(dbresolver.Read)
}

func (w winCoinRewardDo) WriteDB() *winCoinRewardDo {
	return w.Clauses(dbresolver.Write)
}

func (w winCoinRewardDo) Session(config *gorm.Session) *winCoinRewardDo {
	return w.withDO(w.DO.Session(config))
}

func (w winCoinRewardDo) Clauses(conds ...clause.Expression) *winCoinRewardDo {
	return w.withDO(w.DO.Clauses(conds...))
}

func (w winCoinRewardDo) Returning(value interface{}, columns ...string) *winCoinRewardDo {
	return w.withDO(w.DO.Returning(value, columns...))
}

func (w winCoinRewardDo) Not(conds ...gen.Condition) *winCoinRewardDo {
	return w.withDO(w.DO.Not(conds...))
}

func (w winCoinRewardDo) Or(conds ...gen.Condition) *winCoinRewardDo {
	return w.withDO(w.DO.Or(conds...))
}

func (w winCoinRewardDo) Select(conds ...field.Expr) *winCoinRewardDo {
	return w.withDO(w.DO.Select(conds...))
}

func (w winCoinRewardDo) Where(conds ...gen.Condition) *winCoinRewardDo {
	return w.withDO(w.DO.Where(conds...))
}

func (w winCoinRewardDo) Order(conds ...field.Expr) *winCoinRewardDo {
	return w.withDO(w.DO.Order(conds...))
}

func (w winCoinRewardDo) Distinct(cols ...field.Expr) *winCoinRewardDo {
	return w.withDO(w.DO.Distinct(cols...))
}

func (w winCoinRewardDo) Omit(cols ...field.Expr) *winCoinRewardDo {
	return w.withDO(w.DO.Omit(cols...))
}

func (w winCoinRewardDo) Join(table schema.Tabler, on ...field.Expr) *winCoinRewardDo {
	return w.withDO(w.DO.Join(table, on...))
}

func (w winCoinRewardDo) LeftJoin(table schema.Tabler, on ...field.Expr) *winCoinRewardDo {
	return w.withDO(w.DO.LeftJoin(table, on...))
}

func (w winCoinRewardDo) RightJoin(table schema.Tabler, on ...field.Expr) *winCoinRewardDo {
	return w.withDO(w.DO.RightJoin(table, on...))
}

func (w winCoinRewardDo) Group(cols ...field.Expr) *winCoinRewardDo {
	return w.withDO(w.DO.Group(cols...))
}

func (w winCoinRewardDo) Having(conds ...gen.Condition) *winCoinRewardDo {
	return w.withDO(w.DO.Having(conds...))
}

func (w winCoinRewardDo) Limit(limit int) *winCoinRewardDo {
	return w.withDO(w.DO.Limit(limit))
}

func (w winCoinRewardDo) Offset(offset int) *winCoinRewardDo {
	return w.withDO(w.DO.Offset(offset))
}

func (w winCoinRewardDo) Scopes(funcs ...func(gen.Dao) gen.Dao) *winCoinRewardDo {
	return w.withDO(w.DO.Scopes(funcs...))
}

func (w winCoinRewardDo) Unscoped() *winCoinRewardDo {
	return w.withDO(w.DO.Unscoped())
}

func (w winCoinRewardDo) Create(values ...*model.WinCoinReward) error {
	if len(values) == 0 {
		return nil
	}
	return w.DO.Create(values)
}

func (w winCoinRewardDo) CreateInBatches(values []*model.WinCoinReward, batchSize int) error {
	return w.DO.CreateInBatches(values, batchSize)
}

// Save : !!! underlying implementation is different with GORM
// The method is equivalent to executing the statement: db.Clauses(clause.OnConflict{UpdateAll: true}).Create(values)
func (w winCoinRewardDo) Save(values ...*model.WinCoinReward) error {
	if len(values) == 0 {
		return nil
	}
	return w.DO.Save(values)
}

func (w winCoinRewardDo) First() (*model.WinCoinReward, error) {
	if result, err := w.DO.First(); err != nil {
		return nil, err
	} else {
		return result.(*model.WinCoinReward), nil
	}
}

func (w winCoinRewardDo) Take() (*model.WinCoinReward, error) {
	if result, err := w.DO.Take(); err != nil {
		return nil, err
	} else {
		return result.(*model.WinCoinReward), nil
	}
}

func (w winCoinRewardDo) Last() (*model.WinCoinReward, error) {
	if result, err := w.DO.Last(); err != nil {
		return nil, err
	} else {
		return result.(*model.WinCoinReward), nil
	}
}

func (w winCoinRewardDo) Find() ([]*model.WinCoinReward, error) {
	result, err := w.DO.Find()
	return result.([]*model.WinCoinReward), err
}

func (w winCoinRewardDo) FindInBatch(batchSize int, fc func(tx gen.Dao, batch int) error) (results []*model.WinCoinReward, err error) {
	buf := make([]*model.WinCoinReward, 0, batchSize)
	err = w.DO.FindInBatches(&buf, batchSize, func(tx gen.Dao, batch int) error {
		defer func() { results = append(results, buf...) }()
		return fc(tx, batch)
	})
	return results, err
}

func (w winCoinRewardDo) FindInBatches(result *[]*model.WinCoinReward, batchSize int, fc func(tx gen.Dao, batch int) error) error {
	return w.DO.FindInBatches(result, batchSize, fc)
}

func (w winCoinRewardDo) Attrs(attrs ...field.AssignExpr) *winCoinRewardDo {
	return w.withDO(w.DO.Attrs(attrs...))
}

func (w winCoinRewardDo) Assign(attrs ...field.AssignExpr) *winCoinRewardDo {
	return w.withDO(w.DO.Assign(attrs...))
}

func (w winCoinRewardDo) Joins(fields ...field.RelationField) *winCoinRewardDo {
	for _, _f := range fields {
		w = *w.withDO(w.DO.Joins(_f))
	}
	return &w
}

func (w winCoinRewardDo) Preload(fields ...field.RelationField) *winCoinRewardDo {
	for _, _f := range fields {
		w = *w.withDO(w.DO.Preload(_f))
	}
	return &w
}

func (w winCoinRewardDo) FirstOrInit() (*model.WinCoinReward, error) {
	if result, err := w.DO.FirstOrInit(); err != nil {
		return nil, err
	} else {
		return result.(*model.WinCoinReward), nil
	}
}

func (w winCoinRewardDo) FirstOrCreate() (*model.WinCoinReward, error) {
	if result, err := w.DO.FirstOrCreate(); err != nil {
		return nil, err
	} else {
		return result.(*model.WinCoinReward), nil
	}
}

func (w winCoinRewardDo) FindByPage(offset int, limit int) (result []*model.WinCoinReward, count int64, err error) {
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

func (w winCoinRewardDo) ScanByPage(result interface{}, offset int, limit int) (count int64, err error) {
	count, err = w.Count()
	if err != nil {
		return
	}

	err = w.Offset(offset).Limit(limit).Scan(result)
	return
}

func (w winCoinRewardDo) Scan(result interface{}) (err error) {
	return w.DO.Scan(result)
}

func (w winCoinRewardDo) Delete(models ...*model.WinCoinReward) (result gen.ResultInfo, err error) {
	return w.DO.Delete(models)
}

func (w *winCoinRewardDo) withDO(do gen.Dao) *winCoinRewardDo {
	w.DO = *do.(*gen.DO)
	return w
}
