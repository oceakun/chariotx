package app

import (
    "time"
    "fmt"
)

// ConvertUnixMillisToTime converts milliseconds to time.Time
func ConvertUnixMillisToTime(ms int64) time.Time {
    return time.Unix(0, ms*int64(time.Millisecond))
}

// ConvertTimeToUnixMillis converts time.Time to milliseconds since epoch
func ConvertTimeToUnixMillis(t time.Time) int64 {
    return t.UnixNano() / int64(time.Millisecond)
}

// ValidateLatitude checks if latitude is within valid range
func ValidateLatitude(lat float64) bool {
    return lat >= -90 && lat <= 90
}

// ValidateLongitude checks if longitude is within valid range
func ValidateLongitude(lng float64) bool {
    return lng >= -180 && lng <= 180
}

// FormatError adds context to an error
func FormatError(context string, err error) error {
    return fmt.Errorf("%s: %v", context, err)
}
