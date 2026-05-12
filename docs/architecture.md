# Architecture

## Execution Abstraction

This repository defines a provider-agnostic execution contract for QuickSpaces.
The package exposes only execution interfaces, DTOs, and lifecycle models.

Adapters implement the execution contract and map provider-specific behavior to the shared API.
Control-plane components depend on this package to manage workspace lifecycle without embedding infrastructure concerns.

## Separation from Control Plane

The contract layer is intentionally minimal:

- `Workspace` describes the workspace identity and execution profile.
- `ExecutionProfile` carries provider-agnostic runtime settings.
- `ExecutionAdapter` defines the lifecycle operations any execution provider must support.
- `WorkspaceState` enumerates lifecycle states common across providers.

This package MUST NOT contain control-plane logic, orchestration state machines, or provider-specific libraries.
Those responsibilities belong to the control-plane and adapter implementations.

## Lifecycle Semantics

Workspaces are represented as a logical execution unit.
The execution contract defines the following states:

- `pending`: workspace creation is in progress.
- `running`: workspace is active and ready.
- `stopped`: workspace has been stopped or terminated cleanly.
- `error`: workspace is in a failed state and requires attention.

Each adapter must interpret those states in a provider-consistent way while preserving the contract semantics.
