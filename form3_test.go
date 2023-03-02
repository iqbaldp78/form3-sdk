package form3

import (
	"errors"
	"reflect"
	"testing"

	"github.com/iqbaldp78/form3/mocks"
	"github.com/iqbaldp78/form3/models"
	"github.com/stretchr/testify/mock"
)

func TestNewClient(t *testing.T) {
	svc := newAccountSvc()
	type args struct {
		hostUrl string
		svc     IrepoAccountService
	}
	tests := []struct {
		name string
		args args
		want *Client
	}{
		{
			name: "Test New Client",
			args: args{
				hostUrl: "testSuccess",
				svc:     svc,
			},
			want: &Client{
				HostUrl: "testSuccess",
				svc:     svc,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewClient(tt.args.hostUrl); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewClient() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestClient_FetchAccount(t *testing.T) {
	type fields struct {
		HostUrl string
		svc     IrepoAccountService
	}
	type args struct {
		form3id  string
		fullPath string
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    models.AccountData
		wantErr bool
	}{
		{
			name: "SUCCESS_FETCH_DATA",
			fields: fields{
				HostUrl: "http://localhost:8030",
			},
			args: args{form3id: "testAccountID", fullPath: "http://localhost:8030/v1/organisation/accounts/testAccountID"},
			want: models.AccountData{
				Type:           "accounts",
				ID:             "testID",
				OrganisationID: "testOrgID",
				Attributes: &models.AccountAttributes{
					BankID: "GB",
				},
			},
			wantErr: false,
		},
		{
			name: "FAILED_FETCH_DATA_WRONG_HOST",
			fields: fields{
				HostUrl: "http://wrongHost:8030",
			},
			args: args{form3id: "testAccountID", fullPath: "http://wrongHost:8030/v1/organisation/accounts/testAccountID"},
			want: models.AccountData{
				Type: "accounts",
				Attributes: &models.AccountAttributes{
					BankID: "GB",
				},
			},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			mocksvc := new(mocks.IrepoAccountService)

			if !tt.wantErr {
				mocksvc.On("FetchAccount", tt.args.fullPath).Return(tt.want, nil)
			} else {
				mocksvc.On("FetchAccount", tt.args.fullPath).Return(tt.want, errors.New("failed to fetch account"))
			}

			tt.fields.svc = mocksvc

			c := &Client{
				HostUrl: tt.fields.HostUrl,
				svc:     tt.fields.svc,
			}
			got, err := c.FetchAccount(tt.args.form3id)
			if (err != nil) != tt.wantErr {
				t.Errorf("Client.FetchAccount() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Client.FetchAccount() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestClient_CreateAccount(t *testing.T) {
	type fields struct {
		HostUrl string
		svc     IrepoAccountService
	}
	type args struct {
		payload  models.PayloadCreateAccount
		fullPath string
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    models.AccountData
		wantErr bool
	}{
		{
			name: "SUCCESS_CREATE_ACCOUNT",
			fields: fields{
				HostUrl: "http://correctHOST:8030",
			},
			args: args{
				fullPath: "http://correctHOST:8030/v1/organisation/accounts",
				payload: models.PayloadCreateAccount{
					Attributes: &models.AccountAttributes{
						AccountNumber: "testAccNumber",
					},
				},
			},
			want: models.AccountData{
				Attributes: &models.AccountAttributes{
					AccountNumber: "testAccNumber",
				},
			},
			wantErr: false,
		},

		{
			name: "FAIL_MARSHALL_PAYLOAD_CREATE_ACCOUNT",
			fields: fields{
				HostUrl: "http://correctHOST:8030",
			},
			args: args{
				fullPath: "http://correctHOST:8030/v1/organisation/accounts",
				payload:  models.PayloadCreateAccount{},
			},
			want:    models.AccountData{},
			wantErr: true,
		},
		{
			name: "FAIL_UNMARSHALL_PAYLOAD_CREATE_ACCOUNT",
			fields: fields{
				HostUrl: "http://correctHOST:8030",
			},
			args: args{
				fullPath: "http://correctHOST:8030/v1/organisation/accounts",
				payload:  models.PayloadCreateAccount{},
			},
			want:    models.AccountData{},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			mocksvc := new(mocks.IrepoAccountService)
			if !tt.wantErr {
				mocksvc.On("CreateAccount", tt.args.fullPath, mock.Anything).Return(tt.want, nil)
			} else {
				mocksvc.On("CreateAccount", tt.args.fullPath, mock.Anything).Return(tt.want, errors.New("failed to create account"))
			}

			tt.fields.svc = mocksvc

			c := &Client{
				HostUrl: tt.fields.HostUrl,
				svc:     tt.fields.svc,
			}
			got, err := c.CreateAccount(tt.args.payload)

			if (err != nil) != tt.wantErr {
				t.Errorf("Client.CreateAccount() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Client.CreateAccount() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestClient_DeleteAccount(t *testing.T) {
	type fields struct {
		HostUrl string
		svc     IrepoAccountService
	}
	type args struct {
		fullPath string
		form3id  string
		version  int64
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		wantErr bool
	}{
		{
			name: "SUCCESS_DELETE_ACCOUNT",
			fields: fields{
				HostUrl: "http://localhost:8030",
			},
			args: args{
				form3id:  "testAccountID",
				fullPath: "http://localhost:8030/v1/organisation/accounts/testAccountID?version=0",
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			mocksvc := new(mocks.IrepoAccountService)
			if !tt.wantErr {
				mocksvc.On("DeleteAccount", tt.args.fullPath).Return(nil)
			} else {
				mocksvc.On("DeleteAccount", tt.args.fullPath).Return(errors.New("failed to fetch account"))
			}
			tt.fields.svc = mocksvc

			c := &Client{
				HostUrl: tt.fields.HostUrl,
				svc:     tt.fields.svc,
			}
			if err := c.DeleteAccount(tt.args.form3id, tt.args.version); (err != nil) != tt.wantErr {
				t.Errorf("Client.DeleteAccount() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
