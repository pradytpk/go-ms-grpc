package account

import (
	"context"
	"database/sql"

	_ "github.com/lib/pq"
)

type Repository interface {
	Close()
	Ping() error
	PutAccount(ctx context.Context, a Account) error
	GetAccountByID(ctx context.Context, id string) (*Account, error)
	ListAccounts(ctx context.Context, skip uint64, take uint64) ([]Account, error)
}

type postgresRepository struct {
	db *sql.DB
}

func NewPostgresRepository(url string) (Repository, error) {
	db, err := sql.Open("postgres", url)
	if err != nil {
		return nil, err
	}
	err = db.Ping()
	if err != nil {
		return nil, err
	}

	return &postgresRepository{db}, nil
}

// Close implements Repositoy.
func (p *postgresRepository) Close() {
	p.db.Close()
}

// Ping implements Repositoy.
func (p *postgresRepository) Ping() error {
	return p.db.Ping()
}

// GetAccountByID implements Repositoy.
func (p *postgresRepository) GetAccountByID(ctx context.Context, id string) (*Account, error) {
	row := p.db.QueryRowContext(ctx, "SELECT id,name from accounts WHERE id =$1", id)
	a := &Account{}
	if err := row.Scan(&a.ID, &a.Name); err != nil {
		return nil, err
	}
	return a, nil
}

// ListAccounts implements Repositoy.
func (p *postgresRepository) ListAccounts(ctx context.Context, skip uint64, take uint64) ([]Account, error) {
	rows, err := p.db.QueryContext(ctx,
		"SELECT id,name FROM accounts ORDER BY id DESC OFFSET $1 LIMIT $2",
		skip,
		take)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	account := []Account{}
	for rows.Next() {
		a := &Account{}
		if err = rows.Scan(&a.ID, &a.Name); err == nil {
			account = append(account, *a)
		}
	}
	if err = rows.Err(); err != nil {
		return nil, err
	}
	return account, nil
}

// PutAccount implements Repositoy.
func (p *postgresRepository) PutAccount(ctx context.Context, a Account) error {
	_, err := p.db.ExecContext(ctx, "INSERT INTO accounts(id,name) VALUES($1,$2)", a.ID, a.Name)
	return err
}
