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

func newActivityBlackInfo(db *gorm.DB, opts ...gen.DOOption) activityBlackInfo {
	_activityBlackInfo := activityBlackInfo{}

	_activityBlackInfo.activityBlackInfoDo.UseDB(db, opts...)
	_activityBlackInfo.activityBlackInfoDo.UseModel(&model.ActivityBlackInfo{})

	tableName := _activityBlackInfo.activityBlackInfoDo.TableName()
	_activityBlackInfo.ALL = field.NewAsterisk(tableName)
	_activityBlackInfo.ID = field.NewInt32(tableName, "id")
	_activityBlackInfo.BlackType = field.NewInt32(tableName, "black_type")
	_activityBlackInfo.BlackValue = field.NewString(tableName, "black_value")
	_activityBlackInfo.BlackStatus = field.NewInt32(tableName, "black_status")
	_activityBlackInfo.CreateAt = field.NewInt32(tableName, "create_at")
	_activityBlackInfo.UpdateAt = field.NewInt32(tableName, "update_at")
	_activityBlackInfo.OpUser = field.NewString(tableName, "op_user")

	_activityBlackInfo.fillFieldMap()

	return _activityBlackInfo
}

// activityBlackInfo 活动黑名单信息表
type activityBlackInfo struct {
	activityBlackInfoDo

	ALL         field.Asterisk
	ID          field.Int32  // 规则ID，自增主键
	BlackType   field.Int32  // 黑名单用户类型 (0: 用户名, 1: IP, 2: 真实姓名, 3: 银行卡)
	BlackValue  field.String // 黑名单用户的值（例如，用户名、IP地址、真实姓名、银行卡号）
	BlackStatus field.Int32  // 状态：1-启用,2-禁用
	CreateAt    field.Int32  // 创建时间
	UpdateAt    field.Int32  // 修改时间
	OpUser      field.String // 操作人

	fieldMap map[string]field.Expr
}

func (a activityBlackInfo) Table(newTableName string) *activityBlackInfo {
	a.activityBlackInfoDo.UseTable(newTableName)
	return a.updateTableName(newTableName)
}

func (a activityBlackInfo) As(alias string) *activityBlackInfo {
	a.activityBlackInfoDo.DO = *(a.activityBlackInfoDo.As(alias).(*gen.DO))
	return a.updateTableName(alias)
}

func (a *activityBlackInfo) updateTableName(table string) *activityBlackInfo {
	a.ALL = field.NewAsterisk(table)
	a.ID = field.NewInt32(table, "id")
	a.BlackType = field.NewInt32(table, "black_type")
	a.BlackValue = field.NewString(table, "black_value")
	a.BlackStatus = field.NewInt32(table, "black_status")
	a.CreateAt = field.NewInt32(table, "create_at")
	a.UpdateAt = field.NewInt32(table, "update_at")
	a.OpUser = field.NewString(table, "op_user")

	a.fillFieldMap()

	return a
}

func (a *activityBlackInfo) GetFieldByName(fieldName string) (field.OrderExpr, bool) {
	_f, ok := a.fieldMap[fieldName]
	if !ok || _f == nil {
		return nil, false
	}
	_oe, ok := _f.(field.OrderExpr)
	return _oe, ok
}

func (a *activityBlackInfo) fillFieldMap() {
	a.fieldMap = make(map[string]field.Expr, 7)
	a.fieldMap["id"] = a.ID
	a.fieldMap["black_type"] = a.BlackType
	a.fieldMap["black_value"] = a.BlackValue
	a.fieldMap["black_status"] = a.BlackStatus
	a.fieldMap["create_at"] = a.CreateAt
	a.fieldMap["update_at"] = a.UpdateAt
	a.fieldMap["op_user"] = a.OpUser
}

func (a activityBlackInfo) clone(db *gorm.DB) activityBlackInfo {
	a.activityBlackInfoDo.ReplaceConnPool(db.Statement.ConnPool)
	return a
}

func (a activityBlackInfo) replaceDB(db *gorm.DB) activityBlackInfo {
	a.activityBlackInfoDo.ReplaceDB(db)
	return a
}

type activityBlackInfoDo struct{ gen.DO }

func (a activityBlackInfoDo) Debug() *activityBlackInfoDo {
	return a.withDO(a.DO.Debug())
}

func (a activityBlackInfoDo) WithContext(ctx context.Context) *activityBlackInfoDo {
	return a.withDO(a.DO.WithContext(ctx))
}

func (a activityBlackInfoDo) ReadDB() *activityBlackInfoDo {
	return a.Clauses(dbresolver.Read)
}

func (a activityBlackInfoDo) WriteDB() *activityBlackInfoDo {
	return a.Clauses(dbresolver.Write)
}

func (a activityBlackInfoDo) Session(config *gorm.Session) *activityBlackInfoDo {
	return a.withDO(a.DO.Session(config))
}

func (a activityBlackInfoDo) Clauses(conds ...clause.Expression) *activityBlackInfoDo {
	return a.withDO(a.DO.Clauses(conds...))
}

func (a activityBlackInfoDo) Returning(value interface{}, columns ...string) *activityBlackInfoDo {
	return a.withDO(a.DO.Returning(value, columns...))
}

func (a activityBlackInfoDo) Not(conds ...gen.Condition) *activityBlackInfoDo {
	return a.withDO(a.DO.Not(conds...))
}

func (a activityBlackInfoDo) Or(conds ...gen.Condition) *activityBlackInfoDo {
	return a.withDO(a.DO.Or(conds...))
}

func (a activityBlackInfoDo) Select(conds ...field.Expr) *activityBlackInfoDo {
	return a.withDO(a.DO.Select(conds...))
}

func (a activityBlackInfoDo) Where(conds ...gen.Condition) *activityBlackInfoDo {
	return a.withDO(a.DO.Where(conds...))
}

func (a activityBlackInfoDo) Order(conds ...field.Expr) *activityBlackInfoDo {
	return a.withDO(a.DO.Order(conds...))
}

func (a activityBlackInfoDo) Distinct(cols ...field.Expr) *activityBlackInfoDo {
	return a.withDO(a.DO.Distinct(cols...))
}

func (a activityBlackInfoDo) Omit(cols ...field.Expr) *activityBlackInfoDo {
	return a.withDO(a.DO.Omit(cols...))
}

func (a activityBlackInfoDo) Join(table schema.Tabler, on ...field.Expr) *activityBlackInfoDo {
	return a.withDO(a.DO.Join(table, on...))
}

func (a activityBlackInfoDo) LeftJoin(table schema.Tabler, on ...field.Expr) *activityBlackInfoDo {
	return a.withDO(a.DO.LeftJoin(table, on...))
}

func (a activityBlackInfoDo) RightJoin(table schema.Tabler, on ...field.Expr) *activityBlackInfoDo {
	return a.withDO(a.DO.RightJoin(table, on...))
}

func (a activityBlackInfoDo) Group(cols ...field.Expr) *activityBlackInfoDo {
	return a.withDO(a.DO.Group(cols...))
}

func (a activityBlackInfoDo) Having(conds ...gen.Condition) *activityBlackInfoDo {
	return a.withDO(a.DO.Having(conds...))
}

func (a activityBlackInfoDo) Limit(limit int) *activityBlackInfoDo {
	return a.withDO(a.DO.Limit(limit))
}

func (a activityBlackInfoDo) Offset(offset int) *activityBlackInfoDo {
	return a.withDO(a.DO.Offset(offset))
}

func (a activityBlackInfoDo) Scopes(funcs ...func(gen.Dao) gen.Dao) *activityBlackInfoDo {
	return a.withDO(a.DO.Scopes(funcs...))
}

func (a activityBlackInfoDo) Unscoped() *activityBlackInfoDo {
	return a.withDO(a.DO.Unscoped())
}

func (a activityBlackInfoDo) Create(values ...*model.ActivityBlackInfo) error {
	if len(values) == 0 {
		return nil
	}
	return a.DO.Create(values)
}

func (a activityBlackInfoDo) CreateInBatches(values []*model.ActivityBlackInfo, batchSize int) error {
	return a.DO.CreateInBatches(values, batchSize)
}

// Save : !!! underlying implementation is different with GORM
// The method is equivalent to executing the statement: db.Clauses(clause.OnConflict{UpdateAll: true}).Create(values)
func (a activityBlackInfoDo) Save(values ...*model.ActivityBlackInfo) error {
	if len(values) == 0 {
		return nil
	}
	return a.DO.Save(values)
}

func (a activityBlackInfoDo) First() (*model.ActivityBlackInfo, error) {
	if result, err := a.DO.First(); err != nil {
		return nil, err
	} else {
		return result.(*model.ActivityBlackInfo), nil
	}
}

func (a activityBlackInfoDo) Take() (*model.ActivityBlackInfo, error) {
	if result, err := a.DO.Take(); err != nil {
		return nil, err
	} else {
		return result.(*model.ActivityBlackInfo), nil
	}
}

func (a activityBlackInfoDo) Last() (*model.ActivityBlackInfo, error) {
	if result, err := a.DO.Last(); err != nil {
		return nil, err
	} else {
		return result.(*model.ActivityBlackInfo), nil
	}
}

func (a activityBlackInfoDo) Find() ([]*model.ActivityBlackInfo, error) {
	result, err := a.DO.Find()
	return result.([]*model.ActivityBlackInfo), err
}

func (a activityBlackInfoDo) FindInBatch(batchSize int, fc func(tx gen.Dao, batch int) error) (results []*model.ActivityBlackInfo, err error) {
	buf := make([]*model.ActivityBlackInfo, 0, batchSize)
	err = a.DO.FindInBatches(&buf, batchSize, func(tx gen.Dao, batch int) error {
		defer func() { results = append(results, buf...) }()
		return fc(tx, batch)
	})
	return results, err
}

func (a activityBlackInfoDo) FindInBatches(result *[]*model.ActivityBlackInfo, batchSize int, fc func(tx gen.Dao, batch int) error) error {
	return a.DO.FindInBatches(result, batchSize, fc)
}

func (a activityBlackInfoDo) Attrs(attrs ...field.AssignExpr) *activityBlackInfoDo {
	return a.withDO(a.DO.Attrs(attrs...))
}

func (a activityBlackInfoDo) Assign(attrs ...field.AssignExpr) *activityBlackInfoDo {
	return a.withDO(a.DO.Assign(attrs...))
}

func (a activityBlackInfoDo) Joins(fields ...field.RelationField) *activityBlackInfoDo {
	for _, _f := range fields {
		a = *a.withDO(a.DO.Joins(_f))
	}
	return &a
}

func (a activityBlackInfoDo) Preload(fields ...field.RelationField) *activityBlackInfoDo {
	for _, _f := range fields {
		a = *a.withDO(a.DO.Preload(_f))
	}
	return &a
}

func (a activityBlackInfoDo) FirstOrInit() (*model.ActivityBlackInfo, error) {
	if result, err := a.DO.FirstOrInit(); err != nil {
		return nil, err
	} else {
		return result.(*model.ActivityBlackInfo), nil
	}
}

func (a activityBlackInfoDo) FirstOrCreate() (*model.ActivityBlackInfo, error) {
	if result, err := a.DO.FirstOrCreate(); err != nil {
		return nil, err
	} else {
		return result.(*model.ActivityBlackInfo), nil
	}
}

func (a activityBlackInfoDo) FindByPage(offset int, limit int) (result []*model.ActivityBlackInfo, count int64, err error) {
	result, err = a.Offset(offset).Limit(limit).Find()
	if err != nil {
		return
	}

	if size := len(result); 0 < limit && 0 < size && size < limit {
		count = int64(size + offset)
		return
	}

	count, err = a.Offset(-1).Limit(-1).Count()
	return
}

func (a activityBlackInfoDo) ScanByPage(result interface{}, offset int, limit int) (count int64, err error) {
	count, err = a.Count()
	if err != nil {
		return
	}

	err = a.Offset(offset).Limit(limit).Scan(result)
	return
}

func (a activityBlackInfoDo) Scan(result interface{}) (err error) {
	return a.DO.Scan(result)
}

func (a activityBlackInfoDo) Delete(models ...*model.ActivityBlackInfo) (result gen.ResultInfo, err error) {
	return a.DO.Delete(models)
}

func (a *activityBlackInfoDo) withDO(do gen.Dao) *activityBlackInfoDo {
	a.DO = *do.(*gen.DO)
	return a
}
