package telegraph

import (
	"context"
	"sync"
	"time"

	"github.com/Z00mZE/telegraph-api/method/account"
	"github.com/Z00mZE/telegraph-api/method/page"
	"github.com/Z00mZE/telegraph-api/model"
)

type Telegraph struct {
	accessToken string
	account     *account.Account
	page        *page.Page
	wg          *sync.WaitGroup
}
type Token string

func newTelegraph(
	token string,
	account *account.Account,
	page *page.Page,
) *Telegraph {
	return &Telegraph{
		wg:          new(sync.WaitGroup),
		accessToken: token,
		account:     account,
		page:        page,
	}
}

func (c *Telegraph) CreateAccount(ctx context.Context, account model.Account) (*model.Account, error) {
	c.wg.Add(1)
	defer c.wg.Done()

	token, acc, accError := c.account.CreateAccount(ctx, account)
	if accError != nil {
		return nil, accError
	}

	c.accessToken = token
	return acc, accError
}
func (c *Telegraph) AccountInfo(ctx context.Context) (*model.Account, error) {
	c.wg.Wait()
	return c.account.AccountInfo(ctx, c.accessToken)
}
func (c *Telegraph) EditAccountInfo(ctx context.Context, account model.Account) (*model.Account, error) {
	c.wg.Wait()
	return c.account.EditAccountInfo(ctx, c.accessToken, account)
}
func (c *Telegraph) RevokeAccessToken(ctx context.Context) (*model.Account, error) {
	c.wg.Add(1)
	defer c.wg.Done()

	token, acc, accError := c.account.RevokeAccessToken(ctx, c.accessToken)
	if accError != nil {
		return nil, accError
	}

	c.accessToken = token
	return acc, accError
}

func (c *Telegraph) CreatePage(ctx context.Context, page model.Page) (*model.Page, error) {
	c.wg.Wait()
	return c.page.CreatePage(ctx, c.accessToken, page)
}
func (c *Telegraph) EditPage(ctx context.Context, account model.Page) (*model.Page, error) {
	c.wg.Wait()
	return c.page.EditPage(ctx, c.accessToken, account)
}
func (c *Telegraph) Page(ctx context.Context, page model.Page) (*model.Page, error) {
	c.wg.Wait()
	return c.page.GetPage(ctx, page.Path)
}
func (c *Telegraph) PageList(ctx context.Context, offset, limit uint) (*model.PageList, error) {
	c.wg.Wait()
	return c.page.GetPageList(ctx, c.accessToken, offset, limit)
}
func (c *Telegraph) Views(ctx context.Context, page model.Page, date time.Time) (*model.PageViews, error) {
	c.wg.Wait()
	return c.page.GetViews(ctx, page, date)
}
