package main

import (
	"errors"
	"fmt"
)

type Transaction struct {
	Type   string
	Amount float64
}

type BankAccount struct {
	Owner        string
	PIN	     string
	Balance      float64
	Transactions []Transaction
}

func (b *BankAccount) Deposit(Amount float64) error {
	if Amount <= 0 {
		return errors.New("Please enter a positive number")
	} else {
		b.Balance += Amount
		b.Transactions = append(b.Transactions, Transaction{"Deposit", Amount})
		return nil
	}
}

func (b *BankAccount) Withdraw(Amount float64) error {
	if Amount <= 0 {
		return errors.New("Please enter a positive number")
	} else if Amount > b.Balance {
		return errors.New("Insufficient Balance")
	} else {
		b.Balance -= Amount
		b.Transactions = append(b.Transactions, Transaction{"Withdraw", Amount})
		return nil
	}
}

func (b *BankAccount) CheckBalance() {
	fmt.Printf("Current Balance for Owner %s is %f\n", b.Owner, b.Balance)
}

func (b *BankAccount) PrintStatement() {
	fmt.Println("Transactions for Owner :", b.Owner)
	for i, t := range b.Transactions {
		fmt.Printf("%d . Type: %s - Amount: %f\n", i, t.Type, t.Amount)
	}

}

func main() {
	accounts := make(map[string]*BankAccount)
	// Login Or Register
	login:
	for {
	var username,pin string
	var amount float64
	fmt.Println("Enter the username: ")
	fmt.Scanln(&username)
	fmt.Println("Enter the PIN: ")
	fmt.Scanln(&pin)
	account, exists := accounts[username]
	if !exists {
		account = &BankAccount{Owner: username, PIN: pin}
		accounts[username] = account
		fmt.Println("User accounts has been created")
	} else {
		if account.PIN != pin {
			fmt.Println("Incorrect PIN")	
			return 
		} else {
			fmt.Println("Logged in successfully")
		}
	}
	for {
		fmt.Println("Choose the option: ")
		fmt.Println("1. CheckBalance")
		fmt.Println("2. Deposit")
		fmt.Println("3. Withdraw")
		fmt.Println("4. PrintStatement")
		fmt.Println("5. Logout")

		var option int
		fmt.Scanln(&option)
		switch option {
		case 1:
			accounts[username].CheckBalance()
		case 2:
			fmt.Println("Enter the amount you want to deposit: ")
			fmt.Scanln(&amount)
			if err := accounts[username].Deposit(amount); err != nil {
				fmt.Println("Error ", err)
			}
		case 3:
			fmt.Println("Enter the amount you want to withdraw: ")
			fmt.Scanln(&amount)
			if err := accounts[username].Withdraw(amount); err != nil {
				fmt.Println("Error ", err)
			}
		case 4:
			accounts[username].PrintStatement()
		case 5:
			fmt.Println("Thank you!!")
			goto login
		default:
			fmt.Println("Select a valid option")
		}

	}
	}
}
