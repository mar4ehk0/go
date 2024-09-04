package book

type Book struct {
	ID     uint
	Title  string
	Author string
	Year   uint
	Size   uint
	Rate   float32
}

func CreateBook(id uint, title string, author string, year uint, size uint, rate float32) Book {
	book := Book{}
	book.SetID(id)
	book.SetTitle(title)
	book.SetAuthor(author)
	book.SetYear(year)
	book.SetSize(size)
	book.SetRate(rate)

	return book
}

func (b *Book) SetID(id uint) {
	b.ID = id
}

func (b Book) GetID() uint {
	return b.ID
}

func (b *Book) SetTitle(title string) {
	b.Title = title
}

func (b Book) GetTitle() string {
	return b.Title
}

func (b *Book) SetAuthor(title string) {
	b.Author = title
}

func (b Book) GetAuthor() string {
	return b.Author
}

func (b *Book) SetYear(year uint) {
	b.Year = year
}

func (b Book) GetYear() uint {
	return b.Year
}

func (b *Book) SetSize(size uint) {
	b.Size = size
}

func (b Book) GetSize() uint {
	return b.Size
}

func (b *Book) SetRate(rate float32) {
	b.Rate = rate
}

func (b Book) GetRate() float32 {
	return b.Rate
}
