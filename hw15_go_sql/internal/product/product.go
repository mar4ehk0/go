package product

type Product struct {
	Id    int    `json:"id"`    //nolint:structtag
	Name  string `json:"name"`  //nolint:structtag
	Price int    `json:"price"` //nolint:structtag
}
