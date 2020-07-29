package cnv

const TemplateHtml = `{{- define "item"}}
<li>{{if .Scope}}{{.Scope }}{{end}}{{.Title}}{{range .Tickets}} {{.}}{{end}} <span>({{.Author}})</span></li>
{{- end -}}
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
{{- if .Reverted}}
<h3>Rolled back</h3>
<ul>
    {{- range .Reverted}}
        {{- template "item" .}}
    {{- end}}
</ul>
{{- end}}
{{- if .Other}}
<h3>Other</h3>
<ul>
    {{- range .Other -}}
<li>{{if .Type}}{{.Type}}: {{end}}{{if .Scope}}{{.Scope }}{{end}}{{.Title}}{{range .Tickets}} {{.}}{{end}} <span>({{.Author}})</span></li>
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
