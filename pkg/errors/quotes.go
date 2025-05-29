package errors

import "fmt"

var ErrQuoteNotFound = fmt.Errorf("quote not found")

var ErrQuoteAlreadyExists = fmt.Errorf("quote already exists")
