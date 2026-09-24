// Copyright 2016-2024, Pulumi Corporation.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package agentops

import (
	"path"

	// Allow embedding bridge-metadata.json in the provider.
	_ "embed"

	"github.com/komodorio/pulumi-agentops/provider/pkg/version"
	agentopsshim "github.com/komodorio/terraform-provider-agentops/shim"

	pfbridge "github.com/pulumi/pulumi-terraform-bridge/v3/pkg/pf/tfbridge"
	"github.com/pulumi/pulumi-terraform-bridge/v3/pkg/tfbridge"
	"github.com/pulumi/pulumi-terraform-bridge/v3/pkg/tfbridge/tokens"
)

// all of the token components used below.
const (
	// This variable controls the default name of the package in the package
	// registries for nodejs and python:
	mainPkg = "agentops"
	// modules:
	mainMod = "index" // the agentops module
)

//go:embed cmd/pulumi-resource-agentops/bridge-metadata.json
var metadata []byte

// Provider returns additional overlaid schema and metadata associated with the provider.
func Provider() tfbridge.ProviderInfo {
	// Upstream github.com/komodorio/terraform-provider-agentops is MPL-2.0 licensed.
	// TFProviderLicense wants a pointer and MPL20LicenseType is a constant.
	upstreamLicense := tfbridge.MPL20LicenseType

	prov := tfbridge.ProviderInfo{
		// The upstream provider is built on the terraform-plugin-framework, so it is shimmed
		// with pfbridge.ShimProvider rather than the SDKv2 shimv2.NewProvider. Its constructor
		// lives in an internal package and is reached through provider/shim.
		P: pfbridge.ShimProvider(agentopsshim.NewProvider(version.Version)),

		Name:              "agentops",
		Version:           version.Version,
		DisplayName:       "AgentOps",
		Publisher:         "Komodor",
		Description:       "A Pulumi package for creating and managing Komodor AgentOps resources.",
		Keywords:          []string{"agentops", "komodor", "category/cloud"},
		License:           "Apache-2.0",
		Homepage:          "https://komodor.com",
		Repository:        "https://github.com/komodorio/pulumi-agentops",
		TFProviderLicense: &upstreamLicense,
		// Must match the upstream module's require directive, not the shim replace directive.
		// The bridge also uses this to locate the upstream docs/ tree for docs generation.
		GitHubOrg:    "komodorio",
		MetadataInfo: tfbridge.NewProviderMetadata(metadata),

		// Upstream reads these environment variables itself, but declaring them here is what
		// surfaces them to `pulumi config` and to the generated SDK documentation.
		Config: map[string]*tfbridge.SchemaInfo{
			"api_key": {
				Secret:  tfbridge.True(),
				Default: &tfbridge.DefaultInfo{EnvVars: []string{"AGENTOPS_API_KEY"}},
			},
			"endpoint": {
				Default: &tfbridge.DefaultInfo{EnvVars: []string{"AGENTOPS_ENDPOINT"}},
			},
		},

		JavaScript: &tfbridge.JavaScriptInfo{
			// RespectSchemaVersion ensures the SDK is generated linking to the correct version of the provider.
			RespectSchemaVersion: true,
		},
		Python: &tfbridge.PythonInfo{
			// RespectSchemaVersion ensures the SDK is generated linking to the correct version of the provider.
			RespectSchemaVersion: true,
			// Enable modern PyProject support in the generated Python SDK.
			PyProject: struct{ Enabled bool }{true},
		},
		Golang: &tfbridge.GolangInfo{
			// Set where the SDK is going to be published to.
			ImportBasePath: path.Join(
				"github.com/komodorio/pulumi-agentops/sdk/",
				tfbridge.GetModuleMajorVersion(version.Version),
				"go",
				mainPkg,
			),
			// Opt in to all available code generation features.
			GenerateResourceContainerTypes: true,
			GenerateExtraInputTypes:        true,
			// RespectSchemaVersion ensures the SDK is generated linking to the correct version of the provider.
			RespectSchemaVersion: true,
		},
		CSharp: &tfbridge.CSharpInfo{
			// RespectSchemaVersion ensures the SDK is generated linking to the correct version of the provider.
			RespectSchemaVersion: true,
			// Use a wildcard import so NuGet will prefer the latest possible version.
			PackageReferences: map[string]string{
				"Pulumi": "3.*",
			},
		},
	}

	// MustComputeTokens maps all resources and datasources from the upstream provider into Pulumi.
	//
	// tokens.SingleModule puts every upstream item into your provider's main module.
	//
	// You shouldn't need to override anything, but if you do, use the [tfbridge.ProviderInfo.Resources]
	// and [tfbridge.ProviderInfo.DataSources].
	prov.MustComputeTokens(tokens.SingleModule("agentops_", mainMod,
		tokens.MakeStandard(mainPkg)))

	prov.MustApplyAutoAliases()
	prov.SetAutonaming(255, "-")

	return prov
}
