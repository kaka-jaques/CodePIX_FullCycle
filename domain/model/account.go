package model

import (
	"github.com/asaskevich/govalidator"
	uuid "github.com/satori/go.uuid"
	"time"
)

type Account struct {
	Base      `valid:"required"` //Base de dados principais de todos os models; "herança"
	OwnerName string             `gorm:"column:owner_name;type:varchar(255);not null" valid:"notnull"`
	Bank      *Bank              `valid:"-"`
	BankId    string             `gorm:"column:bank_id;type:uuid;not null" valid:"-"`
	Number    string             `json:"number" gorm:"type:varchar(20)" valid:"notnull"`
	PixKeys   []*PixKey          `gorm:"ForeignKey:AccountID" valid:"-"`
}

// Utiliza a lib GoValidator para validar dados da conta
func (account *Account) isValid() error {
	_, err := govalidator.ValidateStruct(account)
	if err != nil {
		return err
	}
	return nil
}

// NewAccount Função que cria nova conta, podendo retornar um erro ou a conta em si criada
func NewAccount(bank *Bank, number string, ownerName string) (*Account, error) {

	account := Account{
		OwnerName: ownerName,
		Bank:      bank,
		Number:    number,
	}

	account.ID = uuid.NewV4().String()
	account.CreatedAt = time.Now()

	err := account.isValid()
	if err != nil {
		return nil, err
	}

	return &account, nil

}
