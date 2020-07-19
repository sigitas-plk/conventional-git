package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"strings"
)

type includeTypes []CommitType

func (t *includeTypes) String() string {
	s := make([]string, 0)
	for _, v := range *t {
		s = append(s, v.String())
	}
	return "[" + strings.Join(s, ", ") + "]"
}

func (t *includeTypes) Set(tp string) error {
	ct := GetCommitType(tp)
	if ct == Unconventional {
		allowed := []CommitType{Build, Ci, Chore, Docs, Feat, Fix, Perf, Refactor, Revert, Style, Test}
		return fmt.Errorf("the only allowed types: %s", allowed)
	}
	*t = append(*t, GetCommitType(tp))
	return nil
}

var (
	inc       includeTypes
	fromHash  string
	toHash    string
	path      string
	returnAll bool
)

func main() {
	includedUsage := "commit type to include (e.g feat, fix, chore)"
	flag.Var(&inc, "include", includedUsage)
	flag.Var(&inc, "i", includedUsage+" (shorthand)")
	flag.StringVar(&fromHash, "from", "HEAD", "hash (or tag) to get changelist from. Defaults to HEAD.")
	flag.StringVar(&toHash, "to", "", "hash (or tag) to get changelist until")
	flag.StringVar(&path, "path", ".", "path to .git location")
	flag.BoolVar(&returnAll, "all", false, "return all (conventional, unconventional, filtered) commit types")
	flag.Parse()

	getCommits()

}

func getCommits() {
	comm, err := GetCommits(path, toHash, fromHash)
	if err != nil {
		panic(err)
	}
	c := ParseByType(comm, inc)

	var j []byte

	if returnAll {
		j, err = json.Marshal(c)
	} else {
		j, err = json.Marshal(c.Conventional)
	}
	if err != nil {
		panic(err)
	}
	fmt.Println(string(j))
}
