package form3

import (
	"fmt"
	"net/http"
	"testing"
)

func Test_accountSvc_DeleteAccount(t *testing.T) {
	type args struct {
		fullPath string
		version  int64
	}

	tests := []struct {
		name       string
		svc        *accountSvc
		args       args
		wantErr    bool
		want       string
		statusCode int
	}{
		{
			name: "SVC_DELETE_ACCOUNT_SUCCES",
			args: args{
				fullPath: "/v1/organisation/accounts/form3id",
				version:  0,
			},
			wantErr:    false,
			want:       "",
			statusCode: http.StatusNoContent,
		},

		{
			name: "SVC_DELETE_ACCOUNT_INVALID_VERSION_NUMBER_FAILED",
			args: args{
				fullPath: fmt.Sprintf("/v1/organisation/accounts/%v", "testDelete"),
			},
			wantErr:    true,
			want:       `"error_message": "invalid version number"`,
			statusCode: http.StatusBadRequest,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			srv := httpMock(tt.args.fullPath, tt.statusCode, tt.want)

			defer srv.Close()
			url := srv.URL + tt.args.fullPath

			svc := &accountSvc{}
			err := svc.DeleteAccount(url)
			if (err != nil) != tt.wantErr {
				t.Errorf("accountSvc.DeleteAccount() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
