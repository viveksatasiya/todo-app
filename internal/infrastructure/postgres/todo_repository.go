package postgres

import (
	"database/sql"
	"fmt"

	_ "github.com/lib/pq"
	"github.com/viveksatasiya/todo-app/internal/domain"
)

type ToDoRepository struct {
	db *sql.DB
}

func NewToDoRepository() (*ToDoRepository, error) {
	connStr := "user=your_user dbname=your_db password=your_password host=localhost sslmode=disable"
	db, err := sql.Open("postgres", connStr)
	if err != nil {
		return nil, err
	}

	return &ToDoRepository{db: db}, nil
}

func (r *ToDoRepository) Close() {
	r.db.Close()
}

func (r *ToDoRepository) Create(todo *domain.ToDo) error {
	query := `INSERT INTO todos (title, description, created_at) VALUES ($1, $2, $3) RETURNING id`
	err := r.db.QueryRow(query, todo.Title, todo.Description, todo.CreatedAt).Scan(&todo.ID)
	return err
}

func (r *ToDoRepository) FindAll() ([]*domain.ToDo, error) {
	query := `SELECT id, title, description, created_at, updated_at FROM todos`
	rows, err := r.db.Query(query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	todos := make([]*domain.ToDo, 0)
	for rows.Next() {
		todo := new(domain.ToDo)
		err := rows.Scan(&todo.ID, &todo.Title, &todo.Description, &todo.CreatedAt, &todo.UpdatedAt)
		if err != nil {
			return nil, err
		}
		todos = append(todos, todo)
	}

	return todos, nil
}

func (r *ToDoRepository) FindById(id string) (*domain.ToDo, error) {
	query := `SELECT id, title, description, created_at, updated_at FROM todos WHERE id=$1`
	row := r.db.QueryRow(query, id)

	todo := new(domain.ToDo)
	err := row.Scan(&todo.ID, &todo.Title, &todo.Description, &todo.CreatedAt, &todo.UpdatedAt)
	if err == sql.ErrNoRows {
		return nil, fmt.Errorf("ToDo not found")
	} else if err != nil {
		return nil, err
	}

	return todo, nil
}

func (r *ToDoRepository) Update(todo *domain.ToDo) error {
	query := `UPDATE todos SET title=$1, description=$2, updated_at=$3 WHERE id=$4`
	_, err := r.db.Exec(query, todo.Title, todo.Description, todo.UpdatedAt, todo.ID)
	return err
}

func (r *ToDoRepository) Delete(id string) error {
	query := `DELETE FROM todos WHERE id=$1`
	_, err := r.db.Exec(query, id)
	return err
}
