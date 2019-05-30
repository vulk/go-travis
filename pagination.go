package travis

type Pagination struct {
	Limit   int                      `json:"limit,omitempty"`
	Offset  int                      `json:"offset,omitempty"`
	Count   int                      `json:"count,omitempty"`
	IsFirst bool                     `json:"is_first,omitempty"`
	IsLast  bool                     `json:"is_last,omitempty"`
	Next    *PaginationPageReference `json:"next,omitempty"`
	Prev    *PaginationPageReference `json:"prev,omitempty"`
	First   *PaginationPageReference `json:"first,omitempty"`
	Last    *PaginationPageReference `json:"last,omitempty"`
}

type PaginationPageReference struct {
	HREF   string `json:"@href,omitempty"`
	Offset int    `json:"offset,omitempty"`
	Limit  int    `json:"limit,omitempty"`
}
