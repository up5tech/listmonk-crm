package core

import (
	"strconv"

	"github.com/knadh/listmonk/models"
)

type SearchAccountParams struct {
	Name  string `query:"name"`
	Email string `query:"email"`
}

func (c *Core) QueryAccounts(
	search *SearchAccountParams,
	orderBy string,
	order string,
	limit int,
	offset int,
) ([]models.Account, int, error) {
	out := []models.Account{}
	query := "SELECT * FROM accounts WHERE deleted = false"
	where := ""
	if search != nil {
		if search.Name != "" {
			where += " AND name LIKE '%" + search.Name + "%'"
		}
		if search.Email != "" {
			where += " AND email LIKE '%" + search.Email + "%'"
		}
	}
	// get total
	total := 0
	err := c.db.Get(&total, "SELECT COUNT(*) FROM accounts WHERE deleted = false"+where)
	if err != nil {
		return nil, 0, err
	}
	// get list
	query += where
	if orderBy != "" {
		query += " ORDER BY " + orderBy + " " + order
	}
	if limit > 0 {
		query += " LIMIT " + strconv.Itoa(limit)
	}
	if offset > 0 {
		query += " OFFSET " + strconv.Itoa(offset)
	}
	err = c.db.Select(&out, query)
	return out, total, err
}

func (c *Core) GetAccountByID(id int64) (models.Account, error) {
	out := models.Account{}
	err := c.db.Get(&out, "SELECT * FROM accounts WHERE id = ? AND deleted = false", id)
	return out, err
}

func (c *Core) CreateAccount(acc models.Account) (models.Account, error) {
	err := c.q.CreateAccount.QueryRow(
		acc.Name,
		acc.ShortName,
		acc.AccountType,
		acc.Industry,
		acc.SicCode,
		acc.Website,
		acc.Description,
		acc.Source,
		acc.SourceID,
	).Scan(&acc.ID)
	return acc, err
}

func (c *Core) UpdateAccount(id int64, acc models.Account) error {
	_, err := c.q.UpdateAccount.Exec(
		id,
		acc.Name,
		acc.ShortName,
		acc.AccountType,
		acc.Industry,
		acc.SicCode,
		acc.Website,
		acc.Description,
		acc.Source,
		acc.SourceID,
	)
	return err
}

func (c *Core) DeleteAccount(id int64) error {
	_, err := c.db.Exec("UPDATE accounts SET deleted = true WHERE id = ?", id)
	return err
}
