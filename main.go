package main

import (
	"github.com/hashicorp/terraform/plugin"

	"github.com/geovanisouza92/terraform-provider-dreamfactory/dreamfactory"
)

func main() {
	plugin.Serve(&plugin.ServeOpts{
		ProviderFunc: dreamfactory.Provider,
	})
}
