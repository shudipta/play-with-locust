package mockCore

import (
	"time"
)

// StatusType is the type of a particular order at any given moment
type StatusType string

type (
	// CoreAssignedReq ..
	CoreAssignedReq struct {
		Driver       *int       `json:"driver"`
		Status       StatusType `json:"status"`
		DriversTried string     `json:"drivers_tried"`
		AssignedAt   time.Time  `json:"assigned_at"`
		CancelledBy  *int       `json:"cancelled_by"`
	}

	// ResponseBody ..
	ResponseBody struct {
		OrderHash string `json:"order_hash"`
		Status    string `json:"status"`
		Message   string `json:"message"`
	}
)
