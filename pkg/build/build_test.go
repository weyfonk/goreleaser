package build

import (
	"testing"

	"github.com/weyfonk/goreleaser/pkg/config"
	"github.com/weyfonk/goreleaser/pkg/context"
	"github.com/stretchr/testify/require"
)

type dummy struct{}

func (*dummy) WithDefaults(build config.Build) (config.Build, error) {
	return build, nil
}

func (*dummy) Build(_ *context.Context, _ config.Build, _ Options) error {
	return nil
}

func TestRegisterAndGet(t *testing.T) {
	builder := &dummy{}
	Register("dummy", builder)
	require.Equal(t, builder, For("dummy"))
}
