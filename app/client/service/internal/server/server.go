package server

import "github.com/google/wire"

// ProviderSet is cron server providers.
var ProviderSet = wire.NewSet(NewCronServer)
