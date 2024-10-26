package client

import (
	"github.com/google/uuid"
)

type SecretSaver[T any] interface {
	UpdateOrCreate(obj *T) error
}

type SecretDeleter[T any] interface {
	Delete(obj *T) error
}

type SecretGetter[T any] interface {
	// Get - получить объект по id
	Get(id uuid.UUID, obj *T) error
	// GetAll - получить список объектов
	// Результат записывается в слайс objects
	// Запрашивается не больше чем capacity слайса objects
	// Пропускается от начала offset элементов
	// Важно. Сортировка захардкожена в реализации метода
	GetAll(userID uuid.UUID, offset int, objects []*T) error
}
