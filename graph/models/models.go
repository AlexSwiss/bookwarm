package models

type Author struct {
	AuthorID  *string `json:"authorID" gorm: "primary_key"`
	Firstname string  `json:"firstname"`
	Lastname  string  `json:"lastname"`
	BookID    string  `json:"bookID"`
}

type Book struct {
	ID     *string   `json:"id" gorm: "primary_key"`
	Name   string    `json:"name"`
	Isbn   int       `json:"ISBN"`
	Author []*Author `json:"author" gorm: "foreign_key:bookID"`
}
