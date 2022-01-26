# TDD.sh

A simple tool to enforce the TDD practice.

It allows to:

- easily launch your tests
- automatically commit when your tests are green
- launch a notification when your tests have been red for too long
- have a consistent way to launch your tests across all your projects

## How to install

### Requirements

TODO

### Installation

TODO


## How to use

### The classical workflow

Start a new TDD session:

```bash
tdd new "a clear commit message that explains what I want to achieve"
```

Then, start your classical TDD loop. Write a failing test, write the minimum code needed to make it pass, run all your tests, and refactor. To run your tests, use:

```bash
tdd my_alias
```

If your tests are green, your code is automatically committed. Great! You can now refactor or write a new failing test.
If your tests are red and you have configured a timer for this alias, you'll receive soon a small notification telling you to try to take a smaller step.


### The configuration file

At the root of your project, you must define a `.tdd.yml` configuration file. This configuration defines all the aliases you want to use. For instance:

```yaml

aliases:
    ut: # I use "ut" for Unit Tests. Personally, I define a "ut" alias for all my projects
        command: docker-compose run --rm php vendor/bin/phpspec run -v
        git:
            amend: true # commits will be amended when tests are green
        timer: 120 # you'll receive a small notification if your steps are still red after 120 seconds
    it_invoicing: # another alias
        command: docker-compose run --rm php vendor/bin/phpunit --suite=Integration-Invoicing
        # if no "git" key is configured, commits won't be amended: the previous message will be reused
        # if no "timer" key is defined, no notification will pop
```

### Using the todo list

While you're working on something you can think about fixing or improving something else. To not loose the focus, it's handy to use a todo list for that:

```bash
tdd todo "I should do this other thing"
tdd todo "oh, something else I think about too"
```

When you are ready you can pickup a task from this list:

```bash
tdd do
```

The list of tasks is displayed:

```  
> Here is your todo list:
> Which task do you want to tackle?
>   1 - I should do this other thing
>   2 - oh, something else I think about too
>
```

Now you can enter the number of the task you want to work on, this will start a new TDD session by using this task as a commit message.

If you need to, you can clear this todo list at anytime with:

```bash
tdd done
```

## FAQ

### Why committing every time tests are green is important?

TODO

- stop chaos development (lost in what we have done, CTRL Z, rewrite the same thing several time, check local history etc)
- free your mind
- committed = the code works as expected

### Won't this screw my git history?

TODO

- quite the opposite, the "tdd new" command forces us to explain what we want to achieve, not the how. Which is exactly what should be good a git history.
- besides, we're doing TDD right? so why explain what we have done, after having done it. Let's apply the TDD cycle from the beginning
- if you want to, you can still rewrite the history whenever you want

### Why launching a notification will help me to practice TDD?

TODO

- in the book "TDD by example" Kent Beck keeps talking about how small the changes should be so that we can be smoothly driven by the tests
- sadly, "small" is vague difficult to count. The only metric I found to help me is time. When I take too much time to turn my tests green is a sign I've been taking a too big/difficult step
- not ideal, but it works for me
- if you're a beginner, start with a 10 minutes timer and try to reduce it week after week

## TODO list

- ~~be able to launch a test command with an alias~~
- ~~be able to start a new tdd session (that should create an empty commit I can amend on)~~
- ~~be able to define if I want to amend commits or not~~
- ~~be able to define a timer (duration, bell and/or notification) when the tests have been red for too long~~
- ~~play a sound along with the notification when tests have been red for too long~~
- ~~reset the timer when tests are green~~
- ~~handle a TODO list with~~
  - ~~tdd todo : add an item in the list~~
  - ~~tdd do : pick one item in the list, and do a `tdd new` with it~~
  - tdd do: sanitize todo lists => remove empty lines
  - ~~tdd done : clear the list~~
- log properly debug/info/error, see https://www.honeybadger.io/blog/golang-logging/
- display execution results like:
  - Tests pass ✅
  - Tests do not pass ❌
  - OK ✅ (for non-test commands)
  - Error ❌ (for non-test commands)
- display errors to users(like "alias not found" or "missing commit message" with _tdd new_)
- proper error message when there is no configuration file
- default values for configuration
- handle all cases where configuration is incorrect =>  validate configuration file schema? or simply check at startup?
- be able to initialize a default config file
- be able to define a default alias? 
- be able to get stats on my tdd sessions?
- log commands, exit code and datetime to be able to have stats?
- use a CLI framework => check https://github.com/urfave/cli or https://github.com/spf13/cobra 
