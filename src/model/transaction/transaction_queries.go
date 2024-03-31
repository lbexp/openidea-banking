package transaction_model

import (
	"fmt"
)

func BuildGetAllByUserIdQuery(filters GetAllByUserIdFilters) string {
	query := "SELECT transaction_id, balance, currency, proof_image_url, created_at, bank_account_number, bank_name, COUNT(*) OVER() AS total_items FROM transactions WHERE user_id = $1 ORDER BY created_at DESC"

	// Add limit and offset
	if filters.Limit > 0 {
		query += fmt.Sprintf(" LIMIT %d", filters.Limit)
		query += fmt.Sprintf(" OFFSET %d", filters.Offset)
	}

	return query
}
