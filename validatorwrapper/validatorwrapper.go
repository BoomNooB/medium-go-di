package validatorwrapper

import (
	"context"
	"encoding/csv"
	"errors"
	"fmt"
	"os"
	"sync"
	"time"

	"github.com/go-playground/validator/v10"
)

var (
	ErrValidationFailed = errors.New("validation failed")
)

type validatorWrapper struct {
	validator *validator.Validate
	mu        sync.Mutex
	csvFile   string
}

func NewValidatorWrapper(v *validator.Validate) *validatorWrapper {
	return &validatorWrapper{
		validator: v,
		mu:        sync.Mutex{},
		csvFile:   "validation_errors.csv",
	}
}

func (v *validatorWrapper) StructValidation(ctx context.Context, req any) error {
	err := v.validator.StructCtx(ctx, req)
	if err != nil {
		var validationErrs validator.ValidationErrors
		if errors.As(err, &validationErrs) {
			// Log validation errors to CSV
			err := v.logValidationErrors(validationErrs)
			if err != nil {
				return err
			}
			return ErrValidationFailed
		}
	}
	return err
}

func (v *validatorWrapper) logValidationErrors(vErr validator.ValidationErrors) error {
	v.mu.Lock()
	defer v.mu.Unlock()

	// Check if file exists to determine if we need to write headers
	fileExists := true
	_, err := os.Stat(v.csvFile)
	if os.IsNotExist(err) {
		fileExists = false
	}

	// Open CSV file in append mode
	file, err := os.OpenFile(v.csvFile, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		fmt.Printf("Error opening CSV file: %v\n", err)
		return err
	}
	defer file.Close()

	writer := csv.NewWriter(file)
	defer writer.Flush()

	// Write header if file is new
	if !fileExists {
		header := []string{"timestamp", "struct_and_field_name", "error_tag"}
		if err := writer.Write(header); err != nil {
			fmt.Printf("Error writing CSV header: %v\n", err)
			return err
		}
	}

	timestamp := time.Now().Format(time.RFC3339)

	// Collect all rows
	rows := make([][]string, 0, len(vErr))
	for _, fieldErr := range vErr {
		row := []string{
			timestamp,
			fieldErr.StructNamespace(),
			fieldErr.Tag(),
		}
		rows = append(rows, row)
	}

	// Write all rows at once
	if err := writer.WriteAll(rows); err != nil {
		fmt.Printf("Error writing CSV rows: %v\n", err)
		return err
	}
	return nil
}
