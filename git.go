package cnv

import (
	"bufio"
	"bytes"
	"encoding/xml"
	"fmt"
	"io"
	"os/exec"
	"strings"
	"time"
)

// Commit information from git log
type Commit struct {
	Hash        string
	ShortHash   string
	Author      string
	AuthorEmail string
	Date        time.Time
	Title       string
	Body        string
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
		<body><![CDATA[%B]]></body>
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
		var rc rawCommit
		err := d.Decode(&rc)
		if err != nil {
			if err == io.EOF {
				break
			}
			return nil, err
		}
		title, body, err := toTitleAndBody(rc.Body)
		if err != nil {
			return nil, err
		}
		c := Commit{
			Hash:        rc.Hash,
			ShortHash:   rc.ShortHash,
			Date:        rc.Date,
			Author:      rc.Author,
			AuthorEmail: rc.AuthorEmail,
			Title:       title,
			Body:        body,
		}
		commits = append(commits, c)
	}
	return &commits, nil
}

type rawCommit struct {
	Hash        string    `xml:"hash"`
	ShortHash   string    `xml:"short"`
	Author      string    `xml:"author"`
	AuthorEmail string    `xml:"email"`
	Date        time.Time `xml:"date"`
	Body        string    `xml:"body"`
}

// git log %s considers a subject everything until empty space
// Since we want to consider new lines as well, gotta split it ourselves
func toTitleAndBody(rawBody string) (string, string, error) {
	rawTrimmed := strings.TrimSpace(rawBody)
	sc := bufio.NewScanner(strings.NewReader(rawTrimmed))
	title := ""
	body := ""
	if ok := sc.Scan(); ok {
		title = sc.Text()
		for sc.Scan() {
			body += fmt.Sprintln(sc.Text())
		}
	}
	return title, body, sc.Err()
}
