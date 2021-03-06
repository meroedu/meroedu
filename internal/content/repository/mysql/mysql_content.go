package mysql

import (
	"context"
	"database/sql"
	"fmt"

	"github.com/meroedu/meroedu/internal/domain"
	"github.com/meroedu/meroedu/pkg/log"
)

type mysqlRepository struct {
	conn *sql.DB
}

// Init will create an object that represent the tag's Repository interface
func Init(db *sql.DB) domain.ContentRepository {
	return &mysqlRepository{
		conn: db,
	}
}
func (m *mysqlRepository) fetch(ctx context.Context, query string, args ...interface{}) (result []domain.Content, err error) {
	rows, err := m.conn.QueryContext(ctx, query, args...)
	if err != nil {
		log.Error(err)
		return nil, err
	}

	defer func() {
		errRow := rows.Close()
		if errRow != nil {
			log.Error(errRow)
		}
	}()

	result = make([]domain.Content, 0)
	for rows.Next() {
		t := domain.Content{}
		err = rows.Scan(
			&t.ID,
			&t.Title,
			&t.Description,
			&t.UpdatedAt,
			&t.CreatedAt,
		)

		if err != nil {
			log.Error(err)
			return nil, err
		}
		result = append(result, t)
	}

	return result, nil
}

func (m *mysqlRepository) GetAll(ctx context.Context, start int, limit int) (res []domain.Content, err error) {
	query := `SELECT id,title,description,updated_at,created_at FROM contents ORDER BY created_at DESC LIMIT ?,?`

	res, err = m.fetch(ctx, query, start, limit)
	if err != nil {
		return nil, err
	}
	return res, nil
}
func (m *mysqlRepository) GetByID(ctx context.Context, id int64) (res *domain.Content, err error) {
	query := `SELECT id,title,description,updated_at,created_at FROM contents WHERE ID = ?`

	list, err := m.fetch(ctx, query, id)
	if err != nil {
		return nil, err
	}
	var content domain.Content
	if len(list) > 0 {
		content = list[0]
	} else {
		return &content, domain.ErrNotFound
	}

	return &content, nil
}

func (m *mysqlRepository) CreateContent(ctx context.Context, a *domain.Content) (err error) {
	query := `INSERT contents SET title=?,description=?,fileheader=?,lesson_id=?,updated_at=?,created_at=?`
	stmt, err := m.conn.PrepareContext(ctx, query)
	if err != nil {
		log.Error("Error while preparing statement ", err)
		return
	}
	res, err := stmt.ExecContext(ctx, a.Title, a.Description, a.FileHeader, a.LessonID, a.UpdatedAt, a.CreatedAt)
	if err != nil {
		log.Error("Error while executing statement ", err)
		return
	}
	lastID, err := res.LastInsertId()
	if err != nil {
		log.Error("Got Error from LastInsertId method: ", err)
		return
	}
	a.ID = lastID
	return
}

func (m *mysqlRepository) DeleteContent(ctx context.Context, id int64) (err error) {
	query := "DELETE FROM contents WHERE id = ?"

	stmt, err := m.conn.PrepareContext(ctx, query)
	if err != nil {
		return
	}

	res, err := stmt.ExecContext(ctx, id)
	if err != nil {
		return
	}

	rowsAffected, err := res.RowsAffected()
	if err != nil {
		return
	}

	if rowsAffected != 1 {
		err = fmt.Errorf("Weird  Behavior. Total Affected: %d", rowsAffected)
		return
	}

	return
}
func (m *mysqlRepository) UpdateContent(ctx context.Context, ar *domain.Content) (err error) {
	query := `UPDATE contents set title=?,updated_at=? WHERE ID = ?`

	stmt, err := m.conn.PrepareContext(ctx, query)
	if err != nil {
		return
	}

	res, err := stmt.ExecContext(ctx, ar.Title, ar.UpdatedAt, ar.ID)
	if err != nil {
		return
	}
	affect, err := res.RowsAffected()
	if err != nil {
		return
	}
	if affect != 1 {
		err = fmt.Errorf("Weird  Behavior. Total Affected: %d", affect)
		return
	}

	return
}

func (m *mysqlRepository) GetContentCountByLesson(ctx context.Context, lessonID int64) (int, error) {
	query := `SELECT count(*) FROM contents WHERE lesson_id = ?`

	rows, err := m.conn.QueryContext(ctx, query, lessonID)
	if err != nil {
		log.Error(err)
		return 0, nil
	}
	defer rows.Close()
	var count int
	for rows.Next() {
		err := rows.Scan(&count)
		if err != nil {
			log.Error(err)
			return 0, err
		}
	}
	return count, nil
}

func (m *mysqlRepository) GetContentByLesson(ctx context.Context, lessonID int64) ([]domain.Content, error) {
	query := `SELECT id,title,description,updated_at,created_at FROM contents WHERE lesson_id = ?`
	list, err := m.fetch(ctx, query, lessonID)
	if err != nil {
		return nil, err
	}
	return list, nil
}
