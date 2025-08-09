package main

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"
)

func HandleCreateTesting(args []string) {

	settings := LoadSettings()
	if settings == nil {
		ErrorPrintln("Error loading settings")
		return
	}
	commandsActivity := settings.Commands

	if !commandsActivity.CreateService {
		WarningPrintln("Test creation is disabled in settings.")
		return
	}

	templates := settings.Templates
	directories := settings.GeneratedModuleFileStructure
	tmplFile := templates.Service

	module := args[1]
	name := args[0]

	module, err := ValidateName(module)
	if err != nil {
		ErrorPrintln("Erro: invalid module name:", err.Error())
		return
	}

	nameCamalCase, _, name, err := Normalize(name)
	if err != nil {
		ErrorPrintln("Filename is invalid.")
		return
	}
	data := map[string]string{
		"Name": "Test" + nameCamalCase,
	}

	target := filepath.Join(module, directories.Tests)
	os.MkdirAll(target, 0755)
	outFile := fmt.Sprintf("%s/%s.go", target, strings.ToLower(name+"_test"))
	path, err := ParseTemplate(tmplFile, outFile, data, args)
	if err != nil {
		ErrorPrintln("Error creating "+path+":", err.Error())
		return
	}

	SuccessPrintln("Created test: ", outFile)
}
