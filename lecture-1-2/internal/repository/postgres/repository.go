package repository

import (
	"context"
	"errors"

	"github.com/just-a-developer-man/GO-route256/lecture-1-2/internal/models"
	oms "github.com/just-a-developer-man/GO-route256/lecture-1-2/internal/usecase/OMS"
)

type omsRepository struct {
	/*
		PostgreSQL, MSSQL, MySQL, any you want...
	*/
}

// Check that we implemet contract for usecase
var _ oms.OMSRepository = (*omsRepository)(nil)

// NewOMSRepostiory - returns OMS repository
func NewOMSRepostiory( /* ... */ ) *omsRepository {
	return &omsRepository{
		/* ... */
	}
}

func (r *omsRepository) CreateOrder(ctx context.Context, order models.Order) error {
	/* here your SQL queries */
	return errors.New("unimplemented")
}
