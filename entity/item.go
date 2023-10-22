package entity

import "github.com/google/uuid"

// Item is an entity that represents an item in all domains
type Item struct {
	// ID an identifier of the entity
	ID          uuid.UUID
	Name        string
	Description string
}
