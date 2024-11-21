package gromore

import (
	"fmt"

	"github.com/paiz4ever/go-web-framework/global"
	"github.com/paiz4ever/go-web-framework/model"
)

type GromoreService struct{}

func (s *GromoreService) GenerateSign(csjUserID string) error {
	var account model.CsjAccount
	if err := global.DB.Where(&model.CsjAccount{TtUserID: csjUserID}).First(&account).Error; err != nil {
		return err
	}
	fmt.Printf("account: %+v\n", account.TtSecurityKey)
	return nil
}
