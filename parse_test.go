package cnv

import (
	"reflect"
	"testing"
)

func TestParse(t *testing.T) {

	c := Commit{
		Title: "feat(some-scope): commit title JIR-001",
		Body:  "commit body JIR-002 BREAKING_CHANGE",
	}

	expected := ParsedCommit{
		Title:          "commit title",
		Type:           Feat,
		Scope:          "some-scope",
		Tickets:        []string{"JIR-001", "JIR-002"},
		BreakingChange: true,
	}

	actual := Parse(&c)

	if actual.Title != expected.Title {
		t.Errorf("want title %s, got %s", expected.Title, actual.Title)
	}

	if actual.Scope != expected.Scope {
		t.Errorf("want scope %s, got %s", expected.Scope, actual.Scope)
	}

	if actual.Type != expected.Type {
		t.Errorf("want type %s, got %s", expected.Type, actual.Type)
	}

	if !reflect.DeepEqual(actual.Tickets, expected.Tickets) {
		t.Errorf("want tickets %s, got %s", expected.Tickets, actual.Tickets)
	}
	if !actual.BreakingChange {
		t.Errorf("want breaking change %v, got %v", true, actual.BreakingChange)
	}
}

func TestParseNoBody(t *testing.T) {
	c := Commit{
		Title: "wip:fix!:     title ending with new line\n",
		Body:  "",
	}

	expected := ParsedCommit{
		Title:          "title ending with new line",
		Type:           Fix,
		Scope:          "",
		Tickets:        []string{},
		BreakingChange: true,
		WorkInProgress: true,
	}

	actual := Parse(&c)

	if actual.WorkInProgress != expected.WorkInProgress {
		t.Errorf("want work in progress %v, got %v", expected.WorkInProgress, actual.WorkInProgress)
	}

	if actual.BreakingChange != expected.BreakingChange {
		t.Errorf("want breaking change %v, got %v", expected.BreakingChange, actual.BreakingChange)
	}

	if actual.Title != expected.Title {
		t.Errorf("want title '%s', got '%s'", expected.Title, actual.Title)
	}
}

func TestParseUnconventional(t *testing.T) {
	c := Commit{
		Title: "unconventional title",
		Body:  "whatever description\nBREAKING_CHANGE",
	}

	actual := Parse(&c)

	if actual.Title != "" {
		t.Errorf("want empty title, got '%s'", actual.Title)
	}

	if actual.Type != Unconventional {
		t.Errorf("want commit type %s, got %s", Unconventional, actual.Type)
	}

	if !actual.BreakingChange {
		t.Errorf("unconventional commit with 'BREAKING_CHANGE' in body should have breaking change true, got %v", actual.BreakingChange)
	}
}

func TestParseByType(t *testing.T) {

	c := []Commit{
		Commit{Title: "fix: some bugfix"},
		Commit{Title: "ci: automated commit we don't want in changelog"},
		Commit{Title: "chore: something what just had to be done"},
		Commit{Title: "test: test added but not important for changelist"},
		Commit{Title: "commit not following conventional commit conventions"},
		Commit{Title: "feat(core)!: breaking change feature commit"},
	}

	actual := ParseByType(&c, []CommitType{Feat, Fix})

	for _, v := range actual.Conventional {
		if v.Type != Feat && v.Type != Fix {
			t.Errorf("want only feat and fix in conventional, got %s", v.Type)
		}
	}

	for _, v := range actual.Filtered {
		if v.Type == Fix || v.Type == Feat {
			t.Errorf("don't want to filter fix and feat but got %s in filtered", v.Type)
		}
	}

	if l := len(actual.Filtered); l != 3 {
		t.Errorf("want 3 commits in filtered list, got %v", l)
	}

	if l := len(actual.Unconventional); l != 1 {
		t.Errorf("want 1 commit in unconventional list, got %v", l)
	}

}
