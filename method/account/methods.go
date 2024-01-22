package account

import (
	"context"

	"github.com/google/wire"

	"github.com/Z00mZE/telegraph-api/model"
	"github.com/Z00mZE/telegraph-api/pkg/rest"
)

type Account struct {
	client *rest.Client
}

const (
	shortName  = `short_name,`
	authorName = `author_name`
	authorURL  = `author_url`
	authURL    = `auth_url`
	pageCount  = `page_count`
)

var (
	WireSetAccountMethods = wire.NewSet(NewAccount)
	defaultFields         = []string{
		shortName,
		authorName,
		authorURL,
		authURL,
		pageCount,
	}
)

func NewAccount() *Account {
	return &Account{
		client: rest.NewClient(),
	}
}

// CreateAccount
// Use this method to create a new Telegraph account. Most users only need one account,
// but this can be useful for channel administrators who would like to keep individual author names and profile links
// for each of their channels. On success, returns an Account object with the regular fields and an additional
// access_token field.
func (c *Account) CreateAccount(ctx context.Context, account model.Account) (string, *model.Account, error) {
	req := createAccountRequest{
		ShortName:  account.ShortName,
		AuthorName: account.AuthorName,
		AuthorURL:  account.AuthorURL,
	}

	resp := new(accountResponse)

	if respError := c.client.Send(ctx, createAccountURL, req, resp); respError != nil {
		return "", nil, respError
	}

	acc := &model.Account{
		AuthURL:    resp.AuthURL,
		ShortName:  resp.ShortName,
		AuthorName: resp.AuthorName,
		AuthorURL:  resp.AuthorURL,
		PageCount:  resp.PageCount,
	}
	return resp.AccessToken, acc, nil
}

// EditAccountInfo
// Use this method to update information about a Telegraph account. Pass only the parameters that you want to edit.
// On success, returns an Account object with the default fields.
func (c *Account) EditAccountInfo(ctx context.Context, token string, account model.Account) (*model.Account, error) {
	resp := new(model.Account)

	req := editAccountInfoRequest{
		AccessToken: token,
		ShortName:   account.ShortName,
		AuthorName:  account.AuthorName,
		AuthorURL:   account.AuthorURL,
	}

	if respError := c.client.Send(ctx, editAccountInfoURL, req, resp); respError != nil {
		return nil, respError
	}
	return resp, nil
}

// AccountInfo
// Use this method to get information about a Telegraph account. Returns an Account object on success.
func (c *Account) AccountInfo(ctx context.Context, token string) (*model.Account, error) {
	resp := new(model.Account)

	req := getAccountInfoRequest{
		AccessToken: token,
		Fields:      defaultFields,
	}

	if respError := c.client.Send(ctx, getAccountInfoURL, req, resp); respError != nil {
		return nil, respError
	}

	return resp, nil
}

// RevokeAccessToken
// Use this method to revoke access_token and generate a new one, for example, if the user would like to reset all
// connected sessions, or you have reasons to believe the token was compromised. On success, returns an Account object
// with new access_token and auth_url fields.
func (c *Account) RevokeAccessToken(ctx context.Context, token string) (string, *model.Account, error) {
	req := revokeAccessTokenRequest{AccessToken: token}
	resp := new(accountResponse)

	if respError := c.client.Send(ctx, revokeAccessTokenURL, req, resp); respError != nil {
		return "", nil, respError
	}

	acc := &model.Account{
		AuthURL:    resp.AuthURL,
		ShortName:  resp.ShortName,
		AuthorName: resp.AuthorName,
		AuthorURL:  resp.AuthorURL,
		PageCount:  resp.PageCount,
	}
	return resp.AccessToken, acc, nil
}
