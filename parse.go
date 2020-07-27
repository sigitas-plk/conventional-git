package cnv

import (
	"regexp"
	"strings"
	"time"
)

var (
	jiraRegexp     = regexp.MustCompile(`(\b[A-Z][A-Z0-9]{1,10}-[0-9]+\b)`)
	titleRegexp    = regexp.MustCompile(`(?s)^((?P<wip>wip:)\s?)?(?P<type>\S+?)?(?P<scope>\(\S+\))?(?P<breaking>\!?)?: (?P<title>[^\n\r]+)`)
	breakingRegexp = regexp.MustCompile(`BREAKING[ _\\-]CHANGE(S)?`)
)

// ParsedCommit is commit parsed using Conventional Commit convention
type ParsedCommit struct {
	Hash           string     `json:"hash"`
	ShortHash      string     `json:"short_hash"`
	Author         string     `json:"author"`
	AuthorEmail    string     `json:"author_email"`
	Date           time.Time  `json:"date"`
	FullTitle      string     `json:"full_title"`
	Tickets        []string   `json:"tickets"`
	Title          string     `json:"title"`
	Scope          string     `json:"scope"`
	Type           CommitType `json:"commit_type"`
	WorkInProgress bool       `json:"work_in_progress"`
	BreakingChange bool       `json:"breaking_change"`
}

// Commits grouped ParsedCommits based on provided filters / include list
type Commits struct {
	Conventional   []ParsedCommit `json:"commits"`
	Filtered       []ParsedCommit `json:"commits_filtered"`
	Unconventional []ParsedCommit `json:"unconventional"`
}

// ParseByType takes in list of commits parses them and
// assigns them to one of the three commit groups based on type and filters
func ParseByType(gc *[]Commit, keepCommits []CommitType) *Commits {

	// convert array to map
	keep := make(map[CommitType]bool)
	if len(keepCommits) > 0 {
		for _, k := range keepCommits {
			keep[k] = true
		}
	}
	parsed := []ParsedCommit{}
	for _, c := range *gc {
		parsed = append(parsed, Parse(&c))
	}
	commits := Commits{}
	for _, c := range parsed {
		if c.Type == Unconventional {
			commits.Unconventional = append(commits.Unconventional, c)
			continue
		}
		if _, ok := keep[c.Type]; ok {
			commits.Conventional = append(commits.Conventional, c)
			continue
		}
		commits.Filtered = append(commits.Filtered, c)
	}
	return &commits
}

// Parse parses provided commit according to conventional commit conventions
func Parse(c *Commit) ParsedCommit {
	parsed := ParsedCommit{
		Author:         c.Author,
		AuthorEmail:    c.AuthorEmail,
		Date:           c.Date,
		FullTitle:      c.Title,
		Hash:           c.Hash,
		ShortHash:      c.ShortHash,
		Tickets:        getUniqueJiraTickets(c.Title, c.Body),
		BreakingChange: hasBreakingInBody(c.Body),
		Type:           Unconventional,
	}

	cc := parseCommitTitle(c.Title)

	// empty map? Unconventional commit
	if cc == nil {
		return parsed
	}

	parsed.Scope = cc["scope"]
	parsed.Title = cc["title"]

	if t := GetCommitType(cc["type"]); t != nil {
		parsed.Type = *t
	}

	if _, ok := cc["wip"]; ok {
		parsed.WorkInProgress = true
	}

	if _, ok := cc["breaking"]; ok {
		parsed.BreakingChange = true
	}

	return parsed
}

func parseCommitTitle(title string) map[string]string {
	match := titleRegexp.FindStringSubmatch(title)

	if len(match) == 0 {
		return nil
	}

	groups := make(map[string]string, len(match))

	for i, group := range titleRegexp.SubexpNames() {
		if match[i] != "" {
			groups[group] = match[i]
		}
	}

	groups["title"] = removeJiraTickets(groups["title"])

	// remove parenthesis from scope value if it's set
	if val, ok := groups["scope"]; ok {
		groups["scope"] = strings.Trim(val, "()")
	}
	return groups
}

func hasBreakingInBody(body string) bool {
	if body == "" {
		return false
	}
	return breakingRegexp.FindAllSubmatchIndex([]byte(body), -1) != nil
}

func getUniqueJiraTickets(title, body string) []string {
	ticketsMap := make(map[string]bool)

	t := jiraRegexp.FindAll([]byte(title), -1)
	for _, v := range t {
		ticketsMap[string(v)] = true
	}

	t = jiraRegexp.FindAll([]byte(body), -1)
	for _, v := range t {
		ticketsMap[string(v)] = true
	}

	tickets := []string{}
	for k := range ticketsMap {
		tickets = append(tickets, k)
	}

	return tickets
}

func removeJiraTickets(str string) string {
	return strings.TrimSpace(jiraRegexp.ReplaceAllString(str, ""))
}
