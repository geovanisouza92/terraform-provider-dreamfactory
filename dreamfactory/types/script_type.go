package types

// ScriptTypes represents a collection of supported script types
type ScriptTypes struct {
	Resource []struct {
		Name                    string `json:"name"`
		Label                   string `json:"label"`
		Description             string `json:"description"`
		Sandboxed               bool   `json:"sandboxed"`
		SupportsInlineExecution bool   `json:"supports_inline_execution"`
	} `json:"resource"`
}
