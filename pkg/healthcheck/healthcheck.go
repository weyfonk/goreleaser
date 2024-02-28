// Package healthcheck checks for missing binaries that the user needs to
// install.
package healthcheck

import (
	"fmt"

	"github.com/weyfonk/goreleaser/internal/pipe/chocolatey"
	"github.com/weyfonk/goreleaser/internal/pipe/docker"
	"github.com/weyfonk/goreleaser/internal/pipe/nix"
	"github.com/weyfonk/goreleaser/internal/pipe/sbom"
	"github.com/weyfonk/goreleaser/internal/pipe/sign"
	"github.com/weyfonk/goreleaser/internal/pipe/snapcraft"
	"github.com/weyfonk/goreleaser/pkg/context"
)

// Healthchecker should be implemented by pipes that want checks.
type Healthchecker interface {
	fmt.Stringer

	// Dependencies return the binaries of the dependencies needed.
	Dependencies(ctx *context.Context) []string
}

// Healthcheckers is the list of healthchekers.
// nolint: gochecknoglobals
var Healthcheckers = []Healthchecker{
	system{},
	snapcraft.Pipe{},
	sign.Pipe{},
	sign.DockerPipe{},
	sbom.Pipe{},
	docker.Pipe{},
	docker.ManifestPipe{},
	chocolatey.Pipe{},
	nix.NewPublish(),
}

type system struct{}

func (system) String() string                           { return "system" }
func (system) Dependencies(_ *context.Context) []string { return []string{"git", "go"} }
