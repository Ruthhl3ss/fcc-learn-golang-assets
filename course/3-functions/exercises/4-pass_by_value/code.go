package main

import "fmt"

func main() {
	sendsSoFar := 430
	const sendsToAdd = 25
	fmt.Println("you've sent", incrementSends(sendsSoFar, sendsToAdd), "messages")
}

func incrementSends(sendsSoFar, sendsToAdd int) int {
	sendsSoFar = sendsSoFar + sendsToAdd
	return sendsSoFar
}
