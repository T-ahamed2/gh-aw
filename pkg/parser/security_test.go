package parser

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestGitArgumentInjection(t *testing.T) {
	// Verify that various entry points reject hyphen-prefixed arguments
	// to prevent flag injection vulnerabilities.

	t.Run("getOrCreateListRepoClone rejects hyphenated owner", func(t *testing.T) {
		_, err := getOrCreateListRepoClone("-v", "repo", "main", "")
		require.Error(t, err)
		assert.Contains(t, err.Error(), "arguments cannot start with a hyphen")
	})

	t.Run("getOrCreateListRepoClone rejects hyphenated repo", func(t *testing.T) {
		_, err := getOrCreateListRepoClone("owner", "-v", "main", "")
		require.Error(t, err)
		assert.Contains(t, err.Error(), "arguments cannot start with a hyphen")
	})

	t.Run("getOrCreateListRepoClone rejects hyphenated ref", func(t *testing.T) {
		_, err := getOrCreateListRepoClone("owner", "repo", "-v", "")
		require.Error(t, err)
		assert.Contains(t, err.Error(), "arguments cannot start with a hyphen")
	})

	t.Run("resolveRefToSHAViaGit rejects hyphenated ref", func(t *testing.T) {
		_, err := resolveRefToSHAViaGit("owner", "repo", "-v", "")
		require.Error(t, err)
		assert.Contains(t, err.Error(), "arguments cannot start with a hyphen")
	})
}
