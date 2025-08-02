package main

import (
	"flag"
	"fmt"
	"os"
	"strings"
	"text/template"

	"golang.org/x/text/cases"
	"golang.org/x/text/language"
)

func HandleCreateRepository(args []string) {

	tmplFile := "gerard/templates/repository.tmpl"
	target := "/src/repositories"

	var module string
	flagSet := flag.NewFlagSet("args", flag.ContinueOnError)
	flagSet.Parse(args[2:])

	names := flagSet.Args()
	module = names[1]
	name := names[0]

	splitName := strings.Split(name, "_")
	for i, s := range splitName {
		splitName[i] = cases.Title(language.English).String(s)
	}
	data := map[string]string{
		"Name": strings.Join(splitName, ""),
	}

	os.MkdirAll(module+target, 0755)
	outFile := fmt.Sprintf("%s%s/%s.go", module, target, name)
	tmpl, err := template.ParseFiles(tmplFile)

	if err != nil {
		panic(err)
	}

	file, err := os.Create(outFile)

	if err != nil {
		panic(err)
	}

	defer file.Close()
	err = tmpl.Execute(file, data)

	if err != nil {
		panic(err)
	}

	fmt.Printf("Created repository: %s\n", outFile)
}
