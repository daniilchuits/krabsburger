package main

import (
	"fmt"
	"os"
	"strings"
)

type Account struct {
	Owner     string
	Balance   float64
	Operation []string
}

func (a *Account) Deposit(amount float64) {
	a.Balance += amount
	dep := fmt.Sprintf("аккаунт пополнен на %.2f\n", amount)
	a.Operation = append(a.Operation, dep)
	SaveOperationToFile(dep)
	fmt.Printf("аккаунт пополнен на %.2f\n", amount)
}

func (a *Account) Withdraw(amount float64) {
	if amount >= 100 {
		with := fmt.Sprintf("нельзя сниммать больше 100 за раз\n")
		a.Operation = append(a.Operation, with)
		SaveOperationToFile(with)
		fmt.Println("нельзя сниммать больше 100 за раз")
	} else if amount < 10 {
		with := fmt.Sprintf("нельзя сниммать меньше 10 за раз\n")
		a.Operation = append(a.Operation, with)
		SaveOperationToFile(with)
		fmt.Println("нельзя сниммать меньше 10 за раз")
	} else if a.Balance >= amount {
		with := fmt.Sprintf("с аккаунта снято %.2f\n", amount)
		a.Operation = append(a.Operation, with)
		a.Balance -= amount
		SaveOperationToFile(with)
		fmt.Printf("с аккаунта снято %.2f\n", amount)
	} else {
		with := fmt.Sprintf("не достаточено средств\n")
		a.Operation = append(a.Operation, with)
		SaveOperationToFile(with)
		fmt.Println("не достаточено средств")
	}
}

func (a *Account) ShowBalance() {
	fmt.Printf("На аккаунте %s %.2f средств\n", a.Owner, a.Balance)
	balan := fmt.Sprintf("баланс просмотрен\n")
	SaveOperationToFile(balan)
	a.Operation = append(a.Operation, balan)
}

func (a Account) ShowOperation() {
	if len(a.Operation) == 0 {
		fmt.Println("Нет оппераций")
		return
	}
	fmt.Println("Опперации:")
	for _, op := range a.Operation {
		fmt.Println(op)
	}
}

func Transfer(Accs *[]Account, to string, on string, amount float64) {
	var receiver *Account
	var sender *Account
	found := false
	fnd := false

	for i, acc := range *Accs {
		if strings.ToLower(to) == strings.ToLower(acc.Owner) {
			found = true
			sender = &(*Accs)[i]
		}
	}
	if !found {
		fmt.Println("sender is not found")
		return
	}

	for q, acc := range *Accs {
		if strings.ToLower(on) == strings.ToLower(acc.Owner) {
			fnd = true
			receiver = &(*Accs)[q]
		}
	}
	if !fnd {
		fmt.Println("receiver is not found")
		return
	}

	fmt.Println("commision 1%")
	total := amount * 1.01
	if sender.Balance < total {
		fmt.Printf("sender doesn't have enough money\ncurrent money: %.2f\n", sender.Balance)
		return
	}
	sender.Balance -= total
	receiver.Balance += amount
	tranFrom := fmt.Sprintf("с аккаунта %s переведенно %.2f на аккаунт %s\n", sender.Owner, amount, receiver.Owner)
	sender.Operation = append(sender.Operation, tranFrom)
	SaveOperationToFile(tranFrom)
	tranTo := fmt.Sprintf("с аккаунта %s переведенно %.2f на аккаунт %s\n", sender.Owner, amount, receiver.Owner)
	receiver.Operation = append(receiver.Operation, tranTo)
	SaveOperationToFile(tranTo)
	fmt.Printf("added %.2f\n", amount)
}

func (a *Account) AddInterest(rate float64) {
	if a.Balance <= 0 {
		fmt.Printf("на аккаунте %.2f средств\n", a.Balance)
		return
	}
	end := rate/100 + 1
	a.Balance *= end
	fmt.Printf("На аккаунт %s добавлено %.2f%%\n", a.Owner, rate)
	ra := fmt.Sprintf("На аккаунт %s добавлено %.2f%%\n", a.Owner, rate)
	a.Operation = append(a.Operation, ra)
	SaveOperationToFile(ra)
}

func SaveOperationToFile(op string) {
	f, err := os.OpenFile("operation.txt", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}
	defer f.Close()

	if _, err := f.WriteString(op + "\n"); err != nil {
		fmt.Println("Error:", err)
	}
}

func main() {
	Users := []Account{
		{"vasya", 100, []string{}},
		{"gondoplyas", 200, []string{}},
	}
	Users[0].Deposit(10)
	Users[0].ShowBalance()
	Users[0].Withdraw(20)
	Users[0].ShowBalance()
	Transfer(&Users, "vasya", "gondoplyas", 11)
	Users[0].AddInterest(3)
	Users[0].ShowOperation()
}
