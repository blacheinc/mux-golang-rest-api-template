package helper

// SortDirection returns the sort direction interger
// for the given string direction
func SortDirection(sort string) int64 {
	if sort == "asc" {
		return 1
	}
	return -1
}
