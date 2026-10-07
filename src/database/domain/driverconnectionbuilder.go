package domain

import "context"

type DriverConnectionBuilder interface {
	Connect(ctx context.Context) error
	Disconnect(ctx context.Context) error
	Ping(ctx context.Context) error
}
