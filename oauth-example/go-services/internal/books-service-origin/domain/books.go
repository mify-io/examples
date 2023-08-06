package domain

type BooksStorage interface {
	Add(params CreateBookParams) (Book, error)
	Find(id int64) (*Book, error)
	GetAll() ([]Book, error)
}

type Book struct {
	Id   int64
	Name string
}

type CreateBookParams struct {
	Name string
}

type InMemoryBooks struct {
	Books []Book
}

func (s *InMemoryBooks) Add(params CreateBookParams) (Book, error) {
	var id int64 = 0
	if len(s.Books) > 0 {
		id = s.Books[len(s.Books)-1].Id + 1
	}

	book := Book{Id: id, Name: params.Name}
	s.Books = append(s.Books, book)

	return book, nil
}

func (s *InMemoryBooks) Find(id int64) (*Book, error) {
	for _, b := range s.Books {
		if b.Id == id {
			return &b, nil
		}
	}

	return nil, nil
}

func (s *InMemoryBooks) GetAll() ([]Book, error) {
	return s.Books, nil
}
