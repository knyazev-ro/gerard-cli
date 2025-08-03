package main

import (
	"flag"
	"fmt"
	"os"
	"path/filepath"
	"strings"
)

func HandleCreateConfig(args []string) {

	settings := LoadSettings()
	if settings == nil {
		println("Error loading settings")
		return
	}
	commandsActivity := settings.Commands

	if !commandsActivity.CreateMiddleware {
		fmt.Println("Config creation is disabled in settings.")
		return
	}

	templates := settings.Templates
	directories := settings.GeneratedModuleFileStructure
	tmplFile := templates.Config

	var module string
	flagSet := flag.NewFlagSet("args", flag.ContinueOnError)
	flagSet.Parse(args[2:])

	names := flagSet.Args()
	module = names[1]
	name := names[0]

	nameCamalCase, nameVar := CreateStructNameAndVar(name)
	data := map[string]string{
		"Module":  module,
		"Name":    nameCamalCase,
		"NameVar": nameVar,
	}
	target := filepath.Join(module, directories.Configs)
	os.MkdirAll(target, 0755)
	outFile := fmt.Sprintf("%s/%s.go", target, strings.ToLower(name))
	path, err := ParseTemplate(tmplFile, outFile, data)
	if err != nil {
		println("Error creating "+path+":", err.Error())
		return
	}

	fmt.Printf("Created config: %s\n", outFile)
}
