package services

import (
	"fmt"

	"github.com/leogrocha/go-routines/internal/models"
)

func Withdraw(account *models.Account, value float64) error {
	if account != nil && value > 0.0 {

		if value > account.Balance {
			return fmt.Errorf("%s", "SALDO INSUFICIENTE PARA REALIZAR A TRANSAÇÃO.")
		}

		account.Balance -= value
		return nil
	}

	return fmt.Errorf("%s", "Erro ao realizar saque.")
}

func Deposit(account *models.Account, value float64) error {

	if account != nil && value > 0.0 {

		account.Balance += value
		return nil
	}

	return fmt.Errorf("%s", "Erro ao realizar deposito.")
}

func GetBalance(account *models.Account) (float64, error) {
	return account.Balance, nil
}
