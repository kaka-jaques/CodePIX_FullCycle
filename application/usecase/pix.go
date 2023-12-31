package usecase

import (
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
		return nil, err
	}

	return pixKey, nil

}

func (p *PixUseCase) FindKey(key string, kind string) (*model.PixKey, error) {

	pixKey, err := p.PixKeyRepository.FindKeyByKind(key, kind)

	if err != nil {
		return nil, err
	}
	return pixKey, nil
}
