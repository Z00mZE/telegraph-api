package account

type createAccountRequest struct {
	// Account name, helps users with several accounts remember which they are currently using. Displayed to the
	// user above the "Edit/Publish" button on Telegra.ph, other users don't see this name.
	ShortName string `json:"short_name"`
	// Default author name used when creating new articles.
	AuthorName string `json:"author_name,omitempty"`
	// Default profile link, opened when users click on the author's name below the title. Can be any link, not
	// necessarily to a Telegram profile or channel.
	AuthorURL string `json:"author_url,omitempty"`
}
type accountResponse struct {
	AccessToken string
	// URL to authorize a browser on telegra.ph and connect it to a Telegraph account. This URL is valid
	// for only one use and for 5 minutes only.
	AuthURL string `json:"auth_url,omitempty"`

	// Account name, helps users with several accounts remember which they are currently using. Displayed
	// to the user above the "Edit/Publish" button on Telegra.ph, other users don't see this name.
	ShortName string `json:"short_name"`

	// Default author name used when creating new articles.
	AuthorName string `json:"author_name"`

	// Profile link, opened when users click on the author's name below the title. Can be any link, not
	// necessarily to a Telegram profile or channel.
	AuthorURL string `json:"author_url"`

	// Number of pages belonging to the Telegraph account.
	PageCount int `json:"page_count,omitempty"`
}

type editAccountInfoRequest struct {
	// Access token of the Telegraph account.
	AccessToken string `json:"access_token"`
	// New account name.
	ShortName string `json:"short_name,omitempty"`
	// New default author name used when creating new articles.
	AuthorName string `json:"author_name,omitempty"`
	// New default profile link, opened when users click on the author's name below the title. Can be any link,
	// not necessarily to a Telegram profile or channel.
	AuthorURL string `json:"author_url,omitempty"`
}

type getAccountInfoRequest struct {
	// Access token of the Telegraph account.
	AccessToken string `json:"access_token"`
	// List of account fields to return.
	Fields []string `json:"fields,omitempty"`
}

type revokeAccessTokenRequest struct {
	AccessToken string
}
