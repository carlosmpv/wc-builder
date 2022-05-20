package main

import (
	_ "embed"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"path"
	"strings"
	"text/template"

	"github.com/iancoleman/strcase"
)

//go:embed base-template.js.tmpl
var BaseTemplate string

//go:embed wc-template.js.tmpl
var WcTemplate string

//go:embed user-defined-behavior-template.js.tmpl
var UDBTemplate string

type UDBArgs struct {
	ElementName string
}

func scaffold(name string) {
	os.Mkdir("elements", os.ModePerm)

	tmpl, err := template.New("udb").Parse(UDBTemplate)
	if err != nil {
		log.Fatal(err)
	}

	err = os.Mkdir(path.Join("elements", name), os.ModePerm)
	if err != nil {
		log.Fatal("Component has already been created")
	}

	fileJS, err := os.Create(fmt.Sprintf("./elements/%s/%s.js", name, name))
	if err != nil {
		log.Fatal(err)
	}
	defer fileJS.Close()

	err = tmpl.Execute(fileJS, UDBArgs{strcase.ToCamel(name)})
	if err != nil {
		log.Fatal(err)
	}

	fileHTML, err := os.Create(fmt.Sprintf("./elements/%s/%s.html", name, name))
	if err != nil {
		log.Fatal(err)
	}
	defer fileHTML.Close()

	fileHTML.Write([]byte("<h1>Hello World!</h1>"))
}

type Component struct {
	ElementNameCC string
	ElementNameKC string
	UDB           string
	View          string
}

func build(name string) Component {
	fileHTML, err := os.Open(fmt.Sprintf("./elements/%s/%s.html", name, name))
	if err != nil {
		log.Fatal(err)
	}
	defer fileHTML.Close()

	view, err := ioutil.ReadAll(fileHTML)
	if err != nil {
		log.Fatal(err)
	}

	fileJS, err := os.Open(fmt.Sprintf("./elements/%s/%s.js", name, name))
	if err != nil {
		log.Fatal(err)
	}
	defer fileJS.Close()

	behavior, err := ioutil.ReadAll(fileJS)
	if err != nil {
		log.Fatal(err)
	}

	return Component{
		ElementNameCC: strcase.ToCamel(name),
		ElementNameKC: strcase.ToKebab(name),
		UDB:           string(behavior),
		View:          string(view),
	}
}

func bundle(components []Component) {
	tmpl, err := template.New("bundle").Parse(strings.Join([]string{WcTemplate, BaseTemplate}, "\n"))
	if err != nil {
		log.Fatal(err)
	}

	os.Mkdir("./dist", os.ModePerm)

	file, err := os.Create("./dist/bundle.js")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	tmpl.Execute(file, components)
}

func bundleFolder() {
	elements, err := ioutil.ReadDir("./elements")
	if err != nil {
		log.Fatal(err)
	}

	components := []Component{}
	for _, element := range elements {
		components = append(components, build(element.Name()))
	}
	bundle(components)
}

func showHelp() {
	fmt.Printf("%s [new | build]\n\n\tnew <element name> - creats a new custom element on ./elements/<element name>\n\tbuild - bundles all elements to ./dist/bundle.js\n", os.Args[0])
}

func main() {
	if len(os.Args) < 2 {
		showHelp()
		return
	}

	switch os.Args[1] {
	case "new":
		if len(os.Args) < 3 {
			showHelp()
			return
		}

		scaffold(os.Args[2])

	case "build":
		bundleFolder()
	}
}
