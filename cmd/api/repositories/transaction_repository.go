package repositories

import (
	"database/sql"

	"github.com/jailtonjunior94/fullcycle-kafka/cmd/api/database"
	"github.com/jailtonjunior94/fullcycle-kafka/cmd/api/models"
)

func InsertTransaction(transaction models.Transaction) (id int64, err error) {
	conn, err := database.OpenConnection()
	if err != nil {
		return
	}
	defer conn.Close()

	query := `INSERT INTO
					[Transaction] 
					(
					[CustomerId],
					[TransactionValue],
					[TransactionDate],
					[TransactionDescription]
					)
				VALUES
					(
					@customerId,
					@transactionValue,
					@transactionDate,
					@transactionDescription
					);
				SELECT
					SCOPE_IDENTITY();`

	err = conn.QueryRow(query,
		sql.Named("customerId", transaction.CustomerID),
		sql.Named("transactionValue", transaction.TransactionValue),
		sql.Named("transactionDate", transaction.TransactionDate),
		sql.Named("transactionDescription", transaction.TransactionDescription),
	).Scan(&id)

	return
}
