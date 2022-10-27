package main

import "fmt"

type BankAccountInfo struct {
	HolderName string
	AcccountNumber  uint
	AccountBalance float64
}

func main(){
	account1 := new(BankAccountInfo);
	account1.HolderName = "Alice Women"
	account1.AcccountNumber = 293293198494912
	account1.AccountBalance = 234994
	
	account2 := new(BankAccountInfo);
	account2.HolderName = "Bob Man"
	account2.AcccountNumber = 94495694849549
	account2.AccountBalance = 949598
	checkBalance(account1)
	sendMoney(account1, account2)
}
func checkBalance(b *BankAccountInfo) {
	if b.AccountBalance > 0 {
		fmt.Printf("User %s has an account balance of $%f\n", b.HolderName, b.AccountBalance)
	}
}
func (b *BankAccountInfo) validAccount() bool {
	if b.AcccountNumber == 0 {
		fmt.Println("There is no account.")
	}
	return false
}
func sendMoney(b1 *BankAccountInfo, b2 *BankAccountInfo) {
	fmt.Println("Welcome")
	if b1.AcccountNumber < 0 {
		if b1.AccountBalance > 0 {
			fmt.Println("would you like to send money")
			var answer string;
			fmt.Scanln(answer)
		}
	}
}