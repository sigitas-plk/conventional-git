package cnv

import (
	"testing"
)

// test/sample-git-history git log --all --decorate --oneline --graph
//
// * 716bb05 (HEAD -> master) unconventional commit
// | * dc0d5a4 (branch-b) feat(branch-b): unmerged branch in master JIR-004
// |/
// *   ac1ff15 (tag: v2.0) Merge branch 'branch-a'
// |\
// | * 2793f8e (branch-a) feat: another file added JIR-001
// | * 71b88ce (tag: v1.1) refactor: added merge conflict JIR-001
// * | 58e0c48 feat(file)!: breaking change in file JIR-002
// |/
// * c43346c feat(file): added text to file
// * 69aefaa (tag: v1.0) feat: initial commit
func TestGetCommitsWithTags(t *testing.T) {
	got, err := GetCommits("test/sample-git-history", "v1.0", "v2.0")
	if err != nil {
		t.Errorf("%s", err)
	}
	if len(*got) != 4 {
		t.Errorf("got %d, want 4", len(*got))
	}

	expectedHashes := map[string]bool{"c43346c": true, "58e0c48": true, "71b88ce": true, "2793f8e": true}
	for _, v := range *got {
		if _, ok := expectedHashes[v.ShortHash]; !ok {
			t.Errorf("%s is not in a list of expected hashes", v.ShortHash)
		}
	}
}

func TestGetCommitsWithShortHashes(t *testing.T) {
	got, err := GetCommits("test/sample-git-history", "ac1ff15", "69aefaa")
	if err != nil {
		t.Errorf("%s", err)
	}
	if len(*got) != 4 {
		t.Errorf("got %d, want 4", len(*got))
	}

	expectedHashes := map[string]bool{"c43346c": true, "58e0c48": true, "71b88ce": true, "2793f8e": true}
	for _, v := range *got {
		if _, ok := expectedHashes[v.ShortHash]; !ok {
			t.Errorf("%s is not in a list of expected hashes", v.ShortHash)
		}
	}
}

func TestGetCommitsWithLongHashes(t *testing.T) {
	got, err := GetCommits("test/sample-git-history", "ac1ff15ae0ddd6c0fd5f6fd0e365b9e5260f3fef", "69aefaae7217010f8675d1c2d055cbbfd4ded81d")
	if err != nil {
		t.Errorf("%s", err)
	}
	if len(*got) != 4 {
		t.Errorf("got %d, want 4", len(*got))
	}

	expectedHashes := map[string]bool{"c43346c": true, "58e0c48": true, "71b88ce": true, "2793f8e": true}
	for _, v := range *got {
		if _, ok := expectedHashes[v.ShortHash]; !ok {
			t.Errorf("%s is not in a list of expected hashes", v.ShortHash)
		}
	}
}

func TestGetCommitsNoMergeCommits(t *testing.T) {
	got, err := GetCommits("test/sample-git-history", "v1.0", "HEAD")
	if err != nil {
		t.Errorf("%s", err)
	}
	for _, v := range *got {
		if v.ShortHash == "ac1ff15" {
			t.Errorf("want no merge commits, got %s", v.ShortHash)
		}
	}
}

func TestGetCommitsReturnCommitInfo(t *testing.T) {

	exp := Commit{
		ShortHash:   "716bb05",
		Hash:        "716bb055c6645e67705f4f3464c8709de2a98db1",
		Author:      "sigitas",
		AuthorEmail: "code@pleikys.com",
		Title:       "unconventional commit",
	}
	get, err := GetCommits("test/sample-git-history", "ac1ff15", "")
	if err != nil {
		t.Errorf("%s", err)
	}

	c := (*get)[0]

	if c.ShortHash != exp.ShortHash {
		t.Errorf("want %s short hash, got %s", exp.ShortHash, c.ShortHash)
	}

	if c.Hash != exp.Hash {
		t.Errorf("want %s hash, got %s", exp.Hash, c.Hash)
	}

	if c.Author != exp.Author {
		t.Errorf("want author %s, got %s", exp.Author, c.Author)
	}

	if c.AuthorEmail != exp.AuthorEmail {
		t.Errorf("want %s email, got %s", exp.AuthorEmail, c.AuthorEmail)
	}

	if c.Title != exp.Title {
		t.Errorf("want %s title, got %s", exp.Title, c.Title)
	}
}
