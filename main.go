package main

import (
	"context"
	"flag"
	"log"

	"github.com/hashicorp/terraform-plugin-framework/providerserver"

	validatefxprovider "github.com/The-DevOps-Daily/terraform-provider-validatefx/internal/provider"
)

var version string = "dev"

func main() {
	var debug bool

	flag.BoolVar(&debug, "debug", false, "set to true to run the provider with support for debuggers like delve")
	flag.Parse()

	opts := providerserver.ServeOpts{
		Address: "registry.terraform.io/The-DevOps-Daily/validatefx",
		Debug:   debug,
	}

	if err := providerserver.Serve(context.Background(), validatefxprovider.New(version), opts); err != nil {
		log.Fatal(err.Error())
	}
}
