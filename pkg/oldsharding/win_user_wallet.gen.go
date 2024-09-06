// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.

package oldsharding

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

func newWinUserWallet(db *gorm.DB, opts ...gen.DOOption) winUserWallet {
	_winUserWallet := winUserWallet{}

	_winUserWallet.winUserWalletDo.UseDB(db, opts...)
	_winUserWallet.winUserWalletDo.UseModel(&model.WinUserWallet{})

	tableName := _winUserWallet.winUserWalletDo.TableName()
	_winUserWallet.ALL = field.NewAsterisk(tableName)
	_winUserWallet.ID = field.NewInt32(tableName, "id")
	_winUserWallet.UID = field.NewInt32(tableName, "uid")
	_winUserWallet.Username = field.NewString(tableName, "username")
	_winUserWallet.MerchantID = field.NewInt32(tableName, "merchant_id")
	_winUserWallet.Category = field.NewInt32(tableName, "category")
	_winUserWallet.Currency = field.NewInt32(tableName, "currency")
	_winUserWallet.Coin = field.NewFloat64(tableName, "coin")
	_winUserWallet.Frozen = field.NewFloat64(tableName, "frozen")
	_winUserWallet.Version = field.NewInt32(tableName, "version")
	_winUserWallet.ModifyAt = field.NewInt64(tableName, "modify_at")
	_winUserWallet.CreatedAt = field.NewInt32(tableName, "created_at")
	_winUserWallet.UpdatedAt = field.NewInt32(tableName, "updated_at")

	_winUserWallet.fillFieldMap()

	return _winUserWallet
}

// winUserWallet 用户钱包表
type winUserWallet struct {
	winUserWalletDo

	ALL        field.Asterisk
	ID         field.Int32
	UID        field.Int32   // 用户id
	Username   field.String  // 用户名
	MerchantID field.Int32   // 商户id
	Category   field.Int32   // 钱包类型:支付/游戏/活动/佣金
	Currency   field.Int32   // 币种
	Coin       field.Float64 // 账户余额
	Frozen     field.Float64 // 冻结金额
	Version    field.Int32   // 版本号
	ModifyAt   field.Int64   // 13位时间戳
	CreatedAt  field.Int32
	UpdatedAt  field.Int32

	fieldMap map[string]field.Expr
}

func (w winUserWallet) Table(newTableName string) *winUserWallet {
	w.winUserWalletDo.UseTable(newTableName)
	return w.updateTableName(newTableName)
}

func (w winUserWallet) As(alias string) *winUserWallet {
	w.winUserWalletDo.DO = *(w.winUserWalletDo.As(alias).(*gen.DO))
	return w.updateTableName(alias)
}

func (w *winUserWallet) updateTableName(table string) *winUserWallet {
	w.ALL = field.NewAsterisk(table)
	w.ID = field.NewInt32(table, "id")
	w.UID = field.NewInt32(table, "uid")
	w.Username = field.NewString(table, "username")
	w.MerchantID = field.NewInt32(table, "merchant_id")
	w.Category = field.NewInt32(table, "category")
	w.Currency = field.NewInt32(table, "currency")
	w.Coin = field.NewFloat64(table, "coin")
	w.Frozen = field.NewFloat64(table, "frozen")
	w.Version = field.NewInt32(table, "version")
	w.ModifyAt = field.NewInt64(table, "modify_at")
	w.CreatedAt = field.NewInt32(table, "created_at")
	w.UpdatedAt = field.NewInt32(table, "updated_at")

	w.fillFieldMap()

	return w
}

func (w *winUserWallet) GetFieldByName(fieldName string) (field.OrderExpr, bool) {
	_f, ok := w.fieldMap[fieldName]
	if !ok || _f == nil {
		return nil, false
	}
	_oe, ok := _f.(field.OrderExpr)
	return _oe, ok
}

func (w *winUserWallet) fillFieldMap() {
	w.fieldMap = make(map[string]field.Expr, 12)
	w.fieldMap["id"] = w.ID
	w.fieldMap["uid"] = w.UID
	w.fieldMap["username"] = w.Username
	w.fieldMap["merchant_id"] = w.MerchantID
	w.fieldMap["category"] = w.Category
	w.fieldMap["currency"] = w.Currency
	w.fieldMap["coin"] = w.Coin
	w.fieldMap["frozen"] = w.Frozen
	w.fieldMap["version"] = w.Version
	w.fieldMap["modify_at"] = w.ModifyAt
	w.fieldMap["created_at"] = w.CreatedAt
	w.fieldMap["updated_at"] = w.UpdatedAt
}

func (w winUserWallet) clone(db *gorm.DB) winUserWallet {
	w.winUserWalletDo.ReplaceConnPool(db.Statement.ConnPool)
	return w
}

func (w winUserWallet) replaceDB(db *gorm.DB) winUserWallet {
	w.winUserWalletDo.ReplaceDB(db)
	return w
}

type winUserWalletDo struct{ gen.DO }

func (w winUserWalletDo) Debug() *winUserWalletDo {
	return w.withDO(w.DO.Debug())
}

func (w winUserWalletDo) WithContext(ctx context.Context) *winUserWalletDo {
	return w.withDO(w.DO.WithContext(ctx))
}

func (w winUserWalletDo) ReadDB() *winUserWalletDo {
	return w.Clauses(dbresolver.Read)
}

func (w winUserWalletDo) WriteDB() *winUserWalletDo {
	return w.Clauses(dbresolver.Write)
}

func (w winUserWalletDo) Session(config *gorm.Session) *winUserWalletDo {
	return w.withDO(w.DO.Session(config))
}

func (w winUserWalletDo) Clauses(conds ...clause.Expression) *winUserWalletDo {
	return w.withDO(w.DO.Clauses(conds...))
}

func (w winUserWalletDo) Returning(value interface{}, columns ...string) *winUserWalletDo {
	return w.withDO(w.DO.Returning(value, columns...))
}

func (w winUserWalletDo) Not(conds ...gen.Condition) *winUserWalletDo {
	return w.withDO(w.DO.Not(conds...))
}

func (w winUserWalletDo) Or(conds ...gen.Condition) *winUserWalletDo {
	return w.withDO(w.DO.Or(conds...))
}

func (w winUserWalletDo) Select(conds ...field.Expr) *winUserWalletDo {
	return w.withDO(w.DO.Select(conds...))
}

func (w winUserWalletDo) Where(conds ...gen.Condition) *winUserWalletDo {
	return w.withDO(w.DO.Where(conds...))
}

func (w winUserWalletDo) Order(conds ...field.Expr) *winUserWalletDo {
	return w.withDO(w.DO.Order(conds...))
}

func (w winUserWalletDo) Distinct(cols ...field.Expr) *winUserWalletDo {
	return w.withDO(w.DO.Distinct(cols...))
}

func (w winUserWalletDo) Omit(cols ...field.Expr) *winUserWalletDo {
	return w.withDO(w.DO.Omit(cols...))
}

func (w winUserWalletDo) Join(table schema.Tabler, on ...field.Expr) *winUserWalletDo {
	return w.withDO(w.DO.Join(table, on...))
}

func (w winUserWalletDo) LeftJoin(table schema.Tabler, on ...field.Expr) *winUserWalletDo {
	return w.withDO(w.DO.LeftJoin(table, on...))
}

func (w winUserWalletDo) RightJoin(table schema.Tabler, on ...field.Expr) *winUserWalletDo {
	return w.withDO(w.DO.RightJoin(table, on...))
}

func (w winUserWalletDo) Group(cols ...field.Expr) *winUserWalletDo {
	return w.withDO(w.DO.Group(cols...))
}

func (w winUserWalletDo) Having(conds ...gen.Condition) *winUserWalletDo {
	return w.withDO(w.DO.Having(conds...))
}

func (w winUserWalletDo) Limit(limit int) *winUserWalletDo {
	return w.withDO(w.DO.Limit(limit))
}

func (w winUserWalletDo) Offset(offset int) *winUserWalletDo {
	return w.withDO(w.DO.Offset(offset))
}

func (w winUserWalletDo) Scopes(funcs ...func(gen.Dao) gen.Dao) *winUserWalletDo {
	return w.withDO(w.DO.Scopes(funcs...))
}

func (w winUserWalletDo) Unscoped() *winUserWalletDo {
	return w.withDO(w.DO.Unscoped())
}

func (w winUserWalletDo) Create(values ...*model.WinUserWallet) error {
	if len(values) == 0 {
		return nil
	}
	return w.DO.Create(values)
}

func (w winUserWalletDo) CreateInBatches(values []*model.WinUserWallet, batchSize int) error {
	return w.DO.CreateInBatches(values, batchSize)
}

// Save : !!! underlying implementation is different with GORM
// The method is equivalent to executing the statement: db.Clauses(clause.OnConflict{UpdateAll: true}).Create(values)
func (w winUserWalletDo) Save(values ...*model.WinUserWallet) error {
	if len(values) == 0 {
		return nil
	}
	return w.DO.Save(values)
}

func (w winUserWalletDo) First() (*model.WinUserWallet, error) {
	if result, err := w.DO.First(); err != nil {
		return nil, err
	} else {
		return result.(*model.WinUserWallet), nil
	}
}

func (w winUserWalletDo) Take() (*model.WinUserWallet, error) {
	if result, err := w.DO.Take(); err != nil {
		return nil, err
	} else {
		return result.(*model.WinUserWallet), nil
	}
}

func (w winUserWalletDo) Last() (*model.WinUserWallet, error) {
	if result, err := w.DO.Last(); err != nil {
		return nil, err
	} else {
		return result.(*model.WinUserWallet), nil
	}
}

func (w winUserWalletDo) Find() ([]*model.WinUserWallet, error) {
	result, err := w.DO.Find()
	return result.([]*model.WinUserWallet), err
}

func (w winUserWalletDo) FindInBatch(batchSize int, fc func(tx gen.Dao, batch int) error) (results []*model.WinUserWallet, err error) {
	buf := make([]*model.WinUserWallet, 0, batchSize)
	err = w.DO.FindInBatches(&buf, batchSize, func(tx gen.Dao, batch int) error {
		defer func() { results = append(results, buf...) }()
		return fc(tx, batch)
	})
	return results, err
}

func (w winUserWalletDo) FindInBatches(result *[]*model.WinUserWallet, batchSize int, fc func(tx gen.Dao, batch int) error) error {
	return w.DO.FindInBatches(result, batchSize, fc)
}

func (w winUserWalletDo) Attrs(attrs ...field.AssignExpr) *winUserWalletDo {
	return w.withDO(w.DO.Attrs(attrs...))
}

func (w winUserWalletDo) Assign(attrs ...field.AssignExpr) *winUserWalletDo {
	return w.withDO(w.DO.Assign(attrs...))
}

func (w winUserWalletDo) Joins(fields ...field.RelationField) *winUserWalletDo {
	for _, _f := range fields {
		w = *w.withDO(w.DO.Joins(_f))
	}
	return &w
}

func (w winUserWalletDo) Preload(fields ...field.RelationField) *winUserWalletDo {
	for _, _f := range fields {
		w = *w.withDO(w.DO.Preload(_f))
	}
	return &w
}

func (w winUserWalletDo) FirstOrInit() (*model.WinUserWallet, error) {
	if result, err := w.DO.FirstOrInit(); err != nil {
		return nil, err
	} else {
		return result.(*model.WinUserWallet), nil
	}
}

func (w winUserWalletDo) FirstOrCreate() (*model.WinUserWallet, error) {
	if result, err := w.DO.FirstOrCreate(); err != nil {
		return nil, err
	} else {
		return result.(*model.WinUserWallet), nil
	}
}

func (w winUserWalletDo) FindByPage(offset int, limit int) (result []*model.WinUserWallet, count int64, err error) {
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

func (w winUserWalletDo) ScanByPage(result interface{}, offset int, limit int) (count int64, err error) {
	count, err = w.Count()
	if err != nil {
		return
	}

	err = w.Offset(offset).Limit(limit).Scan(result)
	return
}

func (w winUserWalletDo) Scan(result interface{}) (err error) {
	return w.DO.Scan(result)
}

func (w winUserWalletDo) Delete(models ...*model.WinUserWallet) (result gen.ResultInfo, err error) {
	return w.DO.Delete(models)
}

func (w *winUserWalletDo) withDO(do gen.Dao) *winUserWalletDo {
	w.DO = *do.(*gen.DO)
	return w
}
