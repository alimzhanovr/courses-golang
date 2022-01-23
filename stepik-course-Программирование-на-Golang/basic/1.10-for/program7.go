package main

func NumberOfYears(initPayment, percent, finalAmount int) int {
	if initPayment < 0 || percent < 0 || finalAmount < initPayment {
		return -1
	}
	counter, balance := 0, initPayment
	for balance < finalAmount {
		margin := int(balance * percent / 100.0)
		balance += margin
		counter++
	}
	return counter
}
