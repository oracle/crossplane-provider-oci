// Package hack embeds the main template for sub-packages
package hack

import (
	_ "embed"
)

//go:embed main.go.tmpl
var MainTemplate string

//go:embed config_main.go.tmpl
var ConfigMainTemplate string
