package models

import "time"

type Exercise struct {
	ExerciseID   int64
	ExerciseName string
	Description  *string
	Created      time.Time
	Updated      time.Time
}
