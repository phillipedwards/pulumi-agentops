# Komodor AgentOps Resource Provider

The AgentOps Resource Provider lets you manage [Komodor AgentOps](https://komodor.com)
config-plane resources with Pulumi. It is a [bridged](https://github.com/pulumi/pulumi-terraform-bridge)
wrapper around the upstream Terraform provider
[`komodorio/terraform-provider-agentops`](https://github.com/komodorio/terraform-provider-agentops),
which is built on the Terraform Plugin Framework.

## Status

This provider is **not published yet** — it is built and consumed locally. See
[Local development](#local-development) below.

## Installing

Once published, this package will be available for several languages/platforms:

### Node.js (JavaScript/TypeScript)

```bash
npm install @pulumi/agentops     # or: yarn add @pulumi/agentops
```

### Python

```bash
pip install pulumi_agentops
```

### Go

```bash
go get github.com/komodorio/pulumi-agentops/sdk/go/...
```

### .NET

```bash
dotnet add package Pulumi.Agentops
```

## Configuration

The following configuration points are available for the `agentops` provider:

- `agentops:apiKey` (environment: `AGENTOPS_API_KEY`) - AgentOps API key, sent as a Bearer
  token. Required. Marked secret.
- `agentops:endpoint` (environment: `AGENTOPS_ENDPOINT`) - AgentOps control-plane base URL.
  Defaults to `https://agentops.komodor.com`. Use `https://staging.agentops.komodor.com` for
  staging, or your own URL when self-hosting.

## Example

```typescript
import * as agentops from "@pulumi/agentops";

const policy = new agentops.Policy("deploy", {
    name: "deploy-policy",
    description: "Grants agent invocation capabilities",
    grants: JSON.stringify([{ capability: "agent.invoke", resource_type: "agent" }]),
});

const role = new agentops.Role("operator", {
    name: "operator",
    description: "Can invoke agents",
    policyIds: [policy.id],
});
```

## Local development

```bash
mise trust .config/mise.toml && mise install -y
PULUMI_HOME=$PWD/.pulumi pulumi plugin install converter terraform   # needed by tfgen
make tfgen        # generate schema.json + bridge-metadata.json
make provider     # build bin/pulumi-resource-agentops
make build_sdks   # generate and build all four SDKs
mise x -- make lint
```

To use the locally built provider, put `bin/pulumi-resource-agentops` on your `PATH` and run
`make install_nodejs_sdk` (or install `sdk/python/bin`) before `pulumi up`.

### Tracking upstream

The upstream provider is pinned in `provider/shim/go.mod`. To take a new version, bump it
there and in `provider/go.mod`, then re-run `make tfgen`.

`provider/shim` exists because the upstream constructor lives in an `internal` package. The
shim module declares itself under the upstream module path so Go's internal-import rule
permits the reference; `provider/go.mod` then `replace`s it with the local directory.

## Reference

Upstream resource documentation lives at
[registry.terraform.io/providers/komodorio/agentops](https://registry.terraform.io/providers/komodorio/agentops/latest/docs);
the bridge converts it into the generated SDK docs.
