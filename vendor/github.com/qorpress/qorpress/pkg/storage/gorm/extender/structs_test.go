package extender

///////////////////////////////////////////////////////////
// THIS FILE IS AUTO GENERATED by gormgen, DON'T EDIT IT //
//        ANY CHANGES DONE HERE WILL BE LOST             //
///////////////////////////////////////////////////////////

import (
	"fmt"
	"time"

	"github.com/jinzhu/gorm"
)

func (t *BasicModel) Save(db *gorm.DB) error {
	return db.Save(t).Error
}

func (t *BasicModel) Delete(db *gorm.DB) error {
	return db.Delete(t).Error
}

type BasicModelQueryBuilder struct {
	order []string
	where []struct {
		prefix string
		value  interface{}
	}
	limit  int
	offset int
}

func (qb *BasicModelQueryBuilder) buildQuery(db *gorm.DB) *gorm.DB {
	ret := db
	for _, where := range qb.where {
		ret = ret.Where(where.prefix, where.value)
	}
	for _, order := range qb.order {
		ret = ret.Order(order)
	}
	ret = ret.Limit(qb.limit).Offset(qb.offset)
	return ret
}

func (qb *BasicModelQueryBuilder) Count(db *gorm.DB) (int, error) {
	var c int
	res := qb.buildQuery(db).Model(&BasicModel{}).Count(&c)
	if res.RecordNotFound() {
		c = 0
	}
	return c, res.Error
}

func (qb *BasicModelQueryBuilder) First(db *gorm.DB) (*BasicModel, error) {
	ret := &BasicModel{}
	res := qb.buildQuery(db).First(ret)
	if res.RecordNotFound() {
		ret = nil
	}
	return ret, res.Error
}

func (qb *BasicModelQueryBuilder) QueryOne(db *gorm.DB) (*BasicModel, error) {
	qb.limit = 1
	ret, err := qb.QueryAll(db)
	if len(ret) > 0 {
		return &ret[0], err
	}
	return nil, err
}

func (qb *BasicModelQueryBuilder) QueryAll(db *gorm.DB) ([]BasicModel, error) {
	ret := []BasicModel{}
	err := qb.buildQuery(db).Find(&ret).Error
	return ret, err
}

func (qb *BasicModelQueryBuilder) Limit(limit int) *BasicModelQueryBuilder {
	qb.limit = limit
	return qb
}

func (qb *BasicModelQueryBuilder) Offset(offset int) *BasicModelQueryBuilder {
	qb.offset = offset
	return qb
}

func (qb *BasicModelQueryBuilder) WhereID(p Predicate, value uint) *BasicModelQueryBuilder {
	qb.where = append(qb.where, struct {
		prefix string
		value  interface{}
	}{
		fmt.Sprintf("%v %v ?", "id", p),
		value,
	})
	return qb
}

func (qb *BasicModelQueryBuilder) OrderByID(asc bool) *BasicModelQueryBuilder {
	order := "DESC"
	if asc {
		order = "ASC"
	}

	qb.order = append(qb.order, "id "+order)
	return qb
}

func (qb *BasicModelQueryBuilder) WhereName(p Predicate, value string) *BasicModelQueryBuilder {
	qb.where = append(qb.where, struct {
		prefix string
		value  interface{}
	}{
		fmt.Sprintf("%v %v ?", "name2", p),
		value,
	})
	return qb
}

func (qb *BasicModelQueryBuilder) OrderByName(asc bool) *BasicModelQueryBuilder {
	order := "DESC"
	if asc {
		order = "ASC"
	}

	qb.order = append(qb.order, "name2 "+order)
	return qb
}

func (qb *BasicModelQueryBuilder) WhereAge(p Predicate, value int) *BasicModelQueryBuilder {
	qb.where = append(qb.where, struct {
		prefix string
		value  interface{}
	}{
		fmt.Sprintf("%v %v ?", "age", p),
		value,
	})
	return qb
}

func (qb *BasicModelQueryBuilder) OrderByAge(asc bool) *BasicModelQueryBuilder {
	order := "DESC"
	if asc {
		order = "ASC"
	}

	qb.order = append(qb.order, "age "+order)
	return qb
}

func (t *ComplexModel) Save(db *gorm.DB) error {
	return db.Save(t).Error
}

func (t *ComplexModel) Delete(db *gorm.DB) error {
	return db.Delete(t).Error
}

type ComplexModelQueryBuilder struct {
	order []string
	where []struct {
		prefix string
		value  interface{}
	}
	limit  int
	offset int
}

func (qb *ComplexModelQueryBuilder) buildQuery(db *gorm.DB) *gorm.DB {
	ret := db
	for _, where := range qb.where {
		ret = ret.Where(where.prefix, where.value)
	}
	for _, order := range qb.order {
		ret = ret.Order(order)
	}
	ret = ret.Limit(qb.limit).Offset(qb.offset)
	return ret
}

func (qb *ComplexModelQueryBuilder) Count(db *gorm.DB) (int, error) {
	var c int
	res := qb.buildQuery(db).Model(&ComplexModel{}).Count(&c)
	if res.RecordNotFound() {
		c = 0
	}
	return c, res.Error
}

func (qb *ComplexModelQueryBuilder) First(db *gorm.DB) (*ComplexModel, error) {
	ret := &ComplexModel{}
	res := qb.buildQuery(db).First(ret)
	if res.RecordNotFound() {
		ret = nil
	}
	return ret, res.Error
}

func (qb *ComplexModelQueryBuilder) QueryOne(db *gorm.DB) (*ComplexModel, error) {
	qb.limit = 1
	ret, err := qb.QueryAll(db)
	if len(ret) > 0 {
		return &ret[0], err
	}
	return nil, err
}

func (qb *ComplexModelQueryBuilder) QueryAll(db *gorm.DB) ([]ComplexModel, error) {
	ret := []ComplexModel{}
	err := qb.buildQuery(db).Find(&ret).Error
	return ret, err
}

func (qb *ComplexModelQueryBuilder) Limit(limit int) *ComplexModelQueryBuilder {
	qb.limit = limit
	return qb
}

func (qb *ComplexModelQueryBuilder) Offset(offset int) *ComplexModelQueryBuilder {
	qb.offset = offset
	return qb
}

func (qb *ComplexModelQueryBuilder) WhereID(p Predicate, value uint) *ComplexModelQueryBuilder {
	qb.where = append(qb.where, struct {
		prefix string
		value  interface{}
	}{
		fmt.Sprintf("%v %v ?", "id", p),
		value,
	})
	return qb
}

func (qb *ComplexModelQueryBuilder) OrderByID(asc bool) *ComplexModelQueryBuilder {
	order := "DESC"
	if asc {
		order = "ASC"
	}

	qb.order = append(qb.order, "id "+order)
	return qb
}

func (qb *ComplexModelQueryBuilder) WhereCreatedAt(p Predicate, value time.Time) *ComplexModelQueryBuilder {
	qb.where = append(qb.where, struct {
		prefix string
		value  interface{}
	}{
		fmt.Sprintf("%v %v ?", "created_at", p),
		value,
	})
	return qb
}

func (qb *ComplexModelQueryBuilder) OrderByCreatedAt(asc bool) *ComplexModelQueryBuilder {
	order := "DESC"
	if asc {
		order = "ASC"
	}

	qb.order = append(qb.order, "created_at "+order)
	return qb
}

func (qb *ComplexModelQueryBuilder) WhereUpdatedAt(p Predicate, value time.Time) *ComplexModelQueryBuilder {
	qb.where = append(qb.where, struct {
		prefix string
		value  interface{}
	}{
		fmt.Sprintf("%v %v ?", "updated_at", p),
		value,
	})
	return qb
}

func (qb *ComplexModelQueryBuilder) OrderByUpdatedAt(asc bool) *ComplexModelQueryBuilder {
	order := "DESC"
	if asc {
		order = "ASC"
	}

	qb.order = append(qb.order, "updated_at "+order)
	return qb
}

func (qb *ComplexModelQueryBuilder) WhereDeletedAt(p Predicate, value *time.Time) *ComplexModelQueryBuilder {
	qb.where = append(qb.where, struct {
		prefix string
		value  interface{}
	}{
		fmt.Sprintf("%v %v ?", "deleted_at", p),
		value,
	})
	return qb
}

func (qb *ComplexModelQueryBuilder) OrderByDeletedAt(asc bool) *ComplexModelQueryBuilder {
	order := "DESC"
	if asc {
		order = "ASC"
	}

	qb.order = append(qb.order, "deleted_at "+order)
	return qb
}

func (qb *ComplexModelQueryBuilder) WhereName(p Predicate, value string) *ComplexModelQueryBuilder {
	qb.where = append(qb.where, struct {
		prefix string
		value  interface{}
	}{
		fmt.Sprintf("%v %v ?", "name", p),
		value,
	})
	return qb
}

func (qb *ComplexModelQueryBuilder) OrderByName(asc bool) *ComplexModelQueryBuilder {
	order := "DESC"
	if asc {
		order = "ASC"
	}

	qb.order = append(qb.order, "name "+order)
	return qb
}

func (qb *ComplexModelQueryBuilder) WhereEmbeddedName(p Predicate, value string) *ComplexModelQueryBuilder {
	qb.where = append(qb.where, struct {
		prefix string
		value  interface{}
	}{
		fmt.Sprintf("%v %v ?", "embedded_name", p),
		value,
	})
	return qb
}

func (qb *ComplexModelQueryBuilder) OrderByEmbeddedName(asc bool) *ComplexModelQueryBuilder {
	order := "DESC"
	if asc {
		order = "ASC"
	}

	qb.order = append(qb.order, "embedded_name "+order)
	return qb
}

func (qb *ComplexModelQueryBuilder) WhereTest(p Predicate, value struct{ NoIdea string }) *ComplexModelQueryBuilder {
	qb.where = append(qb.where, struct {
		prefix string
		value  interface{}
	}{
		fmt.Sprintf("%v %v ?", "test", p),
		value,
	})
	return qb
}

func (qb *ComplexModelQueryBuilder) OrderByTest(asc bool) *ComplexModelQueryBuilder {
	order := "DESC"
	if asc {
		order = "ASC"
	}

	qb.order = append(qb.order, "test "+order)
	return qb
}