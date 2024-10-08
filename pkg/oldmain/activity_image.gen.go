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

func newActivityImage(db *gorm.DB, opts ...gen.DOOption) activityImage {
	_activityImage := activityImage{}

	_activityImage.activityImageDo.UseDB(db, opts...)
	_activityImage.activityImageDo.UseModel(&model.ActivityImage{})

	tableName := _activityImage.activityImageDo.TableName()
	_activityImage.ALL = field.NewAsterisk(tableName)
	_activityImage.ID = field.NewInt32(tableName, "id")
	_activityImage.ActivityID = field.NewInt64(tableName, "activity_id")
	_activityImage.IsShow = field.NewInt32(tableName, "is_show")
	_activityImage.Sort = field.NewInt32(tableName, "sort")
	_activityImage.ImagePc = field.NewString(tableName, "image_pc")
	_activityImage.ImageH5 = field.NewString(tableName, "image_h5")
	_activityImage.Href = field.NewString(tableName, "href")
	_activityImage.StartDate = field.NewInt32(tableName, "start_date")
	_activityImage.EndDate = field.NewInt32(tableName, "end_date")
	_activityImage.JumpType = field.NewInt32(tableName, "jump_type")
	_activityImage.GameListCode = field.NewString(tableName, "game_list_code")
	_activityImage.GameProviderSubtypeID = field.NewInt32(tableName, "game_provider_subtype_id")
	_activityImage.GameTypeID = field.NewInt32(tableName, "game_type_id")
	_activityImage.ProviderID = field.NewInt32(tableName, "provider_id")
	_activityImage.IsEnabled = field.NewInt32(tableName, "is_enabled")
	_activityImage.ImageType = field.NewInt32(tableName, "image_type")
	_activityImage.Title = field.NewString(tableName, "title")
	_activityImage.CreateAt = field.NewInt32(tableName, "create_at")
	_activityImage.UpdateAt = field.NewInt32(tableName, "update_at")
	_activityImage.OpUser = field.NewString(tableName, "op_user")
	_activityImage.Content = field.NewString(tableName, "content")

	_activityImage.fillFieldMap()

	return _activityImage
}

// activityImage 活动图片/广告图表
type activityImage struct {
	activityImageDo

	ALL                   field.Asterisk
	ID                    field.Int32  // 活动图片表主键
	ActivityID            field.Int64  // 活动ID
	IsShow                field.Int32  // 是否展示:1-展示,2-不展示
	Sort                  field.Int32  // 只可输入数字必须为整数，数字越大排在越前，范围在0-1000以内。2.轮播排序只针对活动排序内容不涉及弹窗和活动
	ImagePc               field.String // pc图片路径
	ImageH5               field.String // h5和app图片路径
	Href                  field.String // 跳转地址
	StartDate             field.Int32  // 活动开始日期
	EndDate               field.Int32  // 活动结束日期
	JumpType              field.Int32  // 跳转类型：1-跳转链接 2-跳转游戏
	GameListCode          field.String // 游戏启动code
	GameProviderSubtypeID field.Int32  // 游戏厂商分类ID
	GameTypeID            field.Int32  // 游戏大类类型
	ProviderID            field.Int32  // 游戏平台id
	IsEnabled             field.Int32  // 状态:1-启用，2-禁用
	ImageType             field.Int32  // 图片类型：1-活动图,2-轮播图,3-弹窗图,4-广告图,5-旧常规活动图,6-存款图
	Title                 field.String // 标题
	CreateAt              field.Int32  // 创建时间
	UpdateAt              field.Int32  // 修改时间
	OpUser                field.String // 操作人
	Content               field.String // 活动内容/广告内容

	fieldMap map[string]field.Expr
}

func (a activityImage) Table(newTableName string) *activityImage {
	a.activityImageDo.UseTable(newTableName)
	return a.updateTableName(newTableName)
}

func (a activityImage) As(alias string) *activityImage {
	a.activityImageDo.DO = *(a.activityImageDo.As(alias).(*gen.DO))
	return a.updateTableName(alias)
}

func (a *activityImage) updateTableName(table string) *activityImage {
	a.ALL = field.NewAsterisk(table)
	a.ID = field.NewInt32(table, "id")
	a.ActivityID = field.NewInt64(table, "activity_id")
	a.IsShow = field.NewInt32(table, "is_show")
	a.Sort = field.NewInt32(table, "sort")
	a.ImagePc = field.NewString(table, "image_pc")
	a.ImageH5 = field.NewString(table, "image_h5")
	a.Href = field.NewString(table, "href")
	a.StartDate = field.NewInt32(table, "start_date")
	a.EndDate = field.NewInt32(table, "end_date")
	a.JumpType = field.NewInt32(table, "jump_type")
	a.GameListCode = field.NewString(table, "game_list_code")
	a.GameProviderSubtypeID = field.NewInt32(table, "game_provider_subtype_id")
	a.GameTypeID = field.NewInt32(table, "game_type_id")
	a.ProviderID = field.NewInt32(table, "provider_id")
	a.IsEnabled = field.NewInt32(table, "is_enabled")
	a.ImageType = field.NewInt32(table, "image_type")
	a.Title = field.NewString(table, "title")
	a.CreateAt = field.NewInt32(table, "create_at")
	a.UpdateAt = field.NewInt32(table, "update_at")
	a.OpUser = field.NewString(table, "op_user")
	a.Content = field.NewString(table, "content")

	a.fillFieldMap()

	return a
}

func (a *activityImage) GetFieldByName(fieldName string) (field.OrderExpr, bool) {
	_f, ok := a.fieldMap[fieldName]
	if !ok || _f == nil {
		return nil, false
	}
	_oe, ok := _f.(field.OrderExpr)
	return _oe, ok
}

func (a *activityImage) fillFieldMap() {
	a.fieldMap = make(map[string]field.Expr, 21)
	a.fieldMap["id"] = a.ID
	a.fieldMap["activity_id"] = a.ActivityID
	a.fieldMap["is_show"] = a.IsShow
	a.fieldMap["sort"] = a.Sort
	a.fieldMap["image_pc"] = a.ImagePc
	a.fieldMap["image_h5"] = a.ImageH5
	a.fieldMap["href"] = a.Href
	a.fieldMap["start_date"] = a.StartDate
	a.fieldMap["end_date"] = a.EndDate
	a.fieldMap["jump_type"] = a.JumpType
	a.fieldMap["game_list_code"] = a.GameListCode
	a.fieldMap["game_provider_subtype_id"] = a.GameProviderSubtypeID
	a.fieldMap["game_type_id"] = a.GameTypeID
	a.fieldMap["provider_id"] = a.ProviderID
	a.fieldMap["is_enabled"] = a.IsEnabled
	a.fieldMap["image_type"] = a.ImageType
	a.fieldMap["title"] = a.Title
	a.fieldMap["create_at"] = a.CreateAt
	a.fieldMap["update_at"] = a.UpdateAt
	a.fieldMap["op_user"] = a.OpUser
	a.fieldMap["content"] = a.Content
}

func (a activityImage) clone(db *gorm.DB) activityImage {
	a.activityImageDo.ReplaceConnPool(db.Statement.ConnPool)
	return a
}

func (a activityImage) replaceDB(db *gorm.DB) activityImage {
	a.activityImageDo.ReplaceDB(db)
	return a
}

type activityImageDo struct{ gen.DO }

func (a activityImageDo) Debug() *activityImageDo {
	return a.withDO(a.DO.Debug())
}

func (a activityImageDo) WithContext(ctx context.Context) *activityImageDo {
	return a.withDO(a.DO.WithContext(ctx))
}

func (a activityImageDo) ReadDB() *activityImageDo {
	return a.Clauses(dbresolver.Read)
}

func (a activityImageDo) WriteDB() *activityImageDo {
	return a.Clauses(dbresolver.Write)
}

func (a activityImageDo) Session(config *gorm.Session) *activityImageDo {
	return a.withDO(a.DO.Session(config))
}

func (a activityImageDo) Clauses(conds ...clause.Expression) *activityImageDo {
	return a.withDO(a.DO.Clauses(conds...))
}

func (a activityImageDo) Returning(value interface{}, columns ...string) *activityImageDo {
	return a.withDO(a.DO.Returning(value, columns...))
}

func (a activityImageDo) Not(conds ...gen.Condition) *activityImageDo {
	return a.withDO(a.DO.Not(conds...))
}

func (a activityImageDo) Or(conds ...gen.Condition) *activityImageDo {
	return a.withDO(a.DO.Or(conds...))
}

func (a activityImageDo) Select(conds ...field.Expr) *activityImageDo {
	return a.withDO(a.DO.Select(conds...))
}

func (a activityImageDo) Where(conds ...gen.Condition) *activityImageDo {
	return a.withDO(a.DO.Where(conds...))
}

func (a activityImageDo) Order(conds ...field.Expr) *activityImageDo {
	return a.withDO(a.DO.Order(conds...))
}

func (a activityImageDo) Distinct(cols ...field.Expr) *activityImageDo {
	return a.withDO(a.DO.Distinct(cols...))
}

func (a activityImageDo) Omit(cols ...field.Expr) *activityImageDo {
	return a.withDO(a.DO.Omit(cols...))
}

func (a activityImageDo) Join(table schema.Tabler, on ...field.Expr) *activityImageDo {
	return a.withDO(a.DO.Join(table, on...))
}

func (a activityImageDo) LeftJoin(table schema.Tabler, on ...field.Expr) *activityImageDo {
	return a.withDO(a.DO.LeftJoin(table, on...))
}

func (a activityImageDo) RightJoin(table schema.Tabler, on ...field.Expr) *activityImageDo {
	return a.withDO(a.DO.RightJoin(table, on...))
}

func (a activityImageDo) Group(cols ...field.Expr) *activityImageDo {
	return a.withDO(a.DO.Group(cols...))
}

func (a activityImageDo) Having(conds ...gen.Condition) *activityImageDo {
	return a.withDO(a.DO.Having(conds...))
}

func (a activityImageDo) Limit(limit int) *activityImageDo {
	return a.withDO(a.DO.Limit(limit))
}

func (a activityImageDo) Offset(offset int) *activityImageDo {
	return a.withDO(a.DO.Offset(offset))
}

func (a activityImageDo) Scopes(funcs ...func(gen.Dao) gen.Dao) *activityImageDo {
	return a.withDO(a.DO.Scopes(funcs...))
}

func (a activityImageDo) Unscoped() *activityImageDo {
	return a.withDO(a.DO.Unscoped())
}

func (a activityImageDo) Create(values ...*model.ActivityImage) error {
	if len(values) == 0 {
		return nil
	}
	return a.DO.Create(values)
}

func (a activityImageDo) CreateInBatches(values []*model.ActivityImage, batchSize int) error {
	return a.DO.CreateInBatches(values, batchSize)
}

// Save : !!! underlying implementation is different with GORM
// The method is equivalent to executing the statement: db.Clauses(clause.OnConflict{UpdateAll: true}).Create(values)
func (a activityImageDo) Save(values ...*model.ActivityImage) error {
	if len(values) == 0 {
		return nil
	}
	return a.DO.Save(values)
}

func (a activityImageDo) First() (*model.ActivityImage, error) {
	if result, err := a.DO.First(); err != nil {
		return nil, err
	} else {
		return result.(*model.ActivityImage), nil
	}
}

func (a activityImageDo) Take() (*model.ActivityImage, error) {
	if result, err := a.DO.Take(); err != nil {
		return nil, err
	} else {
		return result.(*model.ActivityImage), nil
	}
}

func (a activityImageDo) Last() (*model.ActivityImage, error) {
	if result, err := a.DO.Last(); err != nil {
		return nil, err
	} else {
		return result.(*model.ActivityImage), nil
	}
}

func (a activityImageDo) Find() ([]*model.ActivityImage, error) {
	result, err := a.DO.Find()
	return result.([]*model.ActivityImage), err
}

func (a activityImageDo) FindInBatch(batchSize int, fc func(tx gen.Dao, batch int) error) (results []*model.ActivityImage, err error) {
	buf := make([]*model.ActivityImage, 0, batchSize)
	err = a.DO.FindInBatches(&buf, batchSize, func(tx gen.Dao, batch int) error {
		defer func() { results = append(results, buf...) }()
		return fc(tx, batch)
	})
	return results, err
}

func (a activityImageDo) FindInBatches(result *[]*model.ActivityImage, batchSize int, fc func(tx gen.Dao, batch int) error) error {
	return a.DO.FindInBatches(result, batchSize, fc)
}

func (a activityImageDo) Attrs(attrs ...field.AssignExpr) *activityImageDo {
	return a.withDO(a.DO.Attrs(attrs...))
}

func (a activityImageDo) Assign(attrs ...field.AssignExpr) *activityImageDo {
	return a.withDO(a.DO.Assign(attrs...))
}

func (a activityImageDo) Joins(fields ...field.RelationField) *activityImageDo {
	for _, _f := range fields {
		a = *a.withDO(a.DO.Joins(_f))
	}
	return &a
}

func (a activityImageDo) Preload(fields ...field.RelationField) *activityImageDo {
	for _, _f := range fields {
		a = *a.withDO(a.DO.Preload(_f))
	}
	return &a
}

func (a activityImageDo) FirstOrInit() (*model.ActivityImage, error) {
	if result, err := a.DO.FirstOrInit(); err != nil {
		return nil, err
	} else {
		return result.(*model.ActivityImage), nil
	}
}

func (a activityImageDo) FirstOrCreate() (*model.ActivityImage, error) {
	if result, err := a.DO.FirstOrCreate(); err != nil {
		return nil, err
	} else {
		return result.(*model.ActivityImage), nil
	}
}

func (a activityImageDo) FindByPage(offset int, limit int) (result []*model.ActivityImage, count int64, err error) {
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

func (a activityImageDo) ScanByPage(result interface{}, offset int, limit int) (count int64, err error) {
	count, err = a.Count()
	if err != nil {
		return
	}

	err = a.Offset(offset).Limit(limit).Scan(result)
	return
}

func (a activityImageDo) Scan(result interface{}) (err error) {
	return a.DO.Scan(result)
}

func (a activityImageDo) Delete(models ...*model.ActivityImage) (result gen.ResultInfo, err error) {
	return a.DO.Delete(models)
}

func (a *activityImageDo) withDO(do gen.Dao) *activityImageDo {
	a.DO = *do.(*gen.DO)
	return a
}
