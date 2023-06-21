package main

import (
	"bytes"
	"flag"
	"fmt"
	"github.com/AgentNemo00/sca-deployment/templates"
	"io"
	"log"
	"os"
	"text/template"
)

type Vars struct {
	PrivateKeyPath string
	MainPath       string
	ServiceName    string
	Database       bool
}

func main() {
	templateName := flag.String("template", "docker-compose-template.yml", "path to template")
	serviceName := flag.String("service", "service", "service name and container name")
	mainPath := flag.String("main", "main.go", "path to the main.go")
	privateKeyPath := flag.String("key", "./id_rsa", "path to the private key to use to download private dependencies")
	databaseEnabled := flag.Bool("database", false, "enable database service")
	flag.Parse()
	dockerComposeTemplate, err := template.ParseFS(templates.Docker(), fmt.Sprintf("**/%s", *templateName))
	if err != nil {
		log.Fatal(err)
	}
	dockerComposeFile, err := os.Create("docker-compose.yml")
	if err != nil {
		log.Fatal(err)
	}
	defer dockerComposeFile.Close()
	err = dockerComposeTemplate.Execute(dockerComposeFile, Vars{
		PrivateKeyPath: *privateKeyPath,
		MainPath:       *mainPath,
		ServiceName:    *serviceName,
		Database:       *databaseEnabled,
	})
	if err != nil {
		log.Fatal(err)
	}
	dockerFileTemplate, err := templates.Docker().ReadFile("docker/Dockerfile")
	if err != nil {
		log.Fatal(err)
	}
	dockerFile, err := os.Create("Dockerfile")
	if err != nil {
		log.Fatal(err)
	}
	defer dockerFile.Close()
	_, err = io.Copy(dockerFile, bytes.NewBuffer(dockerFileTemplate))
	if err != nil {
		log.Fatal(err)
	}
}
