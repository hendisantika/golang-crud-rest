package controllers

import (
	"context"
	"database/sql"
	"github.com/gin-gonic/gin"
	"golang-crud-rest/schemas"
	"net/http"
	"time"
)

type ContactController struct {
	db  *db.Queries
	ctx context.Context
}

func NewContactController(db *db.Queries, ctx context.Context) *ContactController {
	return &ContactController{db, ctx}
}

// Create contact  handler
func (cc *ContactController) CreateContact(ctx *gin.Context) {
	var payload *schemas.CreateContact

	if err := ctx.ShouldBindJSON(&payload); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"status": "Failed payload", "error": err.Error()})
		return
	}

	now := time.Now()
	args := &db.CreateContactParams{
		FirstName:   payload.FirstName,
		LastName:    payload.LastName,
		PhoneNumber: payload.PhoneNumber,
		Street:      payload.Street,
		CreatedAt:   now,
		UpdatedAt:   now,
	}

	contact, err := cc.db.CreateContact(ctx, *args)

	if err != nil {
		ctx.JSON(http.StatusBadGateway, gin.H{"status": "Failed retrieving contact", "error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"status": "successfully created contact", "contact": contact})
}

// Update contact handler
func (cc *ContactController) UpdateContact(ctx *gin.Context) {
	var payload *schemas.UpdateContact
	contactId := ctx.Param("contactId")

	if err := ctx.ShouldBindJSON(&payload); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"status": "Failed payload", "error": err.Error()})
		return
	}

	now := time.Now()
	args := &db.UpdateContactParams{
		ContactID:   uuid.MustParse(contactId),
		FirstName:   sql.NullString{String: payload.FirstName, Valid: payload.FirstName != ""},
		LastName:    sql.NullString{String: payload.LastName, Valid: payload.LastName != ""},
		PhoneNumber: sql.NullString{String: payload.PhoneNumber, Valid: payload.PhoneNumber != ""},
		Street:      sql.NullString{String: payload.PhoneNumber, Valid: payload.Street != ""},
		UpdatedAt:   sql.NullTime{Time: now, Valid: true},
	}

	contact, err := cc.db.UpdateContact(ctx, *args)

	if err != nil {
		if err == sql.ErrNoRows {
			ctx.JSON(http.StatusNotFound, gin.H{"status": "failed", "message": "Failed to retrieve contact with this ID"})
			return
		}
		ctx.JSON(http.StatusBadGateway, gin.H{"status": "Failed retrieving contact", "error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"status": "successfully updated contact", "contact": contact})
}
