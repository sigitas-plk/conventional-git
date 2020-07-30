package cnv

import (
	"testing"
)

func TestWriteWithTemplateWithNoChanges(t *testing.T) {
	actual := WriteWithTemplate(nil, "html", "")
	if actual != nil {
		t.Errorf("Expected not to get error if no changes provided")
	}
}

func ExampleWriteWithTemplate() {
	changes := []ParsedCommit{
		ParsedCommit{
			Type:  Fix,
			Title: "fix",
		},
		ParsedCommit{
			Type:  Feat,
			Title: "feat",
		},
		ParsedCommit{
			Type:           Feat,
			Title:          "unfinished feat",
			WorkInProgress: true,
		},
		ParsedCommit{
			Type:           Feat,
			BreakingChange: true,
			Title:          "breaking change",
		},
		ParsedCommit{
			Type:  Ci,
			Title: "should be in other",
		},
		ParsedCommit{
			Type:  Refactor,
			Title: "refactor",
		},
		ParsedCommit{
			Type:  Unconventional,
			Title: "should be skipped",
		},
		ParsedCommit{
			Type:  Perf,
			Title: "performance improvement",
		},
		ParsedCommit{
			Type:  Revert,
			Title: "reverted",
		},
	}
	WriteWithTemplate(&changes, "md", "")
	// Output:

	// 	#### BREAKING CHANGES
	// 	- breaking change
	//    #### Added
	// 	- feat
	//    #### Fixed
	// 	- fix
	//    #### Changed
	// 	- refactor
	// 	- performance improvement
	//    #### Reverted
	// 	- reverted
	//    #### Other
	//    - ci: should be in other
	//    #### Unfinished
	//    - feat: unfinished feat
}
