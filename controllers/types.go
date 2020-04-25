package controllers

type Author struct {
	Age       int    `json:"age"`
	Firstname string `json:"firstname"`
	Lastname  string `json:"lastname"`
}

type Book struct {
	Author *Author `json:"author"`
	ID     string  `json:"id"`
	Isbn   string  `json:"isbn"`
	Title  string  `json:"title"`
}

func getBooksMocks() []Book {
	var books []Book
	books = append(books, Book{
		ID:    "1",
		Isbn:  "99001a",
		Title: "The Dark Tale",
		Author: &Author{
			Firstname: "John",
			Lastname:  "Doe",
			Age:       43,
		}})
	books = append(books, Book{
		ID:    "2",
		Isbn:  "98001b",
		Title: "Forest Rain",
		Author: &Author{
			Firstname: "Jane",
			Lastname:  "Doe",
			Age:       49,
		}})
	books = append(books, Book{
		ID:    "3",
		Isbn:  "98001c",
		Title: "The chocolate river",
		Author: &Author{
			Firstname: "Jim",
			Lastname:  "Moriarty",
			Age:       45,
		}})

	return books
}
