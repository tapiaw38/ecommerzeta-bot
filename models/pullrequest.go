package models

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
	Links       Links  `json:"links,omitempty"`
}

type Pullrequest struct {
	Title       string `json:"title,omitempty"`
	Description string `json:"description,omitempty"`
	Links       Links  `json:"links,omitempty"`
}

type PullrequestResponse struct {
	Actor       Actor       `json:"actor,omitempty"`
	Pullrequest Pullrequest `json:"pullrequest,omitempty"`
}
