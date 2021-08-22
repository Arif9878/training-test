package repository

import (
	"database/sql"
	"time"

	"github.com/Arif9878/stockbit-test/question-2/models"
)

type LogsRepository interface {
	StoreLogSearch(query *models.QueryParams) error
}

type Logs struct {
	*sql.DB
}

func NewLogsRepository(db *sql.DB) Logs {
	return Logs{db}
}

func (db Logs) StoreLogSearch(q *models.QueryParams) error {
	dt := time.Now()

	query := `INSERT  log_search SET datetime=? , search=?`
	stmt, err := db.Prepare(query)
	if err != nil {
		return err
	}

	_, err = stmt.Exec(dt.Format(time.RFC3339), q.Search)

	return err
}
