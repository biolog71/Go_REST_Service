package controllers

import (
	"net/http"
	"strconv"

	"qognify.com/sm/server/app/models"

	"github.com/revel/revel"
	"qognify.com/sm/server/app/repository"
)

type Incidents struct {
	*revel.Controller
}

func (c Incidents) GetIncidents() revel.Result {
	return c.RenderJSON(repository.GetIncidents())
}

func (c Incidents) GetIncident() revel.Result {
	id, err := strconv.Atoi(c.Params.Route.Get("id"))
	if err != nil {
		c.RenderError(err)
	}
	return c.RenderJSON(repository.GetIncident(id))
}

func (c Incidents) CreateIncident() revel.Result {
	var incident models.Incident
	err := c.Params.BindJSON(&incident)
	if err != nil {
		c.RenderError(err)
	}

	repository.CreateIncident(&incident)
	c.Response.Status = http.StatusCreated
	return c.RenderJSON(incident)
}

func (c Incidents) DeleteIncident() revel.Result {
	id, err := strconv.Atoi(c.Params.Route.Get("id"))
	if err != nil {
		c.RenderError(err)
	}
	if err := repository.DeleteIncident(id); err != nil {
		c.RenderError(err)
	}

	c.Response.Status = http.StatusOK
	var t interface{}
	return c.RenderJSON(t)
}
