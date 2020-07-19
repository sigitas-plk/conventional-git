## Changelist 

Very simple implementation of changelog generation based on given filters in go.

### How to use? 

`changelist -path "." -to "v1.0" -from "v1.1" -i feat -i fix -i refactor`

Clone the repo, and build/run this on any git repository using conventional commit conventions:
https://www.conventionalcommits.org/en/v1.0.0-beta.2/

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