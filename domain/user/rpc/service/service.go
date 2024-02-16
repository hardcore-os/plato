package service

import (
	context "context"

	"github.com/hardcore-os/plato/domain/user/storage"
)

var sm *storage.StorageManager

func Init(isTest bool) {
	sm = storage.NewStorageManager(isTest)
}

type Service struct {
}

func (s *Service) QueryUsers(ctx context.Context, req *QueryUsersRequest) (*QueryUsersResponse, error) {
	userOpts := make(map[uint64]*storage.Options)
	for uid, opt := range req.Opts {
		userOpts[uid] = &storage.Options{AllDevice: opt.ActiveDevice}
	}
	users := sm.QueryUsers(ctx, userOpts)
	return &QueryUsersResponse{Users: users}, nil
}
func (s *Service) CreateUsers(ctx context.Context, req *CreateUsersRequest) (*CreateUsersResponse, error) {
	// TODO 后续会承接复杂的聚合业务逻辑
	err := sm.CreateUsers(ctx, req.GetUsers(), nil)
	if err != nil {
		return nil, err
	}
	return &CreateUsersResponse{Code: 0, Msg: "ok"}, nil
}
func (s *Service) UpdateUsers(ctx context.Context, req *UpdateUsersRequest) (*UpdateUsersResponse, error) {
	// TODO update的权限目前很大，后面需要裁剪其字端
	err := sm.UpdateUsers(ctx, req.GetUsers(), nil)
	if err != nil {
		return nil, err
	}
	return &UpdateUsersResponse{Msg: "ok"}, nil
}

func (s *Service) mustEmbedUnimplementedUserServer() {}
