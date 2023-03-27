package gpt

// Embedding represents a document and its associated vector.
type Embedding struct {
	ID        string    `json:"id"`
	Embedding []float64 `json:"embedding"`
}
