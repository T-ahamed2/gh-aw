package parser

import (
	"context"
	"strings"
	"testing"
)

func TestGitArgumentInjection(t *testing.T) {
	badArgs := []string{"-v", "--version", "--upload-pack=touch /tmp/pwned"}

	t.Run("resolveRefToSHAViaGit", func(t *testing.T) {
		for _, arg := range badArgs {
			_, err := resolveRefToSHAViaGit("owner", "repo", arg, "")
			if err == nil || !strings.Contains(err.Error(), "security: invalid git argument") {
				t.Errorf("Expected security error for arg %q, but got %v", arg, err)
			}
		}
	})

	t.Run("getOrCreateListRepoClone", func(t *testing.T) {
		for _, arg := range badArgs {
			_, err := getOrCreateListRepoClone("owner", "repo", arg, "")
			if err == nil || !strings.Contains(err.Error(), "security: invalid git argument") {
				t.Errorf("Expected security error for arg %q, but got %v", arg, err)
			}
		}
	})

	t.Run("downloadFileViaGit", func(t *testing.T) {
		for _, arg := range badArgs {
			_, err := downloadFileViaGit(context.Background(), "owner", "repo", "path", arg, "")
			if err == nil || !strings.Contains(err.Error(), "security: invalid git argument") {
				t.Errorf("Expected security error for ref %q, but got %v", arg, err)
			}
			_, err = downloadFileViaGit(context.Background(), "owner", "repo", arg, "ref", "")
			if err == nil || !strings.Contains(err.Error(), "security: invalid git argument") {
				t.Errorf("Expected security error for path %q, but got %v", arg, err)
			}
		}
	})

	t.Run("downloadFileViaGitClone", func(t *testing.T) {
		for _, arg := range badArgs {
			_, err := downloadFileViaGitClone("owner", "repo", "path", arg, "")
			if err == nil || !strings.Contains(err.Error(), "security: invalid git argument") {
				t.Errorf("Expected security error for ref %q, but got %v", arg, err)
			}
			_, err = downloadFileViaGitClone("owner", "repo", arg, "ref", "")
			if err == nil || !strings.Contains(err.Error(), "security: invalid git argument") {
				t.Errorf("Expected security error for path %q, but got %v", arg, err)
			}
		}
	})

	t.Run("listDirAllFilesViaGitForHost", func(t *testing.T) {
		for _, arg := range badArgs {
			_, err := listDirAllFilesViaGitForHost("owner", "repo", "ref", arg, "")
			if err == nil || !strings.Contains(err.Error(), "security: invalid git argument") {
				t.Errorf("Expected security error for dirPath %q, but got %v", arg, err)
			}
		}
	})

	t.Run("listWorkflowFilesViaGitForHost", func(t *testing.T) {
		for _, arg := range badArgs {
			_, err := listWorkflowFilesViaGitForHost("owner", "repo", arg, "path", "")
			if err == nil || !strings.Contains(err.Error(), "security: invalid git argument") {
				t.Errorf("Expected security error for ref %q, but got %v", arg, err)
			}
			_, err = listWorkflowFilesViaGitForHost("owner", "repo", "ref", arg, "")
			if err == nil || !strings.Contains(err.Error(), "security: invalid git argument") {
				t.Errorf("Expected security error for workflowPath %q, but got %v", arg, err)
			}
		}
	})
}
