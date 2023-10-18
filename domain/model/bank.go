package model

import (
	"github.com/asaskevich/govalidator"
	uuid "github.com/satori/go.uuid"
	"time"
)

// Estrutura de dados do banco
type Bank struct {
	Base     `valid:"required"` //Base de dados principais de todos os models; "herança"
	Code     string             `json:"code" gorm:"type:varchar(20)" valid:"notnull"`
	Name     string             `json:"name"  gorm:"type:varchar(255)" valid:"notnull"`
	Accounts []*Account         `gorm:"ForeignKey:BankId" valid:"-"`
}

// GoValidator para Bancos
func (bank *Bank) isValid() error {
	_, err := govalidator.ValidateStruct(bank)
	if err != nil {
		return err
	}
	return nil
}

// NewBank Função de criação de um novo banco
func NewBank(code string, name string) (*Bank, error) {

	bank := Bank{
		Code: code,
		Name: name,
	}

	bank.ID = uuid.NewV4().String()
	bank.CreatedAt = time.Now()

	err := bank.isValid()
	if err != nil {
		return nil, err
	}

	return &bank, nil

}
