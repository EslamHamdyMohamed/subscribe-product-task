package helpers

type TimeInterval string

const (
	MONTHLY TimeInterval = "MONTHLY"
	YEARLY  TimeInterval = "YEARLY"
)

type Status string

const (
	PENDING Status = "PENDING"
	PAID    Status = "PAID"
)
