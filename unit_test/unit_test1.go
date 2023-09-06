package unit_test

import "time"

func LoadFromTrustedLocationOrNil(location string) *time.Location {
	tz, err := time.LoadLocation(location)

	if err != nil {
		return nil
	}

	return tz
}
