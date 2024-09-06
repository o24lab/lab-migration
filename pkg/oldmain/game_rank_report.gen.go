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

func newGameRankReport(db *gorm.DB, opts ...gen.DOOption) gameRankReport {
	_gameRankReport := gameRankReport{}

	_gameRankReport.gameRankReportDo.UseDB(db, opts...)
	_gameRankReport.gameRankReportDo.UseModel(&model.GameRankReport{})

	tableName := _gameRankReport.gameRankReportDo.TableName()
	_gameRankReport.ALL = field.NewAsterisk(tableName)
	_gameRankReport.ID = field.NewInt32(tableName, "id")
	_gameRankReport.AgentID = field.NewInt32(tableName, "agent_id")
	_gameRankReport.SortType = field.NewInt32(tableName, "sort_type")
	_gameRankReport.Provider = field.NewString(tableName, "provider")
	_gameRankReport.ProviderSubtype = field.NewString(tableName, "provider_subtype")
	_gameRankReport.BetAmount = field.NewFloat64(tableName, "bet_amount")
	_gameRankReport.ProfitAmount = field.NewFloat64(tableName, "profit_amount")
	_gameRankReport.ReportDate = field.NewInt32(tableName, "report_date")
	_gameRankReport.CreatedAt = field.NewInt32(tableName, "created_at")
	_gameRankReport.UpdatedAt = field.NewInt32(tableName, "updated_at")

	_gameRankReport.fillFieldMap()

	return _gameRankReport
}

// gameRankReport 代理:遊戲排行報表
type gameRankReport struct {
	gameRankReportDo

	ALL             field.Asterisk
	ID              field.Int32   // id
	AgentID         field.Int32   // 代理uid
	SortType        field.Int32   // 0:盈虧,1:投注
	Provider        field.String  // 遊戲商
	ProviderSubtype field.String  // 子遊戲
	BetAmount       field.Float64 // 投注金额
	ProfitAmount    field.Float64 // 盈虧金额
	ReportDate      field.Int32   // 報表日期
	CreatedAt       field.Int32   // 創建時間
	UpdatedAt       field.Int32   // 更新時間

	fieldMap map[string]field.Expr
}

func (g gameRankReport) Table(newTableName string) *gameRankReport {
	g.gameRankReportDo.UseTable(newTableName)
	return g.updateTableName(newTableName)
}

func (g gameRankReport) As(alias string) *gameRankReport {
	g.gameRankReportDo.DO = *(g.gameRankReportDo.As(alias).(*gen.DO))
	return g.updateTableName(alias)
}

func (g *gameRankReport) updateTableName(table string) *gameRankReport {
	g.ALL = field.NewAsterisk(table)
	g.ID = field.NewInt32(table, "id")
	g.AgentID = field.NewInt32(table, "agent_id")
	g.SortType = field.NewInt32(table, "sort_type")
	g.Provider = field.NewString(table, "provider")
	g.ProviderSubtype = field.NewString(table, "provider_subtype")
	g.BetAmount = field.NewFloat64(table, "bet_amount")
	g.ProfitAmount = field.NewFloat64(table, "profit_amount")
	g.ReportDate = field.NewInt32(table, "report_date")
	g.CreatedAt = field.NewInt32(table, "created_at")
	g.UpdatedAt = field.NewInt32(table, "updated_at")

	g.fillFieldMap()

	return g
}

func (g *gameRankReport) GetFieldByName(fieldName string) (field.OrderExpr, bool) {
	_f, ok := g.fieldMap[fieldName]
	if !ok || _f == nil {
		return nil, false
	}
	_oe, ok := _f.(field.OrderExpr)
	return _oe, ok
}

func (g *gameRankReport) fillFieldMap() {
	g.fieldMap = make(map[string]field.Expr, 10)
	g.fieldMap["id"] = g.ID
	g.fieldMap["agent_id"] = g.AgentID
	g.fieldMap["sort_type"] = g.SortType
	g.fieldMap["provider"] = g.Provider
	g.fieldMap["provider_subtype"] = g.ProviderSubtype
	g.fieldMap["bet_amount"] = g.BetAmount
	g.fieldMap["profit_amount"] = g.ProfitAmount
	g.fieldMap["report_date"] = g.ReportDate
	g.fieldMap["created_at"] = g.CreatedAt
	g.fieldMap["updated_at"] = g.UpdatedAt
}

func (g gameRankReport) clone(db *gorm.DB) gameRankReport {
	g.gameRankReportDo.ReplaceConnPool(db.Statement.ConnPool)
	return g
}

func (g gameRankReport) replaceDB(db *gorm.DB) gameRankReport {
	g.gameRankReportDo.ReplaceDB(db)
	return g
}

type gameRankReportDo struct{ gen.DO }

func (g gameRankReportDo) Debug() *gameRankReportDo {
	return g.withDO(g.DO.Debug())
}

func (g gameRankReportDo) WithContext(ctx context.Context) *gameRankReportDo {
	return g.withDO(g.DO.WithContext(ctx))
}

func (g gameRankReportDo) ReadDB() *gameRankReportDo {
	return g.Clauses(dbresolver.Read)
}

func (g gameRankReportDo) WriteDB() *gameRankReportDo {
	return g.Clauses(dbresolver.Write)
}

func (g gameRankReportDo) Session(config *gorm.Session) *gameRankReportDo {
	return g.withDO(g.DO.Session(config))
}

func (g gameRankReportDo) Clauses(conds ...clause.Expression) *gameRankReportDo {
	return g.withDO(g.DO.Clauses(conds...))
}

func (g gameRankReportDo) Returning(value interface{}, columns ...string) *gameRankReportDo {
	return g.withDO(g.DO.Returning(value, columns...))
}

func (g gameRankReportDo) Not(conds ...gen.Condition) *gameRankReportDo {
	return g.withDO(g.DO.Not(conds...))
}

func (g gameRankReportDo) Or(conds ...gen.Condition) *gameRankReportDo {
	return g.withDO(g.DO.Or(conds...))
}

func (g gameRankReportDo) Select(conds ...field.Expr) *gameRankReportDo {
	return g.withDO(g.DO.Select(conds...))
}

func (g gameRankReportDo) Where(conds ...gen.Condition) *gameRankReportDo {
	return g.withDO(g.DO.Where(conds...))
}

func (g gameRankReportDo) Order(conds ...field.Expr) *gameRankReportDo {
	return g.withDO(g.DO.Order(conds...))
}

func (g gameRankReportDo) Distinct(cols ...field.Expr) *gameRankReportDo {
	return g.withDO(g.DO.Distinct(cols...))
}

func (g gameRankReportDo) Omit(cols ...field.Expr) *gameRankReportDo {
	return g.withDO(g.DO.Omit(cols...))
}

func (g gameRankReportDo) Join(table schema.Tabler, on ...field.Expr) *gameRankReportDo {
	return g.withDO(g.DO.Join(table, on...))
}

func (g gameRankReportDo) LeftJoin(table schema.Tabler, on ...field.Expr) *gameRankReportDo {
	return g.withDO(g.DO.LeftJoin(table, on...))
}

func (g gameRankReportDo) RightJoin(table schema.Tabler, on ...field.Expr) *gameRankReportDo {
	return g.withDO(g.DO.RightJoin(table, on...))
}

func (g gameRankReportDo) Group(cols ...field.Expr) *gameRankReportDo {
	return g.withDO(g.DO.Group(cols...))
}

func (g gameRankReportDo) Having(conds ...gen.Condition) *gameRankReportDo {
	return g.withDO(g.DO.Having(conds...))
}

func (g gameRankReportDo) Limit(limit int) *gameRankReportDo {
	return g.withDO(g.DO.Limit(limit))
}

func (g gameRankReportDo) Offset(offset int) *gameRankReportDo {
	return g.withDO(g.DO.Offset(offset))
}

func (g gameRankReportDo) Scopes(funcs ...func(gen.Dao) gen.Dao) *gameRankReportDo {
	return g.withDO(g.DO.Scopes(funcs...))
}

func (g gameRankReportDo) Unscoped() *gameRankReportDo {
	return g.withDO(g.DO.Unscoped())
}

func (g gameRankReportDo) Create(values ...*model.GameRankReport) error {
	if len(values) == 0 {
		return nil
	}
	return g.DO.Create(values)
}

func (g gameRankReportDo) CreateInBatches(values []*model.GameRankReport, batchSize int) error {
	return g.DO.CreateInBatches(values, batchSize)
}

// Save : !!! underlying implementation is different with GORM
// The method is equivalent to executing the statement: db.Clauses(clause.OnConflict{UpdateAll: true}).Create(values)
func (g gameRankReportDo) Save(values ...*model.GameRankReport) error {
	if len(values) == 0 {
		return nil
	}
	return g.DO.Save(values)
}

func (g gameRankReportDo) First() (*model.GameRankReport, error) {
	if result, err := g.DO.First(); err != nil {
		return nil, err
	} else {
		return result.(*model.GameRankReport), nil
	}
}

func (g gameRankReportDo) Take() (*model.GameRankReport, error) {
	if result, err := g.DO.Take(); err != nil {
		return nil, err
	} else {
		return result.(*model.GameRankReport), nil
	}
}

func (g gameRankReportDo) Last() (*model.GameRankReport, error) {
	if result, err := g.DO.Last(); err != nil {
		return nil, err
	} else {
		return result.(*model.GameRankReport), nil
	}
}

func (g gameRankReportDo) Find() ([]*model.GameRankReport, error) {
	result, err := g.DO.Find()
	return result.([]*model.GameRankReport), err
}

func (g gameRankReportDo) FindInBatch(batchSize int, fc func(tx gen.Dao, batch int) error) (results []*model.GameRankReport, err error) {
	buf := make([]*model.GameRankReport, 0, batchSize)
	err = g.DO.FindInBatches(&buf, batchSize, func(tx gen.Dao, batch int) error {
		defer func() { results = append(results, buf...) }()
		return fc(tx, batch)
	})
	return results, err
}

func (g gameRankReportDo) FindInBatches(result *[]*model.GameRankReport, batchSize int, fc func(tx gen.Dao, batch int) error) error {
	return g.DO.FindInBatches(result, batchSize, fc)
}

func (g gameRankReportDo) Attrs(attrs ...field.AssignExpr) *gameRankReportDo {
	return g.withDO(g.DO.Attrs(attrs...))
}

func (g gameRankReportDo) Assign(attrs ...field.AssignExpr) *gameRankReportDo {
	return g.withDO(g.DO.Assign(attrs...))
}

func (g gameRankReportDo) Joins(fields ...field.RelationField) *gameRankReportDo {
	for _, _f := range fields {
		g = *g.withDO(g.DO.Joins(_f))
	}
	return &g
}

func (g gameRankReportDo) Preload(fields ...field.RelationField) *gameRankReportDo {
	for _, _f := range fields {
		g = *g.withDO(g.DO.Preload(_f))
	}
	return &g
}

func (g gameRankReportDo) FirstOrInit() (*model.GameRankReport, error) {
	if result, err := g.DO.FirstOrInit(); err != nil {
		return nil, err
	} else {
		return result.(*model.GameRankReport), nil
	}
}

func (g gameRankReportDo) FirstOrCreate() (*model.GameRankReport, error) {
	if result, err := g.DO.FirstOrCreate(); err != nil {
		return nil, err
	} else {
		return result.(*model.GameRankReport), nil
	}
}

func (g gameRankReportDo) FindByPage(offset int, limit int) (result []*model.GameRankReport, count int64, err error) {
	result, err = g.Offset(offset).Limit(limit).Find()
	if err != nil {
		return
	}

	if size := len(result); 0 < limit && 0 < size && size < limit {
		count = int64(size + offset)
		return
	}

	count, err = g.Offset(-1).Limit(-1).Count()
	return
}

func (g gameRankReportDo) ScanByPage(result interface{}, offset int, limit int) (count int64, err error) {
	count, err = g.Count()
	if err != nil {
		return
	}

	err = g.Offset(offset).Limit(limit).Scan(result)
	return
}

func (g gameRankReportDo) Scan(result interface{}) (err error) {
	return g.DO.Scan(result)
}

func (g gameRankReportDo) Delete(models ...*model.GameRankReport) (result gen.ResultInfo, err error) {
	return g.DO.Delete(models)
}

func (g *gameRankReportDo) withDO(do gen.Dao) *gameRankReportDo {
	g.DO = *do.(*gen.DO)
	return g
}
