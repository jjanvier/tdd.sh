package main

func main() {
	conf := Load(".tdd.yml")
	cmd := conf.GetCommand("ut")
	cmd.Execute()
}

func Hello(name string) string {
	return "Hello " + name
}
