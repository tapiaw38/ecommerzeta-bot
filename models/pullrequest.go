package models

type Actor struct {
	DisplayName string            `json:"display_name,omitempty"`
	Type        string            `json:"type,omitempty"`
	NickName    string            `json:"nick_name,omitempty"`
	Avatar      map[string]string `json:"avatar,omitempty"`
}

type Links struct {
	Html map[string]string
}

type Pullrequest struct {
	Title       string `json:"title,omitempty"`
	Description string `json:"description,omitempty"`
}

type PullrequestResponse struct {
	Actor       Actor       `json:"actor,omitempty"`
	Pullrequest Pullrequest `json:"pullrequest,omitempty"`
	Links       Links       `json:"links,omitempty"`
}
