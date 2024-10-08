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

func newWinAgentLink(db *gorm.DB, opts ...gen.DOOption) winAgentLink {
	_winAgentLink := winAgentLink{}

	_winAgentLink.winAgentLinkDo.UseDB(db, opts...)
	_winAgentLink.winAgentLinkDo.UseModel(&model.WinAgentLink{})

	tableName := _winAgentLink.winAgentLinkDo.TableName()
	_winAgentLink.ALL = field.NewAsterisk(tableName)
	_winAgentLink.ID = field.NewInt32(tableName, "id")
	_winAgentLink.Link = field.NewString(tableName, "link")
	_winAgentLink.UID = field.NewInt32(tableName, "uid")
	_winAgentLink.Status = field.NewInt32(tableName, "status")
	_winAgentLink.CreatedAt = field.NewInt32(tableName, "created_at")
	_winAgentLink.UpdatedAt = field.NewInt32(tableName, "updated_at")

	_winAgentLink.fillFieldMap()

	return _winAgentLink
}

// winAgentLink 代理专属域名
type winAgentLink struct {
	winAgentLinkDo

	ALL       field.Asterisk
	ID        field.Int32
	Link      field.String // 推广域名
	UID       field.Int32  // 代理ID
	Status    field.Int32  // 状态:1-启用 0-停用
	CreatedAt field.Int32
	UpdatedAt field.Int32

	fieldMap map[string]field.Expr
}

func (w winAgentLink) Table(newTableName string) *winAgentLink {
	w.winAgentLinkDo.UseTable(newTableName)
	return w.updateTableName(newTableName)
}

func (w winAgentLink) As(alias string) *winAgentLink {
	w.winAgentLinkDo.DO = *(w.winAgentLinkDo.As(alias).(*gen.DO))
	return w.updateTableName(alias)
}

func (w *winAgentLink) updateTableName(table string) *winAgentLink {
	w.ALL = field.NewAsterisk(table)
	w.ID = field.NewInt32(table, "id")
	w.Link = field.NewString(table, "link")
	w.UID = field.NewInt32(table, "uid")
	w.Status = field.NewInt32(table, "status")
	w.CreatedAt = field.NewInt32(table, "created_at")
	w.UpdatedAt = field.NewInt32(table, "updated_at")

	w.fillFieldMap()

	return w
}

func (w *winAgentLink) GetFieldByName(fieldName string) (field.OrderExpr, bool) {
	_f, ok := w.fieldMap[fieldName]
	if !ok || _f == nil {
		return nil, false
	}
	_oe, ok := _f.(field.OrderExpr)
	return _oe, ok
}

func (w *winAgentLink) fillFieldMap() {
	w.fieldMap = make(map[string]field.Expr, 6)
	w.fieldMap["id"] = w.ID
	w.fieldMap["link"] = w.Link
	w.fieldMap["uid"] = w.UID
	w.fieldMap["status"] = w.Status
	w.fieldMap["created_at"] = w.CreatedAt
	w.fieldMap["updated_at"] = w.UpdatedAt
}

func (w winAgentLink) clone(db *gorm.DB) winAgentLink {
	w.winAgentLinkDo.ReplaceConnPool(db.Statement.ConnPool)
	return w
}

func (w winAgentLink) replaceDB(db *gorm.DB) winAgentLink {
	w.winAgentLinkDo.ReplaceDB(db)
	return w
}

type winAgentLinkDo struct{ gen.DO }

func (w winAgentLinkDo) Debug() *winAgentLinkDo {
	return w.withDO(w.DO.Debug())
}

func (w winAgentLinkDo) WithContext(ctx context.Context) *winAgentLinkDo {
	return w.withDO(w.DO.WithContext(ctx))
}

func (w winAgentLinkDo) ReadDB() *winAgentLinkDo {
	return w.Clauses(dbresolver.Read)
}

func (w winAgentLinkDo) WriteDB() *winAgentLinkDo {
	return w.Clauses(dbresolver.Write)
}

func (w winAgentLinkDo) Session(config *gorm.Session) *winAgentLinkDo {
	return w.withDO(w.DO.Session(config))
}

func (w winAgentLinkDo) Clauses(conds ...clause.Expression) *winAgentLinkDo {
	return w.withDO(w.DO.Clauses(conds...))
}

func (w winAgentLinkDo) Returning(value interface{}, columns ...string) *winAgentLinkDo {
	return w.withDO(w.DO.Returning(value, columns...))
}

func (w winAgentLinkDo) Not(conds ...gen.Condition) *winAgentLinkDo {
	return w.withDO(w.DO.Not(conds...))
}

func (w winAgentLinkDo) Or(conds ...gen.Condition) *winAgentLinkDo {
	return w.withDO(w.DO.Or(conds...))
}

func (w winAgentLinkDo) Select(conds ...field.Expr) *winAgentLinkDo {
	return w.withDO(w.DO.Select(conds...))
}

func (w winAgentLinkDo) Where(conds ...gen.Condition) *winAgentLinkDo {
	return w.withDO(w.DO.Where(conds...))
}

func (w winAgentLinkDo) Order(conds ...field.Expr) *winAgentLinkDo {
	return w.withDO(w.DO.Order(conds...))
}

func (w winAgentLinkDo) Distinct(cols ...field.Expr) *winAgentLinkDo {
	return w.withDO(w.DO.Distinct(cols...))
}

func (w winAgentLinkDo) Omit(cols ...field.Expr) *winAgentLinkDo {
	return w.withDO(w.DO.Omit(cols...))
}

func (w winAgentLinkDo) Join(table schema.Tabler, on ...field.Expr) *winAgentLinkDo {
	return w.withDO(w.DO.Join(table, on...))
}

func (w winAgentLinkDo) LeftJoin(table schema.Tabler, on ...field.Expr) *winAgentLinkDo {
	return w.withDO(w.DO.LeftJoin(table, on...))
}

func (w winAgentLinkDo) RightJoin(table schema.Tabler, on ...field.Expr) *winAgentLinkDo {
	return w.withDO(w.DO.RightJoin(table, on...))
}

func (w winAgentLinkDo) Group(cols ...field.Expr) *winAgentLinkDo {
	return w.withDO(w.DO.Group(cols...))
}

func (w winAgentLinkDo) Having(conds ...gen.Condition) *winAgentLinkDo {
	return w.withDO(w.DO.Having(conds...))
}

func (w winAgentLinkDo) Limit(limit int) *winAgentLinkDo {
	return w.withDO(w.DO.Limit(limit))
}

func (w winAgentLinkDo) Offset(offset int) *winAgentLinkDo {
	return w.withDO(w.DO.Offset(offset))
}

func (w winAgentLinkDo) Scopes(funcs ...func(gen.Dao) gen.Dao) *winAgentLinkDo {
	return w.withDO(w.DO.Scopes(funcs...))
}

func (w winAgentLinkDo) Unscoped() *winAgentLinkDo {
	return w.withDO(w.DO.Unscoped())
}

func (w winAgentLinkDo) Create(values ...*model.WinAgentLink) error {
	if len(values) == 0 {
		return nil
	}
	return w.DO.Create(values)
}

func (w winAgentLinkDo) CreateInBatches(values []*model.WinAgentLink, batchSize int) error {
	return w.DO.CreateInBatches(values, batchSize)
}

// Save : !!! underlying implementation is different with GORM
// The method is equivalent to executing the statement: db.Clauses(clause.OnConflict{UpdateAll: true}).Create(values)
func (w winAgentLinkDo) Save(values ...*model.WinAgentLink) error {
	if len(values) == 0 {
		return nil
	}
	return w.DO.Save(values)
}

func (w winAgentLinkDo) First() (*model.WinAgentLink, error) {
	if result, err := w.DO.First(); err != nil {
		return nil, err
	} else {
		return result.(*model.WinAgentLink), nil
	}
}

func (w winAgentLinkDo) Take() (*model.WinAgentLink, error) {
	if result, err := w.DO.Take(); err != nil {
		return nil, err
	} else {
		return result.(*model.WinAgentLink), nil
	}
}

func (w winAgentLinkDo) Last() (*model.WinAgentLink, error) {
	if result, err := w.DO.Last(); err != nil {
		return nil, err
	} else {
		return result.(*model.WinAgentLink), nil
	}
}

func (w winAgentLinkDo) Find() ([]*model.WinAgentLink, error) {
	result, err := w.DO.Find()
	return result.([]*model.WinAgentLink), err
}

func (w winAgentLinkDo) FindInBatch(batchSize int, fc func(tx gen.Dao, batch int) error) (results []*model.WinAgentLink, err error) {
	buf := make([]*model.WinAgentLink, 0, batchSize)
	err = w.DO.FindInBatches(&buf, batchSize, func(tx gen.Dao, batch int) error {
		defer func() { results = append(results, buf...) }()
		return fc(tx, batch)
	})
	return results, err
}

func (w winAgentLinkDo) FindInBatches(result *[]*model.WinAgentLink, batchSize int, fc func(tx gen.Dao, batch int) error) error {
	return w.DO.FindInBatches(result, batchSize, fc)
}

func (w winAgentLinkDo) Attrs(attrs ...field.AssignExpr) *winAgentLinkDo {
	return w.withDO(w.DO.Attrs(attrs...))
}

func (w winAgentLinkDo) Assign(attrs ...field.AssignExpr) *winAgentLinkDo {
	return w.withDO(w.DO.Assign(attrs...))
}

func (w winAgentLinkDo) Joins(fields ...field.RelationField) *winAgentLinkDo {
	for _, _f := range fields {
		w = *w.withDO(w.DO.Joins(_f))
	}
	return &w
}

func (w winAgentLinkDo) Preload(fields ...field.RelationField) *winAgentLinkDo {
	for _, _f := range fields {
		w = *w.withDO(w.DO.Preload(_f))
	}
	return &w
}

func (w winAgentLinkDo) FirstOrInit() (*model.WinAgentLink, error) {
	if result, err := w.DO.FirstOrInit(); err != nil {
		return nil, err
	} else {
		return result.(*model.WinAgentLink), nil
	}
}

func (w winAgentLinkDo) FirstOrCreate() (*model.WinAgentLink, error) {
	if result, err := w.DO.FirstOrCreate(); err != nil {
		return nil, err
	} else {
		return result.(*model.WinAgentLink), nil
	}
}

func (w winAgentLinkDo) FindByPage(offset int, limit int) (result []*model.WinAgentLink, count int64, err error) {
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

func (w winAgentLinkDo) ScanByPage(result interface{}, offset int, limit int) (count int64, err error) {
	count, err = w.Count()
	if err != nil {
		return
	}

	err = w.Offset(offset).Limit(limit).Scan(result)
	return
}

func (w winAgentLinkDo) Scan(result interface{}) (err error) {
	return w.DO.Scan(result)
}

func (w winAgentLinkDo) Delete(models ...*model.WinAgentLink) (result gen.ResultInfo, err error) {
	return w.DO.Delete(models)
}

func (w *winAgentLinkDo) withDO(do gen.Dao) *winAgentLinkDo {
	w.DO = *do.(*gen.DO)
	return w
}
