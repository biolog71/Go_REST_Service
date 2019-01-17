package repository

import (
	"fmt"
	"time"

	"qognify.com/sm/server/app/models"
)

var currentID int
var inicdents models.ColIncidents

func init() {
	CreateIncident(&models.Incident{Name: "Incident #1", Status: models.MEDIUM, CreationTime: time.Time.UTC(time.Now())})
	CreateIncident(&models.Incident{Name: "Incident #2", Status: models.HIGH, CreationTime: time.Time.UTC(time.Now())})
}

func CreateIncident(incident *models.Incident) {
	currentID++
	incident.ID = currentID
	incident.CreationTime = time.Time.UTC(time.Now())
	inicdents = append(inicdents, *incident)

}

func GetIncident(id int) models.Incident {
	for _, v := range inicdents {
		if v.ID == id {
			return v
		}
	}

	return models.Incident{}
}

func GetIncidents() models.ColIncidents {
	return inicdents
}

func DeleteIncident(id int) error {
	for i, v := range inicdents {
		if v.ID == id {
			inicdents = append(inicdents[:i], inicdents[i+1:]...)
			return nil
		}
	}
	return fmt.Errorf("Couldn't find Incident with id of %d to delete", id)
}
