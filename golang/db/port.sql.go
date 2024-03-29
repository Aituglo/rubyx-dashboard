// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.25.0
// source: port.sql

package db

import (
	"context"
)

const countPortsByIp = `-- name: CountPortsByIp :one
SELECT COUNT(*) FROM port WHERE ip_id = $1
`

func (q *Queries) CountPortsByIp(ctx context.Context, ipID int64) (int64, error) {
	row := q.db.QueryRow(ctx, countPortsByIp, ipID)
	var count int64
	err := row.Scan(&count)
	return count, err
}

const createPort = `-- name: CreatePort :one
INSERT INTO port (ip_id, port) VALUES ($1, $2) RETURNING id, ip_id, port, created_at, updated_at
`

type CreatePortParams struct {
	IpID int64 `json:"ip_id"`
	Port int32 `json:"port"`
}

func (q *Queries) CreatePort(ctx context.Context, arg CreatePortParams) (Port, error) {
	row := q.db.QueryRow(ctx, createPort, arg.IpID, arg.Port)
	var i Port
	err := row.Scan(
		&i.ID,
		&i.IpID,
		&i.Port,
		&i.CreatedAt,
		&i.UpdatedAt,
	)
	return i, err
}

const deletePortByIDs = `-- name: DeletePortByIDs :exec
DELETE FROM port WHERE id = $1
`

func (q *Queries) DeletePortByIDs(ctx context.Context, id int64) error {
	_, err := q.db.Exec(ctx, deletePortByIDs, id)
	return err
}

const findPortByIDs = `-- name: FindPortByIDs :one
SELECT id, ip_id, port, created_at, updated_at FROM port WHERE id = $1 LIMIT 1
`

func (q *Queries) FindPortByIDs(ctx context.Context, id int64) (Port, error) {
	row := q.db.QueryRow(ctx, findPortByIDs, id)
	var i Port
	err := row.Scan(
		&i.ID,
		&i.IpID,
		&i.Port,
		&i.CreatedAt,
		&i.UpdatedAt,
	)
	return i, err
}

const findPortsWithIpID = `-- name: FindPortsWithIpID :many
SELECT id, ip_id, port, created_at, updated_at FROM port WHERE ip_id = $1
`

func (q *Queries) FindPortsWithIpID(ctx context.Context, ipID int64) ([]Port, error) {
	rows, err := q.db.Query(ctx, findPortsWithIpID, ipID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	items := []Port{}
	for rows.Next() {
		var i Port
		if err := rows.Scan(
			&i.ID,
			&i.IpID,
			&i.Port,
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

const findPortsWithIpIDAndPort = `-- name: FindPortsWithIpIDAndPort :one
SELECT id, ip_id, port, created_at, updated_at FROM port WHERE ip_id = $1 AND port = $2 LIMIT 1
`

type FindPortsWithIpIDAndPortParams struct {
	IpID int64 `json:"ip_id"`
	Port int32 `json:"port"`
}

func (q *Queries) FindPortsWithIpIDAndPort(ctx context.Context, arg FindPortsWithIpIDAndPortParams) (Port, error) {
	row := q.db.QueryRow(ctx, findPortsWithIpIDAndPort, arg.IpID, arg.Port)
	var i Port
	err := row.Scan(
		&i.ID,
		&i.IpID,
		&i.Port,
		&i.CreatedAt,
		&i.UpdatedAt,
	)
	return i, err
}

const updatePort = `-- name: UpdatePort :one
UPDATE port SET ip_id = $2, port = $3, updated_at = NOW() WHERE id = $1 RETURNING id, ip_id, port, created_at, updated_at
`

type UpdatePortParams struct {
	ID   int64 `json:"id"`
	IpID int64 `json:"ip_id"`
	Port int32 `json:"port"`
}

func (q *Queries) UpdatePort(ctx context.Context, arg UpdatePortParams) (Port, error) {
	row := q.db.QueryRow(ctx, updatePort, arg.ID, arg.IpID, arg.Port)
	var i Port
	err := row.Scan(
		&i.ID,
		&i.IpID,
		&i.Port,
		&i.CreatedAt,
		&i.UpdatedAt,
	)
	return i, err
}
