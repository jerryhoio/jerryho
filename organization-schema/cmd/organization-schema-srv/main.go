package main

import (
	"log"

	"github.com/jerryhoio/jerryho/organization-schema/internal/cmd"
)

func main() {
	if err := cmd.RunOrgianizationSchemaSrv(); err != nil {
		log.Fatal(err)
	}
}
