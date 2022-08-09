package repositories

import (
	"database/sql"

	"github.com/jailtonjunior94/fullcycle-kafka/cmd/api/database"
	"github.com/jailtonjunior94/fullcycle-kafka/cmd/api/models"
)

func InsertCustomer(customer models.Customer) (id int64, err error) {
	conn, err := database.OpenConnection()
	if err != nil {
		return
	}
	defer conn.Close()

	query := `INSERT INTO [Customers] ([Document], [Name]) VALUES (@document, @name); SELECT SCOPE_IDENTITY();`

	err = conn.QueryRow(query, sql.Named("document", customer.Document), sql.Named("name", customer.Name)).Scan(&id)
	return
}
