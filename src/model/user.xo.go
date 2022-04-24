package model

// Code generated by xo. DO NOT EDIT.

import (
	"context"
	"time"
)

// User represents a row from 'bitgin.users'.
type User struct {
	ID        int       `json:"id"`         // id
	Email     string    `json:"email"`      // email
	Password  string    `json:"password"`   // password
	Role      string    `json:"role"`       // role
	UpdatedAt time.Time `json:"updated_at"` // updated_at
	CreatedAt time.Time `json:"created_at"` // created_at
	// xo fields
	_exists, _deleted bool
}

// Exists returns true when the User exists in the database.
func (u *User) Exists() bool {
	return u._exists
}

// Deleted returns true when the User has been marked for deletion from
// the database.
func (u *User) Deleted() bool {
	return u._deleted
}

// Insert inserts the User to the database.
func (u *User) Insert(ctx context.Context, db DB) error {
	switch {
	case u._exists: // already exists
		return logerror(&ErrInsertFailed{ErrAlreadyExists})
	case u._deleted: // deleted
		return logerror(&ErrInsertFailed{ErrMarkedForDeletion})
	}
	// insert (primary key generated and returned by database)
	const sqlstr = `INSERT INTO bitgin.users (` +
		`email, password, role, updated_at, created_at` +
		`) VALUES (` +
		`?, ?, ?, ?, ?` +
		`)`
	// run
	u.CreatedAt = time.Now()
	u.UpdatedAt = time.Now()
	logf(sqlstr, u.Email, u.Password, u.Role, u.UpdatedAt, u.CreatedAt)
	res, err := db.ExecContext(ctx, sqlstr, u.Email, u.Password, u.Role, u.UpdatedAt, u.CreatedAt)
	if err != nil {
		return logerror(err)
	}
	// retrieve id
	id, err := res.LastInsertId()
	if err != nil {
		return logerror(err)
	} // set primary key
	u.ID = int(id)
	// set exists
	u._exists = true
	return nil
}

// Update updates a User in the database.
func (u *User) Update(ctx context.Context, db DB) error {
	switch {
	case !u._exists: // doesn't exist
		return logerror(&ErrUpdateFailed{ErrDoesNotExist})
	case u._deleted: // deleted
		return logerror(&ErrUpdateFailed{ErrMarkedForDeletion})
	}
	// update with primary key
	const sqlstr = `UPDATE bitgin.users SET ` +
		`email = ?, password = ?, role = ?, updated_at = ?, created_at = ? ` +
		`WHERE id = ?`
	// run
	u.UpdatedAt = time.Now()
	logf(sqlstr, u.Email, u.Password, u.Role, u.UpdatedAt, u.CreatedAt, u.ID)
	if _, err := db.ExecContext(ctx, sqlstr, u.Email, u.Password, u.Role, u.UpdatedAt, u.CreatedAt, u.ID); err != nil {
		return logerror(err)
	}
	return nil
}

// Save saves the User to the database.
func (u *User) Save(ctx context.Context, db DB) error {
	if u.Exists() {
		return u.Update(ctx, db)
	}
	return u.Insert(ctx, db)
}

// Upsert performs an upsert for User.
func (u *User) Upsert(ctx context.Context, db DB) error {
	switch {
	case u._deleted: // deleted
		return logerror(&ErrUpsertFailed{ErrMarkedForDeletion})
	}
	// upsert
	const sqlstr = `INSERT INTO bitgin.users (` +
		`id, email, password, role, updated_at, created_at` +
		`) VALUES (` +
		`?, ?, ?, ?, ?, ?` +
		`)` +
		` ON DUPLICATE KEY UPDATE ` +
		`email = VALUES(email), password = VALUES(password), role = VALUES(role), updated_at = VALUES(updated_at), created_at = VALUES(created_at)`
	// run
	logf(sqlstr, u.ID, u.Email, u.Password, u.Role, u.UpdatedAt, u.CreatedAt)
	if _, err := db.ExecContext(ctx, sqlstr, u.ID, u.Email, u.Password, u.Role, u.UpdatedAt, u.CreatedAt); err != nil {
		return logerror(err)
	}
	// set exists
	u._exists = true
	return nil
}

// Delete deletes the User from the database.
func (u *User) Delete(ctx context.Context, db DB) error {
	switch {
	case !u._exists: // doesn't exist
		return nil
	case u._deleted: // deleted
		return nil
	}
	// delete with single primary key
	const sqlstr = `DELETE FROM bitgin.users ` +
		`WHERE id = ?`
	// run
	logf(sqlstr, u.ID)
	if _, err := db.ExecContext(ctx, sqlstr, u.ID); err != nil {
		return logerror(err)
	}
	// set deleted
	u._deleted = true
	return nil
}

// UserByEmail retrieves a row from 'bitgin.users' as a User.
//
// Generated from index 'email'.
func UserByEmail(ctx context.Context, db DB, email string) (*User, error) {
	// query
	const sqlstr = `SELECT ` +
		`id, email, password, role, updated_at, created_at ` +
		`FROM bitgin.users ` +
		`WHERE email = ?`
	// run
	logf(sqlstr, email)
	u := User{
		_exists: true,
	}
	if err := db.QueryRowContext(ctx, sqlstr, email).Scan(&u.ID, &u.Email, &u.Password, &u.Role, &u.UpdatedAt, &u.CreatedAt); err != nil {
		return nil, logerror(err)
	}
	return &u, nil
}

// UserByID retrieves a row from 'bitgin.users' as a User.
//
// Generated from index 'users_id_pkey'.
func UserByID(ctx context.Context, db DB, id int) (*User, error) {
	// query
	const sqlstr = `SELECT ` +
		`id, email, password, role, updated_at, created_at ` +
		`FROM bitgin.users ` +
		`WHERE id = ?`
	// run
	logf(sqlstr, id)
	u := User{
		_exists: true,
	}
	if err := db.QueryRowContext(ctx, sqlstr, id).Scan(&u.ID, &u.Email, &u.Password, &u.Role, &u.UpdatedAt, &u.CreatedAt); err != nil {
		return nil, logerror(err)
	}
	return &u, nil
}
