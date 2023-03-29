package entity

type Product struct {
	ID        uint64
	Name      string
	Deskripsi string
	Harga     uint64
	// Category  Category

	User   User
	UserID uint64

	// Type   []Type
	// Like   []Like
	// Review []Review
}
