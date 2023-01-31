package app

import (
	"fmt"
	"github.com/gocarina/gocsv"
	"github.com/gofrs/uuid"
	"os"
)

var WBCalculation16122022Storage = make(map[uuid.UUID]WBCalculation16122022)

func Storage() error {
	file, err := os.OpenFile("calculation_wb_16_12_2022.csv", os.O_RDWR|os.O_CREATE, os.ModePerm)
	if err != nil {
		return fmt.Errorf("os.OpenFile: %w", err)
	}

	calc := make([]WBCalculation16122022, 0)

	err = gocsv.UnmarshalFile(file, &calc)
	if err != nil {
		return fmt.Errorf("gocsv.UnmarshalFile: %w", err)
	}
	for i := range calc {
		WBCalculation16122022Storage[generateUUID()] = calc[i]
	}

	return nil
}

func generateUUID() uuid.UUID {
	return uuid.Must(uuid.NewV4())
}
