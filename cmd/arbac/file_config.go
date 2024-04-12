package main

type Config struct {
	Perms   []string `json:"perms"`
	Outdir  string   `json:"outdir"`
	Package string   `json:"package"`
}
