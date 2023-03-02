package form3

import (
	"encoding/json"
	"fmt"

	"github.com/iqbaldp78/form3/models"
)

const ()

type Config struct {
}

//CLient used to set Client form3 api
type Client struct {
	HostUrl string
	svc     IrepoAccountService
}

func NewClient(hostUrl string) *Client {
	return &Client{
		HostUrl: hostUrl,
		svc:     newAccountSvc(),
	}
}

func (c *Client) FetchAccount(form3id string) (models.AccountData, error) {
	fullPath := fmt.Sprintf("%v/v1/organisation/accounts/%v", c.HostUrl, form3id)
	return c.svc.FetchAccount(fullPath)
}

func (c *Client) CreateAccount(payload models.PayloadCreateAccount) (models.AccountData, error) {
	accountData := models.AccountData{}
	fullPath := fmt.Sprintf("%v/v1/organisation/accounts", c.HostUrl)
	mars, err := json.Marshal(payload)
	if err != nil {
		return accountData, err
	}

	if err := json.Unmarshal(mars, &accountData); err != nil {
		return accountData, err
	}
	accountData.SetDefault()
	return c.svc.CreateAccount(fullPath, accountData)
}

func (c *Client) DeleteAccount(form3id string, version int64) error {
	fullPath := fmt.Sprintf("%v/v1/organisation/accounts/%v?version=%d", c.HostUrl, form3id, version)
	return c.svc.DeleteAccount(fullPath)
}
