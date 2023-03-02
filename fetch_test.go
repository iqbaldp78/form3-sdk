package form3

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"reflect"
	"testing"
)

var (
	Server *httptest.Server
)

type config struct {
	statusCode int
	response   interface{}
}

func (c *config) mockHandler(w http.ResponseWriter, r *http.Request) {
	var (
		resp []byte
	)

	rt := reflect.TypeOf(c.response)
	if rt.Kind() == reflect.String {
		resp = []byte(c.response.(string))
	} else if rt.Kind() == reflect.Struct || rt.Kind() == reflect.Ptr {
		resp, _ = json.Marshal(c.response)
	} else {
		resp = []byte("{}")
	}

	w.WriteHeader(c.statusCode)
	w.Write(resp)
}

func httpMock(pattern string, statusCode int, response interface{}) *httptest.Server {
	c := &config{statusCode, response}
	handler := http.NewServeMux()
	handler.HandleFunc(pattern, c.mockHandler)
	return httptest.NewServer(handler)
}

// func NewAccountSvcTest() IrepoAccountService {
// 	return &accountSvc{}
// }

func Test_accountSvc_FetchAccount(t *testing.T) {

	type fields struct {
		httpClient http.Client
	}
	type args struct {
		fullPath string
	}
	tests := []struct {
		name       string
		fields     fields
		args       args
		want       string
		statusCode int
		wantErr    bool
	}{
		{
			name: "SVC_FETCH_ACCOUNT_SUCCES",
			fields: fields{
				httpClient: *http.DefaultClient,
			},
			args: args{
				fullPath: "/v1/organisation/accounts/testAccountID",
			},
			want:       `{"msg": "ok"}`,
			statusCode: http.StatusOK,
			wantErr:    false,
		},
		{
			name: "SVC_FETCH_ACCOUNT_400_BAD_REQUEST",
			fields: fields{
				httpClient: *http.DefaultClient,
			},
			args: args{
				fullPath: "/v1/organisation/accounts/testAccountID",
			},
			want:       `{"error_message": "id is not a valid uuid"}`,
			statusCode: http.StatusBadRequest,
			wantErr:    true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			srv := httpMock(tt.args.fullPath, tt.statusCode, tt.want)
			defer srv.Close()

			svc := &accountSvc{}
			_, err := svc.FetchAccount(srv.URL + tt.args.fullPath)

			if (err != nil) != tt.wantErr {
				t.Errorf("accountSvc.FetchAccount() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
		})
	}
}
