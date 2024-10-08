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

func newWinCodeRecordsRequire(db *gorm.DB, opts ...gen.DOOption) winCodeRecordsRequire {
	_winCodeRecordsRequire := winCodeRecordsRequire{}

	_winCodeRecordsRequire.winCodeRecordsRequireDo.UseDB(db, opts...)
	_winCodeRecordsRequire.winCodeRecordsRequireDo.UseModel(&model.WinCodeRecordsRequire{})

	tableName := _winCodeRecordsRequire.winCodeRecordsRequireDo.TableName()
	_winCodeRecordsRequire.ALL = field.NewAsterisk(tableName)
	_winCodeRecordsRequire.ID = field.NewInt32(tableName, "id")
	_winCodeRecordsRequire.ReferID = field.NewInt64(tableName, "refer_id")
	_winCodeRecordsRequire.CodeRequire = field.NewFloat64(tableName, "code_require")
	_winCodeRecordsRequire.Remarks = field.NewString(tableName, "remarks")
	_winCodeRecordsRequire.CreatedAt = field.NewInt32(tableName, "created_at")
	_winCodeRecordsRequire.UpdatedAt = field.NewInt32(tableName, "updated_at")
	_winCodeRecordsRequire.OpUser = field.NewString(tableName, "op_user")
	_winCodeRecordsRequire.CodeRequireNew = field.NewFloat64(tableName, "code_require_new")

	_winCodeRecordsRequire.fillFieldMap()

	return _winCodeRecordsRequire
}

// winCodeRecordsRequire 所需打码量修改记录
type winCodeRecordsRequire struct {
	winCodeRecordsRequireDo

	ALL            field.Asterisk
	ID             field.Int32   // ID
	ReferID        field.Int64   // 关联ID
	CodeRequire    field.Float64 // 所需打码量
	Remarks        field.String  // 备注
	CreatedAt      field.Int32   // 创建日期
	UpdatedAt      field.Int32   // 修改日期
	OpUser         field.String  // 操作人
	CodeRequireNew field.Float64 // 当前所需打码量

	fieldMap map[string]field.Expr
}

func (w winCodeRecordsRequire) Table(newTableName string) *winCodeRecordsRequire {
	w.winCodeRecordsRequireDo.UseTable(newTableName)
	return w.updateTableName(newTableName)
}

func (w winCodeRecordsRequire) As(alias string) *winCodeRecordsRequire {
	w.winCodeRecordsRequireDo.DO = *(w.winCodeRecordsRequireDo.As(alias).(*gen.DO))
	return w.updateTableName(alias)
}

func (w *winCodeRecordsRequire) updateTableName(table string) *winCodeRecordsRequire {
	w.ALL = field.NewAsterisk(table)
	w.ID = field.NewInt32(table, "id")
	w.ReferID = field.NewInt64(table, "refer_id")
	w.CodeRequire = field.NewFloat64(table, "code_require")
	w.Remarks = field.NewString(table, "remarks")
	w.CreatedAt = field.NewInt32(table, "created_at")
	w.UpdatedAt = field.NewInt32(table, "updated_at")
	w.OpUser = field.NewString(table, "op_user")
	w.CodeRequireNew = field.NewFloat64(table, "code_require_new")

	w.fillFieldMap()

	return w
}

func (w *winCodeRecordsRequire) GetFieldByName(fieldName string) (field.OrderExpr, bool) {
	_f, ok := w.fieldMap[fieldName]
	if !ok || _f == nil {
		return nil, false
	}
	_oe, ok := _f.(field.OrderExpr)
	return _oe, ok
}

func (w *winCodeRecordsRequire) fillFieldMap() {
	w.fieldMap = make(map[string]field.Expr, 8)
	w.fieldMap["id"] = w.ID
	w.fieldMap["refer_id"] = w.ReferID
	w.fieldMap["code_require"] = w.CodeRequire
	w.fieldMap["remarks"] = w.Remarks
	w.fieldMap["created_at"] = w.CreatedAt
	w.fieldMap["updated_at"] = w.UpdatedAt
	w.fieldMap["op_user"] = w.OpUser
	w.fieldMap["code_require_new"] = w.CodeRequireNew
}

func (w winCodeRecordsRequire) clone(db *gorm.DB) winCodeRecordsRequire {
	w.winCodeRecordsRequireDo.ReplaceConnPool(db.Statement.ConnPool)
	return w
}

func (w winCodeRecordsRequire) replaceDB(db *gorm.DB) winCodeRecordsRequire {
	w.winCodeRecordsRequireDo.ReplaceDB(db)
	return w
}

type winCodeRecordsRequireDo struct{ gen.DO }

func (w winCodeRecordsRequireDo) Debug() *winCodeRecordsRequireDo {
	return w.withDO(w.DO.Debug())
}

func (w winCodeRecordsRequireDo) WithContext(ctx context.Context) *winCodeRecordsRequireDo {
	return w.withDO(w.DO.WithContext(ctx))
}

func (w winCodeRecordsRequireDo) ReadDB() *winCodeRecordsRequireDo {
	return w.Clauses(dbresolver.Read)
}

func (w winCodeRecordsRequireDo) WriteDB() *winCodeRecordsRequireDo {
	return w.Clauses(dbresolver.Write)
}

func (w winCodeRecordsRequireDo) Session(config *gorm.Session) *winCodeRecordsRequireDo {
	return w.withDO(w.DO.Session(config))
}

func (w winCodeRecordsRequireDo) Clauses(conds ...clause.Expression) *winCodeRecordsRequireDo {
	return w.withDO(w.DO.Clauses(conds...))
}

func (w winCodeRecordsRequireDo) Returning(value interface{}, columns ...string) *winCodeRecordsRequireDo {
	return w.withDO(w.DO.Returning(value, columns...))
}

func (w winCodeRecordsRequireDo) Not(conds ...gen.Condition) *winCodeRecordsRequireDo {
	return w.withDO(w.DO.Not(conds...))
}

func (w winCodeRecordsRequireDo) Or(conds ...gen.Condition) *winCodeRecordsRequireDo {
	return w.withDO(w.DO.Or(conds...))
}

func (w winCodeRecordsRequireDo) Select(conds ...field.Expr) *winCodeRecordsRequireDo {
	return w.withDO(w.DO.Select(conds...))
}

func (w winCodeRecordsRequireDo) Where(conds ...gen.Condition) *winCodeRecordsRequireDo {
	return w.withDO(w.DO.Where(conds...))
}

func (w winCodeRecordsRequireDo) Order(conds ...field.Expr) *winCodeRecordsRequireDo {
	return w.withDO(w.DO.Order(conds...))
}

func (w winCodeRecordsRequireDo) Distinct(cols ...field.Expr) *winCodeRecordsRequireDo {
	return w.withDO(w.DO.Distinct(cols...))
}

func (w winCodeRecordsRequireDo) Omit(cols ...field.Expr) *winCodeRecordsRequireDo {
	return w.withDO(w.DO.Omit(cols...))
}

func (w winCodeRecordsRequireDo) Join(table schema.Tabler, on ...field.Expr) *winCodeRecordsRequireDo {
	return w.withDO(w.DO.Join(table, on...))
}

func (w winCodeRecordsRequireDo) LeftJoin(table schema.Tabler, on ...field.Expr) *winCodeRecordsRequireDo {
	return w.withDO(w.DO.LeftJoin(table, on...))
}

func (w winCodeRecordsRequireDo) RightJoin(table schema.Tabler, on ...field.Expr) *winCodeRecordsRequireDo {
	return w.withDO(w.DO.RightJoin(table, on...))
}

func (w winCodeRecordsRequireDo) Group(cols ...field.Expr) *winCodeRecordsRequireDo {
	return w.withDO(w.DO.Group(cols...))
}

func (w winCodeRecordsRequireDo) Having(conds ...gen.Condition) *winCodeRecordsRequireDo {
	return w.withDO(w.DO.Having(conds...))
}

func (w winCodeRecordsRequireDo) Limit(limit int) *winCodeRecordsRequireDo {
	return w.withDO(w.DO.Limit(limit))
}

func (w winCodeRecordsRequireDo) Offset(offset int) *winCodeRecordsRequireDo {
	return w.withDO(w.DO.Offset(offset))
}

func (w winCodeRecordsRequireDo) Scopes(funcs ...func(gen.Dao) gen.Dao) *winCodeRecordsRequireDo {
	return w.withDO(w.DO.Scopes(funcs...))
}

func (w winCodeRecordsRequireDo) Unscoped() *winCodeRecordsRequireDo {
	return w.withDO(w.DO.Unscoped())
}

func (w winCodeRecordsRequireDo) Create(values ...*model.WinCodeRecordsRequire) error {
	if len(values) == 0 {
		return nil
	}
	return w.DO.Create(values)
}

func (w winCodeRecordsRequireDo) CreateInBatches(values []*model.WinCodeRecordsRequire, batchSize int) error {
	return w.DO.CreateInBatches(values, batchSize)
}

// Save : !!! underlying implementation is different with GORM
// The method is equivalent to executing the statement: db.Clauses(clause.OnConflict{UpdateAll: true}).Create(values)
func (w winCodeRecordsRequireDo) Save(values ...*model.WinCodeRecordsRequire) error {
	if len(values) == 0 {
		return nil
	}
	return w.DO.Save(values)
}

func (w winCodeRecordsRequireDo) First() (*model.WinCodeRecordsRequire, error) {
	if result, err := w.DO.First(); err != nil {
		return nil, err
	} else {
		return result.(*model.WinCodeRecordsRequire), nil
	}
}

func (w winCodeRecordsRequireDo) Take() (*model.WinCodeRecordsRequire, error) {
	if result, err := w.DO.Take(); err != nil {
		return nil, err
	} else {
		return result.(*model.WinCodeRecordsRequire), nil
	}
}

func (w winCodeRecordsRequireDo) Last() (*model.WinCodeRecordsRequire, error) {
	if result, err := w.DO.Last(); err != nil {
		return nil, err
	} else {
		return result.(*model.WinCodeRecordsRequire), nil
	}
}

func (w winCodeRecordsRequireDo) Find() ([]*model.WinCodeRecordsRequire, error) {
	result, err := w.DO.Find()
	return result.([]*model.WinCodeRecordsRequire), err
}

func (w winCodeRecordsRequireDo) FindInBatch(batchSize int, fc func(tx gen.Dao, batch int) error) (results []*model.WinCodeRecordsRequire, err error) {
	buf := make([]*model.WinCodeRecordsRequire, 0, batchSize)
	err = w.DO.FindInBatches(&buf, batchSize, func(tx gen.Dao, batch int) error {
		defer func() { results = append(results, buf...) }()
		return fc(tx, batch)
	})
	return results, err
}

func (w winCodeRecordsRequireDo) FindInBatches(result *[]*model.WinCodeRecordsRequire, batchSize int, fc func(tx gen.Dao, batch int) error) error {
	return w.DO.FindInBatches(result, batchSize, fc)
}

func (w winCodeRecordsRequireDo) Attrs(attrs ...field.AssignExpr) *winCodeRecordsRequireDo {
	return w.withDO(w.DO.Attrs(attrs...))
}

func (w winCodeRecordsRequireDo) Assign(attrs ...field.AssignExpr) *winCodeRecordsRequireDo {
	return w.withDO(w.DO.Assign(attrs...))
}

func (w winCodeRecordsRequireDo) Joins(fields ...field.RelationField) *winCodeRecordsRequireDo {
	for _, _f := range fields {
		w = *w.withDO(w.DO.Joins(_f))
	}
	return &w
}

func (w winCodeRecordsRequireDo) Preload(fields ...field.RelationField) *winCodeRecordsRequireDo {
	for _, _f := range fields {
		w = *w.withDO(w.DO.Preload(_f))
	}
	return &w
}

func (w winCodeRecordsRequireDo) FirstOrInit() (*model.WinCodeRecordsRequire, error) {
	if result, err := w.DO.FirstOrInit(); err != nil {
		return nil, err
	} else {
		return result.(*model.WinCodeRecordsRequire), nil
	}
}

func (w winCodeRecordsRequireDo) FirstOrCreate() (*model.WinCodeRecordsRequire, error) {
	if result, err := w.DO.FirstOrCreate(); err != nil {
		return nil, err
	} else {
		return result.(*model.WinCodeRecordsRequire), nil
	}
}

func (w winCodeRecordsRequireDo) FindByPage(offset int, limit int) (result []*model.WinCodeRecordsRequire, count int64, err error) {
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

func (w winCodeRecordsRequireDo) ScanByPage(result interface{}, offset int, limit int) (count int64, err error) {
	count, err = w.Count()
	if err != nil {
		return
	}

	err = w.Offset(offset).Limit(limit).Scan(result)
	return
}

func (w winCodeRecordsRequireDo) Scan(result interface{}) (err error) {
	return w.DO.Scan(result)
}

func (w winCodeRecordsRequireDo) Delete(models ...*model.WinCodeRecordsRequire) (result gen.ResultInfo, err error) {
	return w.DO.Delete(models)
}

func (w *winCodeRecordsRequireDo) withDO(do gen.Dao) *winCodeRecordsRequireDo {
	w.DO = *do.(*gen.DO)
	return w
}
