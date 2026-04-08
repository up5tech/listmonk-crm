package core

import (
	"strconv"

	"github.com/knadh/listmonk/models"
)

type SearchContactParams struct {
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
	Email     string `json:"email"`
}

func (c *Core) QueryContacts(
	params SearchContactParams,
	orderBy, order string,
	limit, offset int,
) ([]models.Contact, int, error) {
	out := []models.Contact{}
	query := "SELECT * FROM contacts WHERE deleted = false"
	where := ""
	if params.FirstName != "" {
		where += " AND first_name LIKE '%" + params.FirstName + "%'"
	}
	if params.LastName != "" {
		where += " AND last_name LIKE '%" + params.LastName + "%'"
	}
	if params.Email != "" {
		where += " AND email LIKE '%" + params.Email + "%'"
	}
	total := 0
	if err := c.db.Get(&total, "SELECT COUNT(*) FROM contacts WHERE deleted = false"+where); err != nil {
		return nil, 0, err
	}
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
	if err := c.db.Select(&out, query); err != nil {
		return nil, 0, err
	}
	return out, total, nil
}

func (c *Core) GetContactByID(id int64) (*models.Contact, error) {
	var contact models.Contact
	if err := c.db.Get(&contact, "SELECT * FROM contacts WHERE id = ? AND deleted = false", id); err != nil {
		return nil, err
	}
	return &contact, nil
}

func (c *Core) CreateContact(contact models.Contact) (models.Contact, error) {
	err := c.q.CreateContact.QueryRow(
		contact.FirstName,
		contact.LastName,
		contact.Email,
		contact.Phone,
		contact.Description,
		contact.ContactType,
		contact.Status,
		contact.Score,
		contact.SubscriberID,
		contact.AccountID,
		contact.Source,
		contact.SourceID,
	).Scan(&contact.ID)
	if err != nil {
		return contact, err
	}
	return contact, nil
}

func (c *Core) UpdateContact(id int64, contact models.Contact) (models.Contact, error) {
	if _, err := c.q.UpdateContact.Exec(
		id,
		contact.FirstName,
		contact.LastName,
		contact.Email,
		contact.Phone,
		contact.Description,
		contact.ContactType,
		contact.Status,
		contact.Score,
		contact.SubscriberID,
		contact.AccountID,
		contact.Source,
		contact.SourceID,
	); err != nil {
		return contact, err
	}
	return contact, nil
}

func (c *Core) DeleteContact(id int64) error {
	_, err := c.db.Exec("UPDATE contacts SET deleted = true WHERE id = ?", id)
	if err != nil {
		return err
	}
	return nil
}
