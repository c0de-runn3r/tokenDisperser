package utils

import (
	"fmt"

	"github.com/go-numb/go-ftx/auth"
	"github.com/go-numb/go-ftx/rest"
	"github.com/go-numb/go-ftx/rest/private/account"
	"github.com/go-numb/go-ftx/rest/private/subaccount"
	"github.com/go-numb/go-ftx/rest/private/wallet"
)

func Token_send(numOfSubAccs int, ApiKey, SecretKey string) {
	client := rest.New(auth.New(ApiKey, SecretKey)) // create client
	c := client
	info, err := c.Information(&account.RequestForInformation{})
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(info)

	for i := 0; i < numOfSubAccs; i++ {
		subAccName := fmt.Sprintf("SubAccount_%v", i+1)

		subAccParams := subaccount.RequestForCreateSubAccount{
			NickName: subAccName,
		}
		createResponce, errCreate := c.CreateSubAccount(&subAccParams)
		if err != nil {
			fmt.Println(errCreate)
		}
		fmt.Println(createResponce)

		transferSubAccParams := subaccount.RequestForTransferSubAccount{ // can change params here for transfer for subaccs
			Coin:        "USDT",
			Size:        10.0,
			Source:      "main",
			Destination: subAccName,
		}
		transferResponce, errResponce := c.TransferSubAccount(&transferSubAccParams) // transfer from main acc to subacc
		if err != nil {
			fmt.Println(errResponce)
		}
		fmt.Println(transferResponce)
	}

	for i := 0; i < numOfSubAccs; i++ {

		subAccName := fmt.Sprintf("SubAccount_%v", i+1) // relogin with subacc
		clientWithSubAccounts := rest.New(
			auth.New(
				ApiKey,
				SecretKey,
				auth.SubAccount{
					UUID:     i,
					Nickname: subAccName,
				},
			))
		clientWithSubAccounts.Auth.UseSubAccountID(i) // change from main acc to subacc

		withdrawalParams := wallet.RequestForWithdraw{ // can change params here for withdrawal
			Coin:     "USDT",
			Size:     10.0,
			Address:  "0x123123123",
			Methods:  "matic", // matic, bsc, trx, sol, erc20
			Password: "12345678",
		}
		responce, err := c.Withdraw(&withdrawalParams)
		if err != nil {
			fmt.Println(err)
		}
		fmt.Println(responce)

	}
}
