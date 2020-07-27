package cnv

import (
	"bufio"
	"encoding/json"
	"fmt"
	"os"
	"text/template"
)

// WriteGroupedWithTemplate outputs commits grouped by types using predefined templates
// Only two types of templates are supported for now: html and markdown
func WriteGroupedWithTemplate(changes, version, date, outputType, outFile string) error {
	b := []byte(changes)
	if b == nil {
		return fmt.Errorf("nothing to write")
	}

	pc, err := unmarshalCommits(b)

	if err != nil {
		return err
	}

	gc := groupedChangelist{
		Version: version,
		Date:    date,
	}
	groupChanges(&gc, pc)

	layout := TemplateMarkdown
	if outputType == "html" {
		layout = TemplateHtml
	}

	return write(&gc, layout, outFile)
}

func write(changes *groupedChangelist, layout, fileName string) error {
	t, err := template.New("layout").Parse(layout)
	if err != nil {
		return err
	}
	if fileName == "" {
		return t.Execute(os.Stdout, changes)
	}

	f, err := os.Create(fileName)
	if err != nil {
		return err
	}
	defer f.Close()

	w := bufio.NewWriter(f)

	err = t.Execute(w, changes)
	if err != nil {
		return nil
	}
	w.Flush()

	return nil
}

type groupedChangelist struct {
	Version         string
	Date            string
	Added           []ParsedCommit
	Fixed           []ParsedCommit
	Changed         []ParsedCommit
	WorkInProgress  []ParsedCommit
	BreakingChanges []ParsedCommit
}

func unmarshalCommits(b []byte) (*[]ParsedCommit, error) {
	c := &[]ParsedCommit{}
	err := json.Unmarshal(b, c)
	if err != nil {
		return nil, err
	}
	return c, nil
}

//vaguely following groupings from https://keepachangelog.com/en/1.0.0/
// we don't have security, removed or deprecated groups
func groupChanges(gc *groupedChangelist, changes *[]ParsedCommit) {
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
		case Refactor, Perf, Revert:
			gc.Changed = append(gc.Changed, c)

		}
	}
}
