package main

import (
	"fmt"

	"github.com/qpyparty/wdonate"
)

func main() {
	// Replace <API TOKEN> with your actual API token obtained from your wdonate.ru account
	token := "<API TOKEN>"
	// Replace 0 with your actual VK Bot ID, ensuring it doesn't contain a sign
	botId := 0

	client := wdonate.NewClient(token, botId)

	// Execute a request to get the balance
	response, err := client.GetBalance(wdonate.GetBalanceRequest{})
	if err != nil {
		panic(err)
	}

	fmt.Println(response.Balance)
}
