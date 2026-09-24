import * as agentops from "@pulumi/agentops";

// A policy grants capabilities; a role bundles policies so they can be granted
// to holders. Both are cheap to create and destroy, which makes them a good
// smoke test for the bridged provider.
const policy = new agentops.Policy("deploy", {
    name: "deploy-policy",
    description: "Grants agent invocation capabilities",
    // grants is a free-form JSON array of grant definitions.
    grants: JSON.stringify([{ capability: "agent.invoke", resource_type: "agent" }]),
});

const role = new agentops.Role("operator", {
    name: "operator",
    description: "Can invoke agents",
    policyIds: [policy.id],
});

export const policyId = policy.id;
export const roleId = role.id;
