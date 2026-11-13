package workspace

type WorkspacePathConfig struct {
	Trust TrustStatus `json:"trust"`
}

type UserWorkspaceConfig struct {
	ProjectName string                         `json:"project_name"`
	Workspaces  map[string]WorkspacePathConfig `json:"workspaces"`
}
