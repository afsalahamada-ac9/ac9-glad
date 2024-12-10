/*
 * Copyright 2024 AboveCloud9.AI Products and Services Private Limited
 * All rights reserved.
 * This code may not be used, copied, modified, or distributed without explicit permission.
 */

package fcm

import (
	l "ac9/glad/pkg/logger"
	"context"
	"encoding/base64"

	firebase "firebase.google.com/go"
	"firebase.google.com/go/messaging"
	"google.golang.org/api/option"
)

// Firebase service
type Firebase struct {
	app      *firebase.App
	isDryRun bool
}

// NewFirebase
func NewFirebase(ctx context.Context, jsonCfg string, isDryRun bool) (*Firebase, error) {
	app, err := initializeFCMApp(ctx, jsonCfg)
	return &Firebase{app: app, isDryRun: isDryRun}, err
}

func initializeFCMApp(ctx context.Context, jsonCfg string) (*firebase.App, error) {
	// ctx := context.Background()

	json, err := base64.StdEncoding.DecodeString(jsonCfg)
	if err != nil {
		l.Log.Fatalf("Unable to decode FCM configuration. err=%v", err)
		return nil, err
	}

	opt := option.WithCredentialsJSON(json)
	app, err := firebase.NewApp(ctx, nil, opt)
	if err != nil {
		l.Log.Fatalf("Error initializing firebase app. err=%v", err)
		return nil, err
	}

	return app, nil
}

func (fb Firebase) Send(ctx context.Context, token string, header, content string) error {
	client, err := fb.app.Messaging(ctx)
	if err != nil {
		l.Log.Warnf("Error getting messaging client. err=%v", err)
		return err
	}

	// Message to be sent
	message := &messaging.Message{
		Notification: &messaging.Notification{
			Title: header,
			Body:  content,
		},
		Token: token,
	}

	// Send the message
	if fb.isDryRun {
		_, err = client.SendDryRun(ctx, message)
	} else {
		_, err = client.Send(ctx, message)
	}

	if err != nil {
		l.Log.Warnf("Error sending message. err=%v, token=%v", err, token)
		return err
	}

	return nil
}
