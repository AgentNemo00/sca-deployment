package main

import (
	"flag"
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
	templatePath := flag.String("template", "docker-compose-template.yml", "path to template")
	serviceName := flag.String("service", "service", "service name and container name")
	mainPath := flag.String("main", "main.go", "path to the main.go")
	privateKeyPath := flag.String("key", "./id_rsa", "path to the private key to use to download private dependencies")
	databaseEnabled := flag.Bool("database", false, "enable database service")
	dockerCompose, err := template.ParseFiles(*templatePath)
	if err != nil {
		log.Fatal(err)
	}
	f, err := os.Create("docker-compose.yml")
	if err != nil {
		log.Fatal(err)
	}
	err = dockerCompose.Execute(f, Vars{
		PrivateKeyPath: *privateKeyPath,
		MainPath:       *mainPath,
		ServiceName:    *serviceName,
		Database:       *databaseEnabled,
	})
	if err != nil {
		log.Fatal(err)
	}
}
