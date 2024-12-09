/*
 * Copyright 2024 AboveCloud9.AI Products and Services Private Limited
 * All rights reserved.
 * This code may not be used, copied, modified, or distributed without explicit permission.
 */

package handler

import (
	"ac9/glad/config"
	l "ac9/glad/pkg/logger"
	"ac9/glad/services/pushd/presenter"
	"context"
	"encoding/base64"

	firebase "firebase.google.com/go"
	"firebase.google.com/go/messaging"
	"google.golang.org/api/option"
)

type GoogleFCM struct {
	app *firebase.App
}

func GetGoogleFCM() GoogleFCM {
	app := initializeFCMApp()
	return GoogleFCM{app: app}
}

func initializeFCMApp() *firebase.App {
	ctx := context.Background()

	json, err := base64.StdEncoding.DecodeString(config.FCM_JSON)
	if err != nil {
		l.Log.Fatalf("Unable to decode FCM configuration. err=%v", err)
		return nil
	}

	opt := option.WithCredentialsJSON(json)
	app, err := firebase.NewApp(ctx, nil, opt)
	if err != nil {
		l.Log.Fatalf("error initializing FCM app. err=%v", err)
		return nil
	}
	return app
}

func (gFCM GoogleFCM) Send(ctx context.Context, token string, msg presenter.NotificationMessage) error {
	client, err := gFCM.app.Messaging(ctx)
	if err != nil {
		l.Log.Warnf("error getting messaging client. err=%v", err)
		return err
	}

	// Message to be sent
	message := &messaging.Message{
		Notification: &messaging.Notification{
			Title: msg.Header,
			Body:  msg.Content,
		},
		Token: token,
	}

	// TODO: If DRYRUN mode is set, then call client.SendDryRun()
	// Send the message
	_, err = client.Send(ctx, message)
	if err != nil {
		l.Log.Warnf("error sending message. err=%v, token=%v", err, token)
		return err
	}

	return nil
}
