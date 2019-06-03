package sentry

// Issue is an Issue from the sentry api.
type Issue struct {
	ProjectName string
	Title       string
	LastSeen    string
}

// GetIssues returns the list of issues for the projects configured in the settings.
func (widget *Widget) GetIssues() ([]Issue, error) {
	return make([]Issue, 0), nil
}

// GetIssues on the client will get issues for org and project
// c.GetIssues(org, project, stats, shortid, query) ([]Issue, *Link, error)
