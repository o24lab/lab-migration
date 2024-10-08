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

func newXxlJobLogReport(db *gorm.DB, opts ...gen.DOOption) xxlJobLogReport {
	_xxlJobLogReport := xxlJobLogReport{}

	_xxlJobLogReport.xxlJobLogReportDo.UseDB(db, opts...)
	_xxlJobLogReport.xxlJobLogReportDo.UseModel(&model.XxlJobLogReport{})

	tableName := _xxlJobLogReport.xxlJobLogReportDo.TableName()
	_xxlJobLogReport.ALL = field.NewAsterisk(tableName)
	_xxlJobLogReport.ID = field.NewInt32(tableName, "id")
	_xxlJobLogReport.TriggerDay = field.NewTime(tableName, "trigger_day")
	_xxlJobLogReport.RunningCount = field.NewInt32(tableName, "running_count")
	_xxlJobLogReport.SucCount = field.NewInt32(tableName, "suc_count")
	_xxlJobLogReport.FailCount = field.NewInt32(tableName, "fail_count")
	_xxlJobLogReport.UpdateTime = field.NewTime(tableName, "update_time")

	_xxlJobLogReport.fillFieldMap()

	return _xxlJobLogReport
}

type xxlJobLogReport struct {
	xxlJobLogReportDo

	ALL          field.Asterisk
	ID           field.Int32
	TriggerDay   field.Time  // 调度-时间
	RunningCount field.Int32 // 运行中-日志数量
	SucCount     field.Int32 // 执行成功-日志数量
	FailCount    field.Int32 // 执行失败-日志数量
	UpdateTime   field.Time

	fieldMap map[string]field.Expr
}

func (x xxlJobLogReport) Table(newTableName string) *xxlJobLogReport {
	x.xxlJobLogReportDo.UseTable(newTableName)
	return x.updateTableName(newTableName)
}

func (x xxlJobLogReport) As(alias string) *xxlJobLogReport {
	x.xxlJobLogReportDo.DO = *(x.xxlJobLogReportDo.As(alias).(*gen.DO))
	return x.updateTableName(alias)
}

func (x *xxlJobLogReport) updateTableName(table string) *xxlJobLogReport {
	x.ALL = field.NewAsterisk(table)
	x.ID = field.NewInt32(table, "id")
	x.TriggerDay = field.NewTime(table, "trigger_day")
	x.RunningCount = field.NewInt32(table, "running_count")
	x.SucCount = field.NewInt32(table, "suc_count")
	x.FailCount = field.NewInt32(table, "fail_count")
	x.UpdateTime = field.NewTime(table, "update_time")

	x.fillFieldMap()

	return x
}

func (x *xxlJobLogReport) GetFieldByName(fieldName string) (field.OrderExpr, bool) {
	_f, ok := x.fieldMap[fieldName]
	if !ok || _f == nil {
		return nil, false
	}
	_oe, ok := _f.(field.OrderExpr)
	return _oe, ok
}

func (x *xxlJobLogReport) fillFieldMap() {
	x.fieldMap = make(map[string]field.Expr, 6)
	x.fieldMap["id"] = x.ID
	x.fieldMap["trigger_day"] = x.TriggerDay
	x.fieldMap["running_count"] = x.RunningCount
	x.fieldMap["suc_count"] = x.SucCount
	x.fieldMap["fail_count"] = x.FailCount
	x.fieldMap["update_time"] = x.UpdateTime
}

func (x xxlJobLogReport) clone(db *gorm.DB) xxlJobLogReport {
	x.xxlJobLogReportDo.ReplaceConnPool(db.Statement.ConnPool)
	return x
}

func (x xxlJobLogReport) replaceDB(db *gorm.DB) xxlJobLogReport {
	x.xxlJobLogReportDo.ReplaceDB(db)
	return x
}

type xxlJobLogReportDo struct{ gen.DO }

func (x xxlJobLogReportDo) Debug() *xxlJobLogReportDo {
	return x.withDO(x.DO.Debug())
}

func (x xxlJobLogReportDo) WithContext(ctx context.Context) *xxlJobLogReportDo {
	return x.withDO(x.DO.WithContext(ctx))
}

func (x xxlJobLogReportDo) ReadDB() *xxlJobLogReportDo {
	return x.Clauses(dbresolver.Read)
}

func (x xxlJobLogReportDo) WriteDB() *xxlJobLogReportDo {
	return x.Clauses(dbresolver.Write)
}

func (x xxlJobLogReportDo) Session(config *gorm.Session) *xxlJobLogReportDo {
	return x.withDO(x.DO.Session(config))
}

func (x xxlJobLogReportDo) Clauses(conds ...clause.Expression) *xxlJobLogReportDo {
	return x.withDO(x.DO.Clauses(conds...))
}

func (x xxlJobLogReportDo) Returning(value interface{}, columns ...string) *xxlJobLogReportDo {
	return x.withDO(x.DO.Returning(value, columns...))
}

func (x xxlJobLogReportDo) Not(conds ...gen.Condition) *xxlJobLogReportDo {
	return x.withDO(x.DO.Not(conds...))
}

func (x xxlJobLogReportDo) Or(conds ...gen.Condition) *xxlJobLogReportDo {
	return x.withDO(x.DO.Or(conds...))
}

func (x xxlJobLogReportDo) Select(conds ...field.Expr) *xxlJobLogReportDo {
	return x.withDO(x.DO.Select(conds...))
}

func (x xxlJobLogReportDo) Where(conds ...gen.Condition) *xxlJobLogReportDo {
	return x.withDO(x.DO.Where(conds...))
}

func (x xxlJobLogReportDo) Order(conds ...field.Expr) *xxlJobLogReportDo {
	return x.withDO(x.DO.Order(conds...))
}

func (x xxlJobLogReportDo) Distinct(cols ...field.Expr) *xxlJobLogReportDo {
	return x.withDO(x.DO.Distinct(cols...))
}

func (x xxlJobLogReportDo) Omit(cols ...field.Expr) *xxlJobLogReportDo {
	return x.withDO(x.DO.Omit(cols...))
}

func (x xxlJobLogReportDo) Join(table schema.Tabler, on ...field.Expr) *xxlJobLogReportDo {
	return x.withDO(x.DO.Join(table, on...))
}

func (x xxlJobLogReportDo) LeftJoin(table schema.Tabler, on ...field.Expr) *xxlJobLogReportDo {
	return x.withDO(x.DO.LeftJoin(table, on...))
}

func (x xxlJobLogReportDo) RightJoin(table schema.Tabler, on ...field.Expr) *xxlJobLogReportDo {
	return x.withDO(x.DO.RightJoin(table, on...))
}

func (x xxlJobLogReportDo) Group(cols ...field.Expr) *xxlJobLogReportDo {
	return x.withDO(x.DO.Group(cols...))
}

func (x xxlJobLogReportDo) Having(conds ...gen.Condition) *xxlJobLogReportDo {
	return x.withDO(x.DO.Having(conds...))
}

func (x xxlJobLogReportDo) Limit(limit int) *xxlJobLogReportDo {
	return x.withDO(x.DO.Limit(limit))
}

func (x xxlJobLogReportDo) Offset(offset int) *xxlJobLogReportDo {
	return x.withDO(x.DO.Offset(offset))
}

func (x xxlJobLogReportDo) Scopes(funcs ...func(gen.Dao) gen.Dao) *xxlJobLogReportDo {
	return x.withDO(x.DO.Scopes(funcs...))
}

func (x xxlJobLogReportDo) Unscoped() *xxlJobLogReportDo {
	return x.withDO(x.DO.Unscoped())
}

func (x xxlJobLogReportDo) Create(values ...*model.XxlJobLogReport) error {
	if len(values) == 0 {
		return nil
	}
	return x.DO.Create(values)
}

func (x xxlJobLogReportDo) CreateInBatches(values []*model.XxlJobLogReport, batchSize int) error {
	return x.DO.CreateInBatches(values, batchSize)
}

// Save : !!! underlying implementation is different with GORM
// The method is equivalent to executing the statement: db.Clauses(clause.OnConflict{UpdateAll: true}).Create(values)
func (x xxlJobLogReportDo) Save(values ...*model.XxlJobLogReport) error {
	if len(values) == 0 {
		return nil
	}
	return x.DO.Save(values)
}

func (x xxlJobLogReportDo) First() (*model.XxlJobLogReport, error) {
	if result, err := x.DO.First(); err != nil {
		return nil, err
	} else {
		return result.(*model.XxlJobLogReport), nil
	}
}

func (x xxlJobLogReportDo) Take() (*model.XxlJobLogReport, error) {
	if result, err := x.DO.Take(); err != nil {
		return nil, err
	} else {
		return result.(*model.XxlJobLogReport), nil
	}
}

func (x xxlJobLogReportDo) Last() (*model.XxlJobLogReport, error) {
	if result, err := x.DO.Last(); err != nil {
		return nil, err
	} else {
		return result.(*model.XxlJobLogReport), nil
	}
}

func (x xxlJobLogReportDo) Find() ([]*model.XxlJobLogReport, error) {
	result, err := x.DO.Find()
	return result.([]*model.XxlJobLogReport), err
}

func (x xxlJobLogReportDo) FindInBatch(batchSize int, fc func(tx gen.Dao, batch int) error) (results []*model.XxlJobLogReport, err error) {
	buf := make([]*model.XxlJobLogReport, 0, batchSize)
	err = x.DO.FindInBatches(&buf, batchSize, func(tx gen.Dao, batch int) error {
		defer func() { results = append(results, buf...) }()
		return fc(tx, batch)
	})
	return results, err
}

func (x xxlJobLogReportDo) FindInBatches(result *[]*model.XxlJobLogReport, batchSize int, fc func(tx gen.Dao, batch int) error) error {
	return x.DO.FindInBatches(result, batchSize, fc)
}

func (x xxlJobLogReportDo) Attrs(attrs ...field.AssignExpr) *xxlJobLogReportDo {
	return x.withDO(x.DO.Attrs(attrs...))
}

func (x xxlJobLogReportDo) Assign(attrs ...field.AssignExpr) *xxlJobLogReportDo {
	return x.withDO(x.DO.Assign(attrs...))
}

func (x xxlJobLogReportDo) Joins(fields ...field.RelationField) *xxlJobLogReportDo {
	for _, _f := range fields {
		x = *x.withDO(x.DO.Joins(_f))
	}
	return &x
}

func (x xxlJobLogReportDo) Preload(fields ...field.RelationField) *xxlJobLogReportDo {
	for _, _f := range fields {
		x = *x.withDO(x.DO.Preload(_f))
	}
	return &x
}

func (x xxlJobLogReportDo) FirstOrInit() (*model.XxlJobLogReport, error) {
	if result, err := x.DO.FirstOrInit(); err != nil {
		return nil, err
	} else {
		return result.(*model.XxlJobLogReport), nil
	}
}

func (x xxlJobLogReportDo) FirstOrCreate() (*model.XxlJobLogReport, error) {
	if result, err := x.DO.FirstOrCreate(); err != nil {
		return nil, err
	} else {
		return result.(*model.XxlJobLogReport), nil
	}
}

func (x xxlJobLogReportDo) FindByPage(offset int, limit int) (result []*model.XxlJobLogReport, count int64, err error) {
	result, err = x.Offset(offset).Limit(limit).Find()
	if err != nil {
		return
	}

	if size := len(result); 0 < limit && 0 < size && size < limit {
		count = int64(size + offset)
		return
	}

	count, err = x.Offset(-1).Limit(-1).Count()
	return
}

func (x xxlJobLogReportDo) ScanByPage(result interface{}, offset int, limit int) (count int64, err error) {
	count, err = x.Count()
	if err != nil {
		return
	}

	err = x.Offset(offset).Limit(limit).Scan(result)
	return
}

func (x xxlJobLogReportDo) Scan(result interface{}) (err error) {
	return x.DO.Scan(result)
}

func (x xxlJobLogReportDo) Delete(models ...*model.XxlJobLogReport) (result gen.ResultInfo, err error) {
	return x.DO.Delete(models)
}

func (x *xxlJobLogReportDo) withDO(do gen.Dao) *xxlJobLogReportDo {
	x.DO = *do.(*gen.DO)
	return x
}
