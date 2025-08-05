package generator

import (
	"fmt"
	"pipedoc/internal/model"
	"strings"
)

func GenerateMarkdown(p model.Pipeline) (string, string) {
	var md strings.Builder
	md.WriteString(fmt.Sprintf("# %s CI Pipeline\n\n", p.Name))
	md.WriteString("| Job | Runs On | Needs |\n|-----|---------|-------|\n")
	for _, j := range p.Jobs {
		md.WriteString(fmt.Sprintf("| %s | %s | %s |\n",
			j.Name, j.RunsOn, strings.Join(j.Needs, ", ")))
	}

	mermaid := generateMermaid(p)
	md.WriteString("\n## Diagram\n```mermaid\n" + mermaid + "\n```\n")
	return md.String(), mermaid
}

func generateMermaid(p model.Pipeline) string {
	var sb strings.Builder
	sb.WriteString("flowchart TD\n")
	for _, j := range p.Jobs {
		sb.WriteString(fmt.Sprintf("    %s[%s]\n", j.Name, j.Name))
		for _, dep := range j.Needs {
			sb.WriteString(fmt.Sprintf("    %s --> %s\n", dep, j.Name))
		}
	}
	return sb.String()
}
