package main

import (
	"context"
	"fmt"

	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
)

type Storage struct {
	db *sqlx.DB
}

func NewStorage(ctx context.Context, connString string) (*Storage, error) {
	db, err := sqlx.ConnectContext(ctx, "postgres", connString)
	if err != nil {
		return nil, fmt.Errorf("open connection error: %w", err)
	}
	return &Storage{
		db: db,
	}, nil
}

func (s *Storage) Store(ctx context.Context, name string, score int) error {
	query := `INSERT INTO skymarvis (name, score) VALUES ($1, $2) ON CONFLICT (name) DO UPDATE SET score = $3;`
	_, err := s.db.ExecContext(ctx, query, name, score, score)
	if err != nil {
		return fmt.Errorf("insert data error: %w", err)
	}
	return nil
}

func (s *Storage) GetScores(ctx context.Context, limit int) (result []*User, _ error) {
	query := `SELECT name, score FROM skymarvis ORDER BY score DESC LIMIT $1;`
	rows, err := s.db.QueryxContext(ctx, query, limit)
	if err != nil {
		return nil, fmt.Errorf("insert data error: %w", err)
	}
	defer rows.Close()
	for rows.Next() {
		var user User
		if err := rows.StructScan(&user); err != nil {
			return nil, fmt.Errorf("scan row error: %w", err)
		}
		result = append(result, &user)
	}
	if rows.Err() != nil {
		return nil, fmt.Errorf("iterate error: %w", err)
	}
	return result, nil
}
