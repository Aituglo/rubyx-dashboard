// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.15.0
// source: stats.sql

package db

import (
	"context"
	"time"
)

const createStat = `-- name: CreateStat :one
INSERT INTO stats (report_id, report_title, severity, reward, currency, collab, report_status, report_date, platform_id, created_at, updated_at)
VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9, NOW(), NOW())
RETURNING id, report_id, report_title, severity, reward, currency, collab, report_status, report_date, platform_id, created_at, updated_at
`

type CreateStatParams struct {
	ReportID     string    `json:"report_id"`
	ReportTitle  string    `json:"report_title"`
	Severity     string    `json:"severity"`
	Reward       float32   `json:"reward"`
	Currency     string    `json:"currency"`
	Collab       bool      `json:"collab"`
	ReportStatus string    `json:"report_status"`
	ReportDate   time.Time `json:"report_date"`
	PlatformID   int64     `json:"platform_id"`
}

func (q *Queries) CreateStat(ctx context.Context, arg CreateStatParams) (Stat, error) {
	row := q.db.QueryRow(ctx, createStat,
		arg.ReportID,
		arg.ReportTitle,
		arg.Severity,
		arg.Reward,
		arg.Currency,
		arg.Collab,
		arg.ReportStatus,
		arg.ReportDate,
		arg.PlatformID,
	)
	var i Stat
	err := row.Scan(
		&i.ID,
		&i.ReportID,
		&i.ReportTitle,
		&i.Severity,
		&i.Reward,
		&i.Currency,
		&i.Collab,
		&i.ReportStatus,
		&i.ReportDate,
		&i.PlatformID,
		&i.CreatedAt,
		&i.UpdatedAt,
	)
	return i, err
}

const deleteStatByID = `-- name: DeleteStatByID :exec
DELETE FROM stats WHERE id = $1
`

func (q *Queries) DeleteStatByID(ctx context.Context, id int64) error {
	_, err := q.db.Exec(ctx, deleteStatByID, id)
	return err
}

const findStatByID = `-- name: FindStatByID :one
SELECT id, report_id, report_title, severity, reward, currency, collab, report_status, report_date, platform_id, created_at, updated_at FROM stats WHERE id = $1 LIMIT 1
`

func (q *Queries) FindStatByID(ctx context.Context, id int64) (Stat, error) {
	row := q.db.QueryRow(ctx, findStatByID, id)
	var i Stat
	err := row.Scan(
		&i.ID,
		&i.ReportID,
		&i.ReportTitle,
		&i.Severity,
		&i.Reward,
		&i.Currency,
		&i.Collab,
		&i.ReportStatus,
		&i.ReportDate,
		&i.PlatformID,
		&i.CreatedAt,
		&i.UpdatedAt,
	)
	return i, err
}

const findStatByReportID = `-- name: FindStatByReportID :one
SELECT id, report_id, report_title, severity, reward, currency, collab, report_status, report_date, platform_id, created_at, updated_at FROM stats WHERE report_id = $1 LIMIT 1
`

func (q *Queries) FindStatByReportID(ctx context.Context, reportID string) (Stat, error) {
	row := q.db.QueryRow(ctx, findStatByReportID, reportID)
	var i Stat
	err := row.Scan(
		&i.ID,
		&i.ReportID,
		&i.ReportTitle,
		&i.Severity,
		&i.Reward,
		&i.Currency,
		&i.Collab,
		&i.ReportStatus,
		&i.ReportDate,
		&i.PlatformID,
		&i.CreatedAt,
		&i.UpdatedAt,
	)
	return i, err
}

const findStats = `-- name: FindStats :many
SELECT id, report_id, report_title, severity, reward, currency, collab, report_status, report_date, platform_id, created_at, updated_at FROM stats
`

func (q *Queries) FindStats(ctx context.Context) ([]Stat, error) {
	rows, err := q.db.Query(ctx, findStats)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	items := []Stat{}
	for rows.Next() {
		var i Stat
		if err := rows.Scan(
			&i.ID,
			&i.ReportID,
			&i.ReportTitle,
			&i.Severity,
			&i.Reward,
			&i.Currency,
			&i.Collab,
			&i.ReportStatus,
			&i.ReportDate,
			&i.PlatformID,
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

const updateStat = `-- name: UpdateStat :one
UPDATE stats
SET report_id = $2, report_title = $3, severity = $4, reward = $5, currency = $6, collab = $7, report_status = $8, report_date = $9, platform_id = $10, updated_at = NOW()
WHERE id = $1
RETURNING id, report_id, report_title, severity, reward, currency, collab, report_status, report_date, platform_id, created_at, updated_at
`

type UpdateStatParams struct {
	ID           int64     `json:"id"`
	ReportID     string    `json:"report_id"`
	ReportTitle  string    `json:"report_title"`
	Severity     string    `json:"severity"`
	Reward       float32   `json:"reward"`
	Currency     string    `json:"currency"`
	Collab       bool      `json:"collab"`
	ReportStatus string    `json:"report_status"`
	ReportDate   time.Time `json:"report_date"`
	PlatformID   int64     `json:"platform_id"`
}

func (q *Queries) UpdateStat(ctx context.Context, arg UpdateStatParams) (Stat, error) {
	row := q.db.QueryRow(ctx, updateStat,
		arg.ID,
		arg.ReportID,
		arg.ReportTitle,
		arg.Severity,
		arg.Reward,
		arg.Currency,
		arg.Collab,
		arg.ReportStatus,
		arg.ReportDate,
		arg.PlatformID,
	)
	var i Stat
	err := row.Scan(
		&i.ID,
		&i.ReportID,
		&i.ReportTitle,
		&i.Severity,
		&i.Reward,
		&i.Currency,
		&i.Collab,
		&i.ReportStatus,
		&i.ReportDate,
		&i.PlatformID,
		&i.CreatedAt,
		&i.UpdatedAt,
	)
	return i, err
}