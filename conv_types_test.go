package cnv

import "testing"

func TestGetCommitType(t *testing.T) {

	expected := map[string]CommitType{
		"UNCONVENTIONAL": Unconventional,
		"build":          Build,
		"ci":             Ci,
		"chore":          Chore,
		"docs":           Docs,
		"feat":           Feat,
		"fix":            Fix,
		"perf":           Perf,
		"refactor":       Refactor,
		"revert":         Revert,
		"style":          Style,
		"test":           Test,
	}

	for k, v := range expected {
		ct := GetCommitType(k)
		if ct == nil {
			t.Errorf("want commit type for '%s', got nil", k)
			continue
		}
		if actual := expected[k]; actual != *ct {
			t.Errorf("want type %s, got %s", v, actual)
		}
	}
}

func TestGetCommitTypeUnexpected(t *testing.T) {
	if actual := GetCommitType("unexpected_type"); actual != nil {
		t.Errorf("expected %v for unexpected type, got %v", nil, actual)
	}
}
