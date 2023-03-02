package form3

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/iqbaldp78/form3/models"
)

func (svc *accountSvc) CreateAccount(fullPath string, payload models.AccountData) (models.AccountData, error) {

	payloadData := struct {
		Data models.AccountData `json:"data"`
	}{
		Data: payload,
	}
	var buf bytes.Buffer
	err := json.NewEncoder(&buf).Encode(payloadData)
	if err != nil {
		return models.AccountData{}, err
	}

	// log.Println(buf.String())
	request, err := http.NewRequest(http.MethodPost, fullPath, &buf)
	if err != nil {
		return models.AccountData{}, err
	}

	resp, err := doRequest(request)
	if err != nil {
		fmt.Printf("client: could not read response body: %s\n", err)
		return models.AccountData{}, err
	}
	defer resp.Body.Close()

	resBody, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Printf("client: could not read response body: %s\n", err)
		return models.AccountData{}, err
	}

	// log.Println("resp status code ", resp.StatusCode)
	// fmt.Printf("client: response body: %s\n", resBody)
	if resp.StatusCode < http.StatusOK || resp.StatusCode > http.StatusMultipleChoices {
		return models.AccountData{}, fmt.Errorf(string(resBody))
	}

	data := struct {
		AccountData models.AccountData `json:"data"`
	}{}

	if err := json.Unmarshal(resBody, &data); err != nil {
		return models.AccountData{}, err
	}
	// fmt.Printf("client: response data: %+v\n", data.AccountData.Attributes)
	return data.AccountData, nil
}
