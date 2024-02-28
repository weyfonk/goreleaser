// Package defaults make the list of Defaulter implementations available
// so projects extending GoReleaser are able to use it, namely, GoDownloader.
package defaults

import (
	"fmt"

	"github.com/weyfonk/goreleaser/internal/pipe/archive"
	"github.com/weyfonk/goreleaser/internal/pipe/artifactory"
	"github.com/weyfonk/goreleaser/internal/pipe/aur"
	"github.com/weyfonk/goreleaser/internal/pipe/blob"
	"github.com/weyfonk/goreleaser/internal/pipe/brew"
	"github.com/weyfonk/goreleaser/internal/pipe/build"
	"github.com/weyfonk/goreleaser/internal/pipe/checksums"
	"github.com/weyfonk/goreleaser/internal/pipe/chocolatey"
	"github.com/weyfonk/goreleaser/internal/pipe/discord"
	"github.com/weyfonk/goreleaser/internal/pipe/docker"
	"github.com/weyfonk/goreleaser/internal/pipe/gomod"
	"github.com/weyfonk/goreleaser/internal/pipe/ko"
	"github.com/weyfonk/goreleaser/internal/pipe/krew"
	"github.com/weyfonk/goreleaser/internal/pipe/linkedin"
	"github.com/weyfonk/goreleaser/internal/pipe/mastodon"
	"github.com/weyfonk/goreleaser/internal/pipe/mattermost"
	"github.com/weyfonk/goreleaser/internal/pipe/milestone"
	"github.com/weyfonk/goreleaser/internal/pipe/nfpm"
	"github.com/weyfonk/goreleaser/internal/pipe/nix"
	"github.com/weyfonk/goreleaser/internal/pipe/opencollective"
	"github.com/weyfonk/goreleaser/internal/pipe/project"
	"github.com/weyfonk/goreleaser/internal/pipe/reddit"
	"github.com/weyfonk/goreleaser/internal/pipe/release"
	"github.com/weyfonk/goreleaser/internal/pipe/sbom"
	"github.com/weyfonk/goreleaser/internal/pipe/scoop"
	"github.com/weyfonk/goreleaser/internal/pipe/sign"
	"github.com/weyfonk/goreleaser/internal/pipe/slack"
	"github.com/weyfonk/goreleaser/internal/pipe/smtp"
	"github.com/weyfonk/goreleaser/internal/pipe/snapcraft"
	"github.com/weyfonk/goreleaser/internal/pipe/snapshot"
	"github.com/weyfonk/goreleaser/internal/pipe/sourcearchive"
	"github.com/weyfonk/goreleaser/internal/pipe/teams"
	"github.com/weyfonk/goreleaser/internal/pipe/telegram"
	"github.com/weyfonk/goreleaser/internal/pipe/twitter"
	"github.com/weyfonk/goreleaser/internal/pipe/universalbinary"
	"github.com/weyfonk/goreleaser/internal/pipe/upload"
	"github.com/weyfonk/goreleaser/internal/pipe/upx"
	"github.com/weyfonk/goreleaser/internal/pipe/webhook"
	"github.com/weyfonk/goreleaser/internal/pipe/winget"
	"github.com/weyfonk/goreleaser/pkg/context"
)

// Defaulter can be implemented by a Piper to set default values for its
// configuration.
type Defaulter interface {
	fmt.Stringer

	// Default sets the configuration defaults
	Default(ctx *context.Context) error
}

// Defaulters is the list of defaulters.
// nolint: gochecknoglobals
var Defaulters = []Defaulter{
	snapshot.Pipe{},
	release.Pipe{},
	project.Pipe{},
	gomod.Pipe{},
	build.Pipe{},
	universalbinary.Pipe{},
	upx.Pipe{},
	sourcearchive.Pipe{},
	archive.Pipe{},
	nfpm.Pipe{},
	snapcraft.Pipe{},
	checksums.Pipe{},
	sign.Pipe{},
	sign.DockerPipe{},
	sbom.Pipe{},
	docker.Pipe{},
	docker.ManifestPipe{},
	artifactory.Pipe{},
	blob.Pipe{},
	upload.Pipe{},
	aur.Pipe{},
	nix.Pipe{},
	winget.Pipe{},
	brew.Pipe{},
	krew.Pipe{},
	ko.Pipe{},
	scoop.Pipe{},
	discord.Pipe{},
	reddit.Pipe{},
	slack.Pipe{},
	teams.Pipe{},
	twitter.Pipe{},
	smtp.Pipe{},
	mastodon.Pipe{},
	mattermost.Pipe{},
	milestone.Pipe{},
	linkedin.Pipe{},
	telegram.Pipe{},
	webhook.Pipe{},
	chocolatey.Pipe{},
	opencollective.Pipe{},
}
