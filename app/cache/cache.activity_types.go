package cache

// Define a map to hold activity types with calories per minute
var ActivityTypeCache map[string]int

// InitializeCache initializes the cache with static data, including calories per minute
func InitializeCacheActivityTypes() {
	ActivityTypeCache = map[string]int{
		"Walking":    4,
		"Yoga":       4,
		"Stretching": 4,
		"Cycling":    8,
		"Swimming":   8,
		"Dancing":    8,
		"Hiking":     10,
		"Running":    10,
		"HIIT":       10,
		"JumpRope":   10,
	}
}

