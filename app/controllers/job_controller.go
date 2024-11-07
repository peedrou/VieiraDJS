package controllers

import (
	"VieiraDJS/app/services/jobs"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/gocql/gocql"
)

type JobRequest struct {
	IsRecurring bool      `json:"is_recurring"`
	MaxRetries  int       `json:"max_retries"`
	StartTime   time.Time `json:"start_time"`
	Interval    string    `json:"interval"`
}

func CreateJobHandler(c *gin.Context) {
	var jobReq JobRequest

	if err := c.ShouldBindJSON(&jobReq); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid Request"})
		return
	}

	session := c.MustGet("session").(*gocql.Session)

	err := jobs.CreateJob(session, jobReq.IsRecurring, jobReq.MaxRetries, jobReq.StartTime, jobReq.Interval)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, gin.H{"message": "Job created successfully"})
}