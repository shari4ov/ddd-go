package customer

import (
	"errors"
	"github.com/google/uuid"
	"github.com/shari4ov/ddd-go.git/aggregate"
)

var (
	ErrCustomerNotFound    = errors.New("the customer not found in the repository")
	ErrFailedToAddCustomer = errors.New("failed to add the customer")
	ErrUpdateCustomer      = errors.New("failed to update the customer")
)

type CustomerRepository interface {
	Get(uuid.UUID) (aggregate.Customer, error)
	Add(aggregate.Customer) error
	Update(aggregate.Customer) error
}
