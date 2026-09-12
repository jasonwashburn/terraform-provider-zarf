package main

import (
	"context"
	"flag"
	"log"

	"terraform-provider-zarf/internal/provider"

	"github.com/hashicorp/terraform-plugin-framework/providerserver"
)

//go:generate go tool tfplugindocs generate -provider-name zarf

// intended to be set by goreleaser during build.
var version string = "dev"

func main() {
	var debug bool

	flag.BoolVar(&debug, "debug", false, "used for debugging")
	flag.Parse()

	opts := providerserver.ServeOpts{
		Address: "registry.terraform.io/jasonwashburn/zarf",
		Debug:   debug,
	}

	err := providerserver.Serve(context.Background(), provider.New(version), opts)
	if err != nil {
		log.Fatal(err.Error())
	}
}
