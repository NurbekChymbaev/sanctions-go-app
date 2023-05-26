package models

type Entry struct {
	ID        uint   `gorm:"primary_key" json:"id"`
	Firstname string `json:"firstname"`
	Lastname  string `json:"lastname"`
	Title     string `json:"title"`
	Remarks   string `json:"remarks"`
}

type Names struct {
	ID        uint   `gorm:"primary_key" json:"id"`
	EntryID   uint   `json:"entryID"`
	Category  string `json:"category"`
	Firstname string `json:"firstname"`
	Lastname  string `json:"lastname"`
}
