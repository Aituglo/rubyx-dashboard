// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.25.0
// source: scope.sql

package db

import (
	"context"
)

const createScope = `-- name: CreateScope :one
INSERT INTO scopes (program_id, scope, scope_type) VALUES ($1, $2, $3) RETURNING id, program_id, scope, scope_type, created_at, updated_at
`

type CreateScopeParams struct {
	ProgramID int64  `json:"program_id"`
	Scope     string `json:"scope"`
	ScopeType string `json:"scope_type"`
}

func (q *Queries) CreateScope(ctx context.Context, arg CreateScopeParams) (Scope, error) {
	row := q.db.QueryRow(ctx, createScope, arg.ProgramID, arg.Scope, arg.ScopeType)
	var i Scope
	err := row.Scan(
		&i.ID,
		&i.ProgramID,
		&i.Scope,
		&i.ScopeType,
		&i.CreatedAt,
		&i.UpdatedAt,
	)
	return i, err
}

const deleteScopeByID = `-- name: DeleteScopeByID :exec
DELETE FROM scopes WHERE id = $1
`

func (q *Queries) DeleteScopeByID(ctx context.Context, id int64) error {
	_, err := q.db.Exec(ctx, deleteScopeByID, id)
	return err
}

const findProgramByScope = `-- name: FindProgramByScope :one
SELECT program_id FROM scopes WHERE scope LIKE $1 LIMIT 1
`

func (q *Queries) FindProgramByScope(ctx context.Context, scope string) (int64, error) {
	row := q.db.QueryRow(ctx, findProgramByScope, scope)
	var program_id int64
	err := row.Scan(&program_id)
	return program_id, err
}

const findScopeByID = `-- name: FindScopeByID :one
SELECT id, program_id, scope, scope_type, created_at, updated_at FROM scopes WHERE id = $1 LIMIT 1
`

func (q *Queries) FindScopeByID(ctx context.Context, id int64) (Scope, error) {
	row := q.db.QueryRow(ctx, findScopeByID, id)
	var i Scope
	err := row.Scan(
		&i.ID,
		&i.ProgramID,
		&i.Scope,
		&i.ScopeType,
		&i.CreatedAt,
		&i.UpdatedAt,
	)
	return i, err
}

const findScopes = `-- name: FindScopes :many
SELECT id, program_id, scope, scope_type, created_at, updated_at FROM scopes
`

func (q *Queries) FindScopes(ctx context.Context) ([]Scope, error) {
	rows, err := q.db.Query(ctx, findScopes)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	items := []Scope{}
	for rows.Next() {
		var i Scope
		if err := rows.Scan(
			&i.ID,
			&i.ProgramID,
			&i.Scope,
			&i.ScopeType,
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

const findScopesByProgramID = `-- name: FindScopesByProgramID :many
SELECT id, program_id, scope, scope_type, created_at, updated_at FROM scopes WHERE program_id = $1
`

func (q *Queries) FindScopesByProgramID(ctx context.Context, programID int64) ([]Scope, error) {
	rows, err := q.db.Query(ctx, findScopesByProgramID, programID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	items := []Scope{}
	for rows.Next() {
		var i Scope
		if err := rows.Scan(
			&i.ID,
			&i.ProgramID,
			&i.Scope,
			&i.ScopeType,
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

const getScopeByProgramIDAndScope = `-- name: GetScopeByProgramIDAndScope :one
SELECT id, program_id, scope, scope_type, created_at, updated_at FROM scopes WHERE program_id = $1 AND scope = $2 LIMIT 1
`

type GetScopeByProgramIDAndScopeParams struct {
	ProgramID int64  `json:"program_id"`
	Scope     string `json:"scope"`
}

func (q *Queries) GetScopeByProgramIDAndScope(ctx context.Context, arg GetScopeByProgramIDAndScopeParams) (Scope, error) {
	row := q.db.QueryRow(ctx, getScopeByProgramIDAndScope, arg.ProgramID, arg.Scope)
	var i Scope
	err := row.Scan(
		&i.ID,
		&i.ProgramID,
		&i.Scope,
		&i.ScopeType,
		&i.CreatedAt,
		&i.UpdatedAt,
	)
	return i, err
}
