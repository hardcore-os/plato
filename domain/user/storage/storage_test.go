package storage

import (
	"context"
	"fmt"
	"testing"

	"github.com/hardcore-os/plato/common/idl/domain/user"
)

func TestStorageAPI(t *testing.T) {
	s := newMockStorageManager()
	ctx := context.Background()
	req := []*user.UserDTO{{UserID: 111, Information: &user.InformationDTO{Nickname: "test"}}}
	s.CreateUsers(ctx, req, nil)
	user := s.QueryUsers(ctx, map[uint64]*Options{111: {}})
	fmt.Printf("req query1 userID=%d, name=%s\n", user[111].UserID, user[111].Information.Nickname)
	req[0].Information.Nickname = "tes1"
	err := s.UpdateUsers(ctx, req, nil)
	if err != nil {
		panic(err)
	}
	user = s.QueryUsers(ctx, map[uint64]*Options{111: {}})
	fmt.Printf("req query2 userID=%d, name=%s\n", user[111].UserID, user[111].Information.Nickname)
}
