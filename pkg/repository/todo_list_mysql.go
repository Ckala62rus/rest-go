package repository

import (
	"database/sql"
	"fmt"

	"github.com/Ckala62rus/rest-go"
	"github.com/sirupsen/logrus"
)

type TodoListMysql struct {
	db *sql.DB
}

func NewTodoListMysql(db *sql.DB) *TodoListMysql {
	return &TodoListMysql{db: db}
}

func (r *TodoListMysql) Create(userId int, list rest.TodoList) (int, error) {
	tx, err := r.db.Begin()
	if err != nil {
		return 0, err
	}

	var idUserLists int

	createListQuery := fmt.Sprintf("INSERT INTO %s (title, description) VALUES (?,?)", todoListsTable)
	res, err := tx.Prepare(createListQuery)
	if err != nil {
		tx.Rollback()
		return 0, err
	}

	executeQuery, err := res.Exec(list.Title, list.Description)

	if err != nil {
		tx.Rollback()
		return 0, err
	}

	returnId, err := executeQuery.LastInsertId()
	if err != nil {
		tx.Rollback()
		return 0, err
	}

	createUsersListQuery := fmt.Sprintf("INSERT INTO %s (user_id, list_id) VALUES (?,?)", usersListsTable)
	resList, err := tx.Query(createUsersListQuery, userId, returnId)

	if err != nil {
		tx.Rollback()
		return 0, err
	}

	for resList.Next() {
		err = resList.Scan(&idUserLists)

		if err != nil {
			tx.Rollback()
			return 0, err
		}
	}

	return int(returnId), tx.Commit()
}

func (r *TodoListMysql) GetAll(userId int) ([]rest.TodoList, error) {
	var lists []rest.TodoList
	query := fmt.Sprintf("SELECT t1.id, t1.title, t1.description FROM %s t1 INNER JOIN %s ul on t1.id=ul.list_id WHERE ul.user_id = ?", todoListsTable, usersListsTable)
	// res, err := r.db.Query(query, userId)
	res := r.db.QueryRow(query, userId)

	logrus.Info(res)

	// if err != nil {
	// 	return lists, err
	// }

	// if err != nil {
	// 	return lists, err
	// }

	// for res.Next() {
	// 	err = res.Scan(&lists)
	// 	if err != nil {
	// 		return lists, err
	// 	}
	// }

	err := res.Scan(&lists)
	if err != nil {
		return lists, err
	}

	return lists, nil
}
