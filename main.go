package main

import (
	"github.com/hashicorp/terraform-plugin-sdk/v2/plugin"
	"simple-terraform-provider-virtualbox/internal/provider"
)

func main() {
	plugin.Serve(&plugin.ServeOpts{
		ProviderFunc: provider.New,
	})
}
