package main

import (
	"flag"
	"fmt"
	"github.com/AgentNemo00/sca-deployment/templates"
	"log"
	"os"
	"text/template"
)

type Vars struct {
	Name     string
	Replicas int
	Image    string
	Port     int
	Local    bool
}

func main() {
	name := flag.String("name", "service", "service and container name")
	templateName := flag.String("template", "deployment-template.yml", "template name")
	replicas := flag.Int("replicas", 1, "amount of replicas")
	image := flag.String("image", "", "image to use")
	local := flag.Bool("local", false, "local usage")
	port := flag.Int("port", 10001, "container port")
	flag.Parse()
	if *image == "" {
		log.Fatal("no image provided")
	}
	helmTemplate, err := template.ParseFS(templates.Helm(), fmt.Sprintf("**/%s", *templateName))
	if err != nil {
		log.Fatal(err)
	}
	helmFile, err := os.Create("deployment.yml")
	if err != nil {
		log.Fatal(err)
	}
	defer helmFile.Close()
	err = helmTemplate.Execute(helmFile, Vars{
		Name:     *name,
		Replicas: *replicas,
		Image:    *image,
		Port:     *port,
		Local:    *local,
	})
	if err != nil {
		log.Fatal(err)
	}
}
