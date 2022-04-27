package v1

import (
	"chk/internal/chk_apis/v1/handlers"
	"context"
	"fmt"

	"gorm.io/gorm"
)

type repo struct {
	db *gorm.DB
}

func NewRepo(db *gorm.DB) (r *repo) {
	return &repo{db: db}
}

func (r *repo) Create(ctx context.Context, in []*CSV) (out []*CSV, err error) {

	err = r.db.WithContext(ctx).Create(&in).Error
	out = in

	return
}

func (r *repo) List(ctx context.Context, in *handlers.ListRequest) (out []*CSV, err error) {

	out = make([]*CSV, 0)

	qs := r.db.WithContext(ctx).
		Order(fmt.Sprintf("%v %v", in.Sort.GetOrderBy(), in.Sort.GetSort())).
		Offset(in.GetLimit() * (in.GetPage() - 1)).
		Limit(in.GetLimit())

	for _, filter := range in.Filters {
		qs = qs.Where(fmt.Sprintf("%v %v ?", filter.GetKey(), filter.GetMethod()), filter.GetValue())
	}

	if err = qs.Find(&out).Error; err != nil {
		return
	}

	return
}
