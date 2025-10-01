package main

// ======================================================
// 1. BASIC CONSTANTS — UNTYPED & TYPED
// ======================================================

const Pi = 3.14159265358979323846 // untyped constant — flexible!
const appName = "GoMaster"        // untyped string constant
const debug = false               // untyped bool

const maxRetries int = 3 // typed constant — enforced at compile time

// Implement String() to satisfy fmt.Stringer
func (s Status) String() string {
	return [...]string{"Pending", "Approved", "Rejected", "Cancelled"}[s]
}

// ======================================================
// 2. IOTA — AUTO-GENERATED SEQUENCES (ENUM-LIKE)
// ======================================================

// Define a custom type for our "enum"
type Status int

const (
	StatusPending   Status = iota // 0
	StatusApproved                // 1
	StatusRejected                // 2
	StatusCancelled               // 3
)

// ======================================================
// 3. SKIP VALUES WITH UNDERSCORE
// ======================================================

type Priority int

const (
	_        Priority = iota // skip 0
	Low                      // 1
	Medium                   // 2
	High                     // 3
	Critical                 // 4
)

func (p Priority) String() string {
	switch p {
	case Low:
		return "Low"
	case Medium:
		return "Medium"
	case High:
		return "High"
	case Critical:
		return "Critical"
	default:
		return "Unknown"
	}
}
