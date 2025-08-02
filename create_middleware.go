package main

import (
	"flag"
	"fmt"
	"os"
	"strings"
	"text/template"
)

func HandleCreateMiddleware(args []string) {
	tmplFile := "gerard/templates/middleware.tmpl"
	block := "/src/middlewares"

	var module string
	flagSet := flag.NewFlagSet("args", flag.ContinueOnError)
	flagSet.Parse(args[2:])
	allArgs := flagSet.Args()

	if len(allArgs) < 2 {
		fmt.Println("Error: missing fields")
		return
	}
	module = allArgs[1]
	name := allArgs[0]

	data := map[string]string{
		"Package": module,
		"Name":    name,
	}

	// Создаём директорию если нет
	os.MkdirAll(module+block, 0755)
	outFile := fmt.Sprintf("%s%s/%s.go", module, block, strings.ToLower(name))
	tmpl, err := template.ParseFiles(tmplFile) // gerard/templates//middleware.tmpl
	if err != nil {
		panic(err)
	}

	out, err := os.Create(outFile)
	if err != nil {
		panic(err)
	}
	defer out.Close()

	err = tmpl.Execute(out, data)
	if err != nil {
		panic(err)
	}

	fmt.Printf("Created middleware: %s\n", outFile)
}
