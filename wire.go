//go:build wireinject
// +build wireinject

package telegraph

import (
	"github.com/google/wire"

	"github.com/Z00mZE/telegraph-api/method/account"
	"github.com/Z00mZE/telegraph-api/method/page"
)

func NewTelegraphAPI(token string) *Telegraph {
	wire.Build(
		wire.NewSet(
			newTelegraph,
			account.WireSetAccountMethods,
			page.WireSetPageMethods,
		),
	)
	return new(Telegraph)
}
