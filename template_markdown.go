package cnv

// TemplateMarkdown markdown template
const TemplateMarkdown = "{{- define \"item\"}} - {{if .Scope}}**{{.Scope }}** {{end}}{{.Title}}{{range .Tickets}} {{.}}{{end}} {{.ShortHash}}\r\n{{end -}}" +
	"{{- define \"itemWithType\"}}- {{if .Type}}{{.Type}}: {{end}}{{if .Scope}}**{{.Scope }}** {{end}}{{.Title}}{{range .Tickets}} {{.}}{{end}} {{.ShortHash}}\r\n{{end -}}" +
	"{{if .BreakingChanges}}#### BREAKING CHANGES\r\n{{range .BreakingChanges}}{{template \"item\" .}}{{end}}{{end}}" +
	"{{if .Added}}#### Added\r\n{{range .Added}}{{template \"item\" .}}{{end}}{{end}}" +
	"{{if .Fixed}}#### Fixed\r\n{{range .Fixed}}{{template \"item\" .}}{{end}}{{end}}" +
	"{{if .Changed}}#### Changed\r\n{{range .Changed}}{{template \"item\" .}}{{end}}{{end}}" +
	"{{if .Reverted}}#### Reverted\r\n{{range .Reverted}}{{template \"item\" .}}{{end}}{{end}}" +
	"{{if .Other}}#### Other\r\n{{range .Other}}{{template \"itemWithType\" .}}{{end}}{{end}}" +
	"{{if .WorkInProgress}}#### Unfinished\r\n{{range .WorkInProgress}}{{template \"itemWithType\" .}}{{end}}{{end}}"
	// could not use backtick quotes for this due to strange whitespaces even '-' could not eat up..
