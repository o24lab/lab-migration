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

func newWinAuthGroupAccess(db *gorm.DB, opts ...gen.DOOption) winAuthGroupAccess {
	_winAuthGroupAccess := winAuthGroupAccess{}

	_winAuthGroupAccess.winAuthGroupAccessDo.UseDB(db, opts...)
	_winAuthGroupAccess.winAuthGroupAccessDo.UseModel(&model.WinAuthGroupAccess{})

	tableName := _winAuthGroupAccess.winAuthGroupAccessDo.TableName()
	_winAuthGroupAccess.ALL = field.NewAsterisk(tableName)
	_winAuthGroupAccess.ID = field.NewInt32(tableName, "id")
	_winAuthGroupAccess.UID = field.NewInt32(tableName, "uid")
	_winAuthGroupAccess.GroupID = field.NewInt32(tableName, "group_id")
	_winAuthGroupAccess.MenuID = field.NewString(tableName, "menu_id")
	_winAuthGroupAccess.CreatedAt = field.NewInt32(tableName, "created_at")
	_winAuthGroupAccess.UpdatedAt = field.NewInt32(tableName, "updated_at")

	_winAuthGroupAccess.fillFieldMap()

	return _winAuthGroupAccess
}

// winAuthGroupAccess 角色表
type winAuthGroupAccess struct {
	winAuthGroupAccessDo

	ALL       field.Asterisk
	ID        field.Int32
	UID       field.Int32  // UID
	GroupID   field.Int32  // 角色ID
	MenuID    field.String // 补充菜单ID: +新增 -去除
	CreatedAt field.Int32
	UpdatedAt field.Int32

	fieldMap map[string]field.Expr
}

func (w winAuthGroupAccess) Table(newTableName string) *winAuthGroupAccess {
	w.winAuthGroupAccessDo.UseTable(newTableName)
	return w.updateTableName(newTableName)
}

func (w winAuthGroupAccess) As(alias string) *winAuthGroupAccess {
	w.winAuthGroupAccessDo.DO = *(w.winAuthGroupAccessDo.As(alias).(*gen.DO))
	return w.updateTableName(alias)
}

func (w *winAuthGroupAccess) updateTableName(table string) *winAuthGroupAccess {
	w.ALL = field.NewAsterisk(table)
	w.ID = field.NewInt32(table, "id")
	w.UID = field.NewInt32(table, "uid")
	w.GroupID = field.NewInt32(table, "group_id")
	w.MenuID = field.NewString(table, "menu_id")
	w.CreatedAt = field.NewInt32(table, "created_at")
	w.UpdatedAt = field.NewInt32(table, "updated_at")

	w.fillFieldMap()

	return w
}

func (w *winAuthGroupAccess) GetFieldByName(fieldName string) (field.OrderExpr, bool) {
	_f, ok := w.fieldMap[fieldName]
	if !ok || _f == nil {
		return nil, false
	}
	_oe, ok := _f.(field.OrderExpr)
	return _oe, ok
}

func (w *winAuthGroupAccess) fillFieldMap() {
	w.fieldMap = make(map[string]field.Expr, 6)
	w.fieldMap["id"] = w.ID
	w.fieldMap["uid"] = w.UID
	w.fieldMap["group_id"] = w.GroupID
	w.fieldMap["menu_id"] = w.MenuID
	w.fieldMap["created_at"] = w.CreatedAt
	w.fieldMap["updated_at"] = w.UpdatedAt
}

func (w winAuthGroupAccess) clone(db *gorm.DB) winAuthGroupAccess {
	w.winAuthGroupAccessDo.ReplaceConnPool(db.Statement.ConnPool)
	return w
}

func (w winAuthGroupAccess) replaceDB(db *gorm.DB) winAuthGroupAccess {
	w.winAuthGroupAccessDo.ReplaceDB(db)
	return w
}

type winAuthGroupAccessDo struct{ gen.DO }

func (w winAuthGroupAccessDo) Debug() *winAuthGroupAccessDo {
	return w.withDO(w.DO.Debug())
}

func (w winAuthGroupAccessDo) WithContext(ctx context.Context) *winAuthGroupAccessDo {
	return w.withDO(w.DO.WithContext(ctx))
}

func (w winAuthGroupAccessDo) ReadDB() *winAuthGroupAccessDo {
	return w.Clauses(dbresolver.Read)
}

func (w winAuthGroupAccessDo) WriteDB() *winAuthGroupAccessDo {
	return w.Clauses(dbresolver.Write)
}

func (w winAuthGroupAccessDo) Session(config *gorm.Session) *winAuthGroupAccessDo {
	return w.withDO(w.DO.Session(config))
}

func (w winAuthGroupAccessDo) Clauses(conds ...clause.Expression) *winAuthGroupAccessDo {
	return w.withDO(w.DO.Clauses(conds...))
}

func (w winAuthGroupAccessDo) Returning(value interface{}, columns ...string) *winAuthGroupAccessDo {
	return w.withDO(w.DO.Returning(value, columns...))
}

func (w winAuthGroupAccessDo) Not(conds ...gen.Condition) *winAuthGroupAccessDo {
	return w.withDO(w.DO.Not(conds...))
}

func (w winAuthGroupAccessDo) Or(conds ...gen.Condition) *winAuthGroupAccessDo {
	return w.withDO(w.DO.Or(conds...))
}

func (w winAuthGroupAccessDo) Select(conds ...field.Expr) *winAuthGroupAccessDo {
	return w.withDO(w.DO.Select(conds...))
}

func (w winAuthGroupAccessDo) Where(conds ...gen.Condition) *winAuthGroupAccessDo {
	return w.withDO(w.DO.Where(conds...))
}

func (w winAuthGroupAccessDo) Order(conds ...field.Expr) *winAuthGroupAccessDo {
	return w.withDO(w.DO.Order(conds...))
}

func (w winAuthGroupAccessDo) Distinct(cols ...field.Expr) *winAuthGroupAccessDo {
	return w.withDO(w.DO.Distinct(cols...))
}

func (w winAuthGroupAccessDo) Omit(cols ...field.Expr) *winAuthGroupAccessDo {
	return w.withDO(w.DO.Omit(cols...))
}

func (w winAuthGroupAccessDo) Join(table schema.Tabler, on ...field.Expr) *winAuthGroupAccessDo {
	return w.withDO(w.DO.Join(table, on...))
}

func (w winAuthGroupAccessDo) LeftJoin(table schema.Tabler, on ...field.Expr) *winAuthGroupAccessDo {
	return w.withDO(w.DO.LeftJoin(table, on...))
}

func (w winAuthGroupAccessDo) RightJoin(table schema.Tabler, on ...field.Expr) *winAuthGroupAccessDo {
	return w.withDO(w.DO.RightJoin(table, on...))
}

func (w winAuthGroupAccessDo) Group(cols ...field.Expr) *winAuthGroupAccessDo {
	return w.withDO(w.DO.Group(cols...))
}

func (w winAuthGroupAccessDo) Having(conds ...gen.Condition) *winAuthGroupAccessDo {
	return w.withDO(w.DO.Having(conds...))
}

func (w winAuthGroupAccessDo) Limit(limit int) *winAuthGroupAccessDo {
	return w.withDO(w.DO.Limit(limit))
}

func (w winAuthGroupAccessDo) Offset(offset int) *winAuthGroupAccessDo {
	return w.withDO(w.DO.Offset(offset))
}

func (w winAuthGroupAccessDo) Scopes(funcs ...func(gen.Dao) gen.Dao) *winAuthGroupAccessDo {
	return w.withDO(w.DO.Scopes(funcs...))
}

func (w winAuthGroupAccessDo) Unscoped() *winAuthGroupAccessDo {
	return w.withDO(w.DO.Unscoped())
}

func (w winAuthGroupAccessDo) Create(values ...*model.WinAuthGroupAccess) error {
	if len(values) == 0 {
		return nil
	}
	return w.DO.Create(values)
}

func (w winAuthGroupAccessDo) CreateInBatches(values []*model.WinAuthGroupAccess, batchSize int) error {
	return w.DO.CreateInBatches(values, batchSize)
}

// Save : !!! underlying implementation is different with GORM
// The method is equivalent to executing the statement: db.Clauses(clause.OnConflict{UpdateAll: true}).Create(values)
func (w winAuthGroupAccessDo) Save(values ...*model.WinAuthGroupAccess) error {
	if len(values) == 0 {
		return nil
	}
	return w.DO.Save(values)
}

func (w winAuthGroupAccessDo) First() (*model.WinAuthGroupAccess, error) {
	if result, err := w.DO.First(); err != nil {
		return nil, err
	} else {
		return result.(*model.WinAuthGroupAccess), nil
	}
}

func (w winAuthGroupAccessDo) Take() (*model.WinAuthGroupAccess, error) {
	if result, err := w.DO.Take(); err != nil {
		return nil, err
	} else {
		return result.(*model.WinAuthGroupAccess), nil
	}
}

func (w winAuthGroupAccessDo) Last() (*model.WinAuthGroupAccess, error) {
	if result, err := w.DO.Last(); err != nil {
		return nil, err
	} else {
		return result.(*model.WinAuthGroupAccess), nil
	}
}

func (w winAuthGroupAccessDo) Find() ([]*model.WinAuthGroupAccess, error) {
	result, err := w.DO.Find()
	return result.([]*model.WinAuthGroupAccess), err
}

func (w winAuthGroupAccessDo) FindInBatch(batchSize int, fc func(tx gen.Dao, batch int) error) (results []*model.WinAuthGroupAccess, err error) {
	buf := make([]*model.WinAuthGroupAccess, 0, batchSize)
	err = w.DO.FindInBatches(&buf, batchSize, func(tx gen.Dao, batch int) error {
		defer func() { results = append(results, buf...) }()
		return fc(tx, batch)
	})
	return results, err
}

func (w winAuthGroupAccessDo) FindInBatches(result *[]*model.WinAuthGroupAccess, batchSize int, fc func(tx gen.Dao, batch int) error) error {
	return w.DO.FindInBatches(result, batchSize, fc)
}

func (w winAuthGroupAccessDo) Attrs(attrs ...field.AssignExpr) *winAuthGroupAccessDo {
	return w.withDO(w.DO.Attrs(attrs...))
}

func (w winAuthGroupAccessDo) Assign(attrs ...field.AssignExpr) *winAuthGroupAccessDo {
	return w.withDO(w.DO.Assign(attrs...))
}

func (w winAuthGroupAccessDo) Joins(fields ...field.RelationField) *winAuthGroupAccessDo {
	for _, _f := range fields {
		w = *w.withDO(w.DO.Joins(_f))
	}
	return &w
}

func (w winAuthGroupAccessDo) Preload(fields ...field.RelationField) *winAuthGroupAccessDo {
	for _, _f := range fields {
		w = *w.withDO(w.DO.Preload(_f))
	}
	return &w
}

func (w winAuthGroupAccessDo) FirstOrInit() (*model.WinAuthGroupAccess, error) {
	if result, err := w.DO.FirstOrInit(); err != nil {
		return nil, err
	} else {
		return result.(*model.WinAuthGroupAccess), nil
	}
}

func (w winAuthGroupAccessDo) FirstOrCreate() (*model.WinAuthGroupAccess, error) {
	if result, err := w.DO.FirstOrCreate(); err != nil {
		return nil, err
	} else {
		return result.(*model.WinAuthGroupAccess), nil
	}
}

func (w winAuthGroupAccessDo) FindByPage(offset int, limit int) (result []*model.WinAuthGroupAccess, count int64, err error) {
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

func (w winAuthGroupAccessDo) ScanByPage(result interface{}, offset int, limit int) (count int64, err error) {
	count, err = w.Count()
	if err != nil {
		return
	}

	err = w.Offset(offset).Limit(limit).Scan(result)
	return
}

func (w winAuthGroupAccessDo) Scan(result interface{}) (err error) {
	return w.DO.Scan(result)
}

func (w winAuthGroupAccessDo) Delete(models ...*model.WinAuthGroupAccess) (result gen.ResultInfo, err error) {
	return w.DO.Delete(models)
}

func (w *winAuthGroupAccessDo) withDO(do gen.Dao) *winAuthGroupAccessDo {
	w.DO = *do.(*gen.DO)
	return w
}
