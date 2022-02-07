package datastore

type Board struct {
	ID          string `json:"id"`
	Manufacture string `json:"manufacture"`
}

type BoardReader interface {
	List() ([]Board, error)
}
