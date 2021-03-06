package main

import "github.com/hanbang-wang/FakeGit-Go"

func main() {
	addition, name, email, forever := fakegit.ProcArgs()
	cfg := fakegit.NewGitConf(".git/config")
	if forever && name == "" {
		fakegit.Fatal(fakegit.ARGUMENT_ERROR_USERNAME)
	}
	if name != "" {
		cfg.Change(name, email)
	}
	if !forever {
		fakegit.RunCommand(append([]string{"git"}, addition...))
		if name != "" {
			cfg.Recover()
		}
	}
}
