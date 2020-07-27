## Changelist 

Very simple implementation of changelog generation based on given filters in go.

### How to use? 

`go run cmd/changelist/main.go -to "v1.0" -from "v1.1" -i feat -i fix -i refactor`

This will read git history between tags v1.0 to v1.1 and process them according to conventional commits conventions:
[conventionalcommits.org](https://www.conventionalcommits.org/en/v1.0.0-beta.2/)

Only commits with provided types via -i flag will be considered. If you would like to get all the commits provide flag -all.

Main deviation from conventional commits is ability to handle 'wip:' (Work In Progress) prefix. 
This is to allow a little bit more flexibility when adopting trunk based development and merging partial features. 

Binary will simply spit out JSON within given range which are part of the included types supplied by -included (-i shorthand).

Sample output
```json
[
    {
        "hash": "ee77fcd1e8b2bbd0483f2e6c2045a7cdc87399e2",
        "short_hash": "ee77fcd",
        "author": "sigitas",
        "author_email": "mail@mail.com",
        "date": "2020-07-19T20:54:20+01:00",
        "full_title": "feat(core): initial changelog implementation",
        "tickets": [],
        "title": "initial changelog implementation",
        "scope": "core",
        "commit_type": "feat",
        "work_in_progress": false,
        "breaking_change": false
    }
]
```

To see list of flags just type `-h`


## Template

Template binary can be fed with the output of changelist to get a pretty output of either html or markdown which are vaguely grouped by [keepachangelog.com](https://keepachangelog.com/en/1.0.0/) types. This is basically pre-configured opinionated changelist styles. 

### How to use? 

`go run cmd/template/main.go -c '[{"hash":"ee...}]' -v 1.1 -o md -outputFile out.md`

To see list of flags just type `-h`