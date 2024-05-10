package mysql

import (
	"context"

	"github.com/jmoiron/sqlx"
	"github.com/savi2w/pupper/model"
)

type User struct {
	cli *sqlx.DB
}

func (r *User) Insert(ctx context.Context, user *model.User) error {
	query := `INSERT INTO pupper.tab_user (first_name, last_name, document_number, balance) VALUES (?, ?, ?, ?);`

	_, err := r.cli.ExecContext(ctx, query, user.FirstName, user.LastName, user.DocumentNumber, user.Balance)
	if err != nil {
		return err
	}

	return nil
}
