package cnv

import (
	"fmt"
	"strings"
)

// CommitType union type of all conventional commit types allowed
type CommitType string

// Constants of all allowed commit types
const (
	Unconventional CommitType = "UNCONVENTIONAL"
	Build                     = "build"
	Ci                        = "ci"
	Chore                     = "chore"
	Docs                      = "docs"
	Feat                      = "feat"
	Fix                       = "fix"
	Perf                      = "perf"
	Refactor                  = "refactor"
	Revert                    = "revert"
	Style                     = "style"
	Test                      = "test"
)

var (
	types = map[string]CommitType{
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
)

// UnmarshalJSON convert JSON string value of type to CommitType
func (ct *CommitType) UnmarshalJSON(b []byte) error {
	commitType := strings.Trim(string(b), `"`)
	if v, ok := types[commitType]; ok {
		*ct = v
		return nil
	}
	return fmt.Errorf("unexpected commit type received '%s'", commitType)
}

// GetCommitType converts given string to CommitType
// returns Unconventional if unexpected type provided
func GetCommitType(str string) *CommitType {
	if v, ok := types[str]; ok {
		return &v
	}
	return nil
}
