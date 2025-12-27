package auth

import (
	"context"

	"github.com/AndreiShkolnyi/go-auth/internal/repository"
	"github.com/AndreiShkolnyi/go-auth/internal/repository/auth/converter"
	"github.com/AndreiShkolnyi/go-auth/internal/repository/auth/model"
	"github.com/AndreiShkolnyi/go-auth/pkg/auth_v1"
	sq "github.com/Masterminds/squirrel"
	"github.com/jackc/pgx/v5/pgxpool"
)

const (
	tableName = "auth"

	idColumn              = "id"
	nameColumn            = "name"
	emailColumn           = "email"
	passwordColumn        = "password"
	passwordConfirmColumn = "password_confirm"
	roleColumn            = "role"
	createdAtColumn       = "created_at"
	updatedAtColumn       = "updated_at"
)

type repo struct {
	db *pgxpool.Pool
}

func NewRepository(db *pgxpool.Pool) repository.AuthRepository {
	return &repo{db: db}
}

func (r *repo) Create(ctx context.Context, user *auth_v1.CreateRequest) (int64, error) {
	builder := sq.Insert(tableName).
		Columns(nameColumn, emailColumn, passwordColumn, passwordConfirmColumn, roleColumn).
		Values(user.Name, user.Email, user.Password, user.PasswordConfirm, user.Role).
		Suffix("RETURNING id")

	query, args, err := builder.ToSql()
	if err != nil {
		return 0, err
	}

	var id int64
	err = r.db.QueryRow(ctx, query, args...).Scan(&id)
	if err != nil {
		return 0, err
	}
	return id, nil
}

func (r *repo) Get(ctx context.Context, id int64) (*auth_v1.GetResponse, error) {
	builder := sq.Select(idColumn, nameColumn, emailColumn, roleColumn, createdAtColumn, updatedAtColumn).
		From(tableName).
		Where(sq.Eq{idColumn: id}).
		Limit(1)

	query, args, err := builder.ToSql()
	if err != nil {
		return nil, err
	}

	var user model.Auth

	err = r.db.QueryRow(ctx, query, args...).
		Scan(&user.ID, &user.Name, &user.Email, &user.Role, &user.CreatedAt, &user.UpdatedAt)
	if err != nil {
		return nil, err
	}
	return converter.ToUserFromRepo(&user)
}

func (r *repo) Update(ctx context.Context, user *auth_v1.UpdateRequest) (int64, error) {
	builder := sq.Update(tableName).
		Set(nameColumn, user.Name).
		Set(emailColumn, user.Email).
		Where(sq.Eq{idColumn: user.Id}).
		Suffix("RETURNING id")

	query, args, err := builder.ToSql()
	if err != nil {
		return 0, err
	}

	row := r.db.QueryRow(ctx, query, args...)
	var id int64
	err = row.Scan(&id)
	if err != nil {
		return 0, err
	}

	return id, nil
}

func (r *repo) Delete(ctx context.Context, id int64) (int64, error) {
	panic("implement me")
}
