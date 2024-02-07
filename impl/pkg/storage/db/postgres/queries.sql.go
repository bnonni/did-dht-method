// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.25.0
// source: queries.sql

package postgres

import (
	"context"
)

const listAllRecords = `-- name: ListAllRecords :many
SELECT id, key, value, sig, seq FROM pkarr_records
`

func (q *Queries) ListAllRecords(ctx context.Context) ([]PkarrRecord, error) {
	rows, err := q.db.Query(ctx, listAllRecords)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []PkarrRecord
	for rows.Next() {
		var i PkarrRecord
		if err := rows.Scan(
			&i.ID,
			&i.Key,
			&i.Value,
			&i.Sig,
			&i.Seq,
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

const listRecords = `-- name: ListRecords :many
SELECT id, key, value, sig, seq FROM pkarr_records WHERE id > (SELECT id FROM pkarr_records WHERE pkarr_records.key = $1) ORDER BY ID ASC LIMIT $2
`

type ListRecordsParams struct {
	Key   []byte
	Limit int32
}

func (q *Queries) ListRecords(ctx context.Context, arg ListRecordsParams) ([]PkarrRecord, error) {
	rows, err := q.db.Query(ctx, listRecords, arg.Key, arg.Limit)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []PkarrRecord
	for rows.Next() {
		var i PkarrRecord
		if err := rows.Scan(
			&i.ID,
			&i.Key,
			&i.Value,
			&i.Sig,
			&i.Seq,
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

const listRecordsFirstPage = `-- name: ListRecordsFirstPage :many
SELECT id, key, value, sig, seq FROM pkarr_records ORDER BY id ASC LIMIT $1
`

func (q *Queries) ListRecordsFirstPage(ctx context.Context, limit int32) ([]PkarrRecord, error) {
	rows, err := q.db.Query(ctx, listRecordsFirstPage, limit)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []PkarrRecord
	for rows.Next() {
		var i PkarrRecord
		if err := rows.Scan(
			&i.ID,
			&i.Key,
			&i.Value,
			&i.Sig,
			&i.Seq,
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

const readRecord = `-- name: ReadRecord :one
SELECT id, key, value, sig, seq FROM pkarr_records WHERE key = $1 LIMIT 1
`

func (q *Queries) ReadRecord(ctx context.Context, key []byte) (PkarrRecord, error) {
	row := q.db.QueryRow(ctx, readRecord, key)
	var i PkarrRecord
	err := row.Scan(
		&i.ID,
		&i.Key,
		&i.Value,
		&i.Sig,
		&i.Seq,
	)
	return i, err
}

const writeRecord = `-- name: WriteRecord :exec
INSERT INTO pkarr_records(key, value, sig, seq) VALUES($1, $2, $3, $4)
`

type WriteRecordParams struct {
	Key   []byte
	Value []byte
	Sig   []byte
	Seq   int64
}

func (q *Queries) WriteRecord(ctx context.Context, arg WriteRecordParams) error {
	_, err := q.db.Exec(ctx, writeRecord,
		arg.Key,
		arg.Value,
		arg.Sig,
		arg.Seq,
	)
	return err
}