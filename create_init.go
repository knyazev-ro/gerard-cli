package main

import (
	"flag"
	"os"
	"os/exec"
	"text/template"
)

func HandleInit(args []string) {

	templatesFolder := "gerard/templates"
	flagSet := flag.NewFlagSet("args", flag.ContinueOnError)
	flagSet.Parse(args[2:])
	allArgs := flagSet.Args()

	if len(allArgs) < 1 {
		println("Error: missing module name")
		return
	}

	module := allArgs[0]
	if module == "" {
		println("Error: module name cannot be empty")
		return
	}
	println("Initializing module:", module)
	src := module + "/src"
	tests := module + "/tests"
	scripts := module + "/scripts"
	docker := module + "/docker"

	middlewares := src + "/middlewares"
	controllers := src + "/controllers"
	routes := src + "/routes"
	configs := src + "/configs"
	interfaces := src + "/interfaces"
	models := src + "/models"
	services := src + "/services"
	repositories := src + "/repositories"
	utils := src + "/utils"
	enums := src + "/enums"

	os.MkdirAll(src, 0755)
	os.MkdirAll(middlewares, 0755)
	os.MkdirAll(controllers, 0755)
	os.MkdirAll(routes, 0755)
	os.MkdirAll(configs, 0755)
	os.MkdirAll(interfaces, 0755)
	os.MkdirAll(models, 0755)
	os.MkdirAll(services, 0755)
	os.MkdirAll(repositories, 0755)
	os.MkdirAll(utils, 0755)
	os.MkdirAll(tests, 0755)
	os.MkdirAll(scripts, 0755)
	os.MkdirAll(docker, 0755)
	os.MkdirAll(enums, 0755)

	// You can also create a README.md or other initial files here
	readmeFile := module + "/README.md"
	readmeContent := "# " + module + "\n\nThis is the initial setup for the " + module + " module.\n"
	err := os.WriteFile(readmeFile, []byte(readmeContent), 0644)
	if err != nil {
		println("Error creating README.md:", err.Error())
		return
	}

	//create main.go based on template module.tmpl in gerard/templates
	tmplModuleFile := templatesFolder + "/module.tmpl"
	if _, err := os.Stat(tmplModuleFile); os.IsNotExist(err) {
		println("Error: module template not found:", tmplModuleFile)
		return
	}
	outFile := module + "/main.go"
	tmpl, err := template.ParseFiles(tmplModuleFile)

	if err != nil {
		println("Error parsing template:", err.Error())
		return
	}

	file, err := os.Create(outFile)
	if err != nil {
		println("Error creating main.go:", err.Error())
		return
	}

	defer file.Close()
	err = tmpl.Execute(file, map[string]string{"Module": module})

	if err != nil {
		println("Error executing template:", err.Error())
		return
	}

	//create routes.go based on template route.tmpl in gerard/templates
	routesFile := routes + "/routes.go"
	routesTmplFile := templatesFolder + "/route.tmpl"

	if _, err := os.Stat(routesTmplFile); os.IsNotExist(err) {
		println("Error: routes template not found:", routesTmplFile)
		return
	}
	routesTmpl, err := template.ParseFiles(routesTmplFile)
	if err != nil {
		println("Error parsing routes template:", err.Error())
		return
	}
	routesOut, err := os.Create(routesFile)
	if err != nil {
		println("Error creating routes.go:", err.Error())
		return
	}
	defer routesOut.Close()
	err = routesTmpl.Execute(routesOut, map[string]string{"Module": module})
	if err != nil {
		println("Error executing routes template:", err.Error())
		return
	}

	//run go mod init and tidy
	cmd := exec.Command("go", "mod", "init", module)
	cmd.Dir = module
	err = cmd.Run()

	if err != nil {
		println("Error initializing Go module:", err.Error())
		return
	}

	cmd = exec.Command("go", "mod", "tidy")
	cmd.Dir = module
	err = cmd.Run()

	if err != nil {
		println("Error running go mod tidy:", err.Error())
		return
	}

	// create a .gitignore file from the template in gerard/templates
	gitignoreFile := module + "/.gitignore"
	gitignoreTmplFile := templatesFolder + "/gitignore.tmpl"
	if _, err := os.Stat(gitignoreTmplFile); os.IsNotExist(err) {
		println("Error: .gitignore template not found:", gitignoreTmplFile)
		return
	}
	gitignoreTmpl, err := template.ParseFiles(gitignoreTmplFile)
	if err != nil {
		println("Error parsing .gitignore template:", err.Error())
		return
	}
	gitignoreOut, err := os.Create(gitignoreFile)
	if err != nil {
		println("Error creating .gitignore:", err.Error())
		return
	}

	defer gitignoreOut.Close()
	err = gitignoreTmpl.Execute(gitignoreOut, map[string]string{"Module": module})
	if err != nil {
		println("Error executing .gitignore template:", err.Error())
		return
	}

	//create a Dockerfile from the template in gerard/templates
	dockerfile := docker + "/Dockerfile"
	dockerTmplFile := templatesFolder + "/dockerfile.tmpl"
	if _, err := os.Stat(dockerTmplFile); os.IsNotExist(err) {
		println("Error: Dockerfile template not found:", dockerTmplFile)
		return
	}
	dockerTmpl, err := template.ParseFiles(dockerTmplFile)
	if err != nil {
		println("Error parsing Dockerfile template:", err.Error())
		return
	}
	dockerOut, err := os.Create(dockerfile)
	if err != nil {
		println("Error creating Dockerfile:", err.Error())
		return
	}
	defer dockerOut.Close()
	err = dockerTmpl.Execute(dockerOut, map[string]string{"Module": module})
	if err != nil {
		println("Error executing Dockerfile template:", err.Error())
		return
	}

	// create env example file
	envExampleFile := module + "/.env.example"
	envExampleContent := "DB_HOST=localhost\nDB_PORT=5432\nDB_USER=user\nDB_PASSWORD=password\n"
	err = os.WriteFile(envExampleFile, []byte(envExampleContent), 0644)
	if err != nil {
		println("Error creating .env.example:", err.Error())
		return
	}
	// create a config file from the template in gerard/templates
	configFile := configs + "/config.go"
	configTmplFile := templatesFolder + "/config.tmpl"
	if _, err := os.Stat(configTmplFile); os.IsNotExist(err) {
		println("Error: config template not found:", configTmplFile)
		return
	}
	configTmpl, err := template.ParseFiles(configTmplFile)
	if err != nil {
		println("Error parsing config template:", err.Error())
		return
	}
	configOut, err := os.Create(configFile)
	if err != nil {
		println("Error creating config.go:", err.Error())
		return
	}
	defer configOut.Close()
	err = configTmpl.Execute(configOut, map[string]string{"Module": module})
	if err != nil {
		println("Error executing config template:", err.Error())
		return
	}

	cmd = exec.Command("git", "init")
	cmd.Dir = module
	err = cmd.Run()

	if err != nil {
		println("Error initializing git repository:", err.Error())
		return
	}
	// Here you would typically create the module directory structure
	// and possibly generate some initial files or configurations.
	// For now, we just print the module name.
	println("Module initialized successfully:", module)
	println("You can now start adding controllers, middlewares, and routes to your module.")

}
