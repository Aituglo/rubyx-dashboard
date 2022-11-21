// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.15.0
// source: revenue.sql

package db

import (
	"context"
)

const createRevenue = `-- name: CreateRevenue :one
INSERT INTO revenue (program_id, vulnerability_id, money) VALUES ($1, $2, $3) RETURNING id, program_id, vulnerability_id, money, created_at, updated_at
`

type CreateRevenueParams struct {
	ProgramID       int64 `json:"program_id"`
	VulnerabilityID int64 `json:"vulnerability_id"`
	Money           int32 `json:"money"`
}

func (q *Queries) CreateRevenue(ctx context.Context, arg CreateRevenueParams) (Revenue, error) {
	row := q.db.QueryRow(ctx, createRevenue, arg.ProgramID, arg.VulnerabilityID, arg.Money)
	var i Revenue
	err := row.Scan(
		&i.ID,
		&i.ProgramID,
		&i.VulnerabilityID,
		&i.Money,
		&i.CreatedAt,
		&i.UpdatedAt,
	)
	return i, err
}

const deleteRevenueByIDs = `-- name: DeleteRevenueByIDs :exec
DELETE FROM revenue WHERE id = $1
`

func (q *Queries) DeleteRevenueByIDs(ctx context.Context, id int64) error {
	_, err := q.db.Exec(ctx, deleteRevenueByIDs, id)
	return err
}

const findRevenueByIDs = `-- name: FindRevenueByIDs :one
SELECT id, program_id, vulnerability_id, money, created_at, updated_at FROM revenue WHERE id = $1 LIMIT 1
`

func (q *Queries) FindRevenueByIDs(ctx context.Context, id int64) (Revenue, error) {
	row := q.db.QueryRow(ctx, findRevenueByIDs, id)
	var i Revenue
	err := row.Scan(
		&i.ID,
		&i.ProgramID,
		&i.VulnerabilityID,
		&i.Money,
		&i.CreatedAt,
		&i.UpdatedAt,
	)
	return i, err
}

const findRevenues = `-- name: FindRevenues :many
SELECT id, program_id, vulnerability_id, money, created_at, updated_at FROM revenue
`

func (q *Queries) FindRevenues(ctx context.Context) ([]Revenue, error) {
	rows, err := q.db.Query(ctx, findRevenues)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	items := []Revenue{}
	for rows.Next() {
		var i Revenue
		if err := rows.Scan(
			&i.ID,
			&i.ProgramID,
			&i.VulnerabilityID,
			&i.Money,
			&i.CreatedAt,
			&i.UpdatedAt,
		); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}

const updateRevenue = `-- name: UpdateRevenue :one
UPDATE revenue SET program_id = $2, vulnerability_id = $3, money = $4, service = $4, updated_at = NOW() WHERE id = $1 RETURNING id, program_id, vulnerability_id, money, created_at, updated_at
`

type UpdateRevenueParams struct {
	ID              int64 `json:"id"`
	ProgramID       int64 `json:"program_id"`
	VulnerabilityID int64 `json:"vulnerability_id"`
	Money           int32 `json:"money"`
}

func (q *Queries) UpdateRevenue(ctx context.Context, arg UpdateRevenueParams) (Revenue, error) {
	row := q.db.QueryRow(ctx, updateRevenue,
		arg.ID,
		arg.ProgramID,
		arg.VulnerabilityID,
		arg.Money,
	)
	var i Revenue
	err := row.Scan(
		&i.ID,
		&i.ProgramID,
		&i.VulnerabilityID,
		&i.Money,
		&i.CreatedAt,
		&i.UpdatedAt,
	)
	return i, err
}
