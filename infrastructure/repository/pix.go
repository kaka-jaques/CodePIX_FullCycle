package repository

import (
	"fmt"
	"github.com/jinzhu/gorm"
	"github.com/kaka-jaques/CodePIX_FullCycle/codepix-go/domain/model"
)

type PixKeyRepositoryDb struct {
	Db *gorm.DB
}

func (r PixKeyRepositoryDb) AddBank(bank *model.Bank) error {
	err := r.Db.Create(&bank).Error
	if err != nil {
		return err
	}
	return nil
}

func (r PixKeyRepositoryDb) AddAccount(account *model.Account) error {
	err := r.Db.Create(&account).Error
	if err != nil {
		return err
	}
	return nil
}

func (r PixKeyRepositoryDb) RegisterKey(key *model.PixKey) error {
	err := r.Db.Create(&key).Error
	if err != nil {
		return err
	}
	return nil
}

func (r PixKeyRepositoryDb) FinKeyById(key string, kind string) (*model.PixKey, error) {
	var pixKey model.PixKey

	r.Db.Preload("Account.Bank").First(&pixKey, "kind = ? AND key = ?", kind, key)

	if pixKey.ID == "" {
		return nil, fmt.Errorf("no key was found")
	}

	return &pixKey, nil
}

func (r PixKeyRepositoryDb) FindAccount(id string) (*model.Account, error) {
	var account model.Account

	r.Db.Preload("Bank").First(&account, "id = ?", id)

	if account.ID == "" {
		return nil, fmt.Errorf("no account found")
	}

	return &account, nil
}

func (r PixKeyRepositoryDb) FindBank(id string) (*model.Bank, error) {
	var bank model.Bank
	r.Db.First(&bank, "id = ?", id)

	if bank.ID == "" {
		return nil, fmt.Errorf("no bank found")
	}
	return &bank, nil
}
