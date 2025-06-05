package main

// Strukturer för att representera JSON-datan

type Shards struct {
	Total      int `json:"total"`
	Successful int `json:"successful"`
	Skipped    int `json:"skipped"`
	Failed     int `json:"failed"`
}

type Total struct {
	Value    int    `json:"value"`
	Relation string `json:"relation"`
}

type StringFacet struct {
	FacetName  string `json:"facet_name"`
	FacetValue string `json:"facet_value"`
}

type Source struct {
	WebID              int           `json:"webid"`
	Title              string        `json:"title"`
	Description        string        `json:"description"`
	CombinedSearchText string        `json:"combinedsearchtext"`
	Price              float64       `json:"price"`
	CategoryName       string        `json:"categoryName"`
	StockLevel         int           `json:"stockLevel"`
	Color              string        `json:"color"`
	CategoryID         int           `json:"categoryId"`
	StringFacet        []StringFacet `json:"string_facet"`
}

type Hit struct {
	Index  string  `json:"_index"`
	ID     string  `json:"_id"`
	Score  float64 `json:"_score"`
	Source Source  `json:"_source"`
}

type Hits struct {
	Total    Total   `json:"total"`
	MaxScore float64 `json:"max_score"`
	Hits     []Hit   `json:"hits"`
}

type Bucket struct {
	Key      string `json:"key"`
	DocCount int    `json:"doc_count"`
}

type Values struct {
	DocCountErrorUpperBound int      `json:"doc_count_error_upper_bound"`
	SumOtherDocCount        int      `json:"sum_other_doc_count"`
	Buckets                 []Bucket `json:"buckets"`
}

type NameBucket struct {
	Key      string `json:"key"`
	DocCount int    `json:"doc_count"`
	Values   Values `json:"values"`
}

type Names struct {
	DocCountErrorUpperBound int          `json:"doc_count_error_upper_bound"`
	SumOtherDocCount        int          `json:"sum_other_doc_count"`
	Buckets                 []NameBucket `json:"buckets"`
}

type Facets struct {
	DocCount int   `json:"doc_count"`
	Names    Names `json:"names"`
}

type Aggregations struct {
	Facets Facets `json:"facets"`
}

type Response struct {
	Took         int          `json:"took"`
	TimedOut     bool         `json:"timed_out"`
	Shards       Shards       `json:"_shards"`
	Hits         Hits         `json:"hits"`
	Aggregations Aggregations `json:"aggregations"`
}
