package models

import (
	"time"
)

type Item interface {
	ID() string
	ItemName() string
	RemainingStock() float32
	Store() Storage
	Composition() []CompositionUnit
}

type CompositionUnit interface {
	Acount() float32
	Unit() Item
}

type StorageObject struct {
	BaseNomenclatureUnit
	Amount float32

	Item
}

func (s *StorageObject) ID() string {
	return s.Id
}

func (s *StorageObject) ItemName() string {
	return s.Name
}

func (s *StorageObject) RemainingStock() float32 {
	return s.Amount
}

func (s *StorageObject) Store() Storage {
	return s.Storage
}

type BaseNomenclatureUnit struct {
	Id      string
	Name    string
	Storage Storage
}

type Product struct {
	BaseNomenclatureUnit
	Composition []CompositionUnit
	CreatedAt   time.Time
	UpdatedAt   time.Time
	DeletedAt   time.Time
	Barcode     string
}

type StoreObjectCompositionUnit struct {
	Amount float32
	Item   Item
}

type AssemblyUnit struct {
	BaseNomenclatureUnit
	Parts []CompositionUnit
}

type Component struct {
	BaseNomenclatureUnit
}

type Service struct {
	BaseNomenclatureUnit
}
