// Code generated by Wire. DO NOT EDIT.

//go:generate go run github.com/google/wire/cmd/wire
//go:build !wireinject
// +build !wireinject

package telegraph

import (
	"github.com/Z00mZE/telegraph-api/method/account"
	"github.com/Z00mZE/telegraph-api/method/page"
)

// Injectors from wire.go:

func NewTelegraphAPI(token string) *Telegraph {
	accountAccount := account.NewAccount()
	pagePage := page.NewPage()
	telegraph := newTelegraph(token, accountAccount, pagePage)
	return telegraph
}