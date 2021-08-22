package repository_test

import (
	"fmt"
	"testing"
	"time"

	"github.com/Arif9878/stockbit-test/question-2/models"
	"github.com/Arif9878/stockbit-test/question-2/repository"
	"github.com/DATA-DOG/go-sqlmock"
	"github.com/stretchr/testify/assert"
)

func Test_StoreLogSearch(t *testing.T) {

	ar := &models.QueryParams{
		Search: "Batman",
	}
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}
	defer db.Close()

	dt := time.Now()

	query := "INSERT  log_search SET datetime=\\? , search=\\?"
	prep := mock.ExpectPrepare(query)
	prep.ExpectExec().WithArgs(dt.Format(time.RFC3339), ar.Search).WillReturnResult(sqlmock.NewResult(12, 1))
	a := repository.NewLogsRepository(db)

	err = a.StoreLogSearch(ar)
	fmt.Println(err)
	assert.NoError(t, err)
}
