// main package for entry point of app
package main

import (
	"bytes"
	"fmt"
	"html/template"
	"os"
	"path/filepath"

	"gopkg.in/yaml.v3"
)

type (
	// Pkg ...
	Pkg struct {
		Domain string `yaml:"domain"`
	}

	// Godoc ...
	Godoc struct {
		Domain string `yaml:"domain"`
	}

	// PkgInfo ...
	PkgInfo struct {
		Name        string `yaml:"name"`
		Description string `yaml:"description"`
		RepoURL     string `yaml:"repo_url"`
		EnableGodoc bool   `yaml:"enable_godoc"`
	}

	// Info ...
	Info struct {
		PageTitle       string `yaml:"page_title"`
		PageDescription string `yaml:"page_description"`
	}

	// AppConfig ...
	AppConfig struct {
		Info    Info  `yaml:"info"`
		Godoc   Godoc `yaml:"godoc"`
		Pkg     Pkg   `yaml:"pkg"`
		Version string
		Repos   []PkgInfo `yaml:"repos"`
		Debug   bool      `yaml:"debug"`
	}
)

func loadConfig() (cfg AppConfig, err error) {
	yamlFile, err := os.ReadFile("packages.yaml")
	if err != nil {
		return AppConfig{}, err
	}
	err = yaml.Unmarshal(yamlFile, &cfg)
	if err != nil {
		return AppConfig{}, err
	}

	return cfg, nil
}

func shutdown(err error) {
	fmt.Println(err.Error())
	fmt.Println("kthxbye!")
	os.Exit(1)
}

func main() {
	packages, err := loadConfig()
	if err != nil {
		shutdown(err)
	}

	t, err := template.ParseFiles("package.tmpl")
	if err != nil {
		shutdown(err)
	}

	// for _, pkg := range packages.Repos {
	// 	pkgDir := filepath.Join("public", pkg.Name)
	// 	if err := os.MkdirAll(pkgDir, 0o755); err != nil {
	// 		shutdown(err)
	// 	}
	// 	var page bytes.Buffer
	// 	if err := t.Execute(&page, pkg); err != nil {
	// 		shutdown(err)
	// 	}
	//
	// 	pkgFile := filepath.Join(pkgDir, "index.html")
	// 	if err := os.WriteFile(pkgFile, page.Bytes(), 0o755); err != nil {
	// 		shutdown(err)
	// 	}
	// }

	t, err = template.ParseFiles("index.tmpl")
	if err != nil {
		shutdown(err)
	}
	var page bytes.Buffer
	if err := t.Execute(&page, packages); err != nil {
		shutdown(err)
	}

	pkgFile := filepath.Join("public", "index.html")
	if err := os.WriteFile(pkgFile, page.Bytes(), 0o755); err != nil {
		shutdown(err)
	}
}
