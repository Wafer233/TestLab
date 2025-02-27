package dao

type Author struct {
	Name  string
	Email string
}

type BlogV1 struct {
	ID      int
	Author  Author `gorm:"embedded"`
	Upvotes int32
}

// 等效于
type BlogV2 struct {
	ID      int64
	Name    string
	Email   string
	Upvotes int32
}

type BlogV3 struct {
	ID      int
	Author  Author `gorm:"embedded;embeddedPrefix:author_"`
	Upvotes int32
}

// 等效于
type BlogV4 struct {
	ID          int64
	AuthorName  string
	AuthorEmail string
	Upvotes     int32
}
