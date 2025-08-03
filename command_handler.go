package main

import (
	"fmt"
)

func HandleRunCommand(args []string) {
	if len(args) < 1 {
		fmt.Println("Usage: create:{middleware|controller} {Name} {Module}")
		return
	}
	command := args[1]
	switch command {
	case "create:middleware":
		HandleCreateMiddleware(args)
	case "create:controller":
		HandleCreateController(args)
	case "create:repository":
		HandleCreateRepository(args)
	case "create:model":
		HandleCreateModel(args)
	case "create:interface":
		HandleCreateInterface(args)
	case "create:service":
		HandleCreateService(args)
	case "create:config":
		HandleCreateConfig(args)
	case "create:init":
		HandleInit(args)
	case "remove:module":
		HandleRemoveModule(args)
	default:
		fmt.Println("Unknown command:", command)
		return
	}
}
