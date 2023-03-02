package form3

import (
	"net/http"
	"testing"

	"github.com/iqbaldp78/form3/models"
)

func Test_accountSvc_CreateAccount(t *testing.T) {
	country := "GB"
	payloadDummy := models.AccountData{
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
	}
	type args struct {
		fullPath string
		payload  models.AccountData
	}
	tests := []struct {
		name       string
		svc        *accountSvc
		args       args
		want       string
		statusCode int
		wantErr    bool
	}{
		{
			name: "SVC_CREATE_ACCOUNT_SUCCES",
			args: args{
				fullPath: "/v1/organisation/accounts",
				payload:  payloadDummy,
			},
			want:       `{"msg": "ok"}`,
			statusCode: http.StatusOK,
			wantErr:    false,
		},

		{
			name: "SVC_CREATE_ACCOUNT_FAILED_DUPLICATE_ID",
			args: args{
				fullPath: "/v1/organisation/accounts",
				payload:  payloadDummy,
			},
			want:       `"error_message": "Account cannot be created as it violates a duplicate constraint"`,
			statusCode: http.StatusConflict,
			wantErr:    true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			srv := httpMock(tt.args.fullPath, tt.statusCode, tt.want)
			defer srv.Close()

			svc := &accountSvc{}
			url := srv.URL + tt.args.fullPath
			_, err := svc.CreateAccount(url, tt.args.payload)

			if (err != nil) != tt.wantErr {
				t.Errorf("accountSvc.CreateAccount() error = %v, wantErr %v", err, tt.wantErr)
				return
			}

			// if tt.wantErr {
			// 	if !reflect.DeepEqual(err.Error(), tt.want) {
			// 		t.Errorf("accountSvc.CreateAccount() = %v, want %v", err.Error(), tt.want)
			// 		return
			// 	}
			// }

		})
	}
}

// "error_message": "Account cannot be created as it violates a duplicate constraint"
// "error_message": "Account cannot be created as it violates a duplicate constraint"
