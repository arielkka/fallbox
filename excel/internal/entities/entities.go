package entities

type (
	Excel struct {
		ID     int    `db:"id"`
		UserID string `db:"user_id"`
		Body   []byte `db:"body"`
	}
)
