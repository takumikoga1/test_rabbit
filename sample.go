package main

import (
	"fmt"
	"sync"
)

// Personクラス
type Person struct {
	ID        int
	Name      string
	Accounts  []*Account
}

// Accountクラス
type Account struct {
	ID      int
	Balance float64
}

// ロック用のMutex
var mutex sync.Mutex

// Personに新しいAccountを追加するメソッド
func (p *Person) addAccount(account *Account) {
	p.Accounts = append(p.Accounts, account)
}

// Accountから引き出しを行うメソッド
func (a *Account) withdraw(amount float64) {
	mutex.Lock()
	defer mutex.Unlock()
	if a.Balance >= amount {
		a.Balance -= amount
		fmt.Printf("Withdrawal successful. New balance: %.2f\n", a.Balance)
	} else {
		fmt.Println("Insufficient funds.")
	}
}

// クラスを跨いだ複雑な処理の例
func complexProcess(person *Person, account *Account, amount float64) {
	fmt.Printf("Processing transaction for %s (ID: %d)\n", person.Name, person.ID)
	person.addAccount(account)

	fmt.Printf("Account balance before withdrawal: %.2f\n", account.Balance)
	account.withdraw(amount)
	fmt.Printf("Account balance after withdrawal: %.2f\n", account.Balance)
}

func main() {
	// サンプルデータの作成
	person := &Person{ID: 1, Name: "John Doe"}
	account := &Account{ID: 101, Balance: 1000.0}

	// クラスを跨いだ複雑な処理の呼び出し
	complexProcess(person, account, 500.0)
}
