package model

import (
	"gorm.io/datatypes"
)

type CsjAccount struct {
	ID              uint           `gorm:"primarykey"`
	Name            string         `gorm:"type:varchar(64);default:''"`
	TtUserID        string         `gorm:"type:varchar(64);default:''"`
	TtSecurityKey   string         `gorm:"type:varchar(64);default:''"`
	YlhCnf          datatypes.JSON `gorm:"type:json;default:'{}'"`
	BdCnf           datatypes.JSON `gorm:"type:json;default:'{}'"`
	KsCnf           datatypes.JSON `gorm:"type:json;default:'{}'"`
	AdNetworkCnf    datatypes.JSON `gorm:"type:json;default:'[]'"`
	ToponPublishKey string         `gorm:"type:varchar(128);"`
	Owner           string         `gorm:"type:varchar(64);default:'1'"`
	BdAccessKey     string         `gorm:"type:varchar(64);"`
	BdPrivateKey    string         `gorm:"type:text;"`
}

func (s *CsjAccount) TableName() string {
	return "ad_utils_csjaccountmodel"
}
