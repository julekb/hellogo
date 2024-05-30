package api

import (
	"github.com/google/uuid"
)

type ProductRepositoryInterface interface {
	add(obj DomainModel) DomainModel
	list() []DomainModel
	remove(id uuid.UUID)
}

type InMemoryRepository struct {
	objects *[]DomainModel
}

func (r InMemoryRepository) _getObjectsSlice() []DomainModel {
	if r.objects == nil {
		return []DomainModel{}
	} else {
		return *r.objects
	}
}

func (r *InMemoryRepository) add(obj DomainModel) DomainModel {
	obj.ID = uuid.New()
	objects := r._getObjectsSlice()
	updated_objects := append(objects, obj)
	r.objects = &updated_objects
	return obj
}

func (r InMemoryRepository) list() []DomainModel {
	return r._getObjectsSlice()
}

func (r *InMemoryRepository) remove(id uuid.UUID) {
	objects := r._getObjectsSlice()
	for i, item := range objects {
		if item.ID == id {
			updated_objects := append(objects[:i], objects[i+1:]...)
			r.objects = &updated_objects
			break
		}
	}
}
