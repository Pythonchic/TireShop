package storage


type Data struct {
	Products      []Product
	MinPrice      uint16
	MaxPrice      uint16
	Widths        []int
	Profiles      []int
	Diameters     []int
	StoragePath   string
	StorageImages string
}
