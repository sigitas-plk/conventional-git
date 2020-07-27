package cnv

import (
	"bytes"
	"encoding/xml"
	"fmt"
	"io"
	"os/exec"
	"time"
)

// Commit information from git log
type Commit struct {
	Hash        string    `xml:"hash"`
	ShortHash   string    `xml:"short"`
	Author      string    `xml:"author"`
	AuthorEmail string    `xml:"email"`
	Date        time.Time `xml:"date"`
	Title       string    `xml:"title"`
	Body        string    `xml:"body"`
}

// GetCommits retrieves list of commits between to hashes
// uses range git log hash...hash (hash can also be a tag)
// https://www.git-scm.com/docs/git-log
//
// given empty string as to argument, will go through all git history reachable from given commit
// providing empty string as from argument will fallback to HEAD
//
// excludes merge commits
func GetCommits(path, to, from string) (*[]Commit, error) {
	if from == "" {
		from = "HEAD"
	}
	logTo := ""
	if to != "" {
		logTo = "..." + to
	}

	// using XML not json since JSON breaks with commint body containing new lines
	// <![CDATA[]]> is to escape special characters in commit title and body e.g. &"<'
	format := `
	<commit>
		<hash>%H</hash>
		<short>%h</short>
		<author>%an</author>
		<email>%ae</email>
		<date>%aI</date>
		<title><![CDATA[%s]]></title>
		<body><![CDATA[%b]]></body>
	</commit>
	`

	// 'git log hash...hash', if toHash is empty 'git log hash'
	cmd := exec.Command("git", "log", from+logTo, "--pretty="+format, "--no-merges")
	cmd.Dir = path

	out, err := cmd.Output()
	if err != nil {
		return nil, fmt.Errorf("error running git log '%s%s': %s", from, logTo, err)
	}

	d := xml.NewDecoder(bytes.NewBuffer(out))
	var commits []Commit
	for {
		var c Commit
		err := d.Decode(&c)
		if err != nil {
			if err == io.EOF {
				break
			}
			return nil, err
		}
		commits = append(commits, c)
	}
	return &commits, nil
}
