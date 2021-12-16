package bank_test

import (
	"fmt"
	"testing"

	bank "github.com/uu64/gpl-book/ch09/ex01/bank"
)

func TestDeposit(t *testing.T) {
	done := make(chan struct{})

	// Alice
	go func() {
		bank.Deposit(200)
		fmt.Println("=", bank.Balance())
		done <- struct{}{}
	}()

	// Bob
	go func() {
		bank.Deposit(100)
		done <- struct{}{}
	}()

	// Wait for both transactions.
	<-done
	<-done

	if got, want := bank.Balance(), 300; got != want {
		t.Errorf("Balance = %d, want %d", got, want)
	}
}

func TestWithdraw(t *testing.T) {
	if result := bank.Withdraw(230); result != true {
		t.Errorf("Balance = %d, withdraw %d, result %v", bank.Balance(), 230, result)
	}

	if got, want := bank.Balance(), 70; got != want {
		t.Errorf("Balance = %d, want %d", got, want)
	}

	if result := bank.Withdraw(3000); result != false {
		t.Errorf("Balance = %d, withdraw %d, result %v", bank.Balance(), 3000, result)
	}

	if got, want := bank.Balance(), 70; got != want {
		t.Errorf("Balance = %d, want %d", got, want)
	}
}
