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

func newWinUserValidRebate(db *gorm.DB, opts ...gen.DOOption) winUserValidRebate {
	_winUserValidRebate := winUserValidRebate{}

	_winUserValidRebate.winUserValidRebateDo.UseDB(db, opts...)
	_winUserValidRebate.winUserValidRebateDo.UseModel(&model.WinUserValidRebate{})

	tableName := _winUserValidRebate.winUserValidRebateDo.TableName()
	_winUserValidRebate.ALL = field.NewAsterisk(tableName)
	_winUserValidRebate.ID = field.NewInt64(tableName, "id")
	_winUserValidRebate.ReportDate = field.NewInt32(tableName, "report_date")
	_winUserValidRebate.UID = field.NewInt32(tableName, "uid")
	_winUserValidRebate.YesterdayValidAmount = field.NewFloat64(tableName, "yesterday_valid_amount")
	_winUserValidRebate.ValidAmount = field.NewFloat64(tableName, "valid_amount")
	_winUserValidRebate.CalcValidAmount = field.NewFloat64(tableName, "calc_valid_amount")
	_winUserValidRebate.CreatedAt = field.NewInt32(tableName, "created_at")
	_winUserValidRebate.UpdatedAt = field.NewInt32(tableName, "updated_at")

	_winUserValidRebate.fillFieldMap()

	return _winUserValidRebate
}

// winUserValidRebate 用户有效流水返水记录表
type winUserValidRebate struct {
	winUserValidRebateDo

	ALL                  field.Asterisk
	ID                   field.Int64   // 自增主键
	ReportDate           field.Int32   // 统计日期
	UID                  field.Int32   // 用户ID
	YesterdayValidAmount field.Float64 // 前一日有效投注额
	ValidAmount          field.Float64 // 当前有效投注额
	CalcValidAmount      field.Float64 // 计算日有效投注额:前一日有效投注额-当前有效投注额
	CreatedAt            field.Int32   // 创建时间
	UpdatedAt            field.Int32   // 修改时间

	fieldMap map[string]field.Expr
}

func (w winUserValidRebate) Table(newTableName string) *winUserValidRebate {
	w.winUserValidRebateDo.UseTable(newTableName)
	return w.updateTableName(newTableName)
}

func (w winUserValidRebate) As(alias string) *winUserValidRebate {
	w.winUserValidRebateDo.DO = *(w.winUserValidRebateDo.As(alias).(*gen.DO))
	return w.updateTableName(alias)
}

func (w *winUserValidRebate) updateTableName(table string) *winUserValidRebate {
	w.ALL = field.NewAsterisk(table)
	w.ID = field.NewInt64(table, "id")
	w.ReportDate = field.NewInt32(table, "report_date")
	w.UID = field.NewInt32(table, "uid")
	w.YesterdayValidAmount = field.NewFloat64(table, "yesterday_valid_amount")
	w.ValidAmount = field.NewFloat64(table, "valid_amount")
	w.CalcValidAmount = field.NewFloat64(table, "calc_valid_amount")
	w.CreatedAt = field.NewInt32(table, "created_at")
	w.UpdatedAt = field.NewInt32(table, "updated_at")

	w.fillFieldMap()

	return w
}

func (w *winUserValidRebate) GetFieldByName(fieldName string) (field.OrderExpr, bool) {
	_f, ok := w.fieldMap[fieldName]
	if !ok || _f == nil {
		return nil, false
	}
	_oe, ok := _f.(field.OrderExpr)
	return _oe, ok
}

func (w *winUserValidRebate) fillFieldMap() {
	w.fieldMap = make(map[string]field.Expr, 8)
	w.fieldMap["id"] = w.ID
	w.fieldMap["report_date"] = w.ReportDate
	w.fieldMap["uid"] = w.UID
	w.fieldMap["yesterday_valid_amount"] = w.YesterdayValidAmount
	w.fieldMap["valid_amount"] = w.ValidAmount
	w.fieldMap["calc_valid_amount"] = w.CalcValidAmount
	w.fieldMap["created_at"] = w.CreatedAt
	w.fieldMap["updated_at"] = w.UpdatedAt
}

func (w winUserValidRebate) clone(db *gorm.DB) winUserValidRebate {
	w.winUserValidRebateDo.ReplaceConnPool(db.Statement.ConnPool)
	return w
}

func (w winUserValidRebate) replaceDB(db *gorm.DB) winUserValidRebate {
	w.winUserValidRebateDo.ReplaceDB(db)
	return w
}

type winUserValidRebateDo struct{ gen.DO }

func (w winUserValidRebateDo) Debug() *winUserValidRebateDo {
	return w.withDO(w.DO.Debug())
}

func (w winUserValidRebateDo) WithContext(ctx context.Context) *winUserValidRebateDo {
	return w.withDO(w.DO.WithContext(ctx))
}

func (w winUserValidRebateDo) ReadDB() *winUserValidRebateDo {
	return w.Clauses(dbresolver.Read)
}

func (w winUserValidRebateDo) WriteDB() *winUserValidRebateDo {
	return w.Clauses(dbresolver.Write)
}

func (w winUserValidRebateDo) Session(config *gorm.Session) *winUserValidRebateDo {
	return w.withDO(w.DO.Session(config))
}

func (w winUserValidRebateDo) Clauses(conds ...clause.Expression) *winUserValidRebateDo {
	return w.withDO(w.DO.Clauses(conds...))
}

func (w winUserValidRebateDo) Returning(value interface{}, columns ...string) *winUserValidRebateDo {
	return w.withDO(w.DO.Returning(value, columns...))
}

func (w winUserValidRebateDo) Not(conds ...gen.Condition) *winUserValidRebateDo {
	return w.withDO(w.DO.Not(conds...))
}

func (w winUserValidRebateDo) Or(conds ...gen.Condition) *winUserValidRebateDo {
	return w.withDO(w.DO.Or(conds...))
}

func (w winUserValidRebateDo) Select(conds ...field.Expr) *winUserValidRebateDo {
	return w.withDO(w.DO.Select(conds...))
}

func (w winUserValidRebateDo) Where(conds ...gen.Condition) *winUserValidRebateDo {
	return w.withDO(w.DO.Where(conds...))
}

func (w winUserValidRebateDo) Order(conds ...field.Expr) *winUserValidRebateDo {
	return w.withDO(w.DO.Order(conds...))
}

func (w winUserValidRebateDo) Distinct(cols ...field.Expr) *winUserValidRebateDo {
	return w.withDO(w.DO.Distinct(cols...))
}

func (w winUserValidRebateDo) Omit(cols ...field.Expr) *winUserValidRebateDo {
	return w.withDO(w.DO.Omit(cols...))
}

func (w winUserValidRebateDo) Join(table schema.Tabler, on ...field.Expr) *winUserValidRebateDo {
	return w.withDO(w.DO.Join(table, on...))
}

func (w winUserValidRebateDo) LeftJoin(table schema.Tabler, on ...field.Expr) *winUserValidRebateDo {
	return w.withDO(w.DO.LeftJoin(table, on...))
}

func (w winUserValidRebateDo) RightJoin(table schema.Tabler, on ...field.Expr) *winUserValidRebateDo {
	return w.withDO(w.DO.RightJoin(table, on...))
}

func (w winUserValidRebateDo) Group(cols ...field.Expr) *winUserValidRebateDo {
	return w.withDO(w.DO.Group(cols...))
}

func (w winUserValidRebateDo) Having(conds ...gen.Condition) *winUserValidRebateDo {
	return w.withDO(w.DO.Having(conds...))
}

func (w winUserValidRebateDo) Limit(limit int) *winUserValidRebateDo {
	return w.withDO(w.DO.Limit(limit))
}

func (w winUserValidRebateDo) Offset(offset int) *winUserValidRebateDo {
	return w.withDO(w.DO.Offset(offset))
}

func (w winUserValidRebateDo) Scopes(funcs ...func(gen.Dao) gen.Dao) *winUserValidRebateDo {
	return w.withDO(w.DO.Scopes(funcs...))
}

func (w winUserValidRebateDo) Unscoped() *winUserValidRebateDo {
	return w.withDO(w.DO.Unscoped())
}

func (w winUserValidRebateDo) Create(values ...*model.WinUserValidRebate) error {
	if len(values) == 0 {
		return nil
	}
	return w.DO.Create(values)
}

func (w winUserValidRebateDo) CreateInBatches(values []*model.WinUserValidRebate, batchSize int) error {
	return w.DO.CreateInBatches(values, batchSize)
}

// Save : !!! underlying implementation is different with GORM
// The method is equivalent to executing the statement: db.Clauses(clause.OnConflict{UpdateAll: true}).Create(values)
func (w winUserValidRebateDo) Save(values ...*model.WinUserValidRebate) error {
	if len(values) == 0 {
		return nil
	}
	return w.DO.Save(values)
}

func (w winUserValidRebateDo) First() (*model.WinUserValidRebate, error) {
	if result, err := w.DO.First(); err != nil {
		return nil, err
	} else {
		return result.(*model.WinUserValidRebate), nil
	}
}

func (w winUserValidRebateDo) Take() (*model.WinUserValidRebate, error) {
	if result, err := w.DO.Take(); err != nil {
		return nil, err
	} else {
		return result.(*model.WinUserValidRebate), nil
	}
}

func (w winUserValidRebateDo) Last() (*model.WinUserValidRebate, error) {
	if result, err := w.DO.Last(); err != nil {
		return nil, err
	} else {
		return result.(*model.WinUserValidRebate), nil
	}
}

func (w winUserValidRebateDo) Find() ([]*model.WinUserValidRebate, error) {
	result, err := w.DO.Find()
	return result.([]*model.WinUserValidRebate), err
}

func (w winUserValidRebateDo) FindInBatch(batchSize int, fc func(tx gen.Dao, batch int) error) (results []*model.WinUserValidRebate, err error) {
	buf := make([]*model.WinUserValidRebate, 0, batchSize)
	err = w.DO.FindInBatches(&buf, batchSize, func(tx gen.Dao, batch int) error {
		defer func() { results = append(results, buf...) }()
		return fc(tx, batch)
	})
	return results, err
}

func (w winUserValidRebateDo) FindInBatches(result *[]*model.WinUserValidRebate, batchSize int, fc func(tx gen.Dao, batch int) error) error {
	return w.DO.FindInBatches(result, batchSize, fc)
}

func (w winUserValidRebateDo) Attrs(attrs ...field.AssignExpr) *winUserValidRebateDo {
	return w.withDO(w.DO.Attrs(attrs...))
}

func (w winUserValidRebateDo) Assign(attrs ...field.AssignExpr) *winUserValidRebateDo {
	return w.withDO(w.DO.Assign(attrs...))
}

func (w winUserValidRebateDo) Joins(fields ...field.RelationField) *winUserValidRebateDo {
	for _, _f := range fields {
		w = *w.withDO(w.DO.Joins(_f))
	}
	return &w
}

func (w winUserValidRebateDo) Preload(fields ...field.RelationField) *winUserValidRebateDo {
	for _, _f := range fields {
		w = *w.withDO(w.DO.Preload(_f))
	}
	return &w
}

func (w winUserValidRebateDo) FirstOrInit() (*model.WinUserValidRebate, error) {
	if result, err := w.DO.FirstOrInit(); err != nil {
		return nil, err
	} else {
		return result.(*model.WinUserValidRebate), nil
	}
}

func (w winUserValidRebateDo) FirstOrCreate() (*model.WinUserValidRebate, error) {
	if result, err := w.DO.FirstOrCreate(); err != nil {
		return nil, err
	} else {
		return result.(*model.WinUserValidRebate), nil
	}
}

func (w winUserValidRebateDo) FindByPage(offset int, limit int) (result []*model.WinUserValidRebate, count int64, err error) {
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

func (w winUserValidRebateDo) ScanByPage(result interface{}, offset int, limit int) (count int64, err error) {
	count, err = w.Count()
	if err != nil {
		return
	}

	err = w.Offset(offset).Limit(limit).Scan(result)
	return
}

func (w winUserValidRebateDo) Scan(result interface{}) (err error) {
	return w.DO.Scan(result)
}

func (w winUserValidRebateDo) Delete(models ...*model.WinUserValidRebate) (result gen.ResultInfo, err error) {
	return w.DO.Delete(models)
}

func (w *winUserValidRebateDo) withDO(do gen.Dao) *winUserValidRebateDo {
	w.DO = *do.(*gen.DO)
	return w
}
