package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"path/filepath"

	"github.com/urfave/cli/v2"
)

var app = cli.NewApp()

func info() {
	app.Name = "Boring React CLI"
	app.Usage = "Will try to scaffold something in react"
	app.Version = "0.0.1"
}

func main() {
	info()
	setComponentScaffoldCommand()
	err := app.Run(os.Args)
	if err != nil {
		log.Fatal(err)
	}
}

func setComponentScaffoldCommand() {
	var name string
	var style string
	componentScaffold := []*cli.Command{
		{
			Name:    "scaffold",
			Aliases: []string{"scf"},
			Flags: []cli.Flag{
				&cli.StringFlag{
					Name:        "name",
					Usage:       "Name of the component",
					Required:    true,
					Destination: &name,
				},
				&cli.StringFlag{
					Name:        "style",
					Value:       "scss",
					Usage:       "type of the style [css, scss]",
					Destination: &style,
				},
			},
			Usage: "Create component folder, file, style and index",
			Action: func(c *cli.Context) error {
				createComponetFile(name, style)
				return nil
			},
		},
	}

	app.Commands = append(app.Commands, componentScaffold...)
}

func createComponetFile(name string, style string) bool {
	dir, err := os.Getwd()

	if err != nil {
		log.Fatal(err)
	}
	componentFile := filepath.Join(dir, name, fmt.Sprintf("%s.js", name))
	styleFile := filepath.Join(dir, name, fmt.Sprintf("%s.%s", name, style))
	indexFile := filepath.Join(dir, name, "index.js")

	fmt.Printf("Scaffloding Component %s with style %s\n", name, style)
	
	dir, err = filepath.Abs(filepath.Dir(componentFile));

	err = os.MkdirAll(dir, os.ModePerm)
	if err != nil {
		log.Fatal(err)
	}

	err = ioutil.WriteFile(componentFile, []byte(getComponentContent(name, style)), 0644)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("Component created @ %s\n", componentFile)
	err = ioutil.WriteFile(styleFile, []byte(getStyleContent(name)), 0644)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("Style created @ %s\n", styleFile)
	err = ioutil.WriteFile(indexFile, []byte(getIndexContent(name)), 0644)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("Index created @ %s\n", indexFile)
	return true
}
