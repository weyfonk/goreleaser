// Package announce contains the announcing pipe.
package announce

import (
	"fmt"

	"github.com/weyfonk/goreleaser/internal/middleware/errhandler"
	"github.com/weyfonk/goreleaser/internal/middleware/logging"
	"github.com/weyfonk/goreleaser/internal/middleware/skip"
	"github.com/weyfonk/goreleaser/internal/pipe/discord"
	"github.com/weyfonk/goreleaser/internal/pipe/linkedin"
	"github.com/weyfonk/goreleaser/internal/pipe/mastodon"
	"github.com/weyfonk/goreleaser/internal/pipe/mattermost"
	"github.com/weyfonk/goreleaser/internal/pipe/opencollective"
	"github.com/weyfonk/goreleaser/internal/pipe/reddit"
	"github.com/weyfonk/goreleaser/internal/pipe/slack"
	"github.com/weyfonk/goreleaser/internal/pipe/smtp"
	"github.com/weyfonk/goreleaser/internal/pipe/teams"
	"github.com/weyfonk/goreleaser/internal/pipe/telegram"
	"github.com/weyfonk/goreleaser/internal/pipe/twitter"
	"github.com/weyfonk/goreleaser/internal/pipe/webhook"
	"github.com/weyfonk/goreleaser/internal/skips"
	"github.com/weyfonk/goreleaser/internal/tmpl"
	"github.com/weyfonk/goreleaser/pkg/context"
)

// Announcer should be implemented by pipes that want to announce releases.
type Announcer interface {
	fmt.Stringer
	Announce(ctx *context.Context) error
}

// nolint: gochecknoglobals
var announcers = []Announcer{
	// XXX: keep asc sorting
	discord.Pipe{},
	linkedin.Pipe{},
	mastodon.Pipe{},
	mattermost.Pipe{},
	opencollective.Pipe{},
	reddit.Pipe{},
	slack.Pipe{},
	smtp.Pipe{},
	teams.Pipe{},
	telegram.Pipe{},
	twitter.Pipe{},
	webhook.Pipe{},
}

// Pipe that announces releases.
type Pipe struct{}

func (Pipe) String() string { return "announcing" }

func (Pipe) Skip(ctx *context.Context) (bool, error) {
	if skips.Any(ctx, skips.Announce) {
		return true, nil
	}
	return tmpl.New(ctx).Bool(ctx.Config.Announce.Skip)
}

// Run the pipe.
func (Pipe) Run(ctx *context.Context) error {
	memo := errhandler.Memo{}
	for _, announcer := range announcers {
		_ = skip.Maybe(
			announcer,
			logging.PadLog(announcer.String(), memo.Wrap(announcer.Announce)),
		)(ctx)
	}
	if memo.Error() != nil {
		return fmt.Errorf("failed to announce release: %w", memo.Error())
	}
	return nil
}
