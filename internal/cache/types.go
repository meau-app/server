package cache

import "github.com/meau-app/server/internal/dao"

type CacheType string

const (
	TypePet  CacheType = "PET"
	TypeUser CacheType = "USER"
)

type CacheItemType interface {
	dao.Pet | dao.User
}
