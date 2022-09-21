package models

type Author struct {
	DisplayName string `json:"display_name,omitempty"`
	Type        string `json:"type,omitempty"`
	NickName    string `json:"nick_name,omitempty"`
}

type Pullrequest struct {
	Title       string `json:"title,omitempty"`
	Description string `json:"description,omitempty"`
}

type PullrequestResponse struct {
	Author      Author      `json:"author,omitempty"`
	Pullrequest Pullrequest `json:"pullrequest,omitempty"`
}

func (p PullrequestResponse) PullRequestFormat() string {
	return "🎉​New pull request: " +
		p.Pullrequest.Title + "\n" +
		"📝​Description: " +
		p.Pullrequest.Description + "\n" +
		"👀​Author: " + p.Author.DisplayName
}
