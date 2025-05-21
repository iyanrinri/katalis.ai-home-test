package entity

type Book struct {
	Title       string `json:"title"`
	Author      string `json:"author"`
	ISBN        string `json:"isbn"`
	ReleaseDate string `json:"release_date"`
}
