package account

import (
	"context"
	"reflect"
	"testing"
	"time"

	"github.com/Z00mZE/telegraph-api/model"
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

	ctx := context.Background()

	t.Run("CreateAccount", func(t *testing.T) {
		account := model.Account{
			ShortName:  "ShortName",
			AuthorName: "AuthorName",
		}

		c := NewAccount()
		token, _, err := c.CreateAccount(ctx, account)
		if token == "" {
			t.Error("CreateAccount() want not empty token, got empty")
			return
		}
		if err != nil {
			t.Errorf("CreateAccount() error = %v", err)
			return
		}
	})
}

func TestAccount_EditAccountInfo(t *testing.T) {
	ctx := context.Background()

	c := NewAccount()

	t.Run("EditAccountInfo", func(t *testing.T) {
		account := model.Account{
			ShortName:  "ShortName",
			AuthorName: "AuthorName",
		}

		token, _, hasError := c.CreateAccount(ctx, account)
		if hasError != nil {
			t.Errorf("CreateAccount() error = %v", hasError)
			return
		}

		account.AuthorName = time.Now().String()

		got, err := c.EditAccountInfo(ctx, token, account)
		if err != nil {
			t.Errorf("EditAccountInfo() error = %v", err)
			return
		}
		if !reflect.DeepEqual(got.AuthorName, account.AuthorName) {
			t.Errorf("EditAccountInfo() got = %v, want %v", got, account.AuthorName)
		}
	})
}

func TestAccount_RevokeAccessToken(t *testing.T) {
	ctx := context.Background()

	c := NewAccount()

	t.Run("RevokeAccessToken", func(t *testing.T) {
		user := model.Account{
			ShortName:  "ShortName",
			AuthorName: "AuthorName",
		}

		oldToken, _, hasError := c.CreateAccount(ctx, user)
		if hasError != nil {
			t.Errorf("CreateAccount() error = %v", hasError)
			return
		}

		if _, _, err := c.RevokeAccessToken(ctx, oldToken); err != nil {
			t.Errorf("RevokeAccessToken() error = %v", err)
			return
		}

		if _, _, err := c.RevokeAccessToken(ctx, oldToken); err == nil {
			t.Error("RevokeAccessToken() required error? but not got")
			return
		}
	})
}
