package tests

import (
	"encoding/json"
	"net/http"
	"strings"

	"github.com/revel/revel/testing"
	"qognify.com/sm/server/app/models"
)

type IncidentsTest struct {
	testing.TestSuite
}

func (t *IncidentsTest) TestGetIncidents() {
	t.Get("/incidents")
	t.AssertOk()
	t.AssertContentType("application/json; charset=utf-8")
}

func (t *IncidentsTest) TestGetIncident() {
	t.Get("/incidents/2")
	t.AssertOk()
	t.AssertContentType("application/json; charset=utf-8")
}

func (t *IncidentsTest) TestDeleteIncident() {
	t.Delete("/incidents/1")
	t.AssertOk()
}

func (t *IncidentsTest) TestCreateIncident() {
	incident := models.Incident{Name: "incident new 3", Status: models.URGENT}
	inc, _ := json.Marshal(incident)

	t.Post("/incidents", "application/json; charset=utf-8", strings.NewReader(string(inc)))
	t.AssertStatus(http.StatusCreated)
	t.AssertContentType("application/json; charset=utf-8")
}
