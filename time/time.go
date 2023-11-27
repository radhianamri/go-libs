package time

import "time"

func Reformat(timestamp, inputLayout, outputLayout string) (string, error) {
	parsedTime, err := time.Parse(inputLayout, timestamp)
	if err != nil {
		return "", err
	}
	return parsedTime.Format(outputLayout), nil
}
