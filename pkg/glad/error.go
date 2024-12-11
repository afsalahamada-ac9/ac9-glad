/*
 * Copyright 2024 AboveCloud9.AI Products and Services Private Limited
 * All rights reserved.
 * This code may not be used, copied, modified, or distributed without explicit permission.
 */

package glad

import "errors"

// ErrNotFound not found
var ErrNotFound = errors.New("not found")

// glad.ErrInvalidEntity invalid entity
var ErrInvalidEntity = errors.New("invalid entity")

// ErrTokenMismatch auth token invalid (error)
var ErrTokenMismatch = errors.New("auth token invalid")

// ErrAuthFailure credentials doesn't match (error)
var ErrAuthFailure = errors.New("authentication failure: invalid credentials")

// ErrCreateToken token creation error
var ErrCreateToken = errors.New("create auth token failed")

// ErrTooManyLabels too many labels
var ErrTooManyLabels = errors.New("too many labels")

// ErrLabelAlreadySet label already set
var ErrLabelAlreadySet = errors.New("label already set")

// ErrInvalidValue value is not value
var ErrInvalidValue = errors.New("invalid value")

// ErrAlreadyExists already exists
var ErrAlreadyExists = errors.New("already exists")

// ErrMissingParam missing parameter
var ErrMissingParam = errors.New("missing parameter")
