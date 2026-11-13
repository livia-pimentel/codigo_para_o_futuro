package models

import (
	"time"

	"github.com/google/uuid"
)

// Opportunity representa a entidade central de domínio que descreve uma oportunidade.
type Opportunity struct {
	ID        uuid.UUID `json:"id"`
	Title     string    `json:"title"`
	Company   string    `json:"company"`
	Category  string    `json:"category"`
	Active    bool      `json:"active"`
	CreatedAt time.Time `json:"createdAt"`
	UpdatedAt time.Time `json:"updatedAt"`
}