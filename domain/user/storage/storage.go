package storage

import (
	"context"
	"fmt"

	"github.com/hardcore-os/plato/common/bus/event"
	"github.com/hardcore-os/plato/common/cache"
	"github.com/hardcore-os/plato/common/config"
	"github.com/hardcore-os/plato/common/idl/domain/user"
	"gorm.io/driver/mysql"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

const (
	userDomainCacheKey = "userDomainCacheKey:%d"
)

// 保持聚合的一致性
type StorageManager struct {
	cm    *cache.Manager
	db    *gorm.DB
	event *event.Manager
}

func NewStorageManager(isTest bool) *StorageManager {
	if isTest {
		return newMockStorageManager()
	}
	var err error
	cacheOpt := []*cache.Options{{Mode: cache.Local}, {Mode: cache.Remote}}
	channelOpt := map[event.Channel]*event.Options{event.UserEvent: {}}
	sm := &StorageManager{cm: cache.NewManager(cacheOpt), event: event.NewManager(channelOpt)}
	sm.db, err = gorm.Open(mysql.Open(config.GetDomainUserDBDNS()), &gorm.Config{})
	if err != nil {
		panic(err)
	}
	return sm
}

func newMockStorageManager() *StorageManager {
	var err error
	cacheOpt := []*cache.Options{{Mode: cache.Local}, {Mode: cache.Remote}}
	channelOpt := map[event.Channel]*event.Options{event.UserEvent: {}}
	sm := &StorageManager{cm: cache.NewManager(cacheOpt), event: event.NewManager(channelOpt)}
	sm.db, err = gorm.Open(sqlite.Open(":memory:"), &gorm.Config{})
	sm.db = sm.db.Debug()
	if err != nil {
		panic(err)
	}
	err = sm.db.AutoMigrate(&UserDAO{})
	if err != nil {
		panic(err)
	}
	return sm
}

func (s *StorageManager) QueryUsers(ctx context.Context, querys map[uint64]*Options) map[uint64]*user.UserDTO {
	res := make(map[uint64]*user.UserDTO, len(querys))
	miss := make([]uint64, 0)
	keys := make([]string, 0, len(querys))
	// 从cache中获取
	for uid, _ := range querys {
		keys = append(keys, fmt.Sprintf(userDomainCacheKey, uid))
	}
	cacheTable := s.cm.MGet(keys)
	for uid, _ := range querys {
		key := fmt.Sprintf(userDomainCacheKey, uid)
		if data, ok := cacheTable[key]; ok {
			userDTO, _ := data.(*user.UserDTO)
			res[uid] = userDTO
		} else {
			miss = append(miss, uid)
		}
	}
	// 从DB获取
	if len(miss) == 0 {
		return res
	}
	var users []UserDAO
	missUsers := make(map[string]interface{}, 0)
	// TODO: 这里没有查询device信息，这部分内容在存储层设计的时候会详细讲解
	// device信息作为一个实体，会封装在user领域的聚合中，会提供单独的缓存处理
	s.db.WithContext(ctx).Where("user_id IN (?)", miss).Find(&users)
	for _, userDAO := range users {
		userDTO := convertUserDAOToDTO(&userDAO)
		res[userDAO.UserID] = userDTO
		key := fmt.Sprintf(userDomainCacheKey, userDAO.UserID)
		missUsers[key] = userDTO
	}
	// 写入缓存
	s.cm.MSet(missUsers)
	// 创建事件
	// TODO: 领域事件后面会在异步任务框架处完善业务逻辑
	s.event.Send(event.UserEvent, nil)
	return res
}

func (s *StorageManager) CreateUsers(ctx context.Context, users []*user.UserDTO, opt *Options) error {
	// TODO 这里要有一个限流组件，在后面的基础设施层会详细讲解
	// 写DB
	userDAOList := make([]UserDAO, 0, len(users))
	keys := make([]string, 0, len(users))
	for _, userDTO := range users {
		userDAO := convertUserDTOToDAO(userDTO)
		userDAOList = append(userDAOList, *userDAO)
		key := fmt.Sprintf(userDomainCacheKey, userDAO.UserID)
		keys = append(keys, key)
	}
	ret := s.db.Create(&userDAOList)

	if ret.Error != nil {
		return ret.Error
	}
	// 删除缓存
	s.cm.MDel(keys)
	// 创建事件
	// TODO: 领域事件后面会在异步任务框架处完善业务逻辑
	s.event.Send(event.UserEvent, nil)
	return nil
}

func (s *StorageManager) UpdateUsers(ctx context.Context, users []*user.UserDTO, opt *Options) error {
	// TODO 这里要有一个限流组件，在后面的基础设施层会详细讲解
	// 写DB
	userDAOList := make([]UserDAO, 0, len(users))
	keys := make([]string, 0, len(users))
	for _, userDTO := range users {
		userDAO := convertUserDTOToDAO(userDTO)
		userDAOList = append(userDAOList, *userDAO)
		key := fmt.Sprintf(userDomainCacheKey, userDAO.UserID)
		keys = append(keys, key)
	}

	// TODO 这里是需要优化DB的性能的
	err := s.batchUpdateUsers(userDAOList)
	if err != nil {
		return err
	}
	// 删除缓存
	s.cm.MDel(keys)
	return nil
}

func convertUserDAOToDTO(dao *UserDAO) *user.UserDTO {
	// 创建一个新的UserDTO实例
	dto := &user.UserDTO{
		UserID: dao.UserID,
		Setting: &user.SettingDTO{
			FontSize:            dao.FontSize,
			DarkMode:            dao.DarkMode,
			ReceiveNotification: dao.ReceiveNotification,
			Language:            dao.Language,
			Notifications:       dao.Notifications,
		},
		Information: &user.InformationDTO{
			Nickname:  dao.Nickname,
			Avatar:    dao.Avatar,
			Signature: dao.Signature,
		},
		Pprofile: &user.ProfileDTO{
			Location: dao.Location,
			Age:      int32(dao.Age),
			Gender:   dao.Gender,
			Tags:     dao.Tags,
		},
	}
	return dto
}
func convertUserDTOToDAO(dto *user.UserDTO) *UserDAO {
	dao := &UserDAO{}

	if dto != nil {
		dao.UserID = dto.UserID

		if dto.GetSetting() != nil {
			dao.FontSize = dto.GetSetting().FontSize
			dao.DarkMode = dto.GetSetting().DarkMode
		}

		if dto.Setting != nil {
			dao.ReceiveNotification = dto.Setting.GetReceiveNotification()
			dao.Language = dto.GetSetting().Language
			dao.Notifications = dto.GetSetting().Notifications
		}

		if dto.GetInformation() != nil {
			dao.Nickname = dto.GetInformation().Nickname
			dao.Avatar = dto.GetInformation().Avatar
			dao.Signature = dto.GetInformation().Signature
		}

		if dto.GetPprofile() != nil {
			dao.Location = dto.GetPprofile().Location
			dao.Age = int(dto.GetPprofile().Age)
			dao.Gender = dto.GetPprofile().Gender
		}

		if dto.Pprofile != nil {
			dao.Tags = dto.Pprofile.GetTags()
		}
	}

	return dao
}

// 执行批量更新操作
func (sm *StorageManager) batchUpdateUsers(users []UserDAO) error {
	// 开始数据库事务
	tx := sm.db.Begin()
	defer tx.Commit()
	// 遍历用户列表，逐个更新
	for _, user := range users {
		// 更新用户信息
		// TODO 这里在实现存储基础组件是
		result := tx.Model(&UserDAO{}).Where("user_id = ?", user.UserID).Updates(user)
		if result.Error != nil {
			tx.Rollback()
			return result.Error
		}
	}
	return nil
}
