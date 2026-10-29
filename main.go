package main

func main() {
	cfg := &config{
		Command: getCommands(),
	}
	startRepl(cfg)
}
