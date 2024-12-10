/*
 * Copyright 2024 AboveCloud9.AI Products and Services Private Limited
 * All rights reserved.
 * This code may not be used, copied, modified, or distributed without explicit permission.
 */

package device

import (
	l "ac9/glad/pkg/logger"
	"context"
)

// fakePush mimics push notification interface for testing
type fakePush struct {
	isDryRun bool
}

// newFakePNS creates new fake push notification service
func newFakePNS(isDryRun bool) *fakePush {
	return &fakePush{
		isDryRun: isDryRun,
	}
}

// Send sends push notification
func (fp *fakePush) Send(ctx context.Context,
	token string,
	header string,
	content string,
) error {
	if fp.isDryRun {
		l.Log.Debugf("DryRun: Log push notification message (%v:%v) for %v token", header, content, token)
	} else {
		l.Log.Debugf("Log push notification message (%v:%v) for %v token", header, content, token)
	}
	return nil
}
