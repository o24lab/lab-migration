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

func newWinCoinCommission(db *gorm.DB, opts ...gen.DOOption) winCoinCommission {
	_winCoinCommission := winCoinCommission{}

	_winCoinCommission.winCoinCommissionDo.UseDB(db, opts...)
	_winCoinCommission.winCoinCommissionDo.UseModel(&model.WinCoinCommission{})

	tableName := _winCoinCommission.winCoinCommissionDo.TableName()
	_winCoinCommission.ALL = field.NewAsterisk(tableName)
	_winCoinCommission.ID = field.NewInt64(tableName, "id")
	_winCoinCommission.UID = field.NewInt32(tableName, "uid")
	_winCoinCommission.Username = field.NewString(tableName, "username")
	_winCoinCommission.AgentLevel = field.NewInt32(tableName, "agent_level")
	_winCoinCommission.Riqi = field.NewInt32(tableName, "riqi")
	_winCoinCommission.Coin = field.NewFloat64(tableName, "coin")
	_winCoinCommission.SubUID = field.NewInt32(tableName, "sub_uid")
	_winCoinCommission.SubUsername = field.NewString(tableName, "sub_username")
	_winCoinCommission.SubBetTrunover = field.NewFloat64(tableName, "sub_bet_trunover")
	_winCoinCommission.Rate = field.NewFloat64(tableName, "rate")
	_winCoinCommission.CoinBefore = field.NewFloat64(tableName, "coin_before")
	_winCoinCommission.Status = field.NewInt32(tableName, "status")
	_winCoinCommission.CreatedAt = field.NewInt32(tableName, "created_at")
	_winCoinCommission.UpdatedAt = field.NewInt32(tableName, "updated_at")

	_winCoinCommission.fillFieldMap()

	return _winCoinCommission
}

// winCoinCommission 佣金表
type winCoinCommission struct {
	winCoinCommissionDo

	ALL            field.Asterisk
	ID             field.Int64
	UID            field.Int32   // 代理UID
	Username       field.String  // 用户名
	AgentLevel     field.Int32   // 代理层级
	Riqi           field.Int32   // 佣金时间
	Coin           field.Float64 // 佣金金额
	SubUID         field.Int32   // 下级UID
	SubUsername    field.String  // 下级用户名
	SubBetTrunover field.Float64 // 下级流水总额
	Rate           field.Float64 // 佣金比例
	CoinBefore     field.Float64 // 即时余额
	Status         field.Int32   // 状态:0-未发放 1-已发放
	CreatedAt      field.Int32
	UpdatedAt      field.Int32

	fieldMap map[string]field.Expr
}

func (w winCoinCommission) Table(newTableName string) *winCoinCommission {
	w.winCoinCommissionDo.UseTable(newTableName)
	return w.updateTableName(newTableName)
}

func (w winCoinCommission) As(alias string) *winCoinCommission {
	w.winCoinCommissionDo.DO = *(w.winCoinCommissionDo.As(alias).(*gen.DO))
	return w.updateTableName(alias)
}

func (w *winCoinCommission) updateTableName(table string) *winCoinCommission {
	w.ALL = field.NewAsterisk(table)
	w.ID = field.NewInt64(table, "id")
	w.UID = field.NewInt32(table, "uid")
	w.Username = field.NewString(table, "username")
	w.AgentLevel = field.NewInt32(table, "agent_level")
	w.Riqi = field.NewInt32(table, "riqi")
	w.Coin = field.NewFloat64(table, "coin")
	w.SubUID = field.NewInt32(table, "sub_uid")
	w.SubUsername = field.NewString(table, "sub_username")
	w.SubBetTrunover = field.NewFloat64(table, "sub_bet_trunover")
	w.Rate = field.NewFloat64(table, "rate")
	w.CoinBefore = field.NewFloat64(table, "coin_before")
	w.Status = field.NewInt32(table, "status")
	w.CreatedAt = field.NewInt32(table, "created_at")
	w.UpdatedAt = field.NewInt32(table, "updated_at")

	w.fillFieldMap()

	return w
}

func (w *winCoinCommission) GetFieldByName(fieldName string) (field.OrderExpr, bool) {
	_f, ok := w.fieldMap[fieldName]
	if !ok || _f == nil {
		return nil, false
	}
	_oe, ok := _f.(field.OrderExpr)
	return _oe, ok
}

func (w *winCoinCommission) fillFieldMap() {
	w.fieldMap = make(map[string]field.Expr, 14)
	w.fieldMap["id"] = w.ID
	w.fieldMap["uid"] = w.UID
	w.fieldMap["username"] = w.Username
	w.fieldMap["agent_level"] = w.AgentLevel
	w.fieldMap["riqi"] = w.Riqi
	w.fieldMap["coin"] = w.Coin
	w.fieldMap["sub_uid"] = w.SubUID
	w.fieldMap["sub_username"] = w.SubUsername
	w.fieldMap["sub_bet_trunover"] = w.SubBetTrunover
	w.fieldMap["rate"] = w.Rate
	w.fieldMap["coin_before"] = w.CoinBefore
	w.fieldMap["status"] = w.Status
	w.fieldMap["created_at"] = w.CreatedAt
	w.fieldMap["updated_at"] = w.UpdatedAt
}

func (w winCoinCommission) clone(db *gorm.DB) winCoinCommission {
	w.winCoinCommissionDo.ReplaceConnPool(db.Statement.ConnPool)
	return w
}

func (w winCoinCommission) replaceDB(db *gorm.DB) winCoinCommission {
	w.winCoinCommissionDo.ReplaceDB(db)
	return w
}

type winCoinCommissionDo struct{ gen.DO }

func (w winCoinCommissionDo) Debug() *winCoinCommissionDo {
	return w.withDO(w.DO.Debug())
}

func (w winCoinCommissionDo) WithContext(ctx context.Context) *winCoinCommissionDo {
	return w.withDO(w.DO.WithContext(ctx))
}

func (w winCoinCommissionDo) ReadDB() *winCoinCommissionDo {
	return w.Clauses(dbresolver.Read)
}

func (w winCoinCommissionDo) WriteDB() *winCoinCommissionDo {
	return w.Clauses(dbresolver.Write)
}

func (w winCoinCommissionDo) Session(config *gorm.Session) *winCoinCommissionDo {
	return w.withDO(w.DO.Session(config))
}

func (w winCoinCommissionDo) Clauses(conds ...clause.Expression) *winCoinCommissionDo {
	return w.withDO(w.DO.Clauses(conds...))
}

func (w winCoinCommissionDo) Returning(value interface{}, columns ...string) *winCoinCommissionDo {
	return w.withDO(w.DO.Returning(value, columns...))
}

func (w winCoinCommissionDo) Not(conds ...gen.Condition) *winCoinCommissionDo {
	return w.withDO(w.DO.Not(conds...))
}

func (w winCoinCommissionDo) Or(conds ...gen.Condition) *winCoinCommissionDo {
	return w.withDO(w.DO.Or(conds...))
}

func (w winCoinCommissionDo) Select(conds ...field.Expr) *winCoinCommissionDo {
	return w.withDO(w.DO.Select(conds...))
}

func (w winCoinCommissionDo) Where(conds ...gen.Condition) *winCoinCommissionDo {
	return w.withDO(w.DO.Where(conds...))
}

func (w winCoinCommissionDo) Order(conds ...field.Expr) *winCoinCommissionDo {
	return w.withDO(w.DO.Order(conds...))
}

func (w winCoinCommissionDo) Distinct(cols ...field.Expr) *winCoinCommissionDo {
	return w.withDO(w.DO.Distinct(cols...))
}

func (w winCoinCommissionDo) Omit(cols ...field.Expr) *winCoinCommissionDo {
	return w.withDO(w.DO.Omit(cols...))
}

func (w winCoinCommissionDo) Join(table schema.Tabler, on ...field.Expr) *winCoinCommissionDo {
	return w.withDO(w.DO.Join(table, on...))
}

func (w winCoinCommissionDo) LeftJoin(table schema.Tabler, on ...field.Expr) *winCoinCommissionDo {
	return w.withDO(w.DO.LeftJoin(table, on...))
}

func (w winCoinCommissionDo) RightJoin(table schema.Tabler, on ...field.Expr) *winCoinCommissionDo {
	return w.withDO(w.DO.RightJoin(table, on...))
}

func (w winCoinCommissionDo) Group(cols ...field.Expr) *winCoinCommissionDo {
	return w.withDO(w.DO.Group(cols...))
}

func (w winCoinCommissionDo) Having(conds ...gen.Condition) *winCoinCommissionDo {
	return w.withDO(w.DO.Having(conds...))
}

func (w winCoinCommissionDo) Limit(limit int) *winCoinCommissionDo {
	return w.withDO(w.DO.Limit(limit))
}

func (w winCoinCommissionDo) Offset(offset int) *winCoinCommissionDo {
	return w.withDO(w.DO.Offset(offset))
}

func (w winCoinCommissionDo) Scopes(funcs ...func(gen.Dao) gen.Dao) *winCoinCommissionDo {
	return w.withDO(w.DO.Scopes(funcs...))
}

func (w winCoinCommissionDo) Unscoped() *winCoinCommissionDo {
	return w.withDO(w.DO.Unscoped())
}

func (w winCoinCommissionDo) Create(values ...*model.WinCoinCommission) error {
	if len(values) == 0 {
		return nil
	}
	return w.DO.Create(values)
}

func (w winCoinCommissionDo) CreateInBatches(values []*model.WinCoinCommission, batchSize int) error {
	return w.DO.CreateInBatches(values, batchSize)
}

// Save : !!! underlying implementation is different with GORM
// The method is equivalent to executing the statement: db.Clauses(clause.OnConflict{UpdateAll: true}).Create(values)
func (w winCoinCommissionDo) Save(values ...*model.WinCoinCommission) error {
	if len(values) == 0 {
		return nil
	}
	return w.DO.Save(values)
}

func (w winCoinCommissionDo) First() (*model.WinCoinCommission, error) {
	if result, err := w.DO.First(); err != nil {
		return nil, err
	} else {
		return result.(*model.WinCoinCommission), nil
	}
}

func (w winCoinCommissionDo) Take() (*model.WinCoinCommission, error) {
	if result, err := w.DO.Take(); err != nil {
		return nil, err
	} else {
		return result.(*model.WinCoinCommission), nil
	}
}

func (w winCoinCommissionDo) Last() (*model.WinCoinCommission, error) {
	if result, err := w.DO.Last(); err != nil {
		return nil, err
	} else {
		return result.(*model.WinCoinCommission), nil
	}
}

func (w winCoinCommissionDo) Find() ([]*model.WinCoinCommission, error) {
	result, err := w.DO.Find()
	return result.([]*model.WinCoinCommission), err
}

func (w winCoinCommissionDo) FindInBatch(batchSize int, fc func(tx gen.Dao, batch int) error) (results []*model.WinCoinCommission, err error) {
	buf := make([]*model.WinCoinCommission, 0, batchSize)
	err = w.DO.FindInBatches(&buf, batchSize, func(tx gen.Dao, batch int) error {
		defer func() { results = append(results, buf...) }()
		return fc(tx, batch)
	})
	return results, err
}

func (w winCoinCommissionDo) FindInBatches(result *[]*model.WinCoinCommission, batchSize int, fc func(tx gen.Dao, batch int) error) error {
	return w.DO.FindInBatches(result, batchSize, fc)
}

func (w winCoinCommissionDo) Attrs(attrs ...field.AssignExpr) *winCoinCommissionDo {
	return w.withDO(w.DO.Attrs(attrs...))
}

func (w winCoinCommissionDo) Assign(attrs ...field.AssignExpr) *winCoinCommissionDo {
	return w.withDO(w.DO.Assign(attrs...))
}

func (w winCoinCommissionDo) Joins(fields ...field.RelationField) *winCoinCommissionDo {
	for _, _f := range fields {
		w = *w.withDO(w.DO.Joins(_f))
	}
	return &w
}

func (w winCoinCommissionDo) Preload(fields ...field.RelationField) *winCoinCommissionDo {
	for _, _f := range fields {
		w = *w.withDO(w.DO.Preload(_f))
	}
	return &w
}

func (w winCoinCommissionDo) FirstOrInit() (*model.WinCoinCommission, error) {
	if result, err := w.DO.FirstOrInit(); err != nil {
		return nil, err
	} else {
		return result.(*model.WinCoinCommission), nil
	}
}

func (w winCoinCommissionDo) FirstOrCreate() (*model.WinCoinCommission, error) {
	if result, err := w.DO.FirstOrCreate(); err != nil {
		return nil, err
	} else {
		return result.(*model.WinCoinCommission), nil
	}
}

func (w winCoinCommissionDo) FindByPage(offset int, limit int) (result []*model.WinCoinCommission, count int64, err error) {
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

func (w winCoinCommissionDo) ScanByPage(result interface{}, offset int, limit int) (count int64, err error) {
	count, err = w.Count()
	if err != nil {
		return
	}

	err = w.Offset(offset).Limit(limit).Scan(result)
	return
}

func (w winCoinCommissionDo) Scan(result interface{}) (err error) {
	return w.DO.Scan(result)
}

func (w winCoinCommissionDo) Delete(models ...*model.WinCoinCommission) (result gen.ResultInfo, err error) {
	return w.DO.Delete(models)
}

func (w *winCoinCommissionDo) withDO(do gen.Dao) *winCoinCommissionDo {
	w.DO = *do.(*gen.DO)
	return w
}
