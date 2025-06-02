package pageutil

func Offset(page, limit int64) int64 {
	return (page - 1) * limit
}
