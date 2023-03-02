package main

import (
	"fmt"
	"log"

	"github.com/iqbaldp78/form3"
	"github.com/iqbaldp78/form3/models"
)

func main() {
	var (
		version int64
	)
	client := form3.NewClient("http://localhost:8030")

	country := "GB"
	version = 1
	payload := models.PayloadCreateAccount{
		Attributes: &models.AccountAttributes{
			BankID:                  "123456",
			AccountNumber:           "123",
			BankIDCode:              "GBDSC",
			BaseCurrency:            "GBP",
			Bic:                     "EXMPLGB2XXX",
			Name:                    []string{"bob"},
			SecondaryIdentification: "SecondaryIdentification",
			Country:                 &country,
		},
		Version: &version,
	}
	resulCreate, err := client.CreateAccount(payload)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("resulCreate %+v \n", resulCreate)

	resultFetch, err := client.FetchAccount(resulCreate.ID)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("resultFetch %+v \n", resultFetch)

	err = client.DeleteAccount(resulCreate.ID, 0)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("delete request client error  %+v \n", err)
}
