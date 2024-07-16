package src

type Book struct {
	BookID   string
	BookName string
	UniqueID string
}

func (b Book) GetBookID() string {
	return b.BookID
}

func (b Book) GetBookName() string {
	return b.BookName
}
