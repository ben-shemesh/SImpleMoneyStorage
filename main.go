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

	account3 := new(BankAccountInfo);
	account3.HolderName = "Tom Hardey"
	account3.AcccountNumber = 04050450;
	account3.AccountBalance = 0
	account3.checkBalance()
	
	sendMoney(account1, account2)
	account1.checkBalance()
	account1.validAccount()

	account2.checkBalance()
}
func (b *BankAccountInfo) checkBalance() bool {
	if b.AccountBalance > 0 {
		fmt.Printf("User %s has an account balance of $%f\n", b.HolderName, b.AccountBalance)
		return true
	} else {
		return false
	}
}
func (b *BankAccountInfo) validAccount() bool {
	if b.AcccountNumber == 0 {
		fmt.Println("There is no account.")
	}else {fmt.Println("There is an account associated with this account number.")}
	return false
}
func sendMoney(b1 *BankAccountInfo, b2 *BankAccountInfo) {
	if b1.checkBalance() == true {
	}
	var answer string;
	fmt.Println("Welcome")
	fmt.Println("would you like to send money")
	fmt.Scan(&answer)
	if answer == "yes" || answer == "YES" || answer == "y" || answer == "Y" {
		fmt.Println("Your answer was yes")
	} else { fmt.Println("Please come back again.")}
	if b1.AcccountNumber < 0 {
		if b1.AccountBalance > 0 {
		}
	}
}