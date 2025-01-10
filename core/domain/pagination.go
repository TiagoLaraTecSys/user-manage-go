package domain

type (
	Pagination struct {
		Page       int
		TotalPages int
		Limit      int
	}

	Data struct {
		Users []User
		Page  Pagination
	}
)
