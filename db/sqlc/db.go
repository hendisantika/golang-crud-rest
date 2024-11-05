package db

import (
	"context"
	"database/sql"
	"fmt"
)

type DBTX interface {
	ExecContext(context.Context, string, ...interface{}) (sql.Result, error)
	PrepareContext(context.Context, string) (*sql.Stmt, error)
	QueryContext(context.Context, string, ...interface{}) (*sql.Rows, error)
	QueryRowContext(context.Context, string, ...interface{}) *sql.Row
}

func New(db DBTX) *Queries {
	return &Queries{db: db}
}

func Prepare(ctx context.Context, db DBTX) (*Queries, error) {
	q := Queries{db: db}
	var err error
	if q.createContactStmt, err = db.PrepareContext(ctx, createContact); err != nil {
		return nil, fmt.Errorf("error preparing query CreateContact: %w", err)
	}
	if q.deleteContactStmt, err = db.PrepareContext(ctx, deleteContact); err != nil {
		return nil, fmt.Errorf("error preparing query DeleteContact: %w", err)
	}
	if q.getContactByIdStmt, err = db.PrepareContext(ctx, getContactById); err != nil {
		return nil, fmt.Errorf("error preparing query GetContactById: %w", err)
	}
	if q.listContactsStmt, err = db.PrepareContext(ctx, listContacts); err != nil {
		return nil, fmt.Errorf("error preparing query ListContacts: %w", err)
	}
	if q.updateContactStmt, err = db.PrepareContext(ctx, updateContact); err != nil {
		return nil, fmt.Errorf("error preparing query UpdateContact: %w", err)
	}
	return &q, nil
}

func (q *Queries) Close() error {
	var err error
	if q.createContactStmt != nil {
		if cerr := q.createContactStmt.Close(); cerr != nil {
			err = fmt.Errorf("error closing createContactStmt: %w", cerr)
		}
	}
	if q.deleteContactStmt != nil {
		if cerr := q.deleteContactStmt.Close(); cerr != nil {
			err = fmt.Errorf("error closing deleteContactStmt: %w", cerr)
		}
	}
	if q.getContactByIdStmt != nil {
		if cerr := q.getContactByIdStmt.Close(); cerr != nil {
			err = fmt.Errorf("error closing getContactByIdStmt: %w", cerr)
		}
	}
	if q.listContactsStmt != nil {
		if cerr := q.listContactsStmt.Close(); cerr != nil {
			err = fmt.Errorf("error closing listContactsStmt: %w", cerr)
		}
	}
	if q.updateContactStmt != nil {
		if cerr := q.updateContactStmt.Close(); cerr != nil {
			err = fmt.Errorf("error closing updateContactStmt: %w", cerr)
		}
	}
	return err
}
