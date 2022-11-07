package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"

	"tkn_mlt/utils"
)

func main() {
	fmt.Println("Hello!")
	reader := bufio.NewReader(os.Stdin)
	fmt.Print("Enter number of subAccounts: ")
	input, _ := reader.ReadString('\n')
	num, _ := strconv.Atoi(input)
	utils.Token_send(num, ApiKey, SecretKey)
	utils.Connect_to_eth()
	utils.Generate_account()
}
