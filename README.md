# TDD.sh

A simple tool to enforce the TDD practice. It follows principles erected by Kent Beck in _"Test Driven Development: By Example"_ and allows to:

- have a simple and consistent way to launch your tests across all your projects, whatever the language
- automatically commit when your tests are green
- launch a notification when your tests have been red for too long
- stay focused in the red/green/refactor loop by using a todo list

## How to install

### Requirements

On Linux:

- `libasound2-dev` to play the light sound bell when the tests have been red for too long

On MacOS, nothing particular is required.


### Installation

Download the latest binary, which is available at [tdd-dot-sh.zip](https://github.com/jjanvier/tdd.sh/releases/latest/download/tdd-dot-sh.zip). Unzip it and install the binary in a directory recognized by your `$PATH` variable.

Yippee, you are now ready to use the `tdd` command!

## How to use

### The classical workflow

Start a new TDD session:

```bash
tdd new "a clear message that explains what I want to achieve"
```

The message of this TDD session will be used to commit your changes.

Then, start your classical TDD loop: write a failing test, write the minimum code needed to make it pass, run all your tests, refactor and ensure all your tests are still green. The only difference, is that to run your tests you'll now use:

```bash
tdd my_alias
```

Where `my_alias` is the command used to launch your tests. It's defined in the configuration file, but we'll get back to that later.

If after having launched `tdd my_alias`, your tests are green, then your code is automatically committed. Great! You can now refactor or write a new failing test. For both those actions, use `tdd my_alias` again. 

If your tests are red and you have configured a timer for this alias, you'll receive soon a small notification encouraging you to take a smaller step.


### The configuration file

At the root of your project, you must create a `.tdd.yaml` configuration file. It defines all the test aliases you want to use. 

You can initialize a new configuration file with:

```bash
tdd config --init
```

Then, change it according to your needs.

You can validate an existing configuration file with:
```bash
tdd config --validate
```

Here is an example of a configuration file:

```yaml
# .tdd.yaml
aliases:
    ut: # I use "ut" for Unit Tests. Personally, I define a "ut" alias for all my projects
        command: docker-compose run --rm php vendor/bin/phpspec run -v
        git:
            amend: true # commits will be amended when tests are green
        timer: 60 # you'll receive a small notification if your steps are still red after 60 seconds
    ut_go: # another alias
        command: go test ./... -v
        # if no "git" key is configured, commits won't be amended: the previous message will be reused
        # if no "timer" key is defined, no notification will pop
```

With this configuration file, you could launch your aliases `ut` and `ut_go` with respectively `tdd ut` and `tdd ut_go`. Please note that for the moment, the `command` key can not handle "complex" commands with `&&`, `||` or `;`.

### Using the todo list

While you're working on something you can think about fixing or improving something else. To not lose the focus, it's handy to note this idea in a todo list:

```bash
tdd todo "I should do this other thing"
tdd todo "oh, something else I think about too"
```

When you are ready, which means when your tests are green, you can pick a task from this list:

```bash
tdd do
```

The list of tasks is displayed:

```  
Here is your todo list, which task do you want to tackle?
  > I should do this other thing
  > oh, something else I thought about too
```

Now you can select the item you want to work on, this will start a new TDD session by using this task as a commit message.

If you need to, you can clear this todo list at anytime with:

```bash
tdd done
```

## FAQ

### Why committing every time tests are green is important?

I can't remember the number of times I've been lost in my developments because I didn't know exactly what I should add or not for my current commit. CTRL-Z on this file, adding this file but not that one, checking local history to understand what happened, going back and forth... Huh, this chaos is exhausting! 

Committing everytime tests are green totally frees the mind. When it's committed, that means the code works. It's simple as that. There is nothing else to remember. 

### Committing before the code has been written, won't this screw my git history?

Surprisingly, I find it's quite the opposite. The `tdd new` command forces us to explain what we want to achieve, not the how. Which is exactly what should be good a git history.

Besides, we're doing TDD right? So why should we explain what we have done, after having done it? Let's apply the TDD cycle from the very beginning: first, we have to think about what to achieve.

To end that topic, keep also in mind that, if you want or need to, you can still rewrite the git history whenever you want.

### Why launching a notification when tests have been red for too long will help me to practice TDD?

One of the most fundamental principle of TDD is to take baby steps. In the book _"Test Driven Development: By Example"_, Kent Beck can't stop talking about how small the changes should be so that we can be smoothly driven by the tests.

Sadly, _small_ is a vague metric which is impossible to measure. The only metric I found helpful is _time_. When I take too much time to turn my tests green that's a sign I've been taking a too big or a too difficult step. In that case, I cancel my current code, and restart with a smaller step. 

If you're a complete TDD beginner, I suggest you to start with a 10 minutes timer. Try to reduce it week by week. After a few months of practice, you should be able to use easily a ~60 seconds timer for a development environment you master. Congratulations! You are now able to take small steps!

### Why using a todo list will help me to practice TDD?

The purpose of this todo list is not to store tasks that will take us months to complete. It's not an alternative to the useless `// TODO: fix this` that we can find sometimes in our codebase. 

Its goal is to minimize the red/green time: when the tests are failing, our only goal is to make them turn green quickly and easily. Anything that is not related to this precise goal should land in this todo list. Therefore, it should have a short lifetime; typically, one day maximum.

## Features list

- ~~handle a TODO list with~~
  - ~~tdd todo : add an item in the list~~
  - ~~tdd do : pick one item in the list, and do a `tdd new` with it~~
  - ~~tdd done : clear the list~~
  - ~~tdd do: sanitize todo lists => remove empty lines~~
- log properly debug/info/error, see [this library](https://www.honeybadger.io/blog/golang-logging/) of [this one](https://github.com/apex/log)
- ~~display execution results like:~~
  - ~~Tests pass ✅~~
  - ~~Tests do not pass ❌~~
  - ~~OK ✅ (for non-test commands)~~
  - ~~Error ❌ (for non-test commands)~~
- check when the test command is unknown => today we have "tests failed."
- be able to define a default alias? 
- be able to get stats on my tdd sessions?
- log commands, exit code and datetime to be able to have stats?
- handle "complex" commands with `&&`, `||` or `;` => check [this parser](https://github.com/alecthomas/participle)
- be able to choose what to add to the git index?
- default values for configuration?
- option `--verbose` to display the commands that are really launched, hide them otherwise
- option `--version`?
- ~~use [goreleaser](https://goreleaser.com/) to build different versions and push release to Github~~
- stats about download => check [github-release-stats](https://tooomm.github.io/github-release-stats/), [grev](https://hanadigital.github.io/grev/), [github-analytics](http://github-analytics.apphb.com/)
- ~~be able to launch a test command with an alias~~
- ~~be able to start a new tdd session (that should create an empty commit I can amend on)~~
- ~~be able to define if I want to amend commits or not~~
- ~~be able to define a timer (duration, bell and/or notification) when the tests have been red for too long~~
- ~~play a sound along with the notification when tests have been red for too long~~
- ~~reset the timer when tests are green~~
- ~~display errors to users(like "alias not found" or "missing commit message" with _tdd new_)~~
- ~~proper error message when there is no configuration file~~
- ~~handle all cases where configuration is incorrect =>  validate configuration file schema? or simply check at startup?~~
- ~~be able to initialize a default config file~~
- ~~use a CLI framework => check https://github.com/urfave/cli or https://github.com/spf13/cobra~~ 
