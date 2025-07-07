
package main

import (
	"fmt"
	"errors"
)

type Transaction struct {
	Type string
	Amount float64
}


type BankAccount struct {
	Owner string
	balance float64
	Transactions []Transaction
}

func (b *BankAccount) Deposit(amount float64) error {
	if amount <= 0 {
		return errors.New("Please enter a positive amount")	
	} else {		
		b.balance += amount
		b.Transactions = append(b.Transactions,Transaction{"Deposit",amount})
		return nil
	}
}

func (b *BankAccount) Withdrawl(amount float64) error {
	if amount < b.balance {
		b.balance -= amount
		b.Transactions = append(b.Transactions,Transaction{"Withdraw",amount})
		return nil
	} else if amount <= 0 {
		return errors.New("Please enter a positive amount")
	} else {
		return errors.New("Insufficient Balance")
	}
}

func (b *BankAccount) CheckBalance(){
	fmt.Println(b.balance)
}

func (b *BankAccount) PrintStatement() {
	fmt.Printf("please find below entries for Owner %s\n",b.Owner)
	for i,val := range b.Transactions {
		fmt.Printf("%d Type: %s, Amount: %f\n",i+1,val.Type,val.Amount)
	}
}

func main() {
	account1 := BankAccount{"user1",2000,[]Transaction{{"Deposit",2000},}}
	if err := account1.Deposit(1000); err != nil {
		fmt.Println("Error in deposit",err)
	}
	account1.CheckBalance()
	if err := account1.Withdrawl(1000); err != nil {
		fmt.Println("Error in withdraw",err)
	}
	account1.CheckBalance()
	account1.PrintStatement()
}
