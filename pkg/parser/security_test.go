package parser

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestGitArgumentInjection(t *testing.T) {
	tests := []struct {
		name    string
		owner   string
		repo    string
		path    string
		ref     string
		wantErr bool
	}{
		{
			name:    "valid inputs",
			owner:   "owner",
			repo:    "repo",
			path:    "path/to/file.md",
			ref:     "main",
			wantErr: false,
		},
		{
			name:    "malicious owner",
			owner:   "-v",
			repo:    "repo",
			path:    "path/to/file.md",
			ref:     "main",
			wantErr: true,
		},
		{
			name:    "malicious repo",
			owner:   "owner",
			repo:    "--version",
			path:    "path/to/file.md",
			ref:     "main",
			wantErr: true,
		},
		{
			name:    "malicious path",
			owner:   "owner",
			repo:    "repo",
			path:    "-oProxyCommand=calc.exe",
			ref:     "main",
			wantErr: true,
		},
		{
			name:    "malicious ref",
			owner:   "owner",
			repo:    "repo",
			path:    "path/to/file.md",
			ref:     "-v",
			wantErr: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			_, err := DownloadFileFromGitHub(tt.owner, tt.repo, tt.path, tt.ref)
			if tt.wantErr {
				require.Error(t, err)
				assert.Contains(t, err.Error(), "flag injection")
			} else {
				// We expect a different error (e.g. auth or network) but NOT a validation error
				if err != nil {
					assert.NotContains(t, err.Error(), "flag injection")
				}
			}
		})
	}
}

func TestWorkflowSpecInjection(t *testing.T) {
	tests := []struct {
		name    string
		spec    string
		wantErr bool
	}{
		{
			name:    "valid spec",
			spec:    "owner/repo/path/to/file.md@main",
			wantErr: false,
		},
		{
			name:    "malicious owner in spec",
			spec:    "-v/repo/path/to/file.md@main",
			wantErr: true,
		},
		{
			name:    "malicious repo in spec",
			spec:    "owner/--version/path/to/file.md@main",
			wantErr: true,
		},
		{
			name:    "malicious path in spec",
			spec:    "owner/repo/-v@main",
			wantErr: true,
		},
		{
			name:    "malicious ref in spec",
			spec:    "owner/repo/path/to/file.md@-v",
			wantErr: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			_, _, _, _, err := parseWorkflowSpecParts(tt.spec)
			if tt.wantErr {
				require.Error(t, err)
				assert.Contains(t, err.Error(), "flag injection")
			} else {
				assert.NoError(t, err)
			}
		})
	}
}
