package repository

import (
	"database/sql"
	"errors"
	"fmt"
	"strings"

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

	query := fmt.Sprintf(`
		SELECT 
			t1.id, t1.title, t1.description 
		FROM 
			%s t1 
			INNER JOIN %s ul on t1.id=ul.list_id 
		WHERE ul.user_id = ?`, todoListsTable, usersListsTable)

	rows, err := r.db.Query(query, userId)
	if err != nil {
		logrus.Fatal(err)
	}
	defer rows.Close()

	for rows.Next() {
		var list rest.TodoList
		if err := rows.Scan(&list.Id, &list.Title, &list.Description); err != nil {
			return lists, err
		}
		lists = append(lists, list)
	}

	if err = rows.Err(); err != nil {
		return lists, err
	}

	return lists, nil
}

func (r *TodoListMysql) GetById(userId int, listId int) (rest.TodoList, error) {
	var list rest.TodoList

	query := fmt.Sprintf(`
		SELECT 
			t1.id, t1.title, t1.description 
		FROM 
			%s t1 
			INNER JOIN %s ul on t1.id=ul.list_id 
		WHERE ul.user_id = ? and ul.list_id = ?`, todoListsTable, usersListsTable)

	rows, err := r.db.Query(query, userId, listId)
	if err != nil {
		logrus.Fatal(err)
	}
	defer rows.Close()

	for rows.Next() {
		if err := rows.Scan(&list.Id, &list.Title, &list.Description); err != nil {
			return list, err
		}
	}

	if list.Id == 0 {
		return list, errors.New(fmt.Sprintf("list not found by id=%d", listId))
	}

	return list, nil
}

func (r *TodoListMysql) Delete(listId int) error {
	query := fmt.Sprintf(`DELETE FROM %s where id=?`, todoListsTable)
	res, err := r.db.Exec(query, listId)
	deleteRow, err := res.RowsAffected()

	logrus.Info(deleteRow)
	return err
}

func (r *TodoListMysql) Update(userId int, listId int, input rest.UpdateListInput) error {
	setValue := make([]string, 0)
	args := make([]interface{}, 0)
	agrId := 1
	
	if input.Title != nil {
		setValue = append(setValue, fmt.Sprintf("title=$%d", agrId))
		args = append(args, *input.Title)
		agrId++
	}
	
	if input.Description != nil {
		setValue = append(setValue, fmt.Sprintf("description=$%d", agrId))
		args = append(args, *input.Description)
		agrId++
	}
	// logrus.Fatal("333333333333333333")
	setQuery := strings.Join(setValue,  ", ")

	query := fmt.Sprintf("UPDATE %s tl SET %s FROM %s ul WHERE tl.id = ul.list_id AND ul.list_id=$%d AND ul.user_id=$%d", todoListsTable, setQuery, usersListsTable, agrId, agrId+1)
	args = append(args, listId, userId)

	logrus.Debugf("updateQuery: %s", query)
	logrus.Debugf("args: %s", args)

	_, err := r.db.Exec(query, args...)
	return err
}
