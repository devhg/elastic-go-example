package service

import (
	"context"

	"github.com/devhg/es/internal/biz"
	"github.com/devhg/es/internal/model"
)

type UserService struct {
	es *biz.UserEsProvider
}

func NewUserService(es *biz.UserEsProvider) *UserService {
	return &UserService{
		es: es,
	}
}

func (s *UserService) BatchAdd(ctx context.Context, user []*model.UserEs) error {
	return s.es.BatchAdd(ctx, user)
}

func (s *UserService) BatchDel(ctx context.Context, user []*model.UserEs) error {
	return s.es.BatchDel(ctx, user)
}

func (s *UserService) BatchUpdate(ctx context.Context, user []*model.UserEs) error {
	return s.es.BatchUpdate(ctx, user)
}

func (s *UserService) MGet(ctx context.Context, IDS []uint64) ([]*model.UserEs, error) {
	return s.es.MGet(ctx, IDS)
}

func (s *UserService) Search(ctx context.Context, req *model.SearchRequest) ([]*model.UserEs, error) {
	return s.es.Search(ctx, req.ToFilter())
}
