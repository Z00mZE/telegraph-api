package model

// Page represents a page on Telegraph.
type Page struct {
	// Path to the page.
	Path string `json:"path"`
	// URL of the page.
	URL string `json:"url"`
	// Title of the page.
	Title string `json:"title"`
	// Description of the page.
	Description string `json:"description"`
	// Name of the author, displayed below the title.
	AuthorName string `json:"author_name,omitempty"`
	// Profile link, opened when users click on the author's name below the title. Can be any link, not
	// necessarily to a Telegram profile or channel.
	AuthorURL string `json:"author_url,omitempty"`
	// Image URL of the page.
	ImageURL string `json:"image_url,omitempty"`
	// Content of the page.
	Content []Node `json:"content,omitempty"`
	// Number of page views for the page.
	Views int `json:"views"`
	// Only returned if access_token passed. True, if the target Telegraph account can edit the page.
	CanEdit bool `json:"can_edit,omitempty"`
}
