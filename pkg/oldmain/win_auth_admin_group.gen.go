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

func newWinAuthAdminGroup(db *gorm.DB, opts ...gen.DOOption) winAuthAdminGroup {
	_winAuthAdminGroup := winAuthAdminGroup{}

	_winAuthAdminGroup.winAuthAdminGroupDo.UseDB(db, opts...)
	_winAuthAdminGroup.winAuthAdminGroupDo.UseModel(&model.WinAuthAdminGroup{})

	tableName := _winAuthAdminGroup.winAuthAdminGroupDo.TableName()
	_winAuthAdminGroup.ALL = field.NewAsterisk(tableName)
	_winAuthAdminGroup.ID = field.NewInt32(tableName, "id")
	_winAuthAdminGroup.Title = field.NewString(tableName, "title")
	_winAuthAdminGroup.Parent = field.NewInt32(tableName, "parent")
	_winAuthAdminGroup.Pid = field.NewInt32(tableName, "pid")
	_winAuthAdminGroup.Rules = field.NewString(tableName, "rules")
	_winAuthAdminGroup.OperateID = field.NewInt32(tableName, "operate_id")
	_winAuthAdminGroup.Status = field.NewInt32(tableName, "status")
	_winAuthAdminGroup.CreatedAt = field.NewInt32(tableName, "created_at")
	_winAuthAdminGroup.UpdatedAt = field.NewInt32(tableName, "updated_at")

	_winAuthAdminGroup.fillFieldMap()

	return _winAuthAdminGroup
}

// winAuthAdminGroup 用户组表
type winAuthAdminGroup struct {
	winAuthAdminGroupDo

	ALL       field.Asterisk
	ID        field.Int32
	Title     field.String // 用户组名
	Parent    field.Int32  // 创建人
	Pid       field.Int32  // 父ID
	Rules     field.String // 菜单ID集合
	OperateID field.Int32  // 用户组ID
	Status    field.Int32  // 状态: 1-启用 0-冻结
	CreatedAt field.Int32
	UpdatedAt field.Int32

	fieldMap map[string]field.Expr
}

func (w winAuthAdminGroup) Table(newTableName string) *winAuthAdminGroup {
	w.winAuthAdminGroupDo.UseTable(newTableName)
	return w.updateTableName(newTableName)
}

func (w winAuthAdminGroup) As(alias string) *winAuthAdminGroup {
	w.winAuthAdminGroupDo.DO = *(w.winAuthAdminGroupDo.As(alias).(*gen.DO))
	return w.updateTableName(alias)
}

func (w *winAuthAdminGroup) updateTableName(table string) *winAuthAdminGroup {
	w.ALL = field.NewAsterisk(table)
	w.ID = field.NewInt32(table, "id")
	w.Title = field.NewString(table, "title")
	w.Parent = field.NewInt32(table, "parent")
	w.Pid = field.NewInt32(table, "pid")
	w.Rules = field.NewString(table, "rules")
	w.OperateID = field.NewInt32(table, "operate_id")
	w.Status = field.NewInt32(table, "status")
	w.CreatedAt = field.NewInt32(table, "created_at")
	w.UpdatedAt = field.NewInt32(table, "updated_at")

	w.fillFieldMap()

	return w
}

func (w *winAuthAdminGroup) GetFieldByName(fieldName string) (field.OrderExpr, bool) {
	_f, ok := w.fieldMap[fieldName]
	if !ok || _f == nil {
		return nil, false
	}
	_oe, ok := _f.(field.OrderExpr)
	return _oe, ok
}

func (w *winAuthAdminGroup) fillFieldMap() {
	w.fieldMap = make(map[string]field.Expr, 9)
	w.fieldMap["id"] = w.ID
	w.fieldMap["title"] = w.Title
	w.fieldMap["parent"] = w.Parent
	w.fieldMap["pid"] = w.Pid
	w.fieldMap["rules"] = w.Rules
	w.fieldMap["operate_id"] = w.OperateID
	w.fieldMap["status"] = w.Status
	w.fieldMap["created_at"] = w.CreatedAt
	w.fieldMap["updated_at"] = w.UpdatedAt
}

func (w winAuthAdminGroup) clone(db *gorm.DB) winAuthAdminGroup {
	w.winAuthAdminGroupDo.ReplaceConnPool(db.Statement.ConnPool)
	return w
}

func (w winAuthAdminGroup) replaceDB(db *gorm.DB) winAuthAdminGroup {
	w.winAuthAdminGroupDo.ReplaceDB(db)
	return w
}

type winAuthAdminGroupDo struct{ gen.DO }

func (w winAuthAdminGroupDo) Debug() *winAuthAdminGroupDo {
	return w.withDO(w.DO.Debug())
}

func (w winAuthAdminGroupDo) WithContext(ctx context.Context) *winAuthAdminGroupDo {
	return w.withDO(w.DO.WithContext(ctx))
}

func (w winAuthAdminGroupDo) ReadDB() *winAuthAdminGroupDo {
	return w.Clauses(dbresolver.Read)
}

func (w winAuthAdminGroupDo) WriteDB() *winAuthAdminGroupDo {
	return w.Clauses(dbresolver.Write)
}

func (w winAuthAdminGroupDo) Session(config *gorm.Session) *winAuthAdminGroupDo {
	return w.withDO(w.DO.Session(config))
}

func (w winAuthAdminGroupDo) Clauses(conds ...clause.Expression) *winAuthAdminGroupDo {
	return w.withDO(w.DO.Clauses(conds...))
}

func (w winAuthAdminGroupDo) Returning(value interface{}, columns ...string) *winAuthAdminGroupDo {
	return w.withDO(w.DO.Returning(value, columns...))
}

func (w winAuthAdminGroupDo) Not(conds ...gen.Condition) *winAuthAdminGroupDo {
	return w.withDO(w.DO.Not(conds...))
}

func (w winAuthAdminGroupDo) Or(conds ...gen.Condition) *winAuthAdminGroupDo {
	return w.withDO(w.DO.Or(conds...))
}

func (w winAuthAdminGroupDo) Select(conds ...field.Expr) *winAuthAdminGroupDo {
	return w.withDO(w.DO.Select(conds...))
}

func (w winAuthAdminGroupDo) Where(conds ...gen.Condition) *winAuthAdminGroupDo {
	return w.withDO(w.DO.Where(conds...))
}

func (w winAuthAdminGroupDo) Order(conds ...field.Expr) *winAuthAdminGroupDo {
	return w.withDO(w.DO.Order(conds...))
}

func (w winAuthAdminGroupDo) Distinct(cols ...field.Expr) *winAuthAdminGroupDo {
	return w.withDO(w.DO.Distinct(cols...))
}

func (w winAuthAdminGroupDo) Omit(cols ...field.Expr) *winAuthAdminGroupDo {
	return w.withDO(w.DO.Omit(cols...))
}

func (w winAuthAdminGroupDo) Join(table schema.Tabler, on ...field.Expr) *winAuthAdminGroupDo {
	return w.withDO(w.DO.Join(table, on...))
}

func (w winAuthAdminGroupDo) LeftJoin(table schema.Tabler, on ...field.Expr) *winAuthAdminGroupDo {
	return w.withDO(w.DO.LeftJoin(table, on...))
}

func (w winAuthAdminGroupDo) RightJoin(table schema.Tabler, on ...field.Expr) *winAuthAdminGroupDo {
	return w.withDO(w.DO.RightJoin(table, on...))
}

func (w winAuthAdminGroupDo) Group(cols ...field.Expr) *winAuthAdminGroupDo {
	return w.withDO(w.DO.Group(cols...))
}

func (w winAuthAdminGroupDo) Having(conds ...gen.Condition) *winAuthAdminGroupDo {
	return w.withDO(w.DO.Having(conds...))
}

func (w winAuthAdminGroupDo) Limit(limit int) *winAuthAdminGroupDo {
	return w.withDO(w.DO.Limit(limit))
}

func (w winAuthAdminGroupDo) Offset(offset int) *winAuthAdminGroupDo {
	return w.withDO(w.DO.Offset(offset))
}

func (w winAuthAdminGroupDo) Scopes(funcs ...func(gen.Dao) gen.Dao) *winAuthAdminGroupDo {
	return w.withDO(w.DO.Scopes(funcs...))
}

func (w winAuthAdminGroupDo) Unscoped() *winAuthAdminGroupDo {
	return w.withDO(w.DO.Unscoped())
}

func (w winAuthAdminGroupDo) Create(values ...*model.WinAuthAdminGroup) error {
	if len(values) == 0 {
		return nil
	}
	return w.DO.Create(values)
}

func (w winAuthAdminGroupDo) CreateInBatches(values []*model.WinAuthAdminGroup, batchSize int) error {
	return w.DO.CreateInBatches(values, batchSize)
}

// Save : !!! underlying implementation is different with GORM
// The method is equivalent to executing the statement: db.Clauses(clause.OnConflict{UpdateAll: true}).Create(values)
func (w winAuthAdminGroupDo) Save(values ...*model.WinAuthAdminGroup) error {
	if len(values) == 0 {
		return nil
	}
	return w.DO.Save(values)
}

func (w winAuthAdminGroupDo) First() (*model.WinAuthAdminGroup, error) {
	if result, err := w.DO.First(); err != nil {
		return nil, err
	} else {
		return result.(*model.WinAuthAdminGroup), nil
	}
}

func (w winAuthAdminGroupDo) Take() (*model.WinAuthAdminGroup, error) {
	if result, err := w.DO.Take(); err != nil {
		return nil, err
	} else {
		return result.(*model.WinAuthAdminGroup), nil
	}
}

func (w winAuthAdminGroupDo) Last() (*model.WinAuthAdminGroup, error) {
	if result, err := w.DO.Last(); err != nil {
		return nil, err
	} else {
		return result.(*model.WinAuthAdminGroup), nil
	}
}

func (w winAuthAdminGroupDo) Find() ([]*model.WinAuthAdminGroup, error) {
	result, err := w.DO.Find()
	return result.([]*model.WinAuthAdminGroup), err
}

func (w winAuthAdminGroupDo) FindInBatch(batchSize int, fc func(tx gen.Dao, batch int) error) (results []*model.WinAuthAdminGroup, err error) {
	buf := make([]*model.WinAuthAdminGroup, 0, batchSize)
	err = w.DO.FindInBatches(&buf, batchSize, func(tx gen.Dao, batch int) error {
		defer func() { results = append(results, buf...) }()
		return fc(tx, batch)
	})
	return results, err
}

func (w winAuthAdminGroupDo) FindInBatches(result *[]*model.WinAuthAdminGroup, batchSize int, fc func(tx gen.Dao, batch int) error) error {
	return w.DO.FindInBatches(result, batchSize, fc)
}

func (w winAuthAdminGroupDo) Attrs(attrs ...field.AssignExpr) *winAuthAdminGroupDo {
	return w.withDO(w.DO.Attrs(attrs...))
}

func (w winAuthAdminGroupDo) Assign(attrs ...field.AssignExpr) *winAuthAdminGroupDo {
	return w.withDO(w.DO.Assign(attrs...))
}

func (w winAuthAdminGroupDo) Joins(fields ...field.RelationField) *winAuthAdminGroupDo {
	for _, _f := range fields {
		w = *w.withDO(w.DO.Joins(_f))
	}
	return &w
}

func (w winAuthAdminGroupDo) Preload(fields ...field.RelationField) *winAuthAdminGroupDo {
	for _, _f := range fields {
		w = *w.withDO(w.DO.Preload(_f))
	}
	return &w
}

func (w winAuthAdminGroupDo) FirstOrInit() (*model.WinAuthAdminGroup, error) {
	if result, err := w.DO.FirstOrInit(); err != nil {
		return nil, err
	} else {
		return result.(*model.WinAuthAdminGroup), nil
	}
}

func (w winAuthAdminGroupDo) FirstOrCreate() (*model.WinAuthAdminGroup, error) {
	if result, err := w.DO.FirstOrCreate(); err != nil {
		return nil, err
	} else {
		return result.(*model.WinAuthAdminGroup), nil
	}
}

func (w winAuthAdminGroupDo) FindByPage(offset int, limit int) (result []*model.WinAuthAdminGroup, count int64, err error) {
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

func (w winAuthAdminGroupDo) ScanByPage(result interface{}, offset int, limit int) (count int64, err error) {
	count, err = w.Count()
	if err != nil {
		return
	}

	err = w.Offset(offset).Limit(limit).Scan(result)
	return
}

func (w winAuthAdminGroupDo) Scan(result interface{}) (err error) {
	return w.DO.Scan(result)
}

func (w winAuthAdminGroupDo) Delete(models ...*model.WinAuthAdminGroup) (result gen.ResultInfo, err error) {
	return w.DO.Delete(models)
}

func (w *winAuthAdminGroupDo) withDO(do gen.Dao) *winAuthAdminGroupDo {
	w.DO = *do.(*gen.DO)
	return w
}
