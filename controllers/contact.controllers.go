package controllers

import (
	"context"
)

type ContactController struct {
	db  *db.Queries
	ctx context.Context
}

func NewContactController(db *db.Queries, ctx context.Context) *ContactController {
	return &ContactController{db, ctx}
}
