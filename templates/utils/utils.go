package utils

import (
	"errors"
	"os"
	"os/exec"
	"strings"

	"github.com/sirupsen/logrus"
)

func WriteFile(file string, content string) {
	PrepareDir(file)
	err := os.WriteFile(file, []byte(content), 0o600)
	checkErr(err)
	logrus.Info("File generated: ", file)
}

func AppendToFile(file string, content string) {
	PrepareDir(file)
	f, err := os.OpenFile(file, os.O_APPEND|os.O_WRONLY|os.O_CREATE, 0600)
	checkErr(err)

	_, err = f.WriteString(content)
	checkErr(err)
	logrus.Info("File modified: ", file)
}

func AppendToCompose(content string) {
	const compose = `docker-compose.yml`
	if _, err := os.Stat(compose); errors.Is(err, os.ErrNotExist) {
		WriteFile(compose, "services:\n")
	}
	AppendToFile(compose, content)
}

func AppendToMakefile(content string) {
	const makefile = `Makefile`
	if _, err := os.Stat(makefile); errors.Is(err, os.ErrNotExist) {
		WriteFile(makefile, "pwd := $(dir $(abspath $(lastword $(MAKEFILE_LIST))))\n\n")
	}
	AppendToFile(makefile, content)
}

func AppendToNginx(content string) {
	const makefile = `nginx/nginx.conf`
	if _, err := os.Stat(makefile); errors.Is(err, os.ErrNotExist) {
		WriteFile(makefile, "\n\n")
	}
	AppendToFile(makefile, content)
}

func AppendToCerts(mail string, domain string) {
	const certs = `certs.sh`
	if _, err := os.Stat(certs); errors.Is(err, os.ErrNotExist) {
		WriteFile(certs, "go install github.com/go-acme/lego/v4/cmd/lego@latest\n")
	}
	AppendToFile(certs, "sudo lego --email="+mail+" --domains="+domain+" --http run\n")
}

func PrepareDir(filePath string) {
	if len(strings.Split(filePath, `/`)) != 1 {
		splitted := strings.Split(filePath, `/`)
		path := strings.Join(splitted[0:len(splitted)-1], `/`)
		err := os.MkdirAll(path, os.ModePerm)
		checkErr(err)
	}
}

func SystemCall(cmd string) {
	logrus.Info("Executing system call: ", cmd)
	if os.Getenv("IN_DOCKER") != "true" {
		cmd = "docker run --rm -v $(pwd):/wd -w /wd dancheg97.ru/templates/gen-tools:latest " + cmd
	}
	commad := exec.Command("bash", "-c", cmd)
	commad.Stdout = logrus.StandardLogger().Writer()
	commad.Stderr = logrus.StandardLogger().Writer()
	checkErr(commad.Run())
}

func checkErr(err error) {
	if err != nil {
		logrus.Error(err)
		os.Exit(1)
	}
}
