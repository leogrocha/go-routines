package mutex

import (
	"fmt"
	"sync"
	"time"

	"github.com/leogrocha/go-routines/internal/models"
	"github.com/leogrocha/go-routines/internal/services"
)

func generateAccountMock() models.Account {

	person := models.Person{
		Name:     "Seu Madruga",
		Document: "12312377756",
	}

	return models.Account{
		Number:   "1325",
		Agency:   "0153",
		BankName: "Nubank",
		BankCode: "001",
		Person:   person,
	}
}

func GenerateTransactionMock(transactionChannel chan<- models.Transaction, account *models.Account, done chan<- bool) {
	var wg sync.WaitGroup

	transaction := models.Transaction{
		Account: *account,
		Value:   100.00,
		DataAt:  time.Now(),
	}

	wg.Add(1)

	go func() {
		defer wg.Done()
		transactionChannel <- transaction
	}()

	wg.Wait()
	close(transactionChannel)

	done <- true
}

func AccountMutex() {
	// var mutex sync.Mutex

	transactionChannel := make(chan models.Transaction)
	done := make(chan bool)

	account := generateAccountMock()

	if err := services.Deposit(&account, 100.00); err != nil {
		fmt.Println(err)
	}

	if err := services.Deposit(&account, 200.00); err != nil {
		fmt.Println(err)
	}

	go GenerateTransactionMock(transactionChannel, &account, done)

	balance, err := services.GetBalance(&account)
	if err != nil {
		fmt.Println(err)
	}

	fmt.Println(balance)
	fmt.Println("-------------------")
	fmt.Println(<-transactionChannel)

	<-done

}
