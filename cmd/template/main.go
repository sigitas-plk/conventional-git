package main

import (
	"flag"
	"fmt"
	"os"
	"time"

	cnv "github.com/sigitas-plk/conventional-git"
)

var (
	changes string
	version string
	date    string
	outType string
	outFile string
)

var t = `[{"hash":"9d7ae55fc5dee924170f9e7f60210ec601927dd0","short_hash":"9d7ae55","author":"sigitas","author_email":"code@pleikys.com","date":"2020-07-27T10:40:54+01:00","full_title":"feat: whatever \" \u003c\u003e{}[]","tickets":[],"title":"whatever \" \u003c\u003e{}[]","scope":"","commit_type":"feat","work_in_progress":false,"breaking_change":false},{"hash":"37036c0c4912db9f5882575b98ad6d41616f4daa","short_hash":"37036c0","author":"sigitas","author_email":"code@pleikys.com","date":"2020-07-23T20:50:02+01:00","full_title":"feat: special characters ^$%\u0026@","tickets":[],"title":"special characters ^$%\u0026@","scope":"","commit_type":"feat","work_in_progress":false,"breaking_change":false},{"hash":"e0d7c427eb726b9dd8fd8e25455032ff80cfbddf","short_hash":"e0d7c42","author":"sigitas","author_email":"code@pleikys.com","date":"2020-07-21T20:55:07+01:00","full_title":"wip:feat: print out changelist based on templates","tickets":[],"title":"print out changelist based on templates","scope":"","commit_type":"feat","work_in_progress":true,"breaking_change":false},{"hash":"4935a3d35aa1276c75e4147d439c5e0cd6b3f760","short_hash":"4935a3d","author":"sigitas","author_email":"code@pleikys.com","date":"2020-07-19T20:54:20+01:00","full_title":"feat(core): initial changelog implementation","tickets":[],"title":"initial changelog implementation","scope":"core","commit_type":"feat","work_in_progress":false,"breaking_change":false}]`
var k = `[{hash:9d7ae55fc5dee924170f9e7f60210ec601927dd0,short_hash:9d7ae55,author:sigitas,author_email:code@pleikys.com,date:2020-07-27T10:40:54+01:00,full_title:feat: whatever " \u003c\u003e{}[],tickets:[],title:whatever " \u003c\u003e{}[],scope:,commit_type:feat,work_in_progress:false,breaking_change:false},{hash:37036c0c4912db9f5882575b98ad6d41616f4daa,short_hash:37036c0,author:sigitas,author_email:code@pleikys.com,date:2020-07-23T20:50:02+01:00,full_title:feat: special characters ^$%\u0026@,tickets:[],title:special characters ^$%\u0026@,scope:,commit_type:feat,work_in_progress:false,breaking_change:false},{hash:e0d7c427eb726b9dd8fd8e25455032ff80cfbddf,short_hash:e0d7c42,author:sigitas,author_email:code@pleikys.com,date:2020-07-21T20:55:07+01:00,full_title:wip:feat: print out changelist based on templates,tickets:[],title:print out changelist based on templates,scope:,commit_type:feat,work_in_progress:true,breaking_change:false},{hash:4935a3d35aa1276c75e4147d439c5e0cd6b3f760,short_hash:4935a3d,author:sigitas,author_email:code@pleikys.com,date:2020-07-19T20:54:20+01:00,full_title:feat(core): initial changelog implementation,tickets:[],title:initial changelog implementation,scope:core,commit_type:feat,work_in_progress:false,breaking_change:false}]`

func main() {

	some := flag.String("some", "", "do it yourself")

	flag.StringVar(&changes, "changes", "", "JSON array of changes to write to template")
	flag.StringVar(&changes, "c", "", "JSON array of changes to write to template (shorthand)")
	flag.StringVar(&version, "version", "", "release version")
	flag.StringVar(&version, "v", "", "release version (shorthand)")
	flag.StringVar(&date, "date", time.Now().Format("2006-01-02"), "release date string")
	flag.StringVar(&outType, "template", "md", "output type 'md' or 'html'")
	flag.StringVar(&outFile, "outputFile", "", "output file (will overwrite existing content)")
	flag.Parse()

	fmt.Println(*some)

	if changes == "" || changes == "null" {
		fmt.Println("Array with changelist not provided. Nothing to do quitting")
		return
	}
	if version == "" {
		fmt.Println("Version is not provided. It is required to proceed. Quitting.")
		return
	}

	if outType != "md" && outType != "html" {
		fmt.Printf("Unsupported template type provided '%s'. Only supported templates are 'md' and 'html'.", outType)
		return
	}

	err := cnv.WriteGroupedWithTemplate(changes, version, date, outType, outFile)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Can't write to template: %s ", err)
		os.Exit(1)
	}

}
