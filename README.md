## Changelist 

Very simple implementation of changelog generation based on given filters in go.

### How to use? 

`go run cmd/changelist/main.go -fromRef 4935a3d3 -i feat -template json`

This will read git history between tags v1.0 to v1.1 and process them according to conventional commits conventions:
[conventionalcommits.org](https://www.conventionalcommits.org/en/v1.0.0-beta.2/)

Only commits with provided types via -i flag will be considered. If you would like to get all the commits provide flag -all.

Main deviation from conventional commits is ability to handle 'wip:' (Work In Progress) prefix. 
This is to allow a little bit more flexibility when adopting trunk based development and merging partial features.

Sample "json" output
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

There are 2 types of predefined templates with opinionated grouping:

Sample "html" output: 
 ```html
<h3>Added</h3>
<ul>
<li>JSON input of parsed commits to predefined templates <span>(sigitas)</span></li>
<li><strong>core</strong> initial changelog implementation <span>(sigitas)</span></li>
</ul>
```

Sample "md" output:
```markdown
#### Added
 - JSON input of parsed commits to predefined templates 3ade40b
 - **core** initial changelog implementation 4935a3d
 ```

To see full list of flags just type `-h`