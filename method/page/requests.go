package page

import (
	"github.com/Z00mZE/telegraph-api/model"
)

type createPageRequest struct {
	// Access token of the Telegraph account.
	AccessToken string `json:"access_token"`
	// Page title.
	Title string `json:"title"`
	// Author name, displayed below the article's title.
	AuthorName string `json:"author_name,omitempty"`
	// Profile link, opened when users click on the author's name below the title. Can be any link, not
	// necessarily to a Telegram profile or channel.
	AuthorURL string `json:"author_url,omitempty"`
	// Content of the page.
	Content []model.Node `json:"content"`
	// If true, a content field will be returned in the Page object.
	ReturnContent bool `json:"return_content,omitempty"`
}

type editPageRequest struct {
	// Access token of the Telegraph account.
	AccessToken string `json:"access_token"`

	// Path to the page.
	Path string `json:"path"`

	// Page title.
	Title string `json:"title"`

	// Content of the page.
	Content []model.Node `json:"content"`

	// Author name, displayed below the article's title.
	AuthorName string `json:"author_name,omitempty"`

	// Profile link, opened when users click on the author's name below the title. Can be any link, not
	// necessarily to a Telegram profile or channel.
	AuthorURL string `json:"author_url,omitempty"`

	// If true, a content field will be returned in the Page object.
	ReturnContent bool `json:"return_content,omitempty"`
}

type getPageRequest struct {
	// Path to the Telegraph page (in the format Title-12-31, i.e. everything that comes after http://telegra.ph/).
	Path string `json:"path"`

	// If true, content field will be returned in Page object.
	ReturnContent bool `json:"return_content,omitempty"`
}

type getPageListRequest struct {
	// Access token of the Telegraph account.
	AccessToken string `json:"access_token"`

	// Sequential number of the first page to be returned.
	Offset uint `json:"offset,omitempty"`

	// Limits the number of pages to be retrieved.
	Limit uint `json:"limit,omitempty"`
}

type getPageViewsRequest struct {
	// Path to the Telegraph page (in the format Title-12-31, where 12 is the month and 31 the day the article was
	// first published).
	Path string `json:"path"`

	// Required if month is passed. If passed, the number of page views for the requested year will be returned.
	Year int `json:"year,omitempty"`

	// Required if day is passed. If passed, the number of page views for the requested month will be returned.
	Month int `json:"month,omitempty"`

	// Required if hour is passed. If passed, the number of page views for the requested day will be returned.
	Day int `json:"day,omitempty"`

	// If passed, the number of page views for the requested hour will be returned.
	Hour int `json:"hour,omitempty"`
}
