package postgres

import (
	"context"
	"github.com/jackc/pgx/v5"
	"social-media-api/internal/models"
	"social-media-api/pkg/postgres"
)

type CommentPostgres struct {
	*postgres.Postgres
}

func NewCommentPostgres(postgres *postgres.Postgres) *CommentPostgres {
	return &CommentPostgres{Postgres: postgres}
}

func (r *CommentPostgres) Save(ctx context.Context, comment models.Comment) (*models.Comment, error) {
	const query = `INSERT INTO comments (user_id, post_id, parent_id, body) 
				   VALUES ($1, $2, $3, $4)
				   RETURNING id, user_id, post_id, parent_id, body`

	rows, err := r.Pool.Query(ctx, query, comment.UserID, comment.PostID, comment.ParentID, comment.Body)
	if err != nil {
		return nil, err
	}

	commentRes, err := pgx.CollectExactlyOneRow(rows, pgx.RowToAddrOfStructByName[models.Comment])
	if err != nil {
		return nil, err
	}

	return commentRes, nil
}

func (r *CommentPostgres) GetByPostID(ctx context.Context, postID int, limit, offset int) ([]*models.Comment, error) {
	const query = `SELECT * 
				   FROM comments
				   WHERE post_id = $1 AND parent_id IS NULL
				   LIMIT $2 OFFSET $3`

	rows, err := r.Pool.Query(ctx, query, postID, limit, offset)
	if err != nil {
		return nil, err
	}

	comments, err := pgx.CollectRows(rows, pgx.RowToAddrOfStructByName[models.Comment])
	if err != nil {
		return nil, err
	}

	return comments, nil
}

func (r *CommentPostgres) GetByParentID(ctx context.Context, parentID int, limit, offset int) ([]*models.Comment, error) {
	const query = `SELECT * 
				   FROM comments
				   WHERE parent_id = $1
				   LIMIT $2 OFFSET $3`

	rows, err := r.Pool.Query(ctx, query, parentID, limit, offset)
	if err != nil {
		return nil, err
	}

	comments, err := pgx.CollectRows(rows, pgx.RowToAddrOfStructByName[models.Comment])
	if err != nil {
		return nil, err
	}

	return comments, nil
}

func (r *CommentPostgres) GetByID(ctx context.Context, id int) (*models.Comment, error) {
	const query = `SELECT * 
				   FROM comments
				   WHERE id = $1`

	rows, err := r.Pool.Query(ctx, query, id)
	if err != nil {
		return nil, err
	}

	comment, err := pgx.CollectExactlyOneRow(rows, pgx.RowToAddrOfStructByName[models.Comment])
	if err != nil {
		return nil, err
	}

	return comment, nil
}