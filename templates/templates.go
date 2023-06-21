package templates

import (
	"embed"
	_ "embed"
)

var (
	//go:embed docker/*
	resDocker embed.FS
)

func Docker() embed.FS {
	return resDocker
}
