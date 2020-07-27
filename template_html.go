package cnv

const TemplateHtml = `{{- define "item"}}
<li>{{if .Scope}}{{.Scope }}{{end}}{{.Title}}{{range .Tickets}} {{.}}{{end}} <span>({{.Author}})</span></li>
{{- end -}}
<h2>[{{.Version}}] - {{.Date}}</h2>
{{- if .BreakingChanges}}
<h3>BREAKING CHANGES</h3>
<ul>
    {{- range .BreakingChanges}}
        {{- template "item" .}}
    {{- end}}
</ul>
{{- end}}
{{- if .Added -}}
<h3>Added</h3>
<ul>
    {{- range .Added}}
        {{- template "item" .}}
    {{- end}}
</ul>
{{- end}}
{{- if .Fixed}}
<h3>Fixed</h3>
<ul>
    {{- range .Fixed}}
        {{- template "item" .}}
    {{- end}}
</ul>
{{- end}}
{{- if .Changed}}
<h3>Changed</h3>
<ul>
    {{- range .Changed}}
        {{- template "item" .}}
    {{- end}}
</ul>
{{- end}}
{{- if .WorkInProgress}}
<h3>Unfinished</h3>
<ul>
    {{- range .WorkInProgress}}
        {{- template "item" .}}
    {{- end}}
</ul>
{{- end}}`
