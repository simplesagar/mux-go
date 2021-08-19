// Mux Go - Copyright 2019 Mux Inc.
// NOTE: This file is auto generated. Do not edit this file manually.

package muxgo

type Error struct {
	// A unique identifier for this error.
	Id int64 `json:"id,omitempty"`
	// The percentage of views that experienced this error.
	Percentage float64 `json:"percentage,omitempty"`
	// Notes that are attached to this error.
	Notes string `json:"notes,omitempty"`
	// The error message.
	Message string `json:"message,omitempty"`
	// The last time this error was seen (ISO 8601 timestamp).
	LastSeen string `json:"last_seen,omitempty"`
	// Description of the error.
	Description string `json:"description,omitempty"`
	// The total number of views that experiend this error.
	Count int64 `json:"count,omitempty"`
	// The error code
	Code int64 `json:"code,omitempty"`
}
