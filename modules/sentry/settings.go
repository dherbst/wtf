package sentry

import (
	"github.com/olebedev/config"
	"github.com/wtfutil/wtf/cfg"
	"os"
)

const defaultTitle = "Sentry"

// Settings holds the information configured in the config.yml file to contact sentry for your projects.
type Settings struct {
	common   *cfg.Common
	apiKey   string
	projects []string
}

// NewSettingsFromYAML returns the sentry settings from the config.yml.
func NewSettingsFromYAML(name string, ymlConfig *config.Config, globalConfig *config.Config) *Settings {

	settings := Settings{
		common: cfg.NewCommonSettingsFromModule(name, defaultTitle, ymlConfig, globalConfig),
		apiKey: ymlConfig.UString("apiKey", os.Getenv("WTF_SENTRY_API_KEY")),
	}

	settings.projects = settings.arrayifyProjects(ymlConfig, globalConfig)

	return &settings
}

// arrayifyProjects figures out if we're dealing with a single project or an array of projects
func (settings *Settings) arrayifyProjects(ymlConfig *config.Config, globalConfig *config.Config) []string {
	projects := []string{}

	// Single project
	project, err := ymlConfig.String("project")
	if err == nil {
		projects = append(projects, project)
		return projects
	}

	// Array of projects
	projectList := ymlConfig.UList("project")
	for _, projectName := range projectList {
		if project, ok := projectName.(string); ok {
			projects = append(projects, project)
		}
	}

	return projects
}
