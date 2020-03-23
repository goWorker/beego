package models

type BookCategory struct {
	Id 	int
	BookId 	int
	CategoryId int
}

func (m *BookCategory)TableName()string{
	return TNBookCategory()
}