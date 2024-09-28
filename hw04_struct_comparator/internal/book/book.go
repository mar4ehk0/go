package book

type Book struct {
	id     uint
	title  string
	author string
	year   uint
	size   uint
	rate   float32
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
	b.id = id
}

func (b *Book) ID() uint {
	return b.id
}

func (b *Book) SetTitle(title string) {
	b.title = title
}

func (b *Book) Title() string {
	return b.title
}

func (b *Book) SetAuthor(title string) {
	b.author = title
}

func (b *Book) Author() string {
	return b.author
}

func (b *Book) SetYear(year uint) {
	b.year = year
}

func (b *Book) Year() uint {
	return b.year
}

func (b *Book) SetSize(size uint) {
	b.size = size
}

func (b *Book) Size() uint {
	return b.size
}

func (b *Book) SetRate(rate float32) {
	b.rate = rate
}

func (b *Book) Rate() float32 {
	return b.rate
}
