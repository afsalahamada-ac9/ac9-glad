/*
 * Copyright 2024 AboveCloud9.AI Products and Services Private Limited
 * All rights reserved.
 * This code may not be used, copied, modified, or distributed without explicit permission.
 */

package uuid

import guuid "github.com/google/uuid"

// ID entity ID
type UUID = guuid.UUID

// NewID create a new entity UUID
func New() UUID {
	return UUID(guuid.New())
}

// FromString convert a string to an entity UUID
func FromString(s string) (UUID, error) {
	id, err := guuid.Parse(s)
	return UUID(id), err
}
