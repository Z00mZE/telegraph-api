package page

import (
	"context"
	"time"

	"github.com/google/wire"

	"github.com/Z00mZE/telegraph-api/model"
	"github.com/Z00mZE/telegraph-api/pkg/rest"
)

type Page struct {
	token  string
	client *rest.Client
}

var WireSetPageMethods = wire.NewSet(NewPage)

func NewPage() *Page {
	return &Page{
		client: rest.NewClient(),
	}
}

// CreatePage create a new Telegraph page. On success, returns a Page object.
func (c *Page) CreatePage(ctx context.Context, token string, page model.Page) (*model.Page, error) {
	resp := new(model.Page)

	req := createPageRequest{
		AccessToken:   token,
		Title:         page.Title,
		AuthorName:    page.AuthorName,
		AuthorURL:     page.AuthorURL,
		Content:       page.Content,
		ReturnContent: true,
	}

	if respError := c.client.Send(ctx, createPageURL, req, resp); respError != nil {
		return nil, respError
	}
	return resp, nil
}

// EditPage edit an existing Telegraph page. On success, returns a Page object.
func (c *Page) EditPage(ctx context.Context, token string, page model.Page) (*model.Page, error) {
	resp := new(model.Page)

	req := editPageRequest{
		AccessToken:   token,
		Path:          page.Path,
		Title:         page.Title,
		Content:       page.Content,
		AuthorName:    page.AuthorName,
		AuthorURL:     page.AuthorURL,
		ReturnContent: true,
	}

	if respError := c.client.Send(ctx, editPageURL, req, resp); respError != nil {
		return nil, respError
	}
	return resp, nil
}

// GetPage get a Telegraph page. Returns a Page object on success.
func (c *Page) GetPage(ctx context.Context, path string) (*model.Page, error) {
	resp := new(model.Page)

	req := getPageRequest{
		Path:          path,
		ReturnContent: true,
	}

	if respError := c.client.Send(ctx, getPageURL, req, resp); respError != nil {
		return nil, respError
	}
	return resp, nil
}

// GetPageList get a list of pages belonging to a Telegraph account. Returns a PageList object, sorted by most
// recently created pages first.
func (c *Page) GetPageList(ctx context.Context, token string, offset, limit uint) (*model.PageList, error) {
	resp := new(model.PageList)

	req := getPageListRequest{
		AccessToken: token,
		Limit:       limit,
		Offset:      offset,
	}

	if respError := c.client.Send(ctx, getPageURL, req, resp); respError != nil {
		return nil, respError
	}
	return resp, nil
}

// GetViews get the number of views for a Telegraph article. By default, the total number of page views will be
// returned. Returns a PageViews object on success.
func (c *Page) GetViews(ctx context.Context, page model.Page, date time.Time) (*model.PageViews, error) {
	resp := new(model.PageViews)

	req := getPageViewsRequest{Path: page.Path}

	if !date.IsZero() {
		req.Year = date.Year()
		req.Month = int(date.Month())
		req.Day = date.Day()
		req.Hour = date.Hour()
	}

	if respError := c.client.Send(ctx, getPageViewURL, req, resp); respError != nil {
		return nil, respError
	}
	return resp, nil
}
