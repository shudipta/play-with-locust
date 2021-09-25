package mocka2

import "time"

type (
	// Meta ..
	Meta struct {
		CityID           int     `json:"city_id"`
		DriverType       int     `json:"driver_type"`
		IsFreelancer     int     `json:"is_freelancer"`
		IsOnline         int     `json:"is_online"`
		IsLoggedIn       int     `json:"is_logged_in"`
		IsSuspended      int     `json:"is_suspended"`
		IsCooldownLocked int     `json:"is_cooldown_locked"`
		IsVerified       int     `json:"is_verified"`
		Radius           string  `json:"radius"`
		CashInHand       float64 `json:"cash_in_hand"`
	}

	// Abilities ..
	Abilities struct {
		CanTakePeople  int `json:"can_take_people"`
		CanTakeParcels int `json:"can_take_parcels"`
	}

	// Preferences ..
	Preferences struct {
		// WantsToTakePeople  int `json:"wants_to_take_people"`
		WantsToTakeFood int `json:"wants_to_take_food"`
		// WantsToTakeMart    int `json:"wants_to_take_mart"`
		// WantsToTakeParcels int `json:"wants_to_take_parcels"`
		// PreferredDestinationZoneID string `json:"preferred_destination_zone_id"`
	}

	// Location ..
	Location struct {
		Near   string `json:"near"`
		Radius int    `json:"radius"`
	}

	// OrderMeta ..
	OrderMeta struct {
		GMV       float64 `json:"gmv"`        // OrderState.GMV
		UserID    int     `json:"user_id"`    // OrderState.UserID
		CreatedAt int64   `json:"created_at"` // OrderState.InitiatedAt
		// PickupLatitude       float64 `json:"pickup_latitude"`
		// PickupLongitude      float64 `json:"pickup_longitude"`
		// DestinationLatitude  float64 `json:"destination_latitude"`
		// DestinationLongitude float64 `json:"destination_longitude"`
		// UserPayableFare      int     `json:"user_payable_fare"`
		// EstimatedDuration         int     `json:"estimated_duration"`
		// EstimatedDistance         int     `json:"estimated_distance"`
		// RecentFailedAttemptsCount int     `json:"recent_failed_attempts_count"`
		// AssignsTillNow            int     `json:"assigns_till_now"`
	}
	// Order ..
	Order struct {
		Hash        string    `json:"hash"`         // OrderState.OrderID
		ServiceType string    `json:"service_type"` // Default: food
		OrderMeta   OrderMeta `json:"order_meta"`
		// ID          int       `json:"id"`
	}
	// Request //
	Request struct {
		ExcludeDrivers string        `json:"exclude_drivers"`
		Timeout        time.Duration `json:"timeout"`        // cfg.AllocationRequestTimeout
		ReleaseDriver  int           `json:"release_driver"` // Default: 0
		GoOffline      int           `json:"go_offline"`     // Default: 1
		ServerTime     int64         `json:"server_time"`
	}

	// Priority ..
	Priority struct {
		Name        string      `json:"-"`
		Meta        Meta        `json:"meta" `
		Abilities   Abilities   `json:"abilities" `
		Preferences Preferences `json:"preferences"`
		Location    Location    `json:"Location"`
	}

	// AllocateReq is the final request body
	AllocateReq struct {
		Limit      int        `json:"limit"`
		Location   Location   `json:"location"`
		Order      Order      `json:"order"`
		Request    Request    `json:"request"`
		Priorities []Priority `json:"priorities"`
		// Meta        Meta        `json:"meta"`
		// Abilities   Abilities   `json:"abilities"`
		// Preferences Preferences `json:"preferences"`
	}

	// CommonBody ...
	CommonBody struct {
		OrderHash   string `json:"hash"`
		UserID      int    `json:"user_id"`
		ServiceType string `json:"service_type"` // default: food
	}

	// ReleaseSpecs is a dummy
	ReleaseSpecs struct {
		CommonBody
		DriverID         int  `json:"driver_id"`
		GoOffline        bool `json:"go_offline"`
		IsOrderCompleted bool `json:"is_order_completed"`
	}

	/*********************
		Cancel request
	*********************/

	// CancelSpecs is a dummy
	CancelSpecs CommonBody

	// ResponseBody ..
	ResponseBody struct {
		OrderHash string `json:"order_hash"`
		Status    string `json:"status"`
		Message   string `json:"message"`
	}

	// CallbackReq of A2
	CallbackReq struct {
		DriverID *int   `json:"driver_id,omitempty"`
		UserID   int    `json:"user_id"`
		OrderID  string `json:"order_id"`
		Location struct {
			Lat float64 `json:"lat"`
			Lon float64 `json:"lon"`
		} `json:"location"`
	}

	// TimeoutReq of A2
	TimeoutReq struct {
		OrderID string        `json:"order_id"`
		Timeout time.Duration `json:"timeout"`
	}
)
