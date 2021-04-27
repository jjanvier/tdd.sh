# tdd.sh
A simple tool to enforce the TDD practice

TODO:

- log properly debug/info/error, see https://www.honeybadger.io/blog/golang-logging/
- display execution results like:
    
    - Tests pass ✅
    - Tests do not pass ❌
    - OK ✅ (for non-test commands)
    - Error ❌ (for non-test commands)

- display errors to users(like "alias not found" or "missing commit message" with _tdd new_)
- todo list:
    
    - todo
    - do
    - done

- proper error message when there is no configuration file
- default values for configuration
- handle all cases where configuration is incorrect =>  validate configuration file schema? or simply check at startup?
- be able to initialize a default config file
- be able to define a default alias? 
- be able to get stats on my tdd sessions?
