package accruals

type status = string

const (
	NEW        status = "NEW"
	PROCESSING status = "PROCESSING"
	REGISTERED status = "REGISTERED"
	INVALID    status = "INVALID"
	PROCESSED  status = "PROCESSED"
)
