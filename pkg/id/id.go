/*
 * Copyright 2024 AboveCloud9.AI Products and Services Private Limited
 * All rights reserved.
 * This code may not be used, copied, modified, or distributed without explicit permission.
 */

package id

import (
	"ac9/glad/pkg/uid"
	"math/rand"
	"strconv"
)

const (
	IDInvalid = 0
)

// ID identifier
type ID uint64

// NewID create a new ID
func New() ID {
	randomShardID := rand.Intn(1024)
	return ID(uid.Get(randomShardID))
}

// NewIDWithShard create a new ID with given Shard ID
func NewIDWithShard(shardID int) ID {
	return ID(uid.Get(shardID))
}

// StringToID convert a string to an ID
func FromString(s string) (ID, error) {
	id, err := strconv.ParseUint(s, 10, 64)
	return ID(id), err
}

func (id ID) String() string {
	return strconv.FormatUint(uint64(id), 10)
}
