//go:build !integration

package parser

import (
	"strings"
	"testing"
)

func TestBuildCommitLookupAPIPath(t *testing.T) {
	t.Run("escapes refs containing slash", func(t *testing.T) {
		got := buildCommitLookupAPIPath("owner", "repo", "feature/github-agentic-workflows")
		want := "/repos/owner/repo/commits/feature%2Fgithub-agentic-workflows"
		if got != want {
			t.Fatalf("buildCommitLookupAPIPath() = %q, want %q", got, want)
		}
	})

	t.Run("keeps plain refs readable", func(t *testing.T) {
		got := buildCommitLookupAPIPath("owner", "repo", "main")
		want := "/repos/owner/repo/commits/main"
		if got != want {
			t.Fatalf("buildCommitLookupAPIPath() = %q, want %q", got, want)
		}
	})
}

func TestBuildContentsAPIPath(t *testing.T) {
	t.Run("escapes refs with reserved query chars", func(t *testing.T) {
		got := buildContentsAPIPath("owner", "repo", ".github/workflows/demo.md", "release+candidate#1")
		want := "repos/owner/repo/contents/.github/workflows/demo.md?ref=release%2Bcandidate%231"
		if got != want {
			t.Fatalf("buildContentsAPIPath() = %q, want %q", got, want)
		}
	})

	t.Run("keeps plain refs readable", func(t *testing.T) {
		got := buildContentsAPIPath("owner", "repo", ".github/workflows/demo.md", "main")
		want := "repos/owner/repo/contents/.github/workflows/demo.md?ref=main"
		if got != want {
			t.Fatalf("buildContentsAPIPath() = %q, want %q", got, want)
		}
	})

	t.Run("escapes path segments with reserved chars", func(t *testing.T) {
		got := buildContentsAPIPath("owner", "repo", "skills/path with spaces/file#100%.md", "main")
		want := "repos/owner/repo/contents/skills/path%20with%20spaces/file%23100%25.md?ref=main"
		if got != want {
			t.Fatalf("buildContentsAPIPath() = %q, want %q", got, want)
		}
	})
}

func TestGitFallbackRequiresNonEmptyRef(t *testing.T) {
	t.Run("all files fallback validates ref", func(t *testing.T) {
		_, err := listDirAllFilesViaGitForHost("owner", "repo", "", "skills/demo", "")
		if err == nil {
			t.Fatal("expected error for empty ref")
		}
		if !strings.Contains(err.Error(), "non-empty ref") {
			t.Fatalf("expected non-empty ref error, got %q", err)
		}
	})

	t.Run("subdirs fallback validates ref", func(t *testing.T) {
		_, err := listDirSubdirsViaGitForHost("owner", "repo", "   ", "skills", "")
		if err == nil {
			t.Fatal("expected error for empty ref")
		}
		if !strings.Contains(err.Error(), "non-empty ref") {
			t.Fatalf("expected non-empty ref error, got %q", err)
		}
	})
}

func TestListContentsRecursivelyWithDepth_MaxDepthGuard(t *testing.T) {
	_, err := listContentsRecursivelyWithDepth(nil, "owner", "repo", "main", "skills/demo/deep", 11, 10)
	if err == nil {
		t.Fatal("expected depth limit error")
	}
	if !strings.Contains(err.Error(), "maximum skill directory recursion depth exceeded") {
		t.Fatalf("expected depth limit error, got %q", err)
	}
}

func TestGitArgumentInjectionProtection(t *testing.T) {
	t.Run("resolveRefToSHAViaGit_rejects_hyphen_ref", func(t *testing.T) {
		_, err := resolveRefToSHAViaGit("owner", "repo", "-v", "")
		if err == nil || !strings.Contains(err.Error(), "must not start with '-'") {
			t.Fatalf("expected validation error for hyphenated ref, got %v", err)
		}
	})

	t.Run("resolveRefToSHA_rejects_hyphen_ref", func(t *testing.T) {
		_, err := resolveRefToSHA("owner", "repo", "-v", "")
		if err == nil || !strings.Contains(err.Error(), "must not start with '-'") {
			t.Fatalf("expected validation error for hyphenated ref, got %v", err)
		}
	})

	t.Run("downloadFileViaGit_rejects_hyphen_ref", func(t *testing.T) {
		_, err := downloadFileViaGit(nil, "owner", "repo", "file.md", "-v", "")
		if err == nil || !strings.Contains(err.Error(), "must not start with '-'") {
			t.Fatalf("expected validation error for hyphenated ref, got %v", err)
		}
	})

	t.Run("downloadFileViaGit_rejects_hyphen_path", func(t *testing.T) {
		_, err := downloadFileViaGit(nil, "owner", "repo", "-path", "main", "")
		if err == nil || !strings.Contains(err.Error(), "must not start with '-'") {
			t.Fatalf("expected validation error for hyphenated path, got %v", err)
		}
	})

	t.Run("downloadFileViaGitClone_rejects_hyphen_ref", func(t *testing.T) {
		// Validating hyphenated ref rejection in git checkout path
		_, err := downloadFileViaGitClone("owner", "repo", "file.md", "-invalid-ref", "")
		if err == nil || !strings.Contains(err.Error(), "must not start with '-'") {
			t.Fatalf("expected validation error, got %v", err)
		}
	})

	t.Run("downloadFileViaGitClone_rejects_hyphen_path", func(t *testing.T) {
		_, err := downloadFileViaGitClone("owner", "repo", "-path", "main", "")
		if err == nil || !strings.Contains(err.Error(), "must not start with '-'") {
			t.Fatalf("expected validation error, got %v", err)
		}
	})
}
