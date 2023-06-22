package templates

import (
	"embed"
	_ "embed"
)

var (
	//go:embed docker/*
	resDocker embed.FS
	//go:embed helm/*
	resHelm embed.FS
)

func Docker() embed.FS {
	return resDocker
}

func Helm() embed.FS {
	return resHelm
}
