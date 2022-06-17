package storage

type Class struct {
	Id          int
	Description string
	Rule        string
	Class       string
}

type Tag struct {
	Id          int
	TagName     string
	Rule        string
	Description string
}
