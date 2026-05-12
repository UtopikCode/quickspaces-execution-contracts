package contracts

import (
	"context"
	"testing"
)

type mockAdapter struct{}

func (m mockAdapter) StartWorkspace(ctx context.Context, ws Workspace) (WorkspaceState, error) {
	return WorkspaceStateRunning, nil
}

func (m mockAdapter) StopWorkspace(ctx context.Context, id string) error {
	return nil
}

func (m mockAdapter) GetWorkspaceStatus(ctx context.Context, id string) (WorkspaceState, error) {
	return WorkspaceStateStopped, nil
}

var _ ExecutionAdapter = (*mockAdapter)(nil)

func TestExecutionProfileInitialization(t *testing.T) {
	profile := ExecutionProfile{
		Provider: "truenas",
		Backend:  "k8s",
		RuntimeConfig: map[string]interface{}{
			"cpu":    2,
			"memory": "4Gi",
		},
	}

	if profile.Provider != "truenas" {
		t.Fatalf("expected Provider to be truenas, got %q", profile.Provider)
	}

	if profile.Backend != "k8s" {
		t.Fatalf("expected Backend to be k8s, got %q", profile.Backend)
	}

	if profile.RuntimeConfig["cpu"] != 2 {
		t.Fatalf("expected RuntimeConfig.cpu to be 2, got %#v", profile.RuntimeConfig["cpu"])
	}
}

func TestWorkspaceInitialization(t *testing.T) {
	workspace := Workspace{
		ID:    "ws-123",
		Owner: "dev@example.com",
		Repo:  "example/repo",
		Ref:   "main",
		ExecutionProfile: ExecutionProfile{
			Provider: "aws",
			Backend:  "ecs",
		},
	}

	if workspace.ID != "ws-123" {
		t.Fatalf("expected ID to be ws-123, got %q", workspace.ID)
	}

	if workspace.ExecutionProfile.Provider != "aws" {
		t.Fatalf("expected ExecutionProfile.Provider to be aws, got %q", workspace.ExecutionProfile.Provider)
	}
}

func TestWorkspaceStateConstants(t *testing.T) {
	if WorkspaceStatePending != "pending" {
		t.Fatalf("WorkspaceStatePending value changed")
	}

	if WorkspaceStateRunning != "running" {
		t.Fatalf("WorkspaceStateRunning value changed")
	}

	if WorkspaceStateStopped != "stopped" {
		t.Fatalf("WorkspaceStateStopped value changed")
	}

	if WorkspaceStateError != "error" {
		t.Fatalf("WorkspaceStateError value changed")
	}
}
