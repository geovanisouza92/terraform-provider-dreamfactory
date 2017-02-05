package types

// Metadata represents bulk request/reponse metadata
type Metadata struct {
	Schema []string `json:"schema,omitempty"`
	Count  int      `json:"count,omitempty"`
}
