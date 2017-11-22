package data

import "runtime"

var PlatformIdentifier = runtime.GOOS + "/" + runtime.GOARCH
var ServiceApiEndpoint = "http://localhost:7080/api" // TODO release task to update this automatically
var VersionIdentifier = "v0.1.0 beta" // TODO release task to update this automatically
