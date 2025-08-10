package helpers

import "time"

func ConvertStringTime(stringTime string) (time.Time, error) {
	parsedTime, err := time.Parse(time.RFC3339, stringTime)
	if err != nil {
		return time.Time{}, err
	}
	return parsedTime, nil
}
