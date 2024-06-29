package signal

import "fmt"

func Deposit(deposit int, rate int, threshold int) int {
	var balance float32 = 0
	count := 0
	balance += float32(deposit)
	println(fmt.Sprintf("%f", balance))
	//balance = float32(deposit)

	for balance < float32(threshold) {
		increaseBalance := float32(balance) * float32(float32(rate)/float32(100))
		balance += increaseBalance
		println(fmt.Sprintf("%f", balance))
		count++
	}

	return count
}
