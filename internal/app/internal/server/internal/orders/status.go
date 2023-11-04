package orders

type status string

const (
	NEW        status = "NEW"
	PROCESSING status = "PROCESSING"
	INVALID    status = "INVALID"
	PROCESSED  status = "PROCESSED"
)
