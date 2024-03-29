// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.25.0
// source: settings.sql

package db

import (
	"context"
)

const addSetting = `-- name: AddSetting :one
INSERT INTO settings (key, value) VALUES ($1, $2) RETURNING id, key, value, created_at, updated_at
`

type AddSettingParams struct {
	Key   string `json:"key"`
	Value string `json:"value"`
}

func (q *Queries) AddSetting(ctx context.Context, arg AddSettingParams) (Setting, error) {
	row := q.db.QueryRow(ctx, addSetting, arg.Key, arg.Value)
	var i Setting
	err := row.Scan(
		&i.ID,
		&i.Key,
		&i.Value,
		&i.CreatedAt,
		&i.UpdatedAt,
	)
	return i, err
}

const getSettingByKey = `-- name: GetSettingByKey :one
SELECT id, key, value, created_at, updated_at FROM settings WHERE key = $1 LIMIT 1
`

func (q *Queries) GetSettingByKey(ctx context.Context, key string) (Setting, error) {
	row := q.db.QueryRow(ctx, getSettingByKey, key)
	var i Setting
	err := row.Scan(
		&i.ID,
		&i.Key,
		&i.Value,
		&i.CreatedAt,
		&i.UpdatedAt,
	)
	return i, err
}

const getSettings = `-- name: GetSettings :many
SELECT id, key, value, created_at, updated_at FROM settings
`

func (q *Queries) GetSettings(ctx context.Context) ([]Setting, error) {
	rows, err := q.db.Query(ctx, getSettings)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	items := []Setting{}
	for rows.Next() {
		var i Setting
		if err := rows.Scan(
			&i.ID,
			&i.Key,
			&i.Value,
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

const updateSetting = `-- name: UpdateSetting :one
UPDATE settings SET value = $2, updated_at = NOW() WHERE key = $1 RETURNING id, key, value, created_at, updated_at
`

type UpdateSettingParams struct {
	Key   string `json:"key"`
	Value string `json:"value"`
}

func (q *Queries) UpdateSetting(ctx context.Context, arg UpdateSettingParams) (Setting, error) {
	row := q.db.QueryRow(ctx, updateSetting, arg.Key, arg.Value)
	var i Setting
	err := row.Scan(
		&i.ID,
		&i.Key,
		&i.Value,
		&i.CreatedAt,
		&i.UpdatedAt,
	)
	return i, err
}
