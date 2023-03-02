package form3

import (
	"net/http"

	"github.com/iqbaldp78/form3/models"
)

type accountSvc struct {
}

type IrepoAccountService interface {
	FetchAccount(fullPath string) (models.AccountData, error)
	CreateAccount(fullPath string, payload models.AccountData) (models.AccountData, error)
	DeleteAccount(fullPath string) error
}

func newAccountSvc() IrepoAccountService {
	return &accountSvc{}
}

func doRequest(req *http.Request) (*http.Response, error) {
	response, err := http.DefaultClient.Do(req)
	if err != nil {
		return nil, err
	}
	return response, nil
}
