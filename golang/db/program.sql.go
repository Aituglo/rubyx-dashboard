// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.15.0
// source: program.sql

package db

import (
	"context"
)

const countPrograms = `-- name: CountPrograms :one
SELECT COUNT(*) FROM program
`

func (q *Queries) CountPrograms(ctx context.Context) (int64, error) {
	row := q.db.QueryRow(ctx, countPrograms)
	var count int64
	err := row.Scan(&count)
	return count, err
}

const countProgramsWithPlatform = `-- name: CountProgramsWithPlatform :one
SELECT COUNT(*) FROM program WHERE platform_id = $1
`

func (q *Queries) CountProgramsWithPlatform(ctx context.Context, platformID int64) (int64, error) {
	row := q.db.QueryRow(ctx, countProgramsWithPlatform, platformID)
	var count int64
	err := row.Scan(&count)
	return count, err
}

const countProgramsWithSearch = `-- name: CountProgramsWithSearch :one
SELECT COUNT(*) FROM program WHERE name LIKE $1
`

func (q *Queries) CountProgramsWithSearch(ctx context.Context, name string) (int64, error) {
	row := q.db.QueryRow(ctx, countProgramsWithSearch, name)
	var count int64
	err := row.Scan(&count)
	return count, err
}

const countProgramsWithSearchAndPlatform = `-- name: CountProgramsWithSearchAndPlatform :one
SELECT COUNT(*) FROM program WHERE name LIKE $1 AND platform_id = $2
`

type CountProgramsWithSearchAndPlatformParams struct {
	Name       string `json:"name"`
	PlatformID int64  `json:"platform_id"`
}

func (q *Queries) CountProgramsWithSearchAndPlatform(ctx context.Context, arg CountProgramsWithSearchAndPlatformParams) (int64, error) {
	row := q.db.QueryRow(ctx, countProgramsWithSearchAndPlatform, arg.Name, arg.PlatformID)
	var count int64
	err := row.Scan(&count)
	return count, err
}

const countProgramsWithSearchAndType = `-- name: CountProgramsWithSearchAndType :one
SELECT COUNT(*) FROM program WHERE name LIKE $1 AND type = $2::program_type
`

type CountProgramsWithSearchAndTypeParams struct {
	Name    string      `json:"name"`
	Column2 ProgramType `json:"column_2"`
}

func (q *Queries) CountProgramsWithSearchAndType(ctx context.Context, arg CountProgramsWithSearchAndTypeParams) (int64, error) {
	row := q.db.QueryRow(ctx, countProgramsWithSearchAndType, arg.Name, arg.Column2)
	var count int64
	err := row.Scan(&count)
	return count, err
}

const countProgramsWithSearchAndTypeAndPlatform = `-- name: CountProgramsWithSearchAndTypeAndPlatform :one
SELECT COUNT(*) FROM program WHERE name LIKE $1 AND type = $2::program_type AND platform_id = $3
`

type CountProgramsWithSearchAndTypeAndPlatformParams struct {
	Name       string      `json:"name"`
	Column2    ProgramType `json:"column_2"`
	PlatformID int64       `json:"platform_id"`
}

func (q *Queries) CountProgramsWithSearchAndTypeAndPlatform(ctx context.Context, arg CountProgramsWithSearchAndTypeAndPlatformParams) (int64, error) {
	row := q.db.QueryRow(ctx, countProgramsWithSearchAndTypeAndPlatform, arg.Name, arg.Column2, arg.PlatformID)
	var count int64
	err := row.Scan(&count)
	return count, err
}

const countProgramsWithType = `-- name: CountProgramsWithType :one
SELECT COUNT(*) FROM program WHERE type = $1::program_type
`

func (q *Queries) CountProgramsWithType(ctx context.Context, dollar_1 ProgramType) (int64, error) {
	row := q.db.QueryRow(ctx, countProgramsWithType, dollar_1)
	var count int64
	err := row.Scan(&count)
	return count, err
}

const countProgramsWithTypeAndPlatform = `-- name: CountProgramsWithTypeAndPlatform :one
SELECT COUNT(*) FROM program WHERE type = $1::program_type AND platform_id = $2
`

type CountProgramsWithTypeAndPlatformParams struct {
	Column1    ProgramType `json:"column_1"`
	PlatformID int64       `json:"platform_id"`
}

func (q *Queries) CountProgramsWithTypeAndPlatform(ctx context.Context, arg CountProgramsWithTypeAndPlatformParams) (int64, error) {
	row := q.db.QueryRow(ctx, countProgramsWithTypeAndPlatform, arg.Column1, arg.PlatformID)
	var count int64
	err := row.Scan(&count)
	return count, err
}

const createProgram = `-- name: CreateProgram :one
INSERT INTO program (platform_id, name, slug, vdp, tag, url, type) VALUES ($1, $2, $3, $4, $5, $6, $7) RETURNING id, platform_id, name, slug, vdp, favourite, tag, url, type, created_at, updated_at
`

type CreateProgramParams struct {
	PlatformID int64       `json:"platform_id"`
	Name       string      `json:"name"`
	Slug       string      `json:"slug"`
	Vdp        bool        `json:"vdp"`
	Tag        string      `json:"tag"`
	Url        string      `json:"url"`
	Type       ProgramType `json:"type"`
}

func (q *Queries) CreateProgram(ctx context.Context, arg CreateProgramParams) (Program, error) {
	row := q.db.QueryRow(ctx, createProgram,
		arg.PlatformID,
		arg.Name,
		arg.Slug,
		arg.Vdp,
		arg.Tag,
		arg.Url,
		arg.Type,
	)
	var i Program
	err := row.Scan(
		&i.ID,
		&i.PlatformID,
		&i.Name,
		&i.Slug,
		&i.Vdp,
		&i.Favourite,
		&i.Tag,
		&i.Url,
		&i.Type,
		&i.CreatedAt,
		&i.UpdatedAt,
	)
	return i, err
}

const deleteProgramByIDs = `-- name: DeleteProgramByIDs :exec
DELETE FROM program WHERE id = $1
`

func (q *Queries) DeleteProgramByIDs(ctx context.Context, id int64) error {
	_, err := q.db.Exec(ctx, deleteProgramByIDs, id)
	return err
}

const favouriteProgram = `-- name: FavouriteProgram :one
UPDATE program SET favourite = $2, updated_at = NOW() WHERE id = $1 RETURNING id, platform_id, name, slug, vdp, favourite, tag, url, type, created_at, updated_at
`

type FavouriteProgramParams struct {
	ID        int64 `json:"id"`
	Favourite bool  `json:"favourite"`
}

func (q *Queries) FavouriteProgram(ctx context.Context, arg FavouriteProgramParams) (Program, error) {
	row := q.db.QueryRow(ctx, favouriteProgram, arg.ID, arg.Favourite)
	var i Program
	err := row.Scan(
		&i.ID,
		&i.PlatformID,
		&i.Name,
		&i.Slug,
		&i.Vdp,
		&i.Favourite,
		&i.Tag,
		&i.Url,
		&i.Type,
		&i.CreatedAt,
		&i.UpdatedAt,
	)
	return i, err
}

const findAllPrograms = `-- name: FindAllPrograms :many
SELECT id, platform_id, name, slug, vdp, favourite, tag, url, type, created_at, updated_at FROM program
`

func (q *Queries) FindAllPrograms(ctx context.Context) ([]Program, error) {
	rows, err := q.db.Query(ctx, findAllPrograms)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	items := []Program{}
	for rows.Next() {
		var i Program
		if err := rows.Scan(
			&i.ID,
			&i.PlatformID,
			&i.Name,
			&i.Slug,
			&i.Vdp,
			&i.Favourite,
			&i.Tag,
			&i.Url,
			&i.Type,
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

const findProgramByIDs = `-- name: FindProgramByIDs :one
SELECT id, platform_id, name, slug, vdp, favourite, tag, url, type, created_at, updated_at FROM program WHERE id = $1 LIMIT 1
`

func (q *Queries) FindProgramByIDs(ctx context.Context, id int64) (Program, error) {
	row := q.db.QueryRow(ctx, findProgramByIDs, id)
	var i Program
	err := row.Scan(
		&i.ID,
		&i.PlatformID,
		&i.Name,
		&i.Slug,
		&i.Vdp,
		&i.Favourite,
		&i.Tag,
		&i.Url,
		&i.Type,
		&i.CreatedAt,
		&i.UpdatedAt,
	)
	return i, err
}

const findProgramBySlug = `-- name: FindProgramBySlug :one
SELECT id, platform_id, name, slug, vdp, favourite, tag, url, type, created_at, updated_at FROM program WHERE slug = $1 LIMIT 1
`

func (q *Queries) FindProgramBySlug(ctx context.Context, slug string) (Program, error) {
	row := q.db.QueryRow(ctx, findProgramBySlug, slug)
	var i Program
	err := row.Scan(
		&i.ID,
		&i.PlatformID,
		&i.Name,
		&i.Slug,
		&i.Vdp,
		&i.Favourite,
		&i.Tag,
		&i.Url,
		&i.Type,
		&i.CreatedAt,
		&i.UpdatedAt,
	)
	return i, err
}

const findPrograms = `-- name: FindPrograms :many
SELECT id, platform_id, name, slug, vdp, favourite, tag, url, type, created_at, updated_at FROM program LIMIT $1 OFFSET $2
`

type FindProgramsParams struct {
	Limit  int32 `json:"limit"`
	Offset int32 `json:"offset"`
}

func (q *Queries) FindPrograms(ctx context.Context, arg FindProgramsParams) ([]Program, error) {
	rows, err := q.db.Query(ctx, findPrograms, arg.Limit, arg.Offset)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	items := []Program{}
	for rows.Next() {
		var i Program
		if err := rows.Scan(
			&i.ID,
			&i.PlatformID,
			&i.Name,
			&i.Slug,
			&i.Vdp,
			&i.Favourite,
			&i.Tag,
			&i.Url,
			&i.Type,
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

const findProgramsWithPlatform = `-- name: FindProgramsWithPlatform :many
SELECT id, platform_id, name, slug, vdp, favourite, tag, url, type, created_at, updated_at FROM program WHERE platform_id = $1 LIMIT $2 OFFSET $3
`

type FindProgramsWithPlatformParams struct {
	PlatformID int64 `json:"platform_id"`
	Limit      int32 `json:"limit"`
	Offset     int32 `json:"offset"`
}

func (q *Queries) FindProgramsWithPlatform(ctx context.Context, arg FindProgramsWithPlatformParams) ([]Program, error) {
	rows, err := q.db.Query(ctx, findProgramsWithPlatform, arg.PlatformID, arg.Limit, arg.Offset)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	items := []Program{}
	for rows.Next() {
		var i Program
		if err := rows.Scan(
			&i.ID,
			&i.PlatformID,
			&i.Name,
			&i.Slug,
			&i.Vdp,
			&i.Favourite,
			&i.Tag,
			&i.Url,
			&i.Type,
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

const findProgramsWithSearch = `-- name: FindProgramsWithSearch :many
SELECT id, platform_id, name, slug, vdp, favourite, tag, url, type, created_at, updated_at FROM program WHERE name LIKE $1 LIMIT $2 OFFSET $3
`

type FindProgramsWithSearchParams struct {
	Name   string `json:"name"`
	Limit  int32  `json:"limit"`
	Offset int32  `json:"offset"`
}

func (q *Queries) FindProgramsWithSearch(ctx context.Context, arg FindProgramsWithSearchParams) ([]Program, error) {
	rows, err := q.db.Query(ctx, findProgramsWithSearch, arg.Name, arg.Limit, arg.Offset)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	items := []Program{}
	for rows.Next() {
		var i Program
		if err := rows.Scan(
			&i.ID,
			&i.PlatformID,
			&i.Name,
			&i.Slug,
			&i.Vdp,
			&i.Favourite,
			&i.Tag,
			&i.Url,
			&i.Type,
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

const findProgramsWithSearchAndPlatform = `-- name: FindProgramsWithSearchAndPlatform :many
SELECT id, platform_id, name, slug, vdp, favourite, tag, url, type, created_at, updated_at FROM program WHERE name LIKE $1 AND platform_id = $2 LIMIT $3 OFFSET $4
`

type FindProgramsWithSearchAndPlatformParams struct {
	Name       string `json:"name"`
	PlatformID int64  `json:"platform_id"`
	Limit      int32  `json:"limit"`
	Offset     int32  `json:"offset"`
}

func (q *Queries) FindProgramsWithSearchAndPlatform(ctx context.Context, arg FindProgramsWithSearchAndPlatformParams) ([]Program, error) {
	rows, err := q.db.Query(ctx, findProgramsWithSearchAndPlatform,
		arg.Name,
		arg.PlatformID,
		arg.Limit,
		arg.Offset,
	)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	items := []Program{}
	for rows.Next() {
		var i Program
		if err := rows.Scan(
			&i.ID,
			&i.PlatformID,
			&i.Name,
			&i.Slug,
			&i.Vdp,
			&i.Favourite,
			&i.Tag,
			&i.Url,
			&i.Type,
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

const findProgramsWithSearchAndType = `-- name: FindProgramsWithSearchAndType :many
SELECT id, platform_id, name, slug, vdp, favourite, tag, url, type, created_at, updated_at FROM program WHERE name LIKE $1 AND type = $2::program_type LIMIT $3 OFFSET $4
`

type FindProgramsWithSearchAndTypeParams struct {
	Name    string      `json:"name"`
	Column2 ProgramType `json:"column_2"`
	Limit   int32       `json:"limit"`
	Offset  int32       `json:"offset"`
}

func (q *Queries) FindProgramsWithSearchAndType(ctx context.Context, arg FindProgramsWithSearchAndTypeParams) ([]Program, error) {
	rows, err := q.db.Query(ctx, findProgramsWithSearchAndType,
		arg.Name,
		arg.Column2,
		arg.Limit,
		arg.Offset,
	)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	items := []Program{}
	for rows.Next() {
		var i Program
		if err := rows.Scan(
			&i.ID,
			&i.PlatformID,
			&i.Name,
			&i.Slug,
			&i.Vdp,
			&i.Favourite,
			&i.Tag,
			&i.Url,
			&i.Type,
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

const findProgramsWithSearchAndTypeAndPlatform = `-- name: FindProgramsWithSearchAndTypeAndPlatform :many
SELECT id, platform_id, name, slug, vdp, favourite, tag, url, type, created_at, updated_at FROM program WHERE name LIKE $1 AND type = $2::program_type AND platform_id = $3 LIMIT $4 OFFSET $5
`

type FindProgramsWithSearchAndTypeAndPlatformParams struct {
	Name       string      `json:"name"`
	Column2    ProgramType `json:"column_2"`
	PlatformID int64       `json:"platform_id"`
	Limit      int32       `json:"limit"`
	Offset     int32       `json:"offset"`
}

func (q *Queries) FindProgramsWithSearchAndTypeAndPlatform(ctx context.Context, arg FindProgramsWithSearchAndTypeAndPlatformParams) ([]Program, error) {
	rows, err := q.db.Query(ctx, findProgramsWithSearchAndTypeAndPlatform,
		arg.Name,
		arg.Column2,
		arg.PlatformID,
		arg.Limit,
		arg.Offset,
	)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	items := []Program{}
	for rows.Next() {
		var i Program
		if err := rows.Scan(
			&i.ID,
			&i.PlatformID,
			&i.Name,
			&i.Slug,
			&i.Vdp,
			&i.Favourite,
			&i.Tag,
			&i.Url,
			&i.Type,
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

const findProgramsWithType = `-- name: FindProgramsWithType :many
SELECT id, platform_id, name, slug, vdp, favourite, tag, url, type, created_at, updated_at FROM program WHERE type = $1::program_type LIMIT $2 OFFSET $3
`

type FindProgramsWithTypeParams struct {
	Column1 ProgramType `json:"column_1"`
	Limit   int32       `json:"limit"`
	Offset  int32       `json:"offset"`
}

func (q *Queries) FindProgramsWithType(ctx context.Context, arg FindProgramsWithTypeParams) ([]Program, error) {
	rows, err := q.db.Query(ctx, findProgramsWithType, arg.Column1, arg.Limit, arg.Offset)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	items := []Program{}
	for rows.Next() {
		var i Program
		if err := rows.Scan(
			&i.ID,
			&i.PlatformID,
			&i.Name,
			&i.Slug,
			&i.Vdp,
			&i.Favourite,
			&i.Tag,
			&i.Url,
			&i.Type,
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

const findProgramsWithTypeAndPlatform = `-- name: FindProgramsWithTypeAndPlatform :many
SELECT id, platform_id, name, slug, vdp, favourite, tag, url, type, created_at, updated_at FROM program WHERE type = $1::program_type AND platform_id = $2 LIMIT $3 OFFSET $4
`

type FindProgramsWithTypeAndPlatformParams struct {
	Column1    ProgramType `json:"column_1"`
	PlatformID int64       `json:"platform_id"`
	Limit      int32       `json:"limit"`
	Offset     int32       `json:"offset"`
}

func (q *Queries) FindProgramsWithTypeAndPlatform(ctx context.Context, arg FindProgramsWithTypeAndPlatformParams) ([]Program, error) {
	rows, err := q.db.Query(ctx, findProgramsWithTypeAndPlatform,
		arg.Column1,
		arg.PlatformID,
		arg.Limit,
		arg.Offset,
	)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	items := []Program{}
	for rows.Next() {
		var i Program
		if err := rows.Scan(
			&i.ID,
			&i.PlatformID,
			&i.Name,
			&i.Slug,
			&i.Vdp,
			&i.Favourite,
			&i.Tag,
			&i.Url,
			&i.Type,
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

const updateProgram = `-- name: UpdateProgram :one
UPDATE program SET platform_id = $2, name = $3, slug = $4, vdp = $5, url = $6, type = $7, updated_at = NOW() WHERE id = $1 RETURNING id, platform_id, name, slug, vdp, favourite, tag, url, type, created_at, updated_at
`

type UpdateProgramParams struct {
	ID         int64       `json:"id"`
	PlatformID int64       `json:"platform_id"`
	Name       string      `json:"name"`
	Slug       string      `json:"slug"`
	Vdp        bool        `json:"vdp"`
	Url        string      `json:"url"`
	Type       ProgramType `json:"type"`
}

func (q *Queries) UpdateProgram(ctx context.Context, arg UpdateProgramParams) (Program, error) {
	row := q.db.QueryRow(ctx, updateProgram,
		arg.ID,
		arg.PlatformID,
		arg.Name,
		arg.Slug,
		arg.Vdp,
		arg.Url,
		arg.Type,
	)
	var i Program
	err := row.Scan(
		&i.ID,
		&i.PlatformID,
		&i.Name,
		&i.Slug,
		&i.Vdp,
		&i.Favourite,
		&i.Tag,
		&i.Url,
		&i.Type,
		&i.CreatedAt,
		&i.UpdatedAt,
	)
	return i, err
}
