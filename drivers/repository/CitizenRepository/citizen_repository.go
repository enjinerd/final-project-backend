package CitizenRepository

import (
	"context"
	"vaccine-app-be/drivers/records"
)

type CitizenRepository interface {
	Register(ctx context.Context, citizens records.Citizen) (records.Citizen, error)
	FindByEmail(ctx context.Context, email string) (records.Citizen, error)
}