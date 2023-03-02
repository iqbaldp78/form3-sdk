package form3

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

func (svc *accountSvc) DeleteAccount(fullPath string) error {

	request, err := http.NewRequest(http.MethodDelete, fullPath, nil)
	if err != nil {
		return err
	}
	// log.Printf("request %+v", request)
	resp, err := doRequest(request)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	resBody, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return err
	}
	// fmt.Println("resp.StatusCode", resp.StatusCode)
	// fmt.Printf("client: response body: %s\n", resBody)
	if resp.StatusCode < http.StatusOK || resp.StatusCode > http.StatusMultipleChoices {
		return fmt.Errorf(string(resBody))
	}
	return nil
}
