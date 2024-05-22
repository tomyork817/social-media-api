package postgres

import (
	"context"
	"github.com/jackc/pgx/v5"
	"social-media-api/internal/models"
	"social-media-api/pkg/postgres"
)

type PostPostgres struct {
	*postgres.Postgres
}

func NewPostPostgres(postgres *postgres.Postgres) *PostPostgres {
	return &PostPostgres{Postgres: postgres}
}

func (r *PostPostgres) Save(ctx context.Context, post models.Post) (*models.Post, error) {
	const query = `INSERT INTO posts (user_id, body, is_open) 
				   VALUES ($1, $2, $3)
				   RETURNING id, user_id, body, is_open`

	rows, err := r.Pool.Query(ctx, query, post.UserID, post.Body, post.IsOpen)
	if err != nil {
		return nil, err
	}

	postRes, err := pgx.CollectExactlyOneRow(rows, pgx.RowToAddrOfStructByName[models.Post])
	if err != nil {
		return nil, err
	}

	return postRes, nil
}

func (r *PostPostgres) GetAll(ctx context.Context, limit, offset int) ([]*models.Post, error) {
	const query = `SELECT * 
				   FROM posts
				   LIMIT $1 OFFSET $2`

	rows, err := r.Pool.Query(ctx, query, limit, offset)
	if err != nil {
		return nil, err
	}

	posts, err := pgx.CollectRows(rows, pgx.RowToAddrOfStructByName[models.Post])
	if err != nil {
		return nil, err
	}

	return posts, nil
}

func (r *PostPostgres) GetByID(ctx context.Context, id int) (*models.Post, error) {
	const query = `SELECT * 
				   FROM posts
				   WHERE id = $1`

	rows, err := r.Pool.Query(ctx, query, id)
	if err != nil {
		return nil, err
	}

	post, err := pgx.CollectExactlyOneRow(rows, pgx.RowToAddrOfStructByName[models.Post])
	if err != nil {
		return nil, err
	}

	return post, nil
}

func (r *PostPostgres) UpdateIsOpenById(ctx context.Context, id int, isOpen bool) (*models.Post, error) {
	const query = `
        UPDATE posts
        SET is_open = $1
        WHERE id = $2
        RETURNING id, user_id, body, is_open`

	rows, err := r.Pool.Query(ctx, query, isOpen, id)
	if err != nil {
		return nil, err
	}

	post, err := pgx.CollectExactlyOneRow(rows, pgx.RowToAddrOfStructByName[models.Post])
	if err != nil {
		return nil, err
	}

	return post, nil
}

func (r *PostPostgres) GetByUserID(ctx context.Context, userID int, limit, offset int) ([]*models.Post, error) {
	const query = `SELECT * 
				   FROM posts
				   WHERE user_id = $1
				   LIMIT $2 OFFSET $3`

	rows, err := r.Pool.Query(ctx, query, userID, limit, offset)
	if err != nil {
		return nil, err
	}

	posts, err := pgx.CollectRows(rows, pgx.RowToAddrOfStructByName[models.Post])
	if err != nil {
		return nil, err
	}

	return posts, nil
}
