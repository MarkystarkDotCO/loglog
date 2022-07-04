package domain

import "time"

type Checkpoint struct {
	FileName   string
	LineNumber int
	LastSeen   time.Time
}

type BlobMapper struct {
	BlobUrl   string    `bson:"blob_url,omitempty"`
	TimeStamp time.Time `bson:"timestamp,omitempty"`
}
