package route

import "github.com/google/wire"

var Set = wire.NewSet(NewRouter)
