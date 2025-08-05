package parser

import (
	"os"
	"pipedoc/internal/model"

	"gopkg.in/yaml.v3"
)

type GitHubWorkflow struct {
	Name string `yaml:"name"`
	Jobs map[string]struct {
		Needs  []string `yaml:"needs"`
		RunsOn string   `yaml:"runs-on"`
		Steps  []struct {
			Name string `yaml:"name"`
		} `yaml:"steps"`
	} `yaml:"jobs"`
}

func ParseGitHubActions(path string) (model.Pipeline, error) {
	data, err := os.ReadFile(path)
	if err != nil {
		return model.Pipeline{}, err
	}

	var wf GitHubWorkflow
	if err := yaml.Unmarshal(data, &wf); err != nil {
		return model.Pipeline{}, err
	}

	pipeline := model.Pipeline{Name: wf.Name}
	for jobName, job := range wf.Jobs {
		steps := []string{}
		for _, s := range job.Steps {
			steps = append(steps, s.Name)
		}
		pipeline.Jobs = append(pipeline.Jobs, model.Job{
			Name:   jobName,
			Needs:  job.Needs,
			RunsOn: job.RunsOn,
			Steps:  steps,
		})
	}
	return pipeline, nil
}
