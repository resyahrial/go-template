package main

import "flag"

type (
	Flag struct {
		Environment string
	}
)

var (
	appFlag Flag
)

func init() {
	flag.StringVar(
		&appFlag.Environment,
		"env",
		"dev",
		"env of deployment, will load the respective yml conf file.",
	)
}

func main() {
	flag.Parse()
}
