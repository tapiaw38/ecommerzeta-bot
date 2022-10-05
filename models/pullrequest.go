package models

type Repository struct {
	FullName string `json:"full_name,omitempty"`
	Name     string `json:"name,omitempty"`
}

type Avatar struct {
	Href string `json:"href,omitempty"`
}

type Html struct {
	Href string `json:"href,omitempty"`
}

type Links struct {
	Avatar Avatar `json:"avatar,omitempty"`
	Html   Html   `json:"html,omitempty"`
}

type Actor struct {
	DisplayName string `json:"display_name,omitempty"`
	Type        string `json:"type,omitempty"`
	NickName    string `json:"nick_name,omitempty"`
	AccountId   string `json:"account_id,omitempty"`
	Links       Links  `json:"links,omitempty"`
}

type Reviewer struct {
	DisplayName string `json:"display_name,omitempty"`
	Links       Links  `json:"links,omitempty"`
	AccountId   string `json:"account_id,omitempty"`
}

type Pullrequest struct {
	Title       string     `json:"title,omitempty"`
	Description string     `json:"description,omitempty"`
	Links       Links      `json:"links,omitempty"`
	Reviewers   []Reviewer `json:"reviewers,omitempty"`
}

type PullrequestResponse struct {
	Repository  Repository  `json:"repository,omitempty"`
	Actor       Actor       `json:"actor,omitempty"`
	Pullrequest Pullrequest `json:"pullrequest,omitempty"`
}
