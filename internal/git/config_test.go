package git_test

import (
	"context"
	"strings"
	"testing"

	"github.com/weyfonk/goreleaser/internal/git"
	"github.com/weyfonk/goreleaser/internal/testlib"
	"github.com/stretchr/testify/require"
)

func TestNotARepo(t *testing.T) {
	testlib.Mktmp(t)
	_, err := git.ExtractRepoFromConfig(context.Background())
	require.EqualError(t, err, `current folder is not a git repository`)
}

func TestNoRemote(t *testing.T) {
	testlib.Mktmp(t)
	testlib.GitInit(t)
	_, err := git.ExtractRepoFromConfig(context.Background())
	require.EqualError(t, err, `no remote configured to list refs from`)
}

func TestRelativeRemote(t *testing.T) {
	ctx := context.Background()
	testlib.Mktmp(t)
	testlib.GitInit(t)
	testlib.GitRemoteAddWithName(t, "upstream", "https://github.com/weyfonk/goreleaser.git")
	_, err := git.Run(ctx, "pull", "upstream", "main")
	require.NoError(t, err)
	_, err = git.Run(ctx, "branch", "--set-upstream-to", "upstream/main")
	require.NoError(t, err)
	_, err = git.Run(ctx, "checkout", "--track", "-b", "relative_branch")
	require.NoError(t, err)
	gitCfg, err := git.Run(ctx, "config", "--local", "--list")
	require.NoError(t, err)
	require.True(t, strings.Contains(gitCfg, "branch.relative_branch.remote=."))
	repo, err := git.ExtractRepoFromConfig(ctx)
	require.NoError(t, err)
	require.Equal(t, "weyfonk/goreleaser", repo.String())
}

func TestRepoName(t *testing.T) {
	testlib.Mktmp(t)
	testlib.GitInit(t)
	testlib.GitRemoteAdd(t, "git@github.com:weyfonk/goreleaser.git")
	repo, err := git.ExtractRepoFromConfig(context.Background())
	require.NoError(t, err)
	require.Equal(t, "weyfonk/goreleaser", repo.String())
}

func TestRepoNameWithDifferentRemote(t *testing.T) {
	ctx := context.Background()
	testlib.Mktmp(t)
	testlib.GitInit(t)
	testlib.GitRemoteAddWithName(t, "upstream", "https://github.com/weyfonk/goreleaser.git")
	_, err := git.Run(ctx, "pull", "upstream", "main")
	require.NoError(t, err)
	_, err = git.Run(ctx, "branch", "--set-upstream-to", "upstream/main")
	require.NoError(t, err)
	repo, err := git.ExtractRepoFromConfig(ctx)
	require.NoError(t, err)
	require.Equal(t, "weyfonk/goreleaser", repo.String())
}

func TestExtractRepoFromURL(t *testing.T) {
	// valid urls
	for _, url := range []string{
		"git@github.com:weyfonk/goreleaser.git",
		"git@custom:weyfonk/goreleaser.git",
		"https://foo@github.com/weyfonk/goreleaser",
		"https://github.com/weyfonk/goreleaser.git",
		"https://something.with.port:8080/weyfonk/goreleaser.git",
		"https://github.enterprise.com/weyfonk/goreleaser.git",
		"https://gitlab-ci-token:SOME_TOKEN@gitlab.yourcompany.com/weyfonk/goreleaser.git",
	} {
		t.Run(url, func(t *testing.T) {
			repo, err := git.ExtractRepoFromURL(url)
			require.NoError(t, err)
			require.Equal(t, "weyfonk/goreleaser", repo.String())
			require.NoError(t, repo.CheckSCM())
			require.Equal(t, url, repo.RawURL)
		})
	}

	// nested urls
	for _, url := range []string{
		"git@custom:group/nested/weyfonk/goreleaser.git",
		"https://gitlab.mycompany.com/group/nested/weyfonk/goreleaser.git",
		"https://gitlab-ci-token:SOME_TOKEN@gitlab.yourcompany.com/group/nested/weyfonk/goreleaser.git",
	} {
		t.Run(url, func(t *testing.T) {
			repo, err := git.ExtractRepoFromURL(url)
			require.NoError(t, err)
			require.Equal(t, "group/nested/weyfonk/goreleaser", repo.String())
			require.NoError(t, repo.CheckSCM())
			require.Equal(t, url, repo.RawURL)
		})
	}

	for _, url := range []string{
		"git@gist.github.com:someid.git",
		"https://gist.github.com/someid.git",
	} {
		t.Run(url, func(t *testing.T) {
			repo, err := git.ExtractRepoFromURL(url)
			require.NoError(t, err)
			require.Equal(t, "someid", repo.String())
			require.Error(t, repo.CheckSCM())
			require.Equal(t, url, repo.RawURL)
		})
	}

	// invalid urls
	for _, url := range []string{
		"git@gist.github.com:",
		"https://gist.github.com/",
	} {
		t.Run(url, func(t *testing.T) {
			repo, err := git.ExtractRepoFromURL(url)
			require.EqualError(t, err, "unsupported repository URL: "+url)
			require.Equal(t, "", repo.String())
			require.Error(t, repo.CheckSCM())
			require.Equal(t, url, repo.RawURL)
		})
	}
}
