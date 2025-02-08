package storage

import (
	"api/pkg/container"
)

type Storage interface {
	AddNewContainerRecord(address string) error
	UpdateContainerRecord(container *container.Container) error
	ShowContainerRecords() ([]*container.Container, error)
}
