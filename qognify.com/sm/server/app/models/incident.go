package models

import (
	"time"
)

type Incident struct {
	ID           int             `json:"id"`
	Name         string          `json:"name"`
	Status       EIncidentStatus `json:"status"`
	CreationTime time.Time       `json:"creationTime"`
}

type ColIncidents []Incident

type EIncidentStatus int

const (
	LOW EIncidentStatus = iota
	NORMAL
	MEDIUM
	HIGH
	URGENT
)

var statusNames = [...]string{
	"LOW",
	"NORMAL",
	"MEDIUM",
	"HIGH",
	"URGENT",
}

func (status EIncidentStatus) String() string {
	if status < LOW || status > URGENT {
		return "NORMAL"
	}
	return statusNames[status]
}
