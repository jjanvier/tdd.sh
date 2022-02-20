# TDD.sh

A simple tool to enforce the TDD practice. It follows principles erected by Kent Beck in _"Test Driven Development: By Example"_ and allows to:

- **reduce the cognitive load**
  - by having a simple and consistent way to launch your tests across all your projects, whatever the language and technical stack
  - by automatically committing when your tests are green
- **reduce the feedback loop** by launching a notification when your tests have been red for too long
- **stay focused** in the red/green/refactor loop by using a todo list

## How to install

### Requirements

Nothing particular is required. _TDD.sh_ embeds all it needs in its binary.

### Installation

Download the latest binary that suits your computer, which is available at [tdd-dot-sh.zip](https://github.com/jjanvier/tdd.sh/releases/latest/download/tdd-dot-sh.zip). Unzip it and install the binary in a directory recognized by your `$PATH` variable. For instance, if you're a _MacOs_ or a _Linux_ user, you can move the `tdd` binary to `/usr/local/bin/`.

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
            add: "*.php doc/*" # all PHP files and all files present in the "doc" folder will be added to the index 
        timer: 60 # you'll receive a small notification if your steps are still red after 60 seconds
    ut_go: # another alias
        command: go test ./... -v
        # if no "git.amend" key is configured, commits won't be amended: the previous message will be reused
        # if no "git.add" key is configured, all files will be added to the index. It's equivalent to "git add ."
        # if no "timer" key is defined, no notification will pop
```

With this configuration file, you could launch your aliases `ut` and `ut_go` with respectively `tdd ut` and `tdd ut_go`. Please note that for the moment, the `command` key can not handle "complex" commands with `&&`, `||` or `;`.

It's recommended to add the configuration files to your `.gitignore_global`:
```bash
echo ".tdd.yaml" >> ~/.gitignore_global
```

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

It's recommended to add the todo list files to your `.gitignore_global`:
```bash
echo ".tdd.todo" >> ~/.gitignore_global
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
