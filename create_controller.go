package main

import (
	"flag"
	"fmt"
	"os"
	"path/filepath"
	"strings"
)

func HandleCreateController(args []string) {

	settings := LoadSettings()
	if settings == nil {
		println("Error loading settings")
		return
	}

	commandsActivity := settings.Commands

	if !commandsActivity.CreateMiddleware {
		fmt.Println("Controller creation is disabled in settings.")
		return
	}

	templates := settings.Templates
	directories := settings.GeneratedModuleFileStructure
	tmplFile := templates.Controller

	var module string
	flagSet := flag.NewFlagSet("args", flag.ContinueOnError)
	flagSet.Parse(args[2:])

	names := flagSet.Args()
	module = names[1]
	name := names[0]

	nameCamalCase, _ := CreateStructNameAndVar(name)
	data := map[string]string{
		"Name": nameCamalCase,
	}

	target := filepath.Join(module, directories.Controllers)
	os.MkdirAll(target, 0755)
	outFile := fmt.Sprintf("%s/%s.go", target, strings.ToLower(name))
	path, err := ParseTemplate(tmplFile, outFile, data)
	if err != nil {
		println("Error creating "+path+":", err.Error())
		return
	}

	fmt.Printf("Created controller: %s\n", outFile)
}
