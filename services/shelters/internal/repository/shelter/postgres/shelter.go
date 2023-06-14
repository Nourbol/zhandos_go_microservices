package postgres

import (
	"context"
	"database/sql"
	"fmt"
	"github.com/google/uuid"
	_filters "github.com/nourbol/zhandos/pkg/tools/filters"
	sqlErrors "github.com/nourbol/zhandos/pkg/type/errors"
	_shelter "github.com/nourbol/zhandos/services/shelters/internal/domain/shelter"
	"github.com/nourbol/zhandos/services/shelters/internal/repository/shelter/postgres/dao"
	"github.com/pkg/errors"
)

func (r *Repository) Insert(ctx context.Context, shelter _shelter.Shelter) (*_shelter.Shelter, error) {
	query := `INSERT INTO movies (name, ownedBreeds)
			  VALUES ($1, $2)
			  RETURNING id, created_at`

	args := []any{shelter.Name(), shelter.OwnedBreeds()}

	shelterDao := dao.Shelter{
		Name:        shelter.Name().String(),
		OwnedBreeds: shelter.OwnedBreeds(),
	}

	if err := r.db.QueryRow(ctx, query, args...).Scan(&shelterDao.ID, &shelterDao.CreatedAt); err != nil {
		return nil, err
	}

	insertedShelter, err := shelterDao.ToDomainShelter()
	if err != nil {
		return nil, err
	}

	return insertedShelter, nil
}

func (r *Repository) Get(ctx context.Context, id uuid.UUID) (*_shelter.Shelter, error) {
	query := `SELECT id, created_at, name, ownedBreeds
			  FROM movies
			  WHERE id = $1`

	var shelterDao dao.Shelter

	if err := r.db.QueryRow(ctx, query, id).
		Scan(&shelterDao.ID,
			&shelterDao.CreatedAt,
			&shelterDao.Name,
			&shelterDao.OwnedBreeds); err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, sqlErrors.ErrRecordNotFound
		}
		return nil, err
	}

	s, err := shelterDao.ToDomainShelter()
	if err != nil {
		return nil, err
	}

	return s, nil
}

func (r *Repository) Update(ctx context.Context, shelter _shelter.Shelter) error {
	query := `UPDATE shelters
			  SET name = $1, ownedBreeds = $2
			  WHERE id = $3`

	args := []any{
		shelter.Name(),
		shelter.OwnedBreeds(),
	}

	if _, err := r.db.Exec(ctx, query, args...); err != nil {
		return err
	}
	return nil
}

func (r *Repository) Delete(ctx context.Context, id uuid.UUID) error {
	query := `DELETE FROM shelters
			  WHERE id = $1`

	result, err := r.db.Exec(ctx, query, id)
	if err != nil {
		return err
	}

	rowsAffected := result.RowsAffected()
	if rowsAffected == 0 {
		return sqlErrors.ErrRecordNotFound
	}

	return nil
}

func (r *Repository) GetAll(
	ctx context.Context,
	name string,
	ownedBreeds []uuid.UUID,
	filters _filters.Filters) ([]*_shelter.Shelter, _filters.Metadata, error) {
	query := fmt.Sprintf(`
			  SELECT COUNT(*) OVER (), id, created_at, name, ownedBreeds
			  FROM shelters 
			  WHERE (to_tsvector('simple', name) @@ plainto_tsquery('simple', $1) OR $1 = '')
			  AND (ownedBreeds @> $2 OR $2 = '{}')
			  ORDER BY %s %s, id ASC
			  LIMIT $3 OFFSET $4`, filters.SortColumn(), filters.SortDirection())

	args := []any{name, ownedBreeds, filters.Limit(), filters.Offset()}

	rows, err := r.db.Query(ctx, query, args...)
	if err != nil {
		return nil, _filters.Metadata{}, err
	}
	defer rows.Close()

	totalRecords := 0
	var shelters []*_shelter.Shelter

	for rows.Next() {
		var shelterDao dao.Shelter

		if err := rows.Scan(
			&totalRecords,
			&shelterDao.ID,
			&shelterDao.CreatedAt,
			&shelterDao.Name,
			&shelterDao.OwnedBreeds); err != nil {
			return nil, _filters.Metadata{}, err
		}

		shelter, err := shelterDao.ToDomainShelter()
		if err != nil {
			return nil, _filters.Metadata{}, err
		}

		shelters = append(shelters, shelter)
	}

	if err = rows.Err(); err != nil {
		return nil, _filters.Metadata{}, err
	}

	metadata := _filters.CalculateMetadata(totalRecords, filters.Page, filters.PageSize)
	return shelters, metadata, nil
}
