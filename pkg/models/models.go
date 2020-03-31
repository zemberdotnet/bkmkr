package models 
import (
	"errors"
	"time"
) 

type Snippet struct {
	ID int
	Title string
	URL string
	Created time.Time
	Expires time.Time
}

