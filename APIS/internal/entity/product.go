package entity

import (
	"errors"

	"github.com/ItaloG/Go-expert/APIS/pkg/entity"
)

var (
	ErrIDIsRequires = errors.New("id is required")
	ErrInvalidID = errors.
)

type Product struct {
	ID        entity.ID `json:"id`
	Name      string    `json:"name`
	Price     int       `json:"price`
	CreatedAt string    `json:"created_at`
}
