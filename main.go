package main

import (
	"fmt"
	"log"
	"os"
	"sample/utils"
	"text/template"
)

type TemplateData struct {
	UserName    string
	ProjectName string
}

func main() {

	data := TemplateData{
		UserName:    "azar-intelops",
		ProjectName: "test",
	}
	projectRoot := "code"
	templateDir := "templates/grpc"

	for _, e := range utils.GetDirectoryList(templateDir) {
		fmt.Println(utils.CheckIsDir(e.Name()))
	}

	if !utils.CheckDirectory(projectRoot, "./") {
		if err := os.Mkdir(projectRoot, os.ModePerm); err != nil {
			log.Fatal(err)
		}
	}

	tmpl := template.Must(template.ParseFiles("templates/grpc/go.mod.tmpl"))
	file, err := os.Create("./" + projectRoot + "/go.mod")
	if err != nil {
		panic(err)
	}
	defer file.Close()

	err = tmpl.Execute(file, data)
	if err != nil {
		panic(err)
	}
}
