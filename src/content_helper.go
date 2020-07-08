package main

import (
	"fmt"
	"strings"
)

// Return the content of the component file
func getComponentContent(name string, style string) string {
	styleFile := fmt.Sprintf("%s.%s", name, style)
	componentName := strings.ReplaceAll(name, "-", " ")
	componentName = strings.Title(componentName)
	componentName = strings.ReplaceAll(componentName, " ", "")
	return fmt.Sprintf(
		`import React from "react";
import "./%s";

const %s = () => {
	return <div className="%s"></div>;
}

export default %s;`, styleFile, componentName, name, componentName)
}

// Return content of the style file
func getStyleContent(name string) string {
	return fmt.Sprintf(
		`.%s {

}`, name)
}

// Return content of the index file
func getIndexContent(name string) string {
	return fmt.Sprintf(`export { default } from "./%s";`, name)
}
