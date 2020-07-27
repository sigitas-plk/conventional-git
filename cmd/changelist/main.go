package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"os"
	"strings"

	cnv "github.com/conventional-git"
)

type includeTypes []cnv.CommitType

func (t *includeTypes) String() string {
	s := make([]string, 0)
	for _, v := range *t {
		s = append(s, string(v))
	}
	return "[" + strings.Join(s, ", ") + "]"
}

func (t *includeTypes) Set(tp string) error {
	ct := cnv.GetCommitType(tp)
	if ct == nil {
		return fmt.Errorf("see https://www.conventionalcommits.org for allowed types")
	}
	*t = append(*t, *ct)
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

	flag.Var(&inc, "include", "commit type to include e.g feat, fix, chore")
	flag.Var(&inc, "i", "commit type to include e.g feat, fix, chore (shorthand)")
	flag.StringVar(&fromHash, "from", "HEAD", "hash (or tag) to get changelist from. Defaults to HEAD.")
	flag.StringVar(&toHash, "to", "", "hash (or tag) to get changelist until")
	flag.StringVar(&path, "path", ".", "path to .git location")
	flag.BoolVar(&returnAll, "all", false, "return all (conventional, unconventional, filtered) commit types. Only in JSON format.")
	flag.Parse()

	getCommits()
}

func getCommits() {
	comm, err := cnv.GetCommits(path, toHash, fromHash)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Can't get commits: %s ", err)
		os.Exit(1)
	}
	c := cnv.ParseByType(comm, inc)

	var j []byte

	if returnAll {
		j, err = json.Marshal(c)
	} else {
		j, err = json.Marshal(c.Conventional)
	}
	if err != nil {
		fmt.Fprintf(os.Stderr, "Failed marshaling to JSON: %s ", err)
		os.Exit(1)
	}
	fmt.Println(string(j))
	os.Exit(0)
}
