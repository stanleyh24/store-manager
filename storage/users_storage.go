package storage

import (
	"context"
	"fmt"
	"log"
	"store-manager/types"

	"github.com/jackc/pgx/v5/pgxpool"
)

type UserStore interface {
	GetAllUsers(context.Context) ([]*types.User, error)

	//role
	GetAllRoles(context.Context) ([]*types.Role, error)
	CreateRole(context.Context, types.RoleCreateParams) (*types.Role, error)
	UpdateRole(context.Context, types.Role) (*types.Role, error)
	DeleteRole(context.Context, int) error
	AddRoleOperation(context.Context, int, []int) error
	DeleteRoleOperation(context.Context, int, []int) error
	GetAllOperations(ctx context.Context) ([]*types.Operation, error)
	GetRole(ctx context.Context, id int) (*types.RoleWithOperations, error)
}

type User struct {
	db *pgxpool.Pool
}

// New returns a new User storage
func NewUserStorage(db *pgxpool.Pool) *User {
	return &User{db}
}

func (u *User) GetAllUsers(context.Context) ([]*types.User, error) {
	panic("unimplemented")
}

func (u *User) CreateRole(ctx context.Context, params types.RoleCreateParams) (*types.Role, error) {
	sql := "INSERT INTO role(name) VALUES($1) RETURNING id;"
	var role types.Role
	row := u.db.QueryRow(ctx, sql, params.Name)

	if err := row.Scan(&role.ID); err != nil {
		return nil, err
	}
	role.Name = params.Name

	return &role, nil
}

func (u *User) GetAllRoles(ctx context.Context) ([]*types.Role, error) {
	sql := "SELECT id,name from role;"

	rows, err := u.db.Query(ctx, sql)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var roles []*types.Role

	for rows.Next() {
		var role types.Role
		err := rows.Scan(&role.ID, &role.Name)
		if err != nil {
			return nil, err
		}
		roles = append(roles, &role)
	}
	return roles, nil
}

func (u *User) GetRole(ctx context.Context, id int) (*types.RoleWithOperations, error) {
	sql := "SELECT id,name FROM role WHERE id=$1;"

	rows := u.db.QueryRow(ctx, sql, id)

	var role types.RoleWithOperations

	err := rows.Scan(&role.ID, &role.Name)
	if err != nil {
		return nil, err
	}

	return &role, nil
}

func (u *User) UpdateRole(ctx context.Context, role types.Role) (*types.Role, error) {
	sql := "UPDATE role SET name=$2 WHERE id=$1"

	_, err := u.db.Exec(ctx, sql, role.ID, role.Name)

	if err != nil {
		return nil, err
	}

	return &role, nil
}

func (u *User) DeleteRole(ctx context.Context, id int) error {
	sql := "DELETE FROM role WHERE id=$1;"
	row, err := u.db.Exec(ctx, sql, id)

	if err != nil {
		return err
	}
	if row.RowsAffected() == 0 {
		return fmt.Errorf("role with id: %d not found", id)
	}

	return nil
}

func (u *User) AddRoleOperation(ctx context.Context, idRole int, idOperations []int) error {
	sql := `INSERT INTO role_operations (id_role, id_operation) VALUES ($1,$2)`
	tx, err := u.db.Begin(ctx)
	if err != nil {
		log.Println(err)
	}

	for _, id := range idOperations {
		_, err := tx.Exec(ctx, sql, idRole, id)

		if err != nil {
			tx.Rollback(ctx)
			return err
		}
	}
	if err := tx.Commit(ctx); err != nil {
		return err
	}

	return nil
}

func (u *User) DeleteRoleOperation(ctx context.Context, idRole int, idOperations []int) error {
	sql := `DELETE FROM role_operations WHERE id_role = $1 AND id_operation= $2`
	tx, err := u.db.Begin(ctx)
	if err != nil {
		log.Println(err)
	}

	for _, id := range idOperations {
		_, err := tx.Exec(ctx, sql, idRole, id)

		if err != nil {
			tx.Rollback(ctx)
			return err
		}
	}
	if err := tx.Commit(ctx); err != nil {
		return err
	}

	return nil
}

func (u *User) GetAllOperations(ctx context.Context) ([]*types.Operation, error) {
	sql := "select id, name, (select name as module from modules where id=id_module) from operations;"

	rows, err := u.db.Query(ctx, sql)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var operations []*types.Operation

	for rows.Next() {
		var operation types.Operation
		err := rows.Scan(&operation.ID, &operation.Name, &operation.Module)
		if err != nil {
			return nil, err
		}
		operations = append(operations, &operation)
	}
	return operations, nil
}
