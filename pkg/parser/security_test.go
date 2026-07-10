package parser

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGitArgumentInjection(t *testing.T) {
	// Verify that various entry points reject hyphen-prefixed arguments
	// to prevent flag injection vulnerabilities.

	t.Run("getOrCreateListRepoClone rejects hyphenated owner", func(t *testing.T) {
		_, err := getOrCreateListRepoClone("-v", "repo", "main", "")
		assert.Error(t, err)
		assert.Contains(t, err.Error(), "arguments cannot start with a hyphen")
	})

	t.Run("getOrCreateListRepoClone rejects hyphenated repo", func(t *testing.T) {
		_, err := getOrCreateListRepoClone("owner", "-v", "main", "")
		assert.Error(t, err)
		assert.Contains(t, err.Error(), "arguments cannot start with a hyphen")
	})

	t.Run("getOrCreateListRepoClone rejects hyphenated ref", func(t *testing.T) {
		_, err := getOrCreateListRepoClone("owner", "repo", "-v", "")
		assert.Error(t, err)
		assert.Contains(t, err.Error(), "arguments cannot start with a hyphen")
	})

	t.Run("resolveRefToSHAViaGit rejects hyphenated ref", func(t *testing.T) {
		_, err := resolveRefToSHAViaGit("owner", "repo", "-v", "")
		assert.Error(t, err)
		assert.Contains(t, err.Error(), "arguments cannot start with a hyphen")
	})
}
