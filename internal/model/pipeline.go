package model

type Pipeline struct {
	Name string
	Jobs []Job
}

type Job struct {
	Name   string
	Needs  []string
	RunsOn string
	Steps  []string
}
