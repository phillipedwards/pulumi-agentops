// Package shim re-exports the upstream AgentOps Terraform provider.
//
// The upstream constructor lives in an internal package, which Go only permits
// importing from within the github.com/komodorio/terraform-provider-agentops
// tree. Declaring this module under that path satisfies the internal rule and
// lets the Pulumi bridge reach the provider. See
// https://github.com/pulumi/pulumi-random/tree/master/provider/shim.
package shim

import (
	tfpf "github.com/hashicorp/terraform-plugin-framework/provider"
	"github.com/komodorio/terraform-provider-agentops/internal/provider"
)

// NewProvider returns the upstream provider. provider.New returns a factory, so
// it is invoked once here.
func NewProvider(version string) tfpf.Provider {
	return provider.New(version)()
}
