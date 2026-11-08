package workspace

type WorkspacePathConfig struct {
	Trust TrustStatus `json:"trust"`
}

type UserWorkspaceConfig struct {
	Username   string                         `json:"username"`
	Workspaces map[string]WorkspacePathConfig `json:"workspaces"`
}
