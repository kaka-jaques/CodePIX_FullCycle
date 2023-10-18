package usecase

import (
	"fmt"
	"github.com/kaka-jaques/CodePIX_FullCycle/codepix-go/domain/model"
)

type PixUseCase struct {
	PixKeyRepository model.PixKeyRepositoryInterface
}

func (p *PixUseCase) RegisterKey(key string, kind string, accountId string) (*model.PixKey, error) {

	account, err := p.PixKeyRepository.FindAccount(accountId)
	if err != nil {
		return nil, err
	}

	pixKey, err := model.NewPixKey(kind, key, account)
	if err != nil {
		return nil, err
	}

	p.PixKeyRepository.Register(pixKey)
	if pixKey.ID == "" {
		return nil, fmt.Errorf("unable to create new key at the moment")
	}

	return pixKey, nil

}
