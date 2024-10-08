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

func newWinCodeRecordsVip(db *gorm.DB, opts ...gen.DOOption) winCodeRecordsVip {
	_winCodeRecordsVip := winCodeRecordsVip{}

	_winCodeRecordsVip.winCodeRecordsVipDo.UseDB(db, opts...)
	_winCodeRecordsVip.winCodeRecordsVipDo.UseModel(&model.WinCodeRecordsVip{})

	tableName := _winCodeRecordsVip.winCodeRecordsVipDo.TableName()
	_winCodeRecordsVip.ALL = field.NewAsterisk(tableName)
	_winCodeRecordsVip.ID = field.NewInt64(tableName, "id")
	_winCodeRecordsVip.Remark = field.NewString(tableName, "remark")
	_winCodeRecordsVip.UID = field.NewInt32(tableName, "uid")
	_winCodeRecordsVip.ValidAmount = field.NewFloat64(tableName, "valid_amount")
	_winCodeRecordsVip.Category = field.NewInt32(tableName, "category")
	_winCodeRecordsVip.Version = field.NewInt32(tableName, "version")
	_winCodeRecordsVip.CreatedAt = field.NewInt32(tableName, "created_at")
	_winCodeRecordsVip.UpdatedAt = field.NewInt32(tableName, "updated_at")
	_winCodeRecordsVip.UpdatedUser = field.NewString(tableName, "updated_user")
	_winCodeRecordsVip.DaysNotLogin = field.NewInt32(tableName, "days_not_login")
	_winCodeRecordsVip.AlreadySmsCount = field.NewInt32(tableName, "already_sms_count")

	_winCodeRecordsVip.fillFieldMap()

	return _winCodeRecordsVip
}

// winCodeRecordsVip VIP会员有效流水表
type winCodeRecordsVip struct {
	winCodeRecordsVipDo

	ALL             field.Asterisk
	ID              field.Int64   // 自增主键
	Remark          field.String  // 备注
	UID             field.Int32   // UID
	ValidAmount     field.Float64 // 有效投注额/赠送金额
	Category        field.Int32   // 类型:1-有效投注额 2-赠送金额
	Version         field.Int32   // 更新版本号
	CreatedAt       field.Int32   // 创建时间
	UpdatedAt       field.Int32   // 修改时间
	UpdatedUser     field.String  // 最后修改人
	DaysNotLogin    field.Int32   // 未上线天数
	AlreadySmsCount field.Int32   // 已发送消息条数

	fieldMap map[string]field.Expr
}

func (w winCodeRecordsVip) Table(newTableName string) *winCodeRecordsVip {
	w.winCodeRecordsVipDo.UseTable(newTableName)
	return w.updateTableName(newTableName)
}

func (w winCodeRecordsVip) As(alias string) *winCodeRecordsVip {
	w.winCodeRecordsVipDo.DO = *(w.winCodeRecordsVipDo.As(alias).(*gen.DO))
	return w.updateTableName(alias)
}

func (w *winCodeRecordsVip) updateTableName(table string) *winCodeRecordsVip {
	w.ALL = field.NewAsterisk(table)
	w.ID = field.NewInt64(table, "id")
	w.Remark = field.NewString(table, "remark")
	w.UID = field.NewInt32(table, "uid")
	w.ValidAmount = field.NewFloat64(table, "valid_amount")
	w.Category = field.NewInt32(table, "category")
	w.Version = field.NewInt32(table, "version")
	w.CreatedAt = field.NewInt32(table, "created_at")
	w.UpdatedAt = field.NewInt32(table, "updated_at")
	w.UpdatedUser = field.NewString(table, "updated_user")
	w.DaysNotLogin = field.NewInt32(table, "days_not_login")
	w.AlreadySmsCount = field.NewInt32(table, "already_sms_count")

	w.fillFieldMap()

	return w
}

func (w *winCodeRecordsVip) GetFieldByName(fieldName string) (field.OrderExpr, bool) {
	_f, ok := w.fieldMap[fieldName]
	if !ok || _f == nil {
		return nil, false
	}
	_oe, ok := _f.(field.OrderExpr)
	return _oe, ok
}

func (w *winCodeRecordsVip) fillFieldMap() {
	w.fieldMap = make(map[string]field.Expr, 11)
	w.fieldMap["id"] = w.ID
	w.fieldMap["remark"] = w.Remark
	w.fieldMap["uid"] = w.UID
	w.fieldMap["valid_amount"] = w.ValidAmount
	w.fieldMap["category"] = w.Category
	w.fieldMap["version"] = w.Version
	w.fieldMap["created_at"] = w.CreatedAt
	w.fieldMap["updated_at"] = w.UpdatedAt
	w.fieldMap["updated_user"] = w.UpdatedUser
	w.fieldMap["days_not_login"] = w.DaysNotLogin
	w.fieldMap["already_sms_count"] = w.AlreadySmsCount
}

func (w winCodeRecordsVip) clone(db *gorm.DB) winCodeRecordsVip {
	w.winCodeRecordsVipDo.ReplaceConnPool(db.Statement.ConnPool)
	return w
}

func (w winCodeRecordsVip) replaceDB(db *gorm.DB) winCodeRecordsVip {
	w.winCodeRecordsVipDo.ReplaceDB(db)
	return w
}

type winCodeRecordsVipDo struct{ gen.DO }

func (w winCodeRecordsVipDo) Debug() *winCodeRecordsVipDo {
	return w.withDO(w.DO.Debug())
}

func (w winCodeRecordsVipDo) WithContext(ctx context.Context) *winCodeRecordsVipDo {
	return w.withDO(w.DO.WithContext(ctx))
}

func (w winCodeRecordsVipDo) ReadDB() *winCodeRecordsVipDo {
	return w.Clauses(dbresolver.Read)
}

func (w winCodeRecordsVipDo) WriteDB() *winCodeRecordsVipDo {
	return w.Clauses(dbresolver.Write)
}

func (w winCodeRecordsVipDo) Session(config *gorm.Session) *winCodeRecordsVipDo {
	return w.withDO(w.DO.Session(config))
}

func (w winCodeRecordsVipDo) Clauses(conds ...clause.Expression) *winCodeRecordsVipDo {
	return w.withDO(w.DO.Clauses(conds...))
}

func (w winCodeRecordsVipDo) Returning(value interface{}, columns ...string) *winCodeRecordsVipDo {
	return w.withDO(w.DO.Returning(value, columns...))
}

func (w winCodeRecordsVipDo) Not(conds ...gen.Condition) *winCodeRecordsVipDo {
	return w.withDO(w.DO.Not(conds...))
}

func (w winCodeRecordsVipDo) Or(conds ...gen.Condition) *winCodeRecordsVipDo {
	return w.withDO(w.DO.Or(conds...))
}

func (w winCodeRecordsVipDo) Select(conds ...field.Expr) *winCodeRecordsVipDo {
	return w.withDO(w.DO.Select(conds...))
}

func (w winCodeRecordsVipDo) Where(conds ...gen.Condition) *winCodeRecordsVipDo {
	return w.withDO(w.DO.Where(conds...))
}

func (w winCodeRecordsVipDo) Order(conds ...field.Expr) *winCodeRecordsVipDo {
	return w.withDO(w.DO.Order(conds...))
}

func (w winCodeRecordsVipDo) Distinct(cols ...field.Expr) *winCodeRecordsVipDo {
	return w.withDO(w.DO.Distinct(cols...))
}

func (w winCodeRecordsVipDo) Omit(cols ...field.Expr) *winCodeRecordsVipDo {
	return w.withDO(w.DO.Omit(cols...))
}

func (w winCodeRecordsVipDo) Join(table schema.Tabler, on ...field.Expr) *winCodeRecordsVipDo {
	return w.withDO(w.DO.Join(table, on...))
}

func (w winCodeRecordsVipDo) LeftJoin(table schema.Tabler, on ...field.Expr) *winCodeRecordsVipDo {
	return w.withDO(w.DO.LeftJoin(table, on...))
}

func (w winCodeRecordsVipDo) RightJoin(table schema.Tabler, on ...field.Expr) *winCodeRecordsVipDo {
	return w.withDO(w.DO.RightJoin(table, on...))
}

func (w winCodeRecordsVipDo) Group(cols ...field.Expr) *winCodeRecordsVipDo {
	return w.withDO(w.DO.Group(cols...))
}

func (w winCodeRecordsVipDo) Having(conds ...gen.Condition) *winCodeRecordsVipDo {
	return w.withDO(w.DO.Having(conds...))
}

func (w winCodeRecordsVipDo) Limit(limit int) *winCodeRecordsVipDo {
	return w.withDO(w.DO.Limit(limit))
}

func (w winCodeRecordsVipDo) Offset(offset int) *winCodeRecordsVipDo {
	return w.withDO(w.DO.Offset(offset))
}

func (w winCodeRecordsVipDo) Scopes(funcs ...func(gen.Dao) gen.Dao) *winCodeRecordsVipDo {
	return w.withDO(w.DO.Scopes(funcs...))
}

func (w winCodeRecordsVipDo) Unscoped() *winCodeRecordsVipDo {
	return w.withDO(w.DO.Unscoped())
}

func (w winCodeRecordsVipDo) Create(values ...*model.WinCodeRecordsVip) error {
	if len(values) == 0 {
		return nil
	}
	return w.DO.Create(values)
}

func (w winCodeRecordsVipDo) CreateInBatches(values []*model.WinCodeRecordsVip, batchSize int) error {
	return w.DO.CreateInBatches(values, batchSize)
}

// Save : !!! underlying implementation is different with GORM
// The method is equivalent to executing the statement: db.Clauses(clause.OnConflict{UpdateAll: true}).Create(values)
func (w winCodeRecordsVipDo) Save(values ...*model.WinCodeRecordsVip) error {
	if len(values) == 0 {
		return nil
	}
	return w.DO.Save(values)
}

func (w winCodeRecordsVipDo) First() (*model.WinCodeRecordsVip, error) {
	if result, err := w.DO.First(); err != nil {
		return nil, err
	} else {
		return result.(*model.WinCodeRecordsVip), nil
	}
}

func (w winCodeRecordsVipDo) Take() (*model.WinCodeRecordsVip, error) {
	if result, err := w.DO.Take(); err != nil {
		return nil, err
	} else {
		return result.(*model.WinCodeRecordsVip), nil
	}
}

func (w winCodeRecordsVipDo) Last() (*model.WinCodeRecordsVip, error) {
	if result, err := w.DO.Last(); err != nil {
		return nil, err
	} else {
		return result.(*model.WinCodeRecordsVip), nil
	}
}

func (w winCodeRecordsVipDo) Find() ([]*model.WinCodeRecordsVip, error) {
	result, err := w.DO.Find()
	return result.([]*model.WinCodeRecordsVip), err
}

func (w winCodeRecordsVipDo) FindInBatch(batchSize int, fc func(tx gen.Dao, batch int) error) (results []*model.WinCodeRecordsVip, err error) {
	buf := make([]*model.WinCodeRecordsVip, 0, batchSize)
	err = w.DO.FindInBatches(&buf, batchSize, func(tx gen.Dao, batch int) error {
		defer func() { results = append(results, buf...) }()
		return fc(tx, batch)
	})
	return results, err
}

func (w winCodeRecordsVipDo) FindInBatches(result *[]*model.WinCodeRecordsVip, batchSize int, fc func(tx gen.Dao, batch int) error) error {
	return w.DO.FindInBatches(result, batchSize, fc)
}

func (w winCodeRecordsVipDo) Attrs(attrs ...field.AssignExpr) *winCodeRecordsVipDo {
	return w.withDO(w.DO.Attrs(attrs...))
}

func (w winCodeRecordsVipDo) Assign(attrs ...field.AssignExpr) *winCodeRecordsVipDo {
	return w.withDO(w.DO.Assign(attrs...))
}

func (w winCodeRecordsVipDo) Joins(fields ...field.RelationField) *winCodeRecordsVipDo {
	for _, _f := range fields {
		w = *w.withDO(w.DO.Joins(_f))
	}
	return &w
}

func (w winCodeRecordsVipDo) Preload(fields ...field.RelationField) *winCodeRecordsVipDo {
	for _, _f := range fields {
		w = *w.withDO(w.DO.Preload(_f))
	}
	return &w
}

func (w winCodeRecordsVipDo) FirstOrInit() (*model.WinCodeRecordsVip, error) {
	if result, err := w.DO.FirstOrInit(); err != nil {
		return nil, err
	} else {
		return result.(*model.WinCodeRecordsVip), nil
	}
}

func (w winCodeRecordsVipDo) FirstOrCreate() (*model.WinCodeRecordsVip, error) {
	if result, err := w.DO.FirstOrCreate(); err != nil {
		return nil, err
	} else {
		return result.(*model.WinCodeRecordsVip), nil
	}
}

func (w winCodeRecordsVipDo) FindByPage(offset int, limit int) (result []*model.WinCodeRecordsVip, count int64, err error) {
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

func (w winCodeRecordsVipDo) ScanByPage(result interface{}, offset int, limit int) (count int64, err error) {
	count, err = w.Count()
	if err != nil {
		return
	}

	err = w.Offset(offset).Limit(limit).Scan(result)
	return
}

func (w winCodeRecordsVipDo) Scan(result interface{}) (err error) {
	return w.DO.Scan(result)
}

func (w winCodeRecordsVipDo) Delete(models ...*model.WinCodeRecordsVip) (result gen.ResultInfo, err error) {
	return w.DO.Delete(models)
}

func (w *winCodeRecordsVipDo) withDO(do gen.Dao) *winCodeRecordsVipDo {
	w.DO = *do.(*gen.DO)
	return w
}
