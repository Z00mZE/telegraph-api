package account

import (
	"context"
	"reflect"
	"testing"

	"github.com/Z00mZE/telegraph-api/model"
	"github.com/Z00mZE/telegraph-api/pkg/rest"
)

func TestAccount_AccountInfo(t *testing.T) {
	tests := []struct {
		name    string
		token   string
		want    *model.Account
		wantErr bool
	}{
		{
			name:    "",
			token:   "d3b25feccb89e508a9114afb82aa421fe2a9712b963b387cc5ad71e58722",
			want:    nil,
			wantErr: false,
		},
	}

	ctx := context.Background()
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := NewAccount()
			if _, err := c.AccountInfo(ctx, tt.token); err != nil != tt.wantErr {
				t.Errorf("AccountInfo() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
		})
	}
}

func TestAccount_CreateAccount(t *testing.T) {
	type fields struct {
		client *rest.Client
	}
	type args struct {
		ctx     context.Context
		account model.Account
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    string
		want1   *model.Account
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := &Account{
				client: tt.fields.client,
			}
			got, got1, err := c.CreateAccount(tt.args.ctx, tt.args.account)
			if (err != nil) != tt.wantErr {
				t.Errorf("CreateAccount() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("CreateAccount() got = %v, want %v", got, tt.want)
			}
			if !reflect.DeepEqual(got1, tt.want1) {
				t.Errorf("CreateAccount() got1 = %v, want %v", got1, tt.want1)
			}
		})
	}
}

func TestAccount_EditAccountInfo(t *testing.T) {
	type fields struct {
		client *rest.Client
	}
	type args struct {
		ctx     context.Context
		token   string
		account model.Account
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    *model.Account
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := &Account{
				client: tt.fields.client,
			}
			got, err := c.EditAccountInfo(tt.args.ctx, tt.args.token, tt.args.account)
			if (err != nil) != tt.wantErr {
				t.Errorf("EditAccountInfo() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("EditAccountInfo() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestAccount_RevokeAccessToken(t *testing.T) {
	type fields struct {
		client *rest.Client
	}
	type args struct {
		ctx   context.Context
		token string
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    string
		want1   *model.Account
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := &Account{
				client: tt.fields.client,
			}
			got, got1, err := c.RevokeAccessToken(tt.args.ctx, tt.args.token)
			if (err != nil) != tt.wantErr {
				t.Errorf("RevokeAccessToken() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("RevokeAccessToken() got = %v, want %v", got, tt.want)
			}
			if !reflect.DeepEqual(got1, tt.want1) {
				t.Errorf("RevokeAccessToken() got1 = %v, want %v", got1, tt.want1)
			}
		})
	}
}

func TestNewAccount(t *testing.T) {
	tests := []struct {
		name string
		want *Account
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewAccount(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewAccount() = %v, want %v", got, tt.want)
			}
		})
	}
}
