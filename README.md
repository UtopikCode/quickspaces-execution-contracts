# quickspaces-execution-contracts

QuickSpaces Execution Contracts is a provider-agnostic Go module that defines the shared execution contract between control plane components and execution adapters.

## Purpose

This package exposes only the execution interface, DTOs, and workspace lifecycle models necessary to coordinate execution providers without embedding any provider-specific dependencies.

## Architecture Role

- Defines the canonical workspace execution contract.
- Keeps control plane logic separate from adapter implementations.
- Enables multiple execution providers to interoperate against a common API.

## Example Usage

```go
import (
    "context"
    "github.com/UtopikCode/quickspaces-execution-contracts"
)

func runExample(adapter contracts.ExecutionAdapter) error {
    workspace := contracts.Workspace{
        ID:    "ws-001",
        Owner: "team@example.com",
        Repo:  "example/repo",
        Ref:   "main",
        ExecutionProfile: contracts.ExecutionProfile{
            Provider: "truenas",
            Backend:  "k8s",
            RuntimeConfig: map[string]interface{}{
                "cpu": 2,
            },
        },
    }

    state, err := adapter.StartWorkspace(context.Background(), workspace)
    if err != nil {
        return err
    }

    if state == contracts.WorkspaceStateRunning {
        // workspace is active
    }

    return nil
}
```

## Contract Guarantees

- No AWS, Docker, or infrastructure-specific SDK imports.
- Purely a shared contract layer.
- All adapters must implement `ExecutionAdapter`.

## Development

### Format

Run the formatter locally:

```bash
make fmt
```

Check formatting without modifying files:

```bash
make check-format
```

### Lint

Run the lint checker:

```bash
make lint
```

Run lint fixers where available:

```bash
make lint-fix
```

### CI

Run the full verification pipeline locally:

```bash
make ci
```
