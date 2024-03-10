package assets

import "embed"

//go:embed templates/*
var Assets embed.FS

//go:embed certs/cert.pem
var Certpem embed.FS

//go:embed certs/key.pem
var Keypem embed.FS
