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

func newWinMonthlyTax(db *gorm.DB, opts ...gen.DOOption) winMonthlyTax {
	_winMonthlyTax := winMonthlyTax{}

	_winMonthlyTax.winMonthlyTaxDo.UseDB(db, opts...)
	_winMonthlyTax.winMonthlyTaxDo.UseModel(&model.WinMonthlyTax{})

	tableName := _winMonthlyTax.winMonthlyTaxDo.TableName()
	_winMonthlyTax.ALL = field.NewAsterisk(tableName)
	_winMonthlyTax.ID = field.NewInt32(tableName, "id")
	_winMonthlyTax.GameID = field.NewInt32(tableName, "game_id")
	_winMonthlyTax.GameName = field.NewString(tableName, "game_name")
	_winMonthlyTax.Year = field.NewInt32(tableName, "year")
	_winMonthlyTax.Month = field.NewInt32(tableName, "month")
	_winMonthlyTax.CoinBet = field.NewFloat64(tableName, "coin_bet")
	_winMonthlyTax.CoinProfit = field.NewFloat64(tableName, "coin_profit")
	_winMonthlyTax.Rate = field.NewFloat64(tableName, "rate")
	_winMonthlyTax.CoinTax = field.NewFloat64(tableName, "coin_tax")
	_winMonthlyTax.CreatedAt = field.NewInt32(tableName, "created_at")
	_winMonthlyTax.UpdatedAt = field.NewInt32(tableName, "updated_at")

	_winMonthlyTax.fillFieldMap()

	return _winMonthlyTax
}

// winMonthlyTax 每月税收统计
type winMonthlyTax struct {
	winMonthlyTaxDo

	ALL        field.Asterisk
	ID         field.Int32   // ID
	GameID     field.Int32   // 游戏ID
	GameName   field.String  // 游戏名称
	Year       field.Int32   // 统计年
	Month      field.Int32   // 统计月
	CoinBet    field.Float64 // 投注金额
	CoinProfit field.Float64 // 负盈利金额
	Rate       field.Float64 // 税收比例
	CoinTax    field.Float64 // 税收金额
	CreatedAt  field.Int32   // 创建时间
	UpdatedAt  field.Int32   // 修改时间

	fieldMap map[string]field.Expr
}

func (w winMonthlyTax) Table(newTableName string) *winMonthlyTax {
	w.winMonthlyTaxDo.UseTable(newTableName)
	return w.updateTableName(newTableName)
}

func (w winMonthlyTax) As(alias string) *winMonthlyTax {
	w.winMonthlyTaxDo.DO = *(w.winMonthlyTaxDo.As(alias).(*gen.DO))
	return w.updateTableName(alias)
}

func (w *winMonthlyTax) updateTableName(table string) *winMonthlyTax {
	w.ALL = field.NewAsterisk(table)
	w.ID = field.NewInt32(table, "id")
	w.GameID = field.NewInt32(table, "game_id")
	w.GameName = field.NewString(table, "game_name")
	w.Year = field.NewInt32(table, "year")
	w.Month = field.NewInt32(table, "month")
	w.CoinBet = field.NewFloat64(table, "coin_bet")
	w.CoinProfit = field.NewFloat64(table, "coin_profit")
	w.Rate = field.NewFloat64(table, "rate")
	w.CoinTax = field.NewFloat64(table, "coin_tax")
	w.CreatedAt = field.NewInt32(table, "created_at")
	w.UpdatedAt = field.NewInt32(table, "updated_at")

	w.fillFieldMap()

	return w
}

func (w *winMonthlyTax) GetFieldByName(fieldName string) (field.OrderExpr, bool) {
	_f, ok := w.fieldMap[fieldName]
	if !ok || _f == nil {
		return nil, false
	}
	_oe, ok := _f.(field.OrderExpr)
	return _oe, ok
}

func (w *winMonthlyTax) fillFieldMap() {
	w.fieldMap = make(map[string]field.Expr, 11)
	w.fieldMap["id"] = w.ID
	w.fieldMap["game_id"] = w.GameID
	w.fieldMap["game_name"] = w.GameName
	w.fieldMap["year"] = w.Year
	w.fieldMap["month"] = w.Month
	w.fieldMap["coin_bet"] = w.CoinBet
	w.fieldMap["coin_profit"] = w.CoinProfit
	w.fieldMap["rate"] = w.Rate
	w.fieldMap["coin_tax"] = w.CoinTax
	w.fieldMap["created_at"] = w.CreatedAt
	w.fieldMap["updated_at"] = w.UpdatedAt
}

func (w winMonthlyTax) clone(db *gorm.DB) winMonthlyTax {
	w.winMonthlyTaxDo.ReplaceConnPool(db.Statement.ConnPool)
	return w
}

func (w winMonthlyTax) replaceDB(db *gorm.DB) winMonthlyTax {
	w.winMonthlyTaxDo.ReplaceDB(db)
	return w
}

type winMonthlyTaxDo struct{ gen.DO }

func (w winMonthlyTaxDo) Debug() *winMonthlyTaxDo {
	return w.withDO(w.DO.Debug())
}

func (w winMonthlyTaxDo) WithContext(ctx context.Context) *winMonthlyTaxDo {
	return w.withDO(w.DO.WithContext(ctx))
}

func (w winMonthlyTaxDo) ReadDB() *winMonthlyTaxDo {
	return w.Clauses(dbresolver.Read)
}

func (w winMonthlyTaxDo) WriteDB() *winMonthlyTaxDo {
	return w.Clauses(dbresolver.Write)
}

func (w winMonthlyTaxDo) Session(config *gorm.Session) *winMonthlyTaxDo {
	return w.withDO(w.DO.Session(config))
}

func (w winMonthlyTaxDo) Clauses(conds ...clause.Expression) *winMonthlyTaxDo {
	return w.withDO(w.DO.Clauses(conds...))
}

func (w winMonthlyTaxDo) Returning(value interface{}, columns ...string) *winMonthlyTaxDo {
	return w.withDO(w.DO.Returning(value, columns...))
}

func (w winMonthlyTaxDo) Not(conds ...gen.Condition) *winMonthlyTaxDo {
	return w.withDO(w.DO.Not(conds...))
}

func (w winMonthlyTaxDo) Or(conds ...gen.Condition) *winMonthlyTaxDo {
	return w.withDO(w.DO.Or(conds...))
}

func (w winMonthlyTaxDo) Select(conds ...field.Expr) *winMonthlyTaxDo {
	return w.withDO(w.DO.Select(conds...))
}

func (w winMonthlyTaxDo) Where(conds ...gen.Condition) *winMonthlyTaxDo {
	return w.withDO(w.DO.Where(conds...))
}

func (w winMonthlyTaxDo) Order(conds ...field.Expr) *winMonthlyTaxDo {
	return w.withDO(w.DO.Order(conds...))
}

func (w winMonthlyTaxDo) Distinct(cols ...field.Expr) *winMonthlyTaxDo {
	return w.withDO(w.DO.Distinct(cols...))
}

func (w winMonthlyTaxDo) Omit(cols ...field.Expr) *winMonthlyTaxDo {
	return w.withDO(w.DO.Omit(cols...))
}

func (w winMonthlyTaxDo) Join(table schema.Tabler, on ...field.Expr) *winMonthlyTaxDo {
	return w.withDO(w.DO.Join(table, on...))
}

func (w winMonthlyTaxDo) LeftJoin(table schema.Tabler, on ...field.Expr) *winMonthlyTaxDo {
	return w.withDO(w.DO.LeftJoin(table, on...))
}

func (w winMonthlyTaxDo) RightJoin(table schema.Tabler, on ...field.Expr) *winMonthlyTaxDo {
	return w.withDO(w.DO.RightJoin(table, on...))
}

func (w winMonthlyTaxDo) Group(cols ...field.Expr) *winMonthlyTaxDo {
	return w.withDO(w.DO.Group(cols...))
}

func (w winMonthlyTaxDo) Having(conds ...gen.Condition) *winMonthlyTaxDo {
	return w.withDO(w.DO.Having(conds...))
}

func (w winMonthlyTaxDo) Limit(limit int) *winMonthlyTaxDo {
	return w.withDO(w.DO.Limit(limit))
}

func (w winMonthlyTaxDo) Offset(offset int) *winMonthlyTaxDo {
	return w.withDO(w.DO.Offset(offset))
}

func (w winMonthlyTaxDo) Scopes(funcs ...func(gen.Dao) gen.Dao) *winMonthlyTaxDo {
	return w.withDO(w.DO.Scopes(funcs...))
}

func (w winMonthlyTaxDo) Unscoped() *winMonthlyTaxDo {
	return w.withDO(w.DO.Unscoped())
}

func (w winMonthlyTaxDo) Create(values ...*model.WinMonthlyTax) error {
	if len(values) == 0 {
		return nil
	}
	return w.DO.Create(values)
}

func (w winMonthlyTaxDo) CreateInBatches(values []*model.WinMonthlyTax, batchSize int) error {
	return w.DO.CreateInBatches(values, batchSize)
}

// Save : !!! underlying implementation is different with GORM
// The method is equivalent to executing the statement: db.Clauses(clause.OnConflict{UpdateAll: true}).Create(values)
func (w winMonthlyTaxDo) Save(values ...*model.WinMonthlyTax) error {
	if len(values) == 0 {
		return nil
	}
	return w.DO.Save(values)
}

func (w winMonthlyTaxDo) First() (*model.WinMonthlyTax, error) {
	if result, err := w.DO.First(); err != nil {
		return nil, err
	} else {
		return result.(*model.WinMonthlyTax), nil
	}
}

func (w winMonthlyTaxDo) Take() (*model.WinMonthlyTax, error) {
	if result, err := w.DO.Take(); err != nil {
		return nil, err
	} else {
		return result.(*model.WinMonthlyTax), nil
	}
}

func (w winMonthlyTaxDo) Last() (*model.WinMonthlyTax, error) {
	if result, err := w.DO.Last(); err != nil {
		return nil, err
	} else {
		return result.(*model.WinMonthlyTax), nil
	}
}

func (w winMonthlyTaxDo) Find() ([]*model.WinMonthlyTax, error) {
	result, err := w.DO.Find()
	return result.([]*model.WinMonthlyTax), err
}

func (w winMonthlyTaxDo) FindInBatch(batchSize int, fc func(tx gen.Dao, batch int) error) (results []*model.WinMonthlyTax, err error) {
	buf := make([]*model.WinMonthlyTax, 0, batchSize)
	err = w.DO.FindInBatches(&buf, batchSize, func(tx gen.Dao, batch int) error {
		defer func() { results = append(results, buf...) }()
		return fc(tx, batch)
	})
	return results, err
}

func (w winMonthlyTaxDo) FindInBatches(result *[]*model.WinMonthlyTax, batchSize int, fc func(tx gen.Dao, batch int) error) error {
	return w.DO.FindInBatches(result, batchSize, fc)
}

func (w winMonthlyTaxDo) Attrs(attrs ...field.AssignExpr) *winMonthlyTaxDo {
	return w.withDO(w.DO.Attrs(attrs...))
}

func (w winMonthlyTaxDo) Assign(attrs ...field.AssignExpr) *winMonthlyTaxDo {
	return w.withDO(w.DO.Assign(attrs...))
}

func (w winMonthlyTaxDo) Joins(fields ...field.RelationField) *winMonthlyTaxDo {
	for _, _f := range fields {
		w = *w.withDO(w.DO.Joins(_f))
	}
	return &w
}

func (w winMonthlyTaxDo) Preload(fields ...field.RelationField) *winMonthlyTaxDo {
	for _, _f := range fields {
		w = *w.withDO(w.DO.Preload(_f))
	}
	return &w
}

func (w winMonthlyTaxDo) FirstOrInit() (*model.WinMonthlyTax, error) {
	if result, err := w.DO.FirstOrInit(); err != nil {
		return nil, err
	} else {
		return result.(*model.WinMonthlyTax), nil
	}
}

func (w winMonthlyTaxDo) FirstOrCreate() (*model.WinMonthlyTax, error) {
	if result, err := w.DO.FirstOrCreate(); err != nil {
		return nil, err
	} else {
		return result.(*model.WinMonthlyTax), nil
	}
}

func (w winMonthlyTaxDo) FindByPage(offset int, limit int) (result []*model.WinMonthlyTax, count int64, err error) {
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

func (w winMonthlyTaxDo) ScanByPage(result interface{}, offset int, limit int) (count int64, err error) {
	count, err = w.Count()
	if err != nil {
		return
	}

	err = w.Offset(offset).Limit(limit).Scan(result)
	return
}

func (w winMonthlyTaxDo) Scan(result interface{}) (err error) {
	return w.DO.Scan(result)
}

func (w winMonthlyTaxDo) Delete(models ...*model.WinMonthlyTax) (result gen.ResultInfo, err error) {
	return w.DO.Delete(models)
}

func (w *winMonthlyTaxDo) withDO(do gen.Dao) *winMonthlyTaxDo {
	w.DO = *do.(*gen.DO)
	return w
}
