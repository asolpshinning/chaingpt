package gpt

import (
	"fmt"
	"math"
	"sort"

	"gonum.org/v1/gonum/floats"
	"gonum.org/v1/gonum/mat"
)

// BySimilarity is a custom sorting type for sorting Embeddings by similarity.
type BySimilarity []Embedding

func (a BySimilarity) Len() int           { return len(a) }
func (a BySimilarity) Swap(i, j int)      { a[i], a[j] = a[j], a[i] }
func (a BySimilarity) Less(i, j int) bool { return a[i].Embedding[0] > a[j].Embedding[0] }

// CosineSimilarity calculates the cosine similarity between two vectors.
func CosineSimilarity(a, b []float64) float64 {
	normA := mat.Norm(mat.NewVecDense(len(a), a), 2)
	normB := mat.Norm(mat.NewVecDense(len(b), b), 2)
	dotProduct := floats.Dot(a, b)

	return dotProduct / (normA * normB)
}

// SimilaritySearch performs a similarity search on the indexed embeddings using the query embedding.
func SimilaritySearch(index []Embedding, query []float64) []Embedding {
	for i := range index {
		similarity := CosineSimilarity(index[i].Embedding, query)
		index[i].Embedding = []float64{similarity}
	}

	sort.Sort(BySimilarity(index))
	return index
}

func TestingSimSearch() {
	// Your indexed embeddings
	index := []Embedding{
		{"doc1", []float64{0.1, 0.2, 0.3, 0.4}},
		{"doc2", []float64{0.5, 0.6, 0.7, 0.8}},
		{"doc3", []float64{0.9, 1.0, 1.1, 1.2}},
	}

	// Your query embedding
	query := []float64{0.1, 0.2, 0.3, 0.4}

	// Perform similarity search
	results := SimilaritySearch(index, query)

	// Print top-k results
	topK := 3
	fmt.Println("Top", topK, "results:")
	for i := 0; i < int(math.Min(float64(topK), float64(len(results)))); i++ {
		fmt.Printf("%s: %f\n", results[i].ID, results[i].Embedding[0])
	}
}
