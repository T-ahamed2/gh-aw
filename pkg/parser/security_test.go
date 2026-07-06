package parser

import (
	"context"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGitArgumentInjection(t *testing.T) {
	ctx := context.Background()

	t.Run("resolveRefToSHAViaGit rejects hyphenated arguments", func(t *testing.T) {
		_, err := resolveRefToSHAViaGit("-oProxyCommand=touch%20/tmp/pwn", "repo", "ref", "")
		assert.Error(t, err)
		assert.Contains(t, err.Error(), "invalid git argument")
	})

	t.Run("resolveRefToSHA rejects hyphenated arguments", func(t *testing.T) {
		_, err := resolveRefToSHA("owner", "-oProxyCommand=touch%20/tmp/pwn", "ref", "")
		assert.Error(t, err)
		assert.Contains(t, err.Error(), "invalid git argument")
	})

	t.Run("downloadFileViaGit rejects hyphenated arguments", func(t *testing.T) {
		_, err := downloadFileViaGit(ctx, "owner", "repo", "path", "-oProxyCommand=touch%20/tmp/pwn", "")
		assert.Error(t, err)
		assert.Contains(t, err.Error(), "invalid git argument")
	})

	t.Run("downloadFileViaGitClone rejects hyphenated arguments", func(t *testing.T) {
		_, err := downloadFileViaGitClone("owner", "repo", "-oProxyCommand=touch%20/tmp/pwn", "ref", "")
		assert.Error(t, err)
		assert.Contains(t, err.Error(), "invalid git argument")
	})

	t.Run("listDirAllFilesViaGitForHost rejects hyphenated arguments", func(t *testing.T) {
		_, err := listDirAllFilesViaGitForHost("owner", "repo", "ref", "dir", "-oProxyCommand=touch%20/tmp/pwn")
		assert.Error(t, err)
		assert.Contains(t, err.Error(), "invalid git argument")
	})
}
