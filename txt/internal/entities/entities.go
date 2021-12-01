package entities

type (
	Txt struct {
		ID     int    `db:"id"`
		UserID string `db:"user_id"`
		Body   []byte `db:"body"`
	}
)
