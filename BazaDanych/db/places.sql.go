// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.21.0
// source: places.sql

package db

import (
	"context"
)

const createCemetery = `-- name: CreateCemetery :one
INSERT INTO cemetery (name, burial_place_id)
VALUES ($1, $2)
RETURNING id
`

type CreateCemeteryParams struct {
	Name          *string `json:"name"`
	BurialPlaceID int32   `json:"burial_place_id"`
}

func (q *Queries) CreateCemetery(ctx context.Context, arg CreateCemeteryParams) (int32, error) {
	row := q.db.QueryRow(ctx, createCemetery, arg.Name, arg.BurialPlaceID)
	var id int32
	err := row.Scan(&id)
	return id, err
}

const createGrave = `-- name: CreateGrave :one
INSERT INTO grave (name, row_id)
VALUES ($1, $2)
RETURNING id
`

type CreateGraveParams struct {
	Name  *string `json:"name"`
	RowID int32   `json:"row_id"`
}

func (q *Queries) CreateGrave(ctx context.Context, arg CreateGraveParams) (int32, error) {
	row := q.db.QueryRow(ctx, createGrave, arg.Name, arg.RowID)
	var id int32
	err := row.Scan(&id)
	return id, err
}

const createPlace = `-- name: CreatePlace :one
INSERT INTO place (name)
VALUES ($1)
RETURNING id
`

func (q *Queries) CreatePlace(ctx context.Context, name string) (int32, error) {
	row := q.db.QueryRow(ctx, createPlace, name)
	var id int32
	err := row.Scan(&id)
	return id, err
}

const createQuarter = `-- name: CreateQuarter :one
INSERT INTO quarter (name, cemetery_id)
VALUES ($1, $2)
RETURNING id
`

type CreateQuarterParams struct {
	Name       *string `json:"name"`
	CemeteryID int32   `json:"cemetery_id"`
}

func (q *Queries) CreateQuarter(ctx context.Context, arg CreateQuarterParams) (int32, error) {
	row := q.db.QueryRow(ctx, createQuarter, arg.Name, arg.CemeteryID)
	var id int32
	err := row.Scan(&id)
	return id, err
}

const createRow = `-- name: CreateRow :one
INSERT INTO row (name, quarter_id)
VALUES ($1, $2)
RETURNING id
`

type CreateRowParams struct {
	Name      *string `json:"name"`
	QuarterID int32   `json:"quarter_id"`
}

func (q *Queries) CreateRow(ctx context.Context, arg CreateRowParams) (int32, error) {
	row := q.db.QueryRow(ctx, createRow, arg.Name, arg.QuarterID)
	var id int32
	err := row.Scan(&id)
	return id, err
}

const deleteCemetery = `-- name: DeleteCemetery :exec
DELETE
FROM cemetery
WHERE id = $1
`

func (q *Queries) DeleteCemetery(ctx context.Context, id int32) error {
	_, err := q.db.Exec(ctx, deleteCemetery, id)
	return err
}

const deleteGrave = `-- name: DeleteGrave :exec
DELETE
FROM grave
WHERE id = $1
`

func (q *Queries) DeleteGrave(ctx context.Context, id int32) error {
	_, err := q.db.Exec(ctx, deleteGrave, id)
	return err
}

const deletePlace = `-- name: DeletePlace :exec
DELETE
FROM place
WHERE id = $1
`

func (q *Queries) DeletePlace(ctx context.Context, id int32) error {
	_, err := q.db.Exec(ctx, deletePlace, id)
	return err
}

const deleteQuarter = `-- name: DeleteQuarter :exec
DELETE
FROM quarter
WHERE id = $1
`

func (q *Queries) DeleteQuarter(ctx context.Context, id int32) error {
	_, err := q.db.Exec(ctx, deleteQuarter, id)
	return err
}

const deleteRow = `-- name: DeleteRow :exec
DELETE
FROM row
WHERE id = $1
`

func (q *Queries) DeleteRow(ctx context.Context, id int32) error {
	_, err := q.db.Exec(ctx, deleteRow, id)
	return err
}

const getCemeteries = `-- name: GetCemeteries :many
SELECT id, name, burial_place_id FROM cemetery
`

func (q *Queries) GetCemeteries(ctx context.Context) ([]Cemetery, error) {
	rows, err := q.db.Query(ctx, getCemeteries)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []Cemetery
	for rows.Next() {
		var i Cemetery
		if err := rows.Scan(&i.ID, &i.Name, &i.BurialPlaceID); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}

const getGraves = `-- name: GetGraves :many
SELECT id, name, row_id FROM grave
`

func (q *Queries) GetGraves(ctx context.Context) ([]Grave, error) {
	rows, err := q.db.Query(ctx, getGraves)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []Grave
	for rows.Next() {
		var i Grave
		if err := rows.Scan(&i.ID, &i.Name, &i.RowID); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}

const getPlaces = `-- name: GetPlaces :many
SELECT id, name FROM place
`

func (q *Queries) GetPlaces(ctx context.Context) ([]Place, error) {
	rows, err := q.db.Query(ctx, getPlaces)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []Place
	for rows.Next() {
		var i Place
		if err := rows.Scan(&i.ID, &i.Name); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}

const getPlacesFull = `-- name: GetPlacesFull :many
SELECT g.id   AS id,
       g.name AS grave,
       r.name AS row,
       q.name AS quarter,
       c.name AS cemetery,
       p.name AS place
FROM place p
         LEFT JOIN cemetery c ON p.id = c.burial_place_id
         LEFT JOIN quarter q ON c.id = q.cemetery_id
         LEFT JOIN row r ON q.id = r.quarter_id
         LEFT JOIN grave g ON r.id = g.row_id
`

type GetPlacesFullRow struct {
	ID       *int32  `json:"id"`
	Grave    *string `json:"grave"`
	Row      *string `json:"row"`
	Quarter  *string `json:"quarter"`
	Cemetery *string `json:"cemetery"`
	Place    string  `json:"place"`
}

func (q *Queries) GetPlacesFull(ctx context.Context) ([]GetPlacesFullRow, error) {
	rows, err := q.db.Query(ctx, getPlacesFull)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []GetPlacesFullRow
	for rows.Next() {
		var i GetPlacesFullRow
		if err := rows.Scan(
			&i.ID,
			&i.Grave,
			&i.Row,
			&i.Quarter,
			&i.Cemetery,
			&i.Place,
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

const getQuarters = `-- name: GetQuarters :many
SELECT id, name, cemetery_id FROM quarter
`

func (q *Queries) GetQuarters(ctx context.Context) ([]Quarter, error) {
	rows, err := q.db.Query(ctx, getQuarters)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []Quarter
	for rows.Next() {
		var i Quarter
		if err := rows.Scan(&i.ID, &i.Name, &i.CemeteryID); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}

const getRows = `-- name: GetRows :many
SELECT id, name, quarter_id FROM row
`

func (q *Queries) GetRows(ctx context.Context) ([]Row, error) {
	rows, err := q.db.Query(ctx, getRows)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []Row
	for rows.Next() {
		var i Row
		if err := rows.Scan(&i.ID, &i.Name, &i.QuarterID); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}

const updatePlace = `-- name: UpdatePlace :one
UPDATE place
SET name = $2
WHERE id = $1
RETURNING id
`

type UpdatePlaceParams struct {
	ID   int32  `json:"id"`
	Name string `json:"name"`
}

func (q *Queries) UpdatePlace(ctx context.Context, arg UpdatePlaceParams) (int32, error) {
	row := q.db.QueryRow(ctx, updatePlace, arg.ID, arg.Name)
	var id int32
	err := row.Scan(&id)
	return id, err
}