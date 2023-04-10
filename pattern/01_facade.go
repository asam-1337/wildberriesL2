package main

import (
	"errors"
	"fmt"
)

/*
	Реализовать паттерн «фасад».
Объяснить применимость паттерна, его плюсы и минусы,а также реальные примеры использования данного примера на практике.
	https://en.wikipedia.org/wiki/Facade_pattern
Фасад необходим когда требуется скрыть сложности реализации, но предоставить простой интерфейс для общения с клиентом.
*/
// account ------------------------------------------------------------------
type account struct {
	ID string
}

func newAccount(ID string) *account {
	return &account{
		ID: ID,
	}
}

func (a account) checkAccount(id string) bool {
	return a.ID == id
}

// wallet -------------------------------------------------------------------
type wallet struct {
	balance int
}

func newWallet() *wallet {
	return &wallet{0}
}

func (w *wallet) debitBalance(amount int) error {
	if w.balance-amount < 0 {
		return fmt.Errorf("not enough money")
	}
	w.balance -= amount
	return nil
}

func (w *wallet) creditBalance(amount int) {
	w.balance += amount
}

// walletFacade под капотом может быть спрятана сложная логика и вызовы различных методов,
// однако клиенту не нужно знать все аспекты реализации, реализует
type walletFacade struct {
	a *account
	w *wallet
}

func newWalletFacade(accountID string) *walletFacade {
	return &walletFacade{
		a: newAccount(accountID),
		w: newWallet(),
	}
}

func (wf *walletFacade) addMoneyToWallet(accountID string, amount int) error {
	if wf.a.checkAccount(accountID) {
		return errors.New("incorrect accountID")
	}
	wf.w.creditBalance(amount)
	return nil
}

func (wf *walletFacade) deductMoneyFromWallet(accountID string, amount int) error {
	if wf.a.checkAccount(accountID) {
		return errors.New("incorrect accountID")
	}

	if err := wf.w.debitBalance(amount); err != nil {
		return err
	}

	return nil
}

// Facade интерфейс для того чтобы полностью абстрагироваться от реализации
type Facade interface {
	addMoneyToWallet(accountID string, amount int) error
	deductMoneyFromWallet(accountID string, amount int) error
}

func main() {
	slice := make([]string, 0, 3)
	slice = append(slice, "a")
	slice = append(slice, "a")
	func(slice []string) {
		slice = append(slice, "a")
		slice[0] = "b"
		slice[1] = "b"
		fmt.Print(slice)
	}(slice)
	fmt.Print(slice, len(slice), cap(slice))
}
