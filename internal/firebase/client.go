package firebase

import (
	"context"
	"firebasectl/internal/config"
	"fmt"

	firebase "github.com/rapido-labs/firebase-admin-go/v4"
	"github.com/rapido-labs/firebase-admin-go/v4/remoteconfig"
	"google.golang.org/api/option"
)

type RemoteConfigClient struct {
	Client remoteconfig.Client
}

func getFirebaseApp(ctx context.Context) (*firebase.App, error) {
	firebaseConfig, err := config.GetFirebaseConfig()
	if err != nil {
		return nil, err
	}
	opts := option.WithCredentialsFile(firebaseConfig.Service_account_json_path)

	app, err := firebase.NewApp(ctx, nil, opts)
	if err != nil {
		return nil, fmt.Errorf("Error while getting firebase app: %s", err)
	}

	return app, nil
}

func GetRemoteConfigClient(ctx context.Context) (*RemoteConfigClient, error) {
	firebaseApp, err := getFirebaseApp(ctx)
	if err != nil {
		return nil, err
	}

	client, err := firebaseApp.RemoteConfig(ctx)
	if err != nil {
		return nil, fmt.Errorf("Error while getting remoteconfig client: %s", err.Error())
	}
	return &RemoteConfigClient{Client: *client}, nil
}
