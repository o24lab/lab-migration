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

func newWinUserPromotionRegistration(db *gorm.DB, opts ...gen.DOOption) winUserPromotionRegistration {
	_winUserPromotionRegistration := winUserPromotionRegistration{}

	_winUserPromotionRegistration.winUserPromotionRegistrationDo.UseDB(db, opts...)
	_winUserPromotionRegistration.winUserPromotionRegistrationDo.UseModel(&model.WinUserPromotionRegistration{})

	tableName := _winUserPromotionRegistration.winUserPromotionRegistrationDo.TableName()
	_winUserPromotionRegistration.ALL = field.NewAsterisk(tableName)
	_winUserPromotionRegistration.ID = field.NewInt64(tableName, "id")
	_winUserPromotionRegistration.ImportTime = field.NewInt64(tableName, "import_time")
	_winUserPromotionRegistration.BatchNumber = field.NewInt64(tableName, "batch_number")
	_winUserPromotionRegistration.PhoneNumber = field.NewInt64(tableName, "phone_number")
	_winUserPromotionRegistration.RegistrationTime = field.NewInt64(tableName, "registration_time")
	_winUserPromotionRegistration.UserID = field.NewInt32(tableName, "user_id")
	_winUserPromotionRegistration.OpUser = field.NewString(tableName, "op_user")
	_winUserPromotionRegistration.CreatedAt = field.NewInt64(tableName, "created_at")
	_winUserPromotionRegistration.UpdatedAt = field.NewInt64(tableName, "updated_at")
	_winUserPromotionRegistration.PayAmount = field.NewFloat64(tableName, "pay_amount")
	_winUserPromotionRegistration.WithdrawalAmount = field.NewFloat64(tableName, "withdrawal_amount")
	_winUserPromotionRegistration.DiffAmount = field.NewFloat64(tableName, "diff_amount")
	_winUserPromotionRegistration.EffectedAt = field.NewInt64(tableName, "effected_at")

	_winUserPromotionRegistration.fillFieldMap()

	return _winUserPromotionRegistration
}

// winUserPromotionRegistration 用户推广注册信息表
type winUserPromotionRegistration struct {
	winUserPromotionRegistrationDo

	ALL              field.Asterisk
	ID               field.Int64   // 主键
	ImportTime       field.Int64   // 导入时间
	BatchNumber      field.Int64   // 批次号
	PhoneNumber      field.Int64   // 手机号
	RegistrationTime field.Int64   // 注册时间
	UserID           field.Int32   // 注册用户id
	OpUser           field.String  // 操作人
	CreatedAt        field.Int64   // 创建时间
	UpdatedAt        field.Int64   // 更新时间
	PayAmount        field.Float64 // 充值金额
	WithdrawalAmount field.Float64 // 提款金额
	DiffAmount       field.Float64 // 存提差额
	EffectedAt       field.Int64   // 导入生效时间

	fieldMap map[string]field.Expr
}

func (w winUserPromotionRegistration) Table(newTableName string) *winUserPromotionRegistration {
	w.winUserPromotionRegistrationDo.UseTable(newTableName)
	return w.updateTableName(newTableName)
}

func (w winUserPromotionRegistration) As(alias string) *winUserPromotionRegistration {
	w.winUserPromotionRegistrationDo.DO = *(w.winUserPromotionRegistrationDo.As(alias).(*gen.DO))
	return w.updateTableName(alias)
}

func (w *winUserPromotionRegistration) updateTableName(table string) *winUserPromotionRegistration {
	w.ALL = field.NewAsterisk(table)
	w.ID = field.NewInt64(table, "id")
	w.ImportTime = field.NewInt64(table, "import_time")
	w.BatchNumber = field.NewInt64(table, "batch_number")
	w.PhoneNumber = field.NewInt64(table, "phone_number")
	w.RegistrationTime = field.NewInt64(table, "registration_time")
	w.UserID = field.NewInt32(table, "user_id")
	w.OpUser = field.NewString(table, "op_user")
	w.CreatedAt = field.NewInt64(table, "created_at")
	w.UpdatedAt = field.NewInt64(table, "updated_at")
	w.PayAmount = field.NewFloat64(table, "pay_amount")
	w.WithdrawalAmount = field.NewFloat64(table, "withdrawal_amount")
	w.DiffAmount = field.NewFloat64(table, "diff_amount")
	w.EffectedAt = field.NewInt64(table, "effected_at")

	w.fillFieldMap()

	return w
}

func (w *winUserPromotionRegistration) GetFieldByName(fieldName string) (field.OrderExpr, bool) {
	_f, ok := w.fieldMap[fieldName]
	if !ok || _f == nil {
		return nil, false
	}
	_oe, ok := _f.(field.OrderExpr)
	return _oe, ok
}

func (w *winUserPromotionRegistration) fillFieldMap() {
	w.fieldMap = make(map[string]field.Expr, 13)
	w.fieldMap["id"] = w.ID
	w.fieldMap["import_time"] = w.ImportTime
	w.fieldMap["batch_number"] = w.BatchNumber
	w.fieldMap["phone_number"] = w.PhoneNumber
	w.fieldMap["registration_time"] = w.RegistrationTime
	w.fieldMap["user_id"] = w.UserID
	w.fieldMap["op_user"] = w.OpUser
	w.fieldMap["created_at"] = w.CreatedAt
	w.fieldMap["updated_at"] = w.UpdatedAt
	w.fieldMap["pay_amount"] = w.PayAmount
	w.fieldMap["withdrawal_amount"] = w.WithdrawalAmount
	w.fieldMap["diff_amount"] = w.DiffAmount
	w.fieldMap["effected_at"] = w.EffectedAt
}

func (w winUserPromotionRegistration) clone(db *gorm.DB) winUserPromotionRegistration {
	w.winUserPromotionRegistrationDo.ReplaceConnPool(db.Statement.ConnPool)
	return w
}

func (w winUserPromotionRegistration) replaceDB(db *gorm.DB) winUserPromotionRegistration {
	w.winUserPromotionRegistrationDo.ReplaceDB(db)
	return w
}

type winUserPromotionRegistrationDo struct{ gen.DO }

func (w winUserPromotionRegistrationDo) Debug() *winUserPromotionRegistrationDo {
	return w.withDO(w.DO.Debug())
}

func (w winUserPromotionRegistrationDo) WithContext(ctx context.Context) *winUserPromotionRegistrationDo {
	return w.withDO(w.DO.WithContext(ctx))
}

func (w winUserPromotionRegistrationDo) ReadDB() *winUserPromotionRegistrationDo {
	return w.Clauses(dbresolver.Read)
}

func (w winUserPromotionRegistrationDo) WriteDB() *winUserPromotionRegistrationDo {
	return w.Clauses(dbresolver.Write)
}

func (w winUserPromotionRegistrationDo) Session(config *gorm.Session) *winUserPromotionRegistrationDo {
	return w.withDO(w.DO.Session(config))
}

func (w winUserPromotionRegistrationDo) Clauses(conds ...clause.Expression) *winUserPromotionRegistrationDo {
	return w.withDO(w.DO.Clauses(conds...))
}

func (w winUserPromotionRegistrationDo) Returning(value interface{}, columns ...string) *winUserPromotionRegistrationDo {
	return w.withDO(w.DO.Returning(value, columns...))
}

func (w winUserPromotionRegistrationDo) Not(conds ...gen.Condition) *winUserPromotionRegistrationDo {
	return w.withDO(w.DO.Not(conds...))
}

func (w winUserPromotionRegistrationDo) Or(conds ...gen.Condition) *winUserPromotionRegistrationDo {
	return w.withDO(w.DO.Or(conds...))
}

func (w winUserPromotionRegistrationDo) Select(conds ...field.Expr) *winUserPromotionRegistrationDo {
	return w.withDO(w.DO.Select(conds...))
}

func (w winUserPromotionRegistrationDo) Where(conds ...gen.Condition) *winUserPromotionRegistrationDo {
	return w.withDO(w.DO.Where(conds...))
}

func (w winUserPromotionRegistrationDo) Order(conds ...field.Expr) *winUserPromotionRegistrationDo {
	return w.withDO(w.DO.Order(conds...))
}

func (w winUserPromotionRegistrationDo) Distinct(cols ...field.Expr) *winUserPromotionRegistrationDo {
	return w.withDO(w.DO.Distinct(cols...))
}

func (w winUserPromotionRegistrationDo) Omit(cols ...field.Expr) *winUserPromotionRegistrationDo {
	return w.withDO(w.DO.Omit(cols...))
}

func (w winUserPromotionRegistrationDo) Join(table schema.Tabler, on ...field.Expr) *winUserPromotionRegistrationDo {
	return w.withDO(w.DO.Join(table, on...))
}

func (w winUserPromotionRegistrationDo) LeftJoin(table schema.Tabler, on ...field.Expr) *winUserPromotionRegistrationDo {
	return w.withDO(w.DO.LeftJoin(table, on...))
}

func (w winUserPromotionRegistrationDo) RightJoin(table schema.Tabler, on ...field.Expr) *winUserPromotionRegistrationDo {
	return w.withDO(w.DO.RightJoin(table, on...))
}

func (w winUserPromotionRegistrationDo) Group(cols ...field.Expr) *winUserPromotionRegistrationDo {
	return w.withDO(w.DO.Group(cols...))
}

func (w winUserPromotionRegistrationDo) Having(conds ...gen.Condition) *winUserPromotionRegistrationDo {
	return w.withDO(w.DO.Having(conds...))
}

func (w winUserPromotionRegistrationDo) Limit(limit int) *winUserPromotionRegistrationDo {
	return w.withDO(w.DO.Limit(limit))
}

func (w winUserPromotionRegistrationDo) Offset(offset int) *winUserPromotionRegistrationDo {
	return w.withDO(w.DO.Offset(offset))
}

func (w winUserPromotionRegistrationDo) Scopes(funcs ...func(gen.Dao) gen.Dao) *winUserPromotionRegistrationDo {
	return w.withDO(w.DO.Scopes(funcs...))
}

func (w winUserPromotionRegistrationDo) Unscoped() *winUserPromotionRegistrationDo {
	return w.withDO(w.DO.Unscoped())
}

func (w winUserPromotionRegistrationDo) Create(values ...*model.WinUserPromotionRegistration) error {
	if len(values) == 0 {
		return nil
	}
	return w.DO.Create(values)
}

func (w winUserPromotionRegistrationDo) CreateInBatches(values []*model.WinUserPromotionRegistration, batchSize int) error {
	return w.DO.CreateInBatches(values, batchSize)
}

// Save : !!! underlying implementation is different with GORM
// The method is equivalent to executing the statement: db.Clauses(clause.OnConflict{UpdateAll: true}).Create(values)
func (w winUserPromotionRegistrationDo) Save(values ...*model.WinUserPromotionRegistration) error {
	if len(values) == 0 {
		return nil
	}
	return w.DO.Save(values)
}

func (w winUserPromotionRegistrationDo) First() (*model.WinUserPromotionRegistration, error) {
	if result, err := w.DO.First(); err != nil {
		return nil, err
	} else {
		return result.(*model.WinUserPromotionRegistration), nil
	}
}

func (w winUserPromotionRegistrationDo) Take() (*model.WinUserPromotionRegistration, error) {
	if result, err := w.DO.Take(); err != nil {
		return nil, err
	} else {
		return result.(*model.WinUserPromotionRegistration), nil
	}
}

func (w winUserPromotionRegistrationDo) Last() (*model.WinUserPromotionRegistration, error) {
	if result, err := w.DO.Last(); err != nil {
		return nil, err
	} else {
		return result.(*model.WinUserPromotionRegistration), nil
	}
}

func (w winUserPromotionRegistrationDo) Find() ([]*model.WinUserPromotionRegistration, error) {
	result, err := w.DO.Find()
	return result.([]*model.WinUserPromotionRegistration), err
}

func (w winUserPromotionRegistrationDo) FindInBatch(batchSize int, fc func(tx gen.Dao, batch int) error) (results []*model.WinUserPromotionRegistration, err error) {
	buf := make([]*model.WinUserPromotionRegistration, 0, batchSize)
	err = w.DO.FindInBatches(&buf, batchSize, func(tx gen.Dao, batch int) error {
		defer func() { results = append(results, buf...) }()
		return fc(tx, batch)
	})
	return results, err
}

func (w winUserPromotionRegistrationDo) FindInBatches(result *[]*model.WinUserPromotionRegistration, batchSize int, fc func(tx gen.Dao, batch int) error) error {
	return w.DO.FindInBatches(result, batchSize, fc)
}

func (w winUserPromotionRegistrationDo) Attrs(attrs ...field.AssignExpr) *winUserPromotionRegistrationDo {
	return w.withDO(w.DO.Attrs(attrs...))
}

func (w winUserPromotionRegistrationDo) Assign(attrs ...field.AssignExpr) *winUserPromotionRegistrationDo {
	return w.withDO(w.DO.Assign(attrs...))
}

func (w winUserPromotionRegistrationDo) Joins(fields ...field.RelationField) *winUserPromotionRegistrationDo {
	for _, _f := range fields {
		w = *w.withDO(w.DO.Joins(_f))
	}
	return &w
}

func (w winUserPromotionRegistrationDo) Preload(fields ...field.RelationField) *winUserPromotionRegistrationDo {
	for _, _f := range fields {
		w = *w.withDO(w.DO.Preload(_f))
	}
	return &w
}

func (w winUserPromotionRegistrationDo) FirstOrInit() (*model.WinUserPromotionRegistration, error) {
	if result, err := w.DO.FirstOrInit(); err != nil {
		return nil, err
	} else {
		return result.(*model.WinUserPromotionRegistration), nil
	}
}

func (w winUserPromotionRegistrationDo) FirstOrCreate() (*model.WinUserPromotionRegistration, error) {
	if result, err := w.DO.FirstOrCreate(); err != nil {
		return nil, err
	} else {
		return result.(*model.WinUserPromotionRegistration), nil
	}
}

func (w winUserPromotionRegistrationDo) FindByPage(offset int, limit int) (result []*model.WinUserPromotionRegistration, count int64, err error) {
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

func (w winUserPromotionRegistrationDo) ScanByPage(result interface{}, offset int, limit int) (count int64, err error) {
	count, err = w.Count()
	if err != nil {
		return
	}

	err = w.Offset(offset).Limit(limit).Scan(result)
	return
}

func (w winUserPromotionRegistrationDo) Scan(result interface{}) (err error) {
	return w.DO.Scan(result)
}

func (w winUserPromotionRegistrationDo) Delete(models ...*model.WinUserPromotionRegistration) (result gen.ResultInfo, err error) {
	return w.DO.Delete(models)
}

func (w *winUserPromotionRegistrationDo) withDO(do gen.Dao) *winUserPromotionRegistrationDo {
	w.DO = *do.(*gen.DO)
	return w
}
