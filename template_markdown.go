package cnv

const TemplateMarkdown = `{{- define "item"}}
- {{if .Scope}}**{{.Scope }}** {{end}}{{.Title}}{{range .Tickets}} {{.}}{{end}} {{.ShortHash}}
{{- end -}}
### [{{.Version}}] - {{.Date}}
{{- if .BreakingChanges}}
#### BREAKING CHANGES
    {{- range .BreakingChanges}}
        {{- template "item" . -}}
    {{end -}}
{{end -}}
{{- if .Added}}
#### Added
   {{- range .Added}}
     {{- template "item" . -}}
   {{end -}}
{{end}}
{{- if .Fixed}}
#### Fixed
    {{- range .Fixed}}
        {{- template "item" . -}}
    {{end -}}
{{end}}
{{- if .Changed}}
#### Changed
    {{- range .Changed}}
        {{- template "item" . -}}
    {{end -}}
{{end}}
{{- if .WorkInProgress}}
#### Unfinished
    {{- range .WorkInProgress}}
        {{- template "item" . -}}
    {{end -}}
{{end}}`
