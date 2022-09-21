package models

type PullRequest struct {
	Title       string `json:"title,omitempty"`
	Description string `json:"description,omitempty"`
	Url         string `json:"url,omitempty"`
	Reviewers   string `json:"reviewers,omitempty"`
}

func (p PullRequest) PullRequestFormat() string {
	return "🎉​New pull request: " +
		p.Title + "\n" +
		"📝​Description: " +
		p.Description + "\n" +
		"👀​Reviewers: " + p.Reviewers
}
