package web

import (
	"embed"
)

//go:generate npm run build

//go:embed dist/*
var Dist embed.FS
