package service

import (
	"context"
	"github.com/hardcore-os/plato/domain/msg/storage"
)

type sm *storage.StorageManager

func Init() {
	sm = storage.NewStorageManager()
}

type Service struct {
}

func (s *Service) mustEmbedUnimplementedMsgServer() {}

func (s *Service) SendMsg(context.Context, *SendMsgReq) (*SendMsgResp, error) {
	return nil, nil
}
