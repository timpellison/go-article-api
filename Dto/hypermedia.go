package Dto

type Hypermedia struct {
	Relation  string `json:"_rel"`
	Reference string `json:"_ref"`
}

func NewHypermedia(relation string, reference string) Hypermedia {
	var result = &Hypermedia{}
	result.Relation = relation
	result.Reference = reference
	return *result
}
