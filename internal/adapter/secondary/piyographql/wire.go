package piyographql

import "github.com/google/wire"

var Set = wire.NewSet(
	NewBaseClient,
	NewSampleClient,
)