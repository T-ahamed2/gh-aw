package parser

import (
	"context"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestGitArgumentInjection(t *testing.T) {
	// These tests verify that hyphen-prefixed arguments are rejected by the
	// hardened functions in remote_fetch.go.

	t.Run("getOrCreateListRepoClone rejects hyphenated ref", func(t *testing.T) {
		_, err := getOrCreateListRepoClone("owner", "repo", "-v", "")
		require.Error(t, err)
		require.Contains(t, err.Error(), "cannot start with a hyphen")
	})

	t.Run("resolveRefToSHAViaGit rejects hyphenated owner", func(t *testing.T) {
		_, err := resolveRefToSHAViaGit("-owner", "repo", "main", "")
		require.Error(t, err)
		require.Contains(t, err.Error(), "cannot start with a hyphen")
	})

	t.Run("downloadFileViaGit rejects hyphenated path", func(t *testing.T) {
		_, err := downloadFileViaGit(context.Background(), "owner", "repo", "--help", "main", "")
		require.Error(t, err)
		require.Contains(t, err.Error(), "cannot start with a hyphen")
	})

	t.Run("downloadFileViaGitClone rejects hyphenated ref", func(t *testing.T) {
		_, err := downloadFileViaGitClone("owner", "repo", "path/to/file", "--upload-pack", "")
		require.Error(t, err)
		require.Contains(t, err.Error(), "cannot start with a hyphen")
	})

	t.Run("listWorkflowFilesViaGitForHost rejects hyphenated workflowPath", func(t *testing.T) {
		_, err := listWorkflowFilesViaGitForHost("owner", "repo", "main", "-v", "")
		require.Error(t, err)
		require.Contains(t, err.Error(), "cannot start with a hyphen")
	})
}
