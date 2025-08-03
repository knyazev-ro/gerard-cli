package main

import (
	"os"

	"gopkg.in/yaml.v2"
)

type GeneratedModuleFileStructure struct {
	Repositories string `yaml:"repositories"`
	Controllers  string `yaml:"controllers"`
	Middlewares  string `yaml:"middlewares"`
	Models       string `yaml:"models"`
	Services     string `yaml:"services"`
	Interfaces   string `yaml:"interfaces"`
	Routes       string `yaml:"routes"`
	Utils        string `yaml:"utils"`
	Enums        string `yaml:"enums"`
	Configs      string `yaml:"configs"`
	Tests        string `yaml:"tests"`
	Scripts      string `yaml:"scripts"`
	Docker       string `yaml:"docker"`
	Docs         string `yaml:"docs"`
	ConfigUtils  string `yaml:"config_utils"`
	Src          string `yaml:"src"`
}

type Template struct {
	Repository     string `yaml:"repository"`
	Controller     string `yaml:"controller"`
	Service        string `yaml:"service"`
	Interface      string `yaml:"interface"`
	Model          string `yaml:"model"`
	Middleware     string `yaml:"middleware"`
	Route          string `yaml:"route"`
	Enum           string `yaml:"enum"`
	Dockerfile     string `yaml:"dockerfile"`
	EnvExample     string `yaml:"env-example"`
	Module         string `yaml:"module"`
	Readme         string `yaml:"readme"`
	GitIgnore      string `yaml:"gitignore"`
	Config         string `yaml:"config"`
	ConfigBase     string `yaml:"config-base"`
	ConfigDatabase string `yaml:"config-database"`
	ConfigServer   string `yaml:"config-server"`
	ConfigUtils    string `yaml:"config-utils"`
}

type Command struct {
	CreateInit       bool `yaml:"create-init"`
	CreateModel      bool `yaml:"create-model"`
	CreateConfig     bool `yaml:"create-config"`
	CreateService    bool `yaml:"create-service"`
	CreateInterface  bool `yaml:"create-interface"`
	CreateController bool `yaml:"create-controller"`
	CreateMiddleware bool `yaml:"create-middleware"`
	CreateRepository bool `yaml:"create-repository"`
	RemoveModule     bool `yaml:"remove-module"`
}

type Settings struct {
	GeneratedModuleFileStructure GeneratedModuleFileStructure `yaml:"generated-file-structure"`
	Templates                    Template                     `yaml:"templates"`
	Commands                     Command                      `yaml:"commands"`
}

func DefaultSettings() *Settings {
	settings := Settings{
		GeneratedModuleFileStructure: GeneratedModuleFileStructure{
			Repositories: "src/repositories",
			Controllers:  "src/controllers",
			Middlewares:  "src/middlewares",
			Models:       "src/models",
			Services:     "src/services",
			Interfaces:   "src/interfaces",
			Routes:       "src/routes",
			Utils:        "src/utils",
			Enums:        "src/enums",
			Configs:      "configs",
			Tests:        "tests",
			Scripts:      "scripts",
			Docker:       "docker",
			Docs:         "docs",
			ConfigUtils:  "configs/utils",
			Src:          "src",
		},
		Templates: Template{
			Repository:     "repository.tmpl",
			Controller:     "controller.tmpl",
			Service:        "service.tmpl",
			Interface:      "interface.tmpl",
			Model:          "model.tmpl",
			Middleware:     "middleware.tmpl",
			Route:          "route.tmpl",
			Enum:           "enum.tmpl",
			Dockerfile:     "dockerfile.tmpl",
			EnvExample:     ".env-example.tmpl",
			GitIgnore:      ".gitignore.tmpl",
			ConfigBase:     "config_base.tmpl",
			ConfigDatabase: "config_database.tmpl",
			ConfigServer:   "config_server.tmpl",
			ConfigUtils:    "config_utils.tmpl",
		},

		Commands: Command{
			CreateInit:       true,
			CreateModel:      true,
			CreateConfig:     true,
			CreateService:    true,
			CreateInterface:  true,
			CreateController: true,
			CreateMiddleware: true,
			CreateRepository: true,
			RemoveModule:     true,
		},
	}
	return &settings
}

func LoadSettings() *Settings {

	settings, err := os.ReadFile("gerard/settings.yaml")

	if err != nil {
		println("Error reading settings.yaml:", err.Error())
		println("Using default settings")
		return DefaultSettings()
	}

	var settingsConfig Settings
	err = yaml.Unmarshal(settings, &settingsConfig)
	if err != nil {
		println("Error unmarshalling settings.yaml:", err.Error())
		println("Using default settings")
		return DefaultSettings()
	}
	return &settingsConfig
}
