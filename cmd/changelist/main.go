package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"os"
	"strings"

	cnv "github.com/sigitas-plk/conventional-git"
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
	inc          includeTypes
	fromHash     string
	toHash       string
	path         string
	templateType string
	out          string
	returnAll    bool
)

func main() {

	flag.Var(&inc, "include", "commit type to include e.g feat, fix, chore")
	flag.Var(&inc, "i", "commit type to include e.g feat, fix, chore (shorthand)")
	flag.StringVar(&fromHash, "fromRef", "HEAD", "hash (or tag) to get changelist from")
	flag.StringVar(&fromHash, "from", "HEAD", "hash (or tag) to get changelist from (shorthand)")
	flag.StringVar(&toHash, "toRef", "", "hash (or tag) to get changelist until")
	flag.StringVar(&toHash, "to", "", "hash (or tag) to get changelist until (shorthand)")
	flag.StringVar(&path, "path", ".", "path to .git location")
	flag.StringVar(&templateType, "template", "json", `output type "md", "html" or "json"`)
	flag.StringVar(&templateType, "t", "json", `output type "md", "html" or "json" (shorthand)`)
	flag.StringVar(&out, "out", "", "output to file")
	flag.StringVar(&out, "o", "", "output to file (shorthand)")
	flag.BoolVar(&returnAll, "all", false, "return all (conventional, unconventional, filtered) commit types. Only in JSON format.")
	flag.Parse()

	if returnAll == true && templateType != "json" {
		fmt.Fprintf(os.Stderr, "Can only return in JSON format with the flag -all true")
		os.Exit(1)
	}

	getCommits()
}

func getCommits() {
	comm, err := cnv.GetCommits(path, toHash, fromHash)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Can't get commits: %s ", err)
		os.Exit(1)
	}
	c := cnv.ParseByType(comm, inc)

	switch templateType {
	case "md", "html":
		err = cnv.WriteWithTemplate(&c.Conventional, templateType, out)
		if err != nil {
			fmt.Fprintf(os.Stderr, "Can't write changes with template '%s': %s ", templateType, err)
			os.Exit(1)
		}
	case "json":
		err = writeJSON(c)
		if err != nil {
			fmt.Fprintf(os.Stderr, "Can't write json: %s", err)
			os.Exit(1)
		}
	default:
		fmt.Fprintf(os.Stderr, "Unsupported type '%s'. Only allwed ones are 'md', 'json' and 'html'", templateType)
		os.Exit(1)
	}
	os.Exit(0)
}

func writeJSON(c *cnv.Commits) error {
	var j []byte
	var err error
	if returnAll {
		j, err = json.Marshal(c)
	} else {
		j, err = json.Marshal(c.Conventional)
	}
	if err != nil {
		return err
	}

	if out == "" {
		fmt.Fprintf(os.Stdout, "%s", j)
		return nil
	}

	f, err := os.OpenFile(out, os.O_RDWR|os.O_CREATE, 0755)
	if err != nil {
		return err
	}
	defer f.Close()

	_, err = f.Write(j)
	if err != nil {
		return err
	}
	return nil
}
