package thing

import (
	"net/http"
	"time"
)

// ThingResponse represents the response body for the Thing services.
type ThingResponse struct {
	// The UUID of a thing
	UUID string `json:"uuid" example:"6204037c-30e6-408b-8aaa-dd8219860b4b" description:"UUID of the thing"`
	// The Name of a thing
	Name string `json:"name" example:"Some name" description:"Name for the thing"`
	// The Value of a thing
	Value string `json:"value" example:"Some value" description:"Value for the thing"`
	// The last time a thing was updated
	Updated time.Time `json:"updated" example:"2021-05-25T00:53:16.535668Z" format:"date-time" description:"Time when the Thing was last updated"`
	// The time a thing was created
	Created time.Time `json:"created" example:"2021-05-25T00:53:16.535668Z" format:"date-time" description:"Time when the Thing was created."`
}

// ErrorResponse represents the error response messages when something has failed.
type ErrorResponse struct {
	// The error message
	Error string `json:"error" example:"An error occurred" description:"Error message"`
}

// GetThing godoc
// @title This is the summary for getting a thing by its UUID
// @description This is the description for getting a thing by its UUID. Which can be longer,
// @description and can continue over multiple lines
// @id get-thing
// @tags Thing
// @resource thing
// @param uuid path string true "The UUID of a thing"
// @success 200 object ThingResponse "Thing requested by ID"
// @failure 400 object ErrorResponse
// @failure 404 object ErrorResponse
// @failure 500 object ErrorResponse
// @route /thing/{uuid} [get]
func (s *Server) GetThing(w http.ResponseWriter, r *http.Request) {
	// Your implementation here
}

type CreateThing struct {
	// The name for a thing
	Name string `json:"name" validate:"required" example:"Some name"`
	// The value for a thing
	Value string `json:"value" validate:"required" example:"Some value"`
}

// CreateThing godoc
// @title This is the summary for creating a thing
// @description This is the description for creating a thing. Which can be longer.
// @id create-thing
// @tags Thing
// @resource thing
// @param newThing body CreateThing true "The body to create a thing"
// @success 200 object ThingResponse
// @failure 404 object ErrorResponse
// @failure 500 object ErrorResponse
// @route /thing/new [post]
func (s *Server) CreateThing(w http.ResponseWriter, r *http.Request) {
	// Your implementation here
}

type UpdateThing struct {
	// The new value for a thing
	Value string `json:"value" validate:"required" example:"Update value"`
}

// UpdateThing godoc
// @title This is the summary for updating a thing
// @description This is the description for updating a thing. Which can be longer.
// @id update-thing
// @tags Thing
// @param uuid path string true "The UUID of a thing"
// @param Body body UpdateThing true "The body to update a thing"
// @success 200 object ThingResponse
// @failure 404,500 object ErrorResponse
// @route /thing/{uuid} [put]
func (s *Server) UpdateThing(w http.ResponseWriter, r *http.Request) {
	// Your implementation here
}

// DeleteThing godoc
// @title This is the summary for deleting a thing
// @description This is the description for deleting a thing. Which can be longer.
// @id delete-thing
// @tags Thing
// @param uuid path string true "The UUID of a thing"
// @success 204 null
// @failure 500 object ErrorResponse
// @route /thing/{uuid} [delete]
func (s *Server) DeleteThing(w http.ResponseWriter, r *http.Request) {
	// Your implementation here
}

type ThingsResponse struct {
	Total  int             `json:"total" format:"int64"`
	Page   int             `json:"page" format:"int64"`
	Limit  int             `json:"limit" format:"int64"`
	Things []ThingResponse `json:"things"`
}

// ListThings godoc
// @title This is the summary for listing things
// @description This is the description for listing things. Which can be longer.
// @id list-things
// @tags Thing
// @param page query int false "Page"
// @param limit query int false "Limit (max 100)"
// @success 200 object ThingsResponse
// @failure 500 object ErrorResponse
// @route /thing [get]
func (s *Server) ListThings(w http.ResponseWriter, r *http.Request) {
	// Your implementation here
}
