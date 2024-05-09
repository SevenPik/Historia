// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.21.0
// source: badges.sql

package db

import (
	"context"
	"encoding/json"
)

const getBadges = `-- name: GetBadges :many
SELECT b.id                                                      AS id,
       b.name                                                    AS name,
       json_agg(json_build_object('id', sb.id, 'name', sb.name)) AS sub_badges
FROM sub_badge sb
         LEFT JOIN badge b on sb.badge_id = b.id
GROUP BY b.id
`

type GetBadgesRow struct {
	ID        *int32          `json:"id"`
	Name      *string         `json:"name"`
	SubBadges json.RawMessage `json:"sub_badges"`
}

func (q *Queries) GetBadges(ctx context.Context) ([]GetBadgesRow, error) {
	rows, err := q.db.Query(ctx, getBadges)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []GetBadgesRow
	for rows.Next() {
		var i GetBadgesRow
		if err := rows.Scan(&i.ID, &i.Name, &i.SubBadges); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}