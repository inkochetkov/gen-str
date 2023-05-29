package gen

import "github.com/google/uuid"

func GenUUID() (string, error) {

	newUUID, err := uuid.NewRandom()
	if err != nil {
		return "", err
	}

	return newUUID.String(), nil
}
