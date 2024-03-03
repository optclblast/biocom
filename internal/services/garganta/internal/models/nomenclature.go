package models

import (
	"time"
)

type Item interface {
	ID() string
	ItemName() string
	Store() Storage
	ItemComposition() []CompositionUnit
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

func (s *StorageObject) Store() Storage {
	return s.Storage
}

func (s *StorageObject) ItemComposition() []CompositionUnit {
	return s.Item.ItemComposition()
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

func (s *Product) ID() string {
	return s.Id
}

func (s *Product) ItemName() string {
	return s.Name
}

func (s *Product) Store() Storage {
	return s.Storage
}

func (s *Product) ItemComposition() []CompositionUnit {
	return s.Composition
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

type NomenclatureType string

var (
	UnknownType   NomenclatureType = "unknown"
	ProductType   NomenclatureType = "product"
	AssemblyType  NomenclatureType = "assembly_unit"
	ComponentType NomenclatureType = "component"
	ServiceType   NomenclatureType = "service"
)

func NewNomenclature(
	id, companyId, name string,
	createdAt, updatedAt, deletedAt time.Time,
	nType NomenclatureType,
) (Item, error) {
	item := &StorageObject{
		BaseNomenclatureUnit: BaseNomenclatureUnit{
			Id:   id,
			Name: name,
		},
	}

	switch nType {
	case ProductType:
		item.Item = &Product{}
	case AssemblyType:
	case ComponentType:
	case ServiceType:
	default:

	}

	return nil, nil
}
