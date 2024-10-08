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

func newWinFbOrderLog(db *gorm.DB, opts ...gen.DOOption) winFbOrderLog {
	_winFbOrderLog := winFbOrderLog{}

	_winFbOrderLog.winFbOrderLogDo.UseDB(db, opts...)
	_winFbOrderLog.winFbOrderLogDo.UseModel(&model.WinFbOrderLog{})

	tableName := _winFbOrderLog.winFbOrderLogDo.TableName()
	_winFbOrderLog.ALL = field.NewAsterisk(tableName)
	_winFbOrderLog.ID = field.NewInt64(tableName, "id")
	_winFbOrderLog.TransactionID = field.NewString(tableName, "transaction_id")
	_winFbOrderLog.UserID = field.NewString(tableName, "user_id")
	_winFbOrderLog.MerchantID = field.NewString(tableName, "merchant_id")
	_winFbOrderLog.MerchantUserID = field.NewString(tableName, "merchant_user_id")
	_winFbOrderLog.BusinessID = field.NewString(tableName, "business_id")
	_winFbOrderLog.TransactionType = field.NewString(tableName, "transaction_type")
	_winFbOrderLog.TransferType = field.NewString(tableName, "transfer_type")
	_winFbOrderLog.CurrencyID = field.NewInt32(tableName, "currency_id")
	_winFbOrderLog.Amount = field.NewFloat64(tableName, "amount")
	_winFbOrderLog.Status = field.NewInt32(tableName, "status")
	_winFbOrderLog.RelatedID = field.NewString(tableName, "related_id")
	_winFbOrderLog.CreatedAt = field.NewInt32(tableName, "created_at")

	_winFbOrderLog.fillFieldMap()

	return _winFbOrderLog
}

// winFbOrderLog Fb体育订单流水表
type winFbOrderLog struct {
	winFbOrderLogDo

	ALL             field.Asterisk
	ID              field.Int64
	TransactionID   field.String  // 交易流水ID，全服唯一
	UserID          field.String  // FB用户ID
	MerchantID      field.String  // 渠道ID
	MerchantUserID  field.String  // 渠道用户ID
	BusinessID      field.String  // 业务ID，即订单ID
	TransactionType field.String  // 交易类型 OUT 转出，IN 转入
	TransferType    field.String  // 转账类型
	CurrencyID      field.Int32   // 币种ID
	Amount          field.Float64 // 流水金额
	Status          field.Int32   // 转账状态,0 取消，1成功
	RelatedID       field.String  // 三方数据关联ID
	CreatedAt       field.Int32   // 创建时间

	fieldMap map[string]field.Expr
}

func (w winFbOrderLog) Table(newTableName string) *winFbOrderLog {
	w.winFbOrderLogDo.UseTable(newTableName)
	return w.updateTableName(newTableName)
}

func (w winFbOrderLog) As(alias string) *winFbOrderLog {
	w.winFbOrderLogDo.DO = *(w.winFbOrderLogDo.As(alias).(*gen.DO))
	return w.updateTableName(alias)
}

func (w *winFbOrderLog) updateTableName(table string) *winFbOrderLog {
	w.ALL = field.NewAsterisk(table)
	w.ID = field.NewInt64(table, "id")
	w.TransactionID = field.NewString(table, "transaction_id")
	w.UserID = field.NewString(table, "user_id")
	w.MerchantID = field.NewString(table, "merchant_id")
	w.MerchantUserID = field.NewString(table, "merchant_user_id")
	w.BusinessID = field.NewString(table, "business_id")
	w.TransactionType = field.NewString(table, "transaction_type")
	w.TransferType = field.NewString(table, "transfer_type")
	w.CurrencyID = field.NewInt32(table, "currency_id")
	w.Amount = field.NewFloat64(table, "amount")
	w.Status = field.NewInt32(table, "status")
	w.RelatedID = field.NewString(table, "related_id")
	w.CreatedAt = field.NewInt32(table, "created_at")

	w.fillFieldMap()

	return w
}

func (w *winFbOrderLog) GetFieldByName(fieldName string) (field.OrderExpr, bool) {
	_f, ok := w.fieldMap[fieldName]
	if !ok || _f == nil {
		return nil, false
	}
	_oe, ok := _f.(field.OrderExpr)
	return _oe, ok
}

func (w *winFbOrderLog) fillFieldMap() {
	w.fieldMap = make(map[string]field.Expr, 13)
	w.fieldMap["id"] = w.ID
	w.fieldMap["transaction_id"] = w.TransactionID
	w.fieldMap["user_id"] = w.UserID
	w.fieldMap["merchant_id"] = w.MerchantID
	w.fieldMap["merchant_user_id"] = w.MerchantUserID
	w.fieldMap["business_id"] = w.BusinessID
	w.fieldMap["transaction_type"] = w.TransactionType
	w.fieldMap["transfer_type"] = w.TransferType
	w.fieldMap["currency_id"] = w.CurrencyID
	w.fieldMap["amount"] = w.Amount
	w.fieldMap["status"] = w.Status
	w.fieldMap["related_id"] = w.RelatedID
	w.fieldMap["created_at"] = w.CreatedAt
}

func (w winFbOrderLog) clone(db *gorm.DB) winFbOrderLog {
	w.winFbOrderLogDo.ReplaceConnPool(db.Statement.ConnPool)
	return w
}

func (w winFbOrderLog) replaceDB(db *gorm.DB) winFbOrderLog {
	w.winFbOrderLogDo.ReplaceDB(db)
	return w
}

type winFbOrderLogDo struct{ gen.DO }

func (w winFbOrderLogDo) Debug() *winFbOrderLogDo {
	return w.withDO(w.DO.Debug())
}

func (w winFbOrderLogDo) WithContext(ctx context.Context) *winFbOrderLogDo {
	return w.withDO(w.DO.WithContext(ctx))
}

func (w winFbOrderLogDo) ReadDB() *winFbOrderLogDo {
	return w.Clauses(dbresolver.Read)
}

func (w winFbOrderLogDo) WriteDB() *winFbOrderLogDo {
	return w.Clauses(dbresolver.Write)
}

func (w winFbOrderLogDo) Session(config *gorm.Session) *winFbOrderLogDo {
	return w.withDO(w.DO.Session(config))
}

func (w winFbOrderLogDo) Clauses(conds ...clause.Expression) *winFbOrderLogDo {
	return w.withDO(w.DO.Clauses(conds...))
}

func (w winFbOrderLogDo) Returning(value interface{}, columns ...string) *winFbOrderLogDo {
	return w.withDO(w.DO.Returning(value, columns...))
}

func (w winFbOrderLogDo) Not(conds ...gen.Condition) *winFbOrderLogDo {
	return w.withDO(w.DO.Not(conds...))
}

func (w winFbOrderLogDo) Or(conds ...gen.Condition) *winFbOrderLogDo {
	return w.withDO(w.DO.Or(conds...))
}

func (w winFbOrderLogDo) Select(conds ...field.Expr) *winFbOrderLogDo {
	return w.withDO(w.DO.Select(conds...))
}

func (w winFbOrderLogDo) Where(conds ...gen.Condition) *winFbOrderLogDo {
	return w.withDO(w.DO.Where(conds...))
}

func (w winFbOrderLogDo) Order(conds ...field.Expr) *winFbOrderLogDo {
	return w.withDO(w.DO.Order(conds...))
}

func (w winFbOrderLogDo) Distinct(cols ...field.Expr) *winFbOrderLogDo {
	return w.withDO(w.DO.Distinct(cols...))
}

func (w winFbOrderLogDo) Omit(cols ...field.Expr) *winFbOrderLogDo {
	return w.withDO(w.DO.Omit(cols...))
}

func (w winFbOrderLogDo) Join(table schema.Tabler, on ...field.Expr) *winFbOrderLogDo {
	return w.withDO(w.DO.Join(table, on...))
}

func (w winFbOrderLogDo) LeftJoin(table schema.Tabler, on ...field.Expr) *winFbOrderLogDo {
	return w.withDO(w.DO.LeftJoin(table, on...))
}

func (w winFbOrderLogDo) RightJoin(table schema.Tabler, on ...field.Expr) *winFbOrderLogDo {
	return w.withDO(w.DO.RightJoin(table, on...))
}

func (w winFbOrderLogDo) Group(cols ...field.Expr) *winFbOrderLogDo {
	return w.withDO(w.DO.Group(cols...))
}

func (w winFbOrderLogDo) Having(conds ...gen.Condition) *winFbOrderLogDo {
	return w.withDO(w.DO.Having(conds...))
}

func (w winFbOrderLogDo) Limit(limit int) *winFbOrderLogDo {
	return w.withDO(w.DO.Limit(limit))
}

func (w winFbOrderLogDo) Offset(offset int) *winFbOrderLogDo {
	return w.withDO(w.DO.Offset(offset))
}

func (w winFbOrderLogDo) Scopes(funcs ...func(gen.Dao) gen.Dao) *winFbOrderLogDo {
	return w.withDO(w.DO.Scopes(funcs...))
}

func (w winFbOrderLogDo) Unscoped() *winFbOrderLogDo {
	return w.withDO(w.DO.Unscoped())
}

func (w winFbOrderLogDo) Create(values ...*model.WinFbOrderLog) error {
	if len(values) == 0 {
		return nil
	}
	return w.DO.Create(values)
}

func (w winFbOrderLogDo) CreateInBatches(values []*model.WinFbOrderLog, batchSize int) error {
	return w.DO.CreateInBatches(values, batchSize)
}

// Save : !!! underlying implementation is different with GORM
// The method is equivalent to executing the statement: db.Clauses(clause.OnConflict{UpdateAll: true}).Create(values)
func (w winFbOrderLogDo) Save(values ...*model.WinFbOrderLog) error {
	if len(values) == 0 {
		return nil
	}
	return w.DO.Save(values)
}

func (w winFbOrderLogDo) First() (*model.WinFbOrderLog, error) {
	if result, err := w.DO.First(); err != nil {
		return nil, err
	} else {
		return result.(*model.WinFbOrderLog), nil
	}
}

func (w winFbOrderLogDo) Take() (*model.WinFbOrderLog, error) {
	if result, err := w.DO.Take(); err != nil {
		return nil, err
	} else {
		return result.(*model.WinFbOrderLog), nil
	}
}

func (w winFbOrderLogDo) Last() (*model.WinFbOrderLog, error) {
	if result, err := w.DO.Last(); err != nil {
		return nil, err
	} else {
		return result.(*model.WinFbOrderLog), nil
	}
}

func (w winFbOrderLogDo) Find() ([]*model.WinFbOrderLog, error) {
	result, err := w.DO.Find()
	return result.([]*model.WinFbOrderLog), err
}

func (w winFbOrderLogDo) FindInBatch(batchSize int, fc func(tx gen.Dao, batch int) error) (results []*model.WinFbOrderLog, err error) {
	buf := make([]*model.WinFbOrderLog, 0, batchSize)
	err = w.DO.FindInBatches(&buf, batchSize, func(tx gen.Dao, batch int) error {
		defer func() { results = append(results, buf...) }()
		return fc(tx, batch)
	})
	return results, err
}

func (w winFbOrderLogDo) FindInBatches(result *[]*model.WinFbOrderLog, batchSize int, fc func(tx gen.Dao, batch int) error) error {
	return w.DO.FindInBatches(result, batchSize, fc)
}

func (w winFbOrderLogDo) Attrs(attrs ...field.AssignExpr) *winFbOrderLogDo {
	return w.withDO(w.DO.Attrs(attrs...))
}

func (w winFbOrderLogDo) Assign(attrs ...field.AssignExpr) *winFbOrderLogDo {
	return w.withDO(w.DO.Assign(attrs...))
}

func (w winFbOrderLogDo) Joins(fields ...field.RelationField) *winFbOrderLogDo {
	for _, _f := range fields {
		w = *w.withDO(w.DO.Joins(_f))
	}
	return &w
}

func (w winFbOrderLogDo) Preload(fields ...field.RelationField) *winFbOrderLogDo {
	for _, _f := range fields {
		w = *w.withDO(w.DO.Preload(_f))
	}
	return &w
}

func (w winFbOrderLogDo) FirstOrInit() (*model.WinFbOrderLog, error) {
	if result, err := w.DO.FirstOrInit(); err != nil {
		return nil, err
	} else {
		return result.(*model.WinFbOrderLog), nil
	}
}

func (w winFbOrderLogDo) FirstOrCreate() (*model.WinFbOrderLog, error) {
	if result, err := w.DO.FirstOrCreate(); err != nil {
		return nil, err
	} else {
		return result.(*model.WinFbOrderLog), nil
	}
}

func (w winFbOrderLogDo) FindByPage(offset int, limit int) (result []*model.WinFbOrderLog, count int64, err error) {
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

func (w winFbOrderLogDo) ScanByPage(result interface{}, offset int, limit int) (count int64, err error) {
	count, err = w.Count()
	if err != nil {
		return
	}

	err = w.Offset(offset).Limit(limit).Scan(result)
	return
}

func (w winFbOrderLogDo) Scan(result interface{}) (err error) {
	return w.DO.Scan(result)
}

func (w winFbOrderLogDo) Delete(models ...*model.WinFbOrderLog) (result gen.ResultInfo, err error) {
	return w.DO.Delete(models)
}

func (w *winFbOrderLogDo) withDO(do gen.Dao) *winFbOrderLogDo {
	w.DO = *do.(*gen.DO)
	return w
}
