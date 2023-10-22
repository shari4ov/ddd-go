package valueobject

import (
	"github.com/google/uuid"
	"time"
)

// Transaction is a value object because it has no an identifier and an immutable
type Transaction struct {
	amount    int
	from      uuid.UUID
	to        uuid.UUID
	createdAt time.Time
}
