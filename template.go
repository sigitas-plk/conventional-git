package cnv

import (
	"bufio"
	"os"
	"text/template"
)

// WriteWithTemplate groups changes loosely based on https://keepachangelog.com/en/1.0.0/ categories
// (we don't have security or deprecated groups and removed is replaced with reverted)
// and writes out changes to HTML or Markdown
func WriteWithTemplate(changes *[]ParsedCommit, layoutType, outFile string) error {

	if changes == nil || len(*changes) == 0 {
		return nil
	}

	layoutTemplate := TemplateMarkdown
	if layoutType == "html" {
		layoutTemplate = TemplateHtml
	}
	gc := groupChanges(changes)
	if gc == nil {
		return nil
	}
	t, err := template.New("layout").Parse(layoutTemplate)
	if err != nil {
		return err
	}
	if outFile == "" {
		return t.Execute(os.Stdout, gc)
	}

	f, err := os.OpenFile(outFile, os.O_RDWR|os.O_CREATE, 0755)
	if err != nil {
		return err
	}
	defer f.Close()

	w := bufio.NewWriter(f)

	err = t.Execute(w, gc)
	if err != nil {
		return err
	}
	w.Flush()

	return nil
}

type groupedChangelist struct {
	Added           []ParsedCommit
	Fixed           []ParsedCommit
	Changed         []ParsedCommit
	Reverted        []ParsedCommit
	Other           []ParsedCommit
	WorkInProgress  []ParsedCommit
	BreakingChanges []ParsedCommit
}

//vaguely following groupings from https://keepachangelog.com/en/1.0.0/
// we don't have security, deprecated groups
func groupChanges(changes *[]ParsedCommit) *groupedChangelist {
	gc := groupedChangelist{}
	for _, c := range *changes {
		if c.WorkInProgress {
			gc.WorkInProgress = append(gc.WorkInProgress, c)
			continue
		}
		if c.BreakingChange {
			gc.BreakingChanges = append(gc.BreakingChanges, c)
			continue
		}
		switch c.Type {
		case Feat:
			gc.Added = append(gc.Added, c)
		case Fix:
			gc.Fixed = append(gc.Fixed, c)
		case Refactor, Perf:
			gc.Changed = append(gc.Changed, c)
		case Revert:
			gc.Reverted = append(gc.Reverted, c)
		case Ci, Build, Docs, Chore, Style, Test:
			gc.Other = append(gc.Other, c)
		}
	}
	return &gc
}
