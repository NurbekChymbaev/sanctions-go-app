package models

type Entry struct {
	ID        uint   `gorm:"primary_key" json:"id"`
	Firstname string `json:"firstname"`
	Lastname  string `json:"lastname"`
}

type Names struct {
	ID        uint   `gorm:"primary_key" json:"id"`
	EntryID   uint   `json:"entryID"`
	Firstname string `json:"firstname"`
	Lastname  string `json:"lastname"`
}
