package model

import "gorm.io/gorm"

type Todomvc struct {
	gorm.Model
	Item   string
	Status uint
}

type TodomvcAdd struct {
	Item string
}

type TodomvcDel struct {
	Id int
}

type TodomvcUpdate struct {
	Id     int
	Item   string
	Status uint
}
type TodomvcFind struct {
	Item   string
	Status uint
}
