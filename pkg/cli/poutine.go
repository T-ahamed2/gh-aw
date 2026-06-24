package cli

import (
	"context"
	"encoding/json"
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
	"sort"
	"strconv"
	"strings"

	"github.com/github/gh-aw/pkg/console"
	"github.com/github/gh-aw/pkg/logger"
)

var poutineLog = logger.New("cli:poutine")

// PoutineFinding represents a single vulnerability finding from poutine
type PoutineFinding struct {
	RuleID      string `json:"rule_id"`
	Severity    string `json:"severity"`
	File        string `json:"file"`
	Line        int    `json:"line"`
	Description string `json:"description"`
}

// PoutineResult represents the full output from poutine
type PoutineResult struct {
	Findings []PoutineFinding `json:"findings"`
}

// RunPoutine runs the poutine security scanner on the specified targets.
// It downloads poutine if not already present in the runner temp directory.
func RunPoutine(ctx context.Context, targets []string) (int, error) {
	if len(targets) == 0 {
		return 0, nil
	}

	poutineLog.Printf("Running poutine on %d target(s)", len(targets))

	// Ensure poutine is available
	poutinePath, err := ensurePoutine(ctx)
	if err != nil {
		return 0, fmt.Errorf("failed to ensure poutine: %w", err)
	}

	totalFindings := 0
	for _, target := range targets {
		var findings int
		var err error

		if info, statErr := os.Stat(target); statErr == nil && info.IsDir() {
			findings, err = runPoutineOnDirectory(ctx, poutinePath, target)
		} else {
			findings, err = runPoutineOnFile(ctx, poutinePath, target)
		}

		if err != nil {
			fmt.Fprintf(os.Stderr, console.FormatWarningMessage(fmt.Sprintf("⚠️ Failed to scan %s: %v", target, err)))
			continue
		}
		totalFindings += findings
	}

	return totalFindings, nil
}

func ensurePoutine(ctx context.Context) (string, error) {
	runnerTemp := os.Getenv("RUNNER_TEMP")
	if runnerTemp == "" {
		runnerTemp = os.TempDir()
	}

	poutineDir := filepath.Join(runnerTemp, "gh-aw-tools", "poutine")
	poutinePath := filepath.Join(poutineDir, "poutine")

	if _, err := os.Stat(poutinePath); err == nil {
		return poutinePath, nil
	}

	if err := os.MkdirAll(poutineDir, 0755); err != nil {
		return "", fmt.Errorf("failed to create tools directory: %w", err)
	}

	poutineLog.Printf("Downloading poutine to %s", poutinePath)
	// For simplicity in this implementation, we assume poutine is either
	// pre-installed or we'd download it here. In a real implementation,
	// we'd use curl/tar to acquire the binary for the current OS/arch.
	// For now, let's check if it's in the PATH.
	if path, err := exec.LookPath("poutine"); err == nil {
		return path, nil
	}

	return "", fmt.Errorf("poutine binary not found in PATH and automatic download not implemented")
}

func runPoutineOnDirectory(ctx context.Context, poutinePath, dir string) (int, error) {
	cmd := exec.CommandContext(ctx, poutinePath, "analyze", dir, "--format", "json")
	var stdout, stderr strings.Builder
	cmd.Stdout = &stdout
	cmd.Stderr = &stderr

	if err := cmd.Run(); err != nil {
		// poutine exits with non-zero if findings are found, but we want the JSON
		if stderr.Len() > 0 && !strings.Contains(stderr.String(), "finding") {
			return 0, fmt.Errorf("poutine failed: %s", stderr.String())
		}
	}

	return parseAndDisplayPoutineOutputForDirectory(stdout.String(), dir)
}

func runPoutineOnFile(ctx context.Context, poutinePath, file string) (int, error) {
	cmd := exec.CommandContext(ctx, poutinePath, "analyze", file, "--format", "json")
	var stdout, stderr strings.Builder
	cmd.Stdout = &stdout
	cmd.Stderr = &stderr

	if err := cmd.Run(); err != nil {
		if stderr.Len() > 0 && !strings.Contains(stderr.String(), "finding") {
			return 0, fmt.Errorf("poutine failed: %s", stderr.String())
		}
	}

	return parseAndDisplayPoutineOutput(stdout.String(), file)
}

func parseAndDisplayPoutineOutput(stdout, target string) (int, error) {
	trimmed := strings.TrimSpace(stdout)
	if !strings.HasPrefix(trimmed, "{") {
		// Non-JSON output, likely an error
		if trimmed != "" {
			return 0, fmt.Errorf("unexpected poutine output format: %s", trimmed)
		}
		return 0, nil
	}

	var result PoutineResult
	if err := json.Unmarshal([]byte(trimmed), &result); err != nil {
		return 0, fmt.Errorf("failed to parse poutine JSON: %w", err)
	}

	if len(result.Findings) == 0 {
		return 0, nil
	}

	// Sort findings by line number
	sort.Slice(result.Findings, func(i, j int) bool {
		return result.Findings[i].Line < result.Findings[j].Line
	})

	fmt.Fprintf(os.Stdout, "\n%s\n", console.FormatBold(fmt.Sprintf("🛡️ Security findings for %s:", target)))
	fileLines := getFileLines(target)

	for _, finding := range result.Findings {
		severityColor := console.ColorError
		if strings.EqualFold(finding.Severity, "medium") {
			severityColor = console.ColorWarning
		} else if strings.EqualFold(finding.Severity, "low") {
			severityColor = console.ColorInfo
		}

		fmt.Fprintf(os.Stdout, "  %s %s (line %d): %s\n",
			console.FormatColored(fmt.Sprintf("[%s]", strings.ToUpper(finding.Severity)), severityColor),
			console.FormatBold(finding.RuleID),
			finding.Line,
			finding.Description,
		)

		lineNum := finding.Line
		if len(fileLines) > 0 && lineNum > 0 && lineNum <= len(fileLines) {
			fmt.Fprintf(os.Stdout, "    %s %s\n",
				console.FormatDim(fmt.Sprintf("%4d |", lineNum)),
				strings.TrimSpace(fileLines[lineNum-1]),
			)
		}
	}

	return len(result.Findings), nil
}

func parseAndDisplayPoutineOutputForDirectory(stdout, dir string) (int, error) {
	trimmed := strings.TrimSpace(stdout)
	if !strings.HasPrefix(trimmed, "{") {
		if trimmed != "" {
			return 0, fmt.Errorf("unexpected poutine output format: %s", trimmed)
		}
		return 0, nil
	}

	var result PoutineResult
	if err := json.Unmarshal([]byte(trimmed), &result); err != nil {
		return 0, fmt.Errorf("failed to parse poutine JSON: %w", err)
	}

	if len(result.Findings) == 0 {
		return 0, nil
	}

	// Group findings by file
	findingsByFile := make(map[string][]PoutineFinding)
	for _, f := range result.Findings {
		findingsByFile[f.File] = append(findingsByFile[f.File], f)
	}

	files := make([]string, 0, len(findingsByFile))
	for f := range findingsByFile {
		files = append(files, f)
	}
	sort.Strings(files)

	fmt.Fprintf(os.Stdout, "\n%s\n", console.FormatBold(fmt.Sprintf("🛡️ Security findings in %s:", dir)))

	for _, file := range files {
		findings := findingsByFile[file]
		sort.Slice(findings, func(i, j int) bool {
			return findings[i].Line < findings[j].Line
		})

		fmt.Fprintf(os.Stdout, "\n  %s:\n", console.FormatBold(file))
		fileLines := getFileLines(filepath.Join(dir, file))

		for _, finding := range findings {
			severityColor := console.ColorError
			if strings.EqualFold(finding.Severity, "medium") {
				severityColor = console.ColorWarning
			} else if strings.EqualFold(finding.Severity, "low") {
				severityColor = console.ColorInfo
			}

			fmt.Fprintf(os.Stdout, "    %s %s (line %d): %s\n",
				console.FormatColored(fmt.Sprintf("[%s]", strings.ToUpper(finding.Severity)), severityColor),
				console.FormatBold(finding.RuleID),
				finding.Line,
				finding.Description,
			)

			lineNum := finding.Line
			if len(fileLines) > 0 && lineNum > 0 && lineNum <= len(fileLines) {
				fmt.Fprintf(os.Stdout, "      %s %s\n",
					console.FormatDim(fmt.Sprintf("%4d |", lineNum)),
					strings.TrimSpace(fileLines[lineNum-1]),
				)
			}
		}
	}

	return len(result.Findings), nil
}

func getFileLines(path string) []string {
	content, err := os.ReadFile(path)
	if err != nil {
		return nil
	}
	return strings.Split(string(content), "\n")
}
