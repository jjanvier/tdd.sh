# Features


## To do

- log properly debug/info/error, see [this library](https://www.honeybadger.io/blog/golang-logging/) of [this one](https://github.com/apex/log)
- be able to define a default alias?
- be able to get stats on my tdd sessions?
- log commands, exit code and datetime to be able to have stats?
- handle "complex" commands with `&&`, `||` or `;` => check [this parser](https://github.com/alecthomas/participle)
- default values for configuration?
- option `--verbose` to display the commands that are really launched, hide them otherwise
- option `--version`?
- stats about download => check [github-release-stats](https://tooomm.github.io/github-release-stats/), [grev](https://hanadigital.github.io/grev/), [github-analytics](http://github-analytics.apphb.com/)
- use only one Dockerized toolchain to build all the packages would be the best
- => that would also allow to use only one `.goreleaser.yaml` config file

## Done

- handle a TODO list with
    - tdd todo : add an item in the list
    - tdd do : pick one item in the list, and do a `tdd new` with it
    - tdd done : clear the list
    - tdd do: sanitize todo lists => remove empty lines
- display execution results like:
    - Tests pass ✅
    - Tests do not pass ❌
    - OK ✅ (for non-test commands)
    - Error ❌ (for non-test commands)
- check when the test command is unknown => today we have "tests failed."
- be able to choose what to add to the git index?
- handle live output for test aliases
- use [goreleaser](https://goreleaser.com/) to build different versions and push release to Github
- be able to launch a test command with an alias
- be able to start a new tdd session (that should create an empty commit I can amend on)
- be able to define if I want to amend commits or not
- be able to define a timer (duration, bell and/or notification) when the tests have been red for too long
- play a sound along with the notification when tests have been red for too long
- reset the timer when tests are green
- display errors to users(like "alias not found" or "missing commit message" with _tdd new_)
- proper error message when there is no configuration file
- handle all cases where configuration is incorrect =>  validate configuration file schema? or simply check at startup?
- be able to initialize a default config file
- use a CLI framework => check https://github.com/urfave/cli or https://github.com/spf13/cobra 



