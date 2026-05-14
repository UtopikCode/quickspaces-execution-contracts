package contracts

import "context"

// ExecutionProfile describes the provider-agnostic settings required to start an execution.
type ExecutionProfile struct {
	Provider      string                 `json:"provider"`
	Backend       string                 `json:"backend"`
	RuntimeConfig map[string]interface{} `json:"runtimeConfig"`
}

// Workspace represents the workspace entity shared between the control plane and execution adapters.
type Workspace struct {
	ID               string           `json:"id"`
	Owner            string           `json:"owner"`
	Repo             string           `json:"repo"`
	Ref              string           `json:"ref"`
	ExecutionProfile ExecutionProfile `json:"executionProfile"`
}

// ExecutionAdapter is the provider-agnostic execution contract for workspace lifecycle operations.
type ExecutionAdapter interface {
	StartWorkspace(ctx context.Context, ws Workspace) (WorkspaceState, error)
	StopWorkspace(ctx context.Context, id string) error
	GetWorkspaceStatus(ctx context.Context, id string) (WorkspaceState, error)
}

// WorkspaceState defines the lifecycle state of a workspace.
type WorkspaceState string

const (
	WorkspaceStatePending WorkspaceState = "pending"
	WorkspaceStateRunning WorkspaceState = "running"
	WorkspaceStateStopped WorkspaceState = "stopped"
	WorkspaceStateError   WorkspaceState = "error"
)
