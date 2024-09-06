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

func newWinSpecialRecordStakeGame(db *gorm.DB, opts ...gen.DOOption) winSpecialRecordStakeGame {
	_winSpecialRecordStakeGame := winSpecialRecordStakeGame{}

	_winSpecialRecordStakeGame.winSpecialRecordStakeGameDo.UseDB(db, opts...)
	_winSpecialRecordStakeGame.winSpecialRecordStakeGameDo.UseModel(&model.WinSpecialRecordStakeGame{})

	tableName := _winSpecialRecordStakeGame.winSpecialRecordStakeGameDo.TableName()
	_winSpecialRecordStakeGame.ALL = field.NewAsterisk(tableName)
	_winSpecialRecordStakeGame.ID = field.NewInt64(tableName, "id")
	_winSpecialRecordStakeGame.GameTypeID = field.NewString(tableName, "game_type_id")
	_winSpecialRecordStakeGame.GameProviderSubtypeID = field.NewString(tableName, "game_provider_subtype_id")
	_winSpecialRecordStakeGame.GameListID = field.NewString(tableName, "game_list_id")
	_winSpecialRecordStakeGame.CreatedAt = field.NewInt32(tableName, "created_at")
	_winSpecialRecordStakeGame.UpdatedAt = field.NewInt32(tableName, "updated_at")
	_winSpecialRecordStakeGame.UpdatedUser = field.NewString(tableName, "updated_user")

	_winSpecialRecordStakeGame.fillFieldMap()

	return _winSpecialRecordStakeGame
}

// winSpecialRecordStakeGame 打码量特殊算法游戏配置表
type winSpecialRecordStakeGame struct {
	winSpecialRecordStakeGameDo

	ALL                   field.Asterisk
	ID                    field.Int64  // 自增主键
	GameTypeID            field.String // 游戏类型：0-全部，默认字符串ex:id1,id2 逗号分隔
	GameProviderSubtypeID field.String // 游戏厂商：0-全部，默认字符串ex:id1,id2 逗号分隔
	GameListID            field.String // 子游戏类型：0-全部，默认字符串ex:id1,id2 逗号分隔
	CreatedAt             field.Int32  // 创建时间
	UpdatedAt             field.Int32  // 修改时间
	UpdatedUser           field.String // 最后修改人

	fieldMap map[string]field.Expr
}

func (w winSpecialRecordStakeGame) Table(newTableName string) *winSpecialRecordStakeGame {
	w.winSpecialRecordStakeGameDo.UseTable(newTableName)
	return w.updateTableName(newTableName)
}

func (w winSpecialRecordStakeGame) As(alias string) *winSpecialRecordStakeGame {
	w.winSpecialRecordStakeGameDo.DO = *(w.winSpecialRecordStakeGameDo.As(alias).(*gen.DO))
	return w.updateTableName(alias)
}

func (w *winSpecialRecordStakeGame) updateTableName(table string) *winSpecialRecordStakeGame {
	w.ALL = field.NewAsterisk(table)
	w.ID = field.NewInt64(table, "id")
	w.GameTypeID = field.NewString(table, "game_type_id")
	w.GameProviderSubtypeID = field.NewString(table, "game_provider_subtype_id")
	w.GameListID = field.NewString(table, "game_list_id")
	w.CreatedAt = field.NewInt32(table, "created_at")
	w.UpdatedAt = field.NewInt32(table, "updated_at")
	w.UpdatedUser = field.NewString(table, "updated_user")

	w.fillFieldMap()

	return w
}

func (w *winSpecialRecordStakeGame) GetFieldByName(fieldName string) (field.OrderExpr, bool) {
	_f, ok := w.fieldMap[fieldName]
	if !ok || _f == nil {
		return nil, false
	}
	_oe, ok := _f.(field.OrderExpr)
	return _oe, ok
}

func (w *winSpecialRecordStakeGame) fillFieldMap() {
	w.fieldMap = make(map[string]field.Expr, 7)
	w.fieldMap["id"] = w.ID
	w.fieldMap["game_type_id"] = w.GameTypeID
	w.fieldMap["game_provider_subtype_id"] = w.GameProviderSubtypeID
	w.fieldMap["game_list_id"] = w.GameListID
	w.fieldMap["created_at"] = w.CreatedAt
	w.fieldMap["updated_at"] = w.UpdatedAt
	w.fieldMap["updated_user"] = w.UpdatedUser
}

func (w winSpecialRecordStakeGame) clone(db *gorm.DB) winSpecialRecordStakeGame {
	w.winSpecialRecordStakeGameDo.ReplaceConnPool(db.Statement.ConnPool)
	return w
}

func (w winSpecialRecordStakeGame) replaceDB(db *gorm.DB) winSpecialRecordStakeGame {
	w.winSpecialRecordStakeGameDo.ReplaceDB(db)
	return w
}

type winSpecialRecordStakeGameDo struct{ gen.DO }

func (w winSpecialRecordStakeGameDo) Debug() *winSpecialRecordStakeGameDo {
	return w.withDO(w.DO.Debug())
}

func (w winSpecialRecordStakeGameDo) WithContext(ctx context.Context) *winSpecialRecordStakeGameDo {
	return w.withDO(w.DO.WithContext(ctx))
}

func (w winSpecialRecordStakeGameDo) ReadDB() *winSpecialRecordStakeGameDo {
	return w.Clauses(dbresolver.Read)
}

func (w winSpecialRecordStakeGameDo) WriteDB() *winSpecialRecordStakeGameDo {
	return w.Clauses(dbresolver.Write)
}

func (w winSpecialRecordStakeGameDo) Session(config *gorm.Session) *winSpecialRecordStakeGameDo {
	return w.withDO(w.DO.Session(config))
}

func (w winSpecialRecordStakeGameDo) Clauses(conds ...clause.Expression) *winSpecialRecordStakeGameDo {
	return w.withDO(w.DO.Clauses(conds...))
}

func (w winSpecialRecordStakeGameDo) Returning(value interface{}, columns ...string) *winSpecialRecordStakeGameDo {
	return w.withDO(w.DO.Returning(value, columns...))
}

func (w winSpecialRecordStakeGameDo) Not(conds ...gen.Condition) *winSpecialRecordStakeGameDo {
	return w.withDO(w.DO.Not(conds...))
}

func (w winSpecialRecordStakeGameDo) Or(conds ...gen.Condition) *winSpecialRecordStakeGameDo {
	return w.withDO(w.DO.Or(conds...))
}

func (w winSpecialRecordStakeGameDo) Select(conds ...field.Expr) *winSpecialRecordStakeGameDo {
	return w.withDO(w.DO.Select(conds...))
}

func (w winSpecialRecordStakeGameDo) Where(conds ...gen.Condition) *winSpecialRecordStakeGameDo {
	return w.withDO(w.DO.Where(conds...))
}

func (w winSpecialRecordStakeGameDo) Order(conds ...field.Expr) *winSpecialRecordStakeGameDo {
	return w.withDO(w.DO.Order(conds...))
}

func (w winSpecialRecordStakeGameDo) Distinct(cols ...field.Expr) *winSpecialRecordStakeGameDo {
	return w.withDO(w.DO.Distinct(cols...))
}

func (w winSpecialRecordStakeGameDo) Omit(cols ...field.Expr) *winSpecialRecordStakeGameDo {
	return w.withDO(w.DO.Omit(cols...))
}

func (w winSpecialRecordStakeGameDo) Join(table schema.Tabler, on ...field.Expr) *winSpecialRecordStakeGameDo {
	return w.withDO(w.DO.Join(table, on...))
}

func (w winSpecialRecordStakeGameDo) LeftJoin(table schema.Tabler, on ...field.Expr) *winSpecialRecordStakeGameDo {
	return w.withDO(w.DO.LeftJoin(table, on...))
}

func (w winSpecialRecordStakeGameDo) RightJoin(table schema.Tabler, on ...field.Expr) *winSpecialRecordStakeGameDo {
	return w.withDO(w.DO.RightJoin(table, on...))
}

func (w winSpecialRecordStakeGameDo) Group(cols ...field.Expr) *winSpecialRecordStakeGameDo {
	return w.withDO(w.DO.Group(cols...))
}

func (w winSpecialRecordStakeGameDo) Having(conds ...gen.Condition) *winSpecialRecordStakeGameDo {
	return w.withDO(w.DO.Having(conds...))
}

func (w winSpecialRecordStakeGameDo) Limit(limit int) *winSpecialRecordStakeGameDo {
	return w.withDO(w.DO.Limit(limit))
}

func (w winSpecialRecordStakeGameDo) Offset(offset int) *winSpecialRecordStakeGameDo {
	return w.withDO(w.DO.Offset(offset))
}

func (w winSpecialRecordStakeGameDo) Scopes(funcs ...func(gen.Dao) gen.Dao) *winSpecialRecordStakeGameDo {
	return w.withDO(w.DO.Scopes(funcs...))
}

func (w winSpecialRecordStakeGameDo) Unscoped() *winSpecialRecordStakeGameDo {
	return w.withDO(w.DO.Unscoped())
}

func (w winSpecialRecordStakeGameDo) Create(values ...*model.WinSpecialRecordStakeGame) error {
	if len(values) == 0 {
		return nil
	}
	return w.DO.Create(values)
}

func (w winSpecialRecordStakeGameDo) CreateInBatches(values []*model.WinSpecialRecordStakeGame, batchSize int) error {
	return w.DO.CreateInBatches(values, batchSize)
}

// Save : !!! underlying implementation is different with GORM
// The method is equivalent to executing the statement: db.Clauses(clause.OnConflict{UpdateAll: true}).Create(values)
func (w winSpecialRecordStakeGameDo) Save(values ...*model.WinSpecialRecordStakeGame) error {
	if len(values) == 0 {
		return nil
	}
	return w.DO.Save(values)
}

func (w winSpecialRecordStakeGameDo) First() (*model.WinSpecialRecordStakeGame, error) {
	if result, err := w.DO.First(); err != nil {
		return nil, err
	} else {
		return result.(*model.WinSpecialRecordStakeGame), nil
	}
}

func (w winSpecialRecordStakeGameDo) Take() (*model.WinSpecialRecordStakeGame, error) {
	if result, err := w.DO.Take(); err != nil {
		return nil, err
	} else {
		return result.(*model.WinSpecialRecordStakeGame), nil
	}
}

func (w winSpecialRecordStakeGameDo) Last() (*model.WinSpecialRecordStakeGame, error) {
	if result, err := w.DO.Last(); err != nil {
		return nil, err
	} else {
		return result.(*model.WinSpecialRecordStakeGame), nil
	}
}

func (w winSpecialRecordStakeGameDo) Find() ([]*model.WinSpecialRecordStakeGame, error) {
	result, err := w.DO.Find()
	return result.([]*model.WinSpecialRecordStakeGame), err
}

func (w winSpecialRecordStakeGameDo) FindInBatch(batchSize int, fc func(tx gen.Dao, batch int) error) (results []*model.WinSpecialRecordStakeGame, err error) {
	buf := make([]*model.WinSpecialRecordStakeGame, 0, batchSize)
	err = w.DO.FindInBatches(&buf, batchSize, func(tx gen.Dao, batch int) error {
		defer func() { results = append(results, buf...) }()
		return fc(tx, batch)
	})
	return results, err
}

func (w winSpecialRecordStakeGameDo) FindInBatches(result *[]*model.WinSpecialRecordStakeGame, batchSize int, fc func(tx gen.Dao, batch int) error) error {
	return w.DO.FindInBatches(result, batchSize, fc)
}

func (w winSpecialRecordStakeGameDo) Attrs(attrs ...field.AssignExpr) *winSpecialRecordStakeGameDo {
	return w.withDO(w.DO.Attrs(attrs...))
}

func (w winSpecialRecordStakeGameDo) Assign(attrs ...field.AssignExpr) *winSpecialRecordStakeGameDo {
	return w.withDO(w.DO.Assign(attrs...))
}

func (w winSpecialRecordStakeGameDo) Joins(fields ...field.RelationField) *winSpecialRecordStakeGameDo {
	for _, _f := range fields {
		w = *w.withDO(w.DO.Joins(_f))
	}
	return &w
}

func (w winSpecialRecordStakeGameDo) Preload(fields ...field.RelationField) *winSpecialRecordStakeGameDo {
	for _, _f := range fields {
		w = *w.withDO(w.DO.Preload(_f))
	}
	return &w
}

func (w winSpecialRecordStakeGameDo) FirstOrInit() (*model.WinSpecialRecordStakeGame, error) {
	if result, err := w.DO.FirstOrInit(); err != nil {
		return nil, err
	} else {
		return result.(*model.WinSpecialRecordStakeGame), nil
	}
}

func (w winSpecialRecordStakeGameDo) FirstOrCreate() (*model.WinSpecialRecordStakeGame, error) {
	if result, err := w.DO.FirstOrCreate(); err != nil {
		return nil, err
	} else {
		return result.(*model.WinSpecialRecordStakeGame), nil
	}
}

func (w winSpecialRecordStakeGameDo) FindByPage(offset int, limit int) (result []*model.WinSpecialRecordStakeGame, count int64, err error) {
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

func (w winSpecialRecordStakeGameDo) ScanByPage(result interface{}, offset int, limit int) (count int64, err error) {
	count, err = w.Count()
	if err != nil {
		return
	}

	err = w.Offset(offset).Limit(limit).Scan(result)
	return
}

func (w winSpecialRecordStakeGameDo) Scan(result interface{}) (err error) {
	return w.DO.Scan(result)
}

func (w winSpecialRecordStakeGameDo) Delete(models ...*model.WinSpecialRecordStakeGame) (result gen.ResultInfo, err error) {
	return w.DO.Delete(models)
}

func (w *winSpecialRecordStakeGameDo) withDO(do gen.Dao) *winSpecialRecordStakeGameDo {
	w.DO = *do.(*gen.DO)
	return w
}
