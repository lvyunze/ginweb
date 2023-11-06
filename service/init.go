package service

import "ginweb/models"

var (
	db, nil = models.ConnectMySQL()
)
