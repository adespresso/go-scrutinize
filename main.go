package main

import (
	"io/ioutil"
	"log"
	"os"
	"os/exec"

	"github.com/phayes/checkstyle"
)

var (
	cstyle = checkstyle.New()
)

func main() {

	homeEnv := os.Getenv("HOME")
	gopathEnv := os.Getenv("GOPATH")
	if gopathEnv == "" {
		gopathEnv = homeEnv + "/go"
	}
	goMetaLinterCmd := gopathEnv + "/bin/gometalinter"

	// Install gometalinter -- no-op if already installed
	cmd := exec.Command("go", "get", "github.com/alecthomas/gometalinter")
	err := cmd.Run()
	if err != nil {
		log.Fatal(err)
	}

	// Install all gometalinter dependancies -- no-op if already installed
	cmd = exec.Command(goMetaLinterCmd, "install")
	err = cmd.Run()
	if err != nil {
		log.Fatal(err)
	}

	// Configure the metalinter
	if _, err := os.Stat("go-scrutinize.config"); os.IsNotExist(err) {
		cmd = exec.Command(goMetaLinterCmd, "--checkstyle")
	} else {
		cmd = exec.Command(goMetaLinterCmd, "--checkstyle", "--config=go-scrutinize.config")
	}

	// Run the metalinter
	out, err := cmd.Output()
	if err != nil {
		exitErr := err.(*exec.ExitError)
		if len(exitErr.Stderr) != 0 {
			log.Fatal(string(exitErr.Stderr))
		}
	}

	// Write the output
	ioutil.WriteFile("checkstyle_report.xml", out, os.ModePerm)

}
