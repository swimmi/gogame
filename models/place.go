package models

type Place struct {
	Id    int
	Name  string
	Type  string
	Sites []Site
}
