package helper

import "fmt"

func CreateErrorForDto(dto any, err error) error {
	return fmt.Errorf("not valid dto - %T: %w", dto, err)
}
