package models

type Crop struct {
	Id         int
	Name       string
	Type       string
	Period     int
	Repeatable bool
}
