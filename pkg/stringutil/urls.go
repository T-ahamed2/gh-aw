package stringutil

import (
	"net/url"
	"strings"

	"github.com/github/gh-aw/pkg/logger"
)

var urlsLog = logger.New("stringutil:urls")

// NormalizeGitHubHostURL ensures the host URL has a scheme (defaulting to https://) and no trailing slashes.
// It is safe to call with URLs that already have an http:// or https:// scheme.
func NormalizeGitHubHostURL(rawHostURL string) string {
	// Remove all trailing slashes
	normalized := strings.TrimRight(rawHostURL, "/")

	// Add https:// scheme if no scheme is present
	if !strings.HasPrefix(normalized, "https://") && !strings.HasPrefix(normalized, "http://") {
		normalized = "https://" + normalized
	}

	return normalized
}

// ExtractDomainFromURL extracts the domain name from a URL string.
// Handles various URL formats including full URLs with protocols, URLs with ports,
// and plain domain names.
//
// This function uses net/url.Parse for proper URL parsing when a protocol is present,
// and falls back to string manipulation for other formats.
//
// Examples:
//
//	ExtractDomainFromURL("https://mcp.tavily.com/mcp/")           // returns "mcp.tavily.com"
//	ExtractDomainFromURL("http://api.example.com:8080/path")      // returns "api.example.com"
//	ExtractDomainFromURL("mcp.example.com")                       // returns "mcp.example.com"
//	ExtractDomainFromURL("github.com:443")                        // returns "github.com"
//	ExtractDomainFromURL("http://sub.domain.com:8080/path")       // returns "sub.domain.com"
//	ExtractDomainFromURL("localhost:8080")                        // returns "localhost"
func ExtractDomainFromURL(urlStr string) string {
	if urlsLog.Enabled() {
		urlsLog.Printf("Extracting domain from URL: %s", urlStr)
	}

	// High-performance fast-path for standard http:// and https:// URLs.
	// This avoids net/url.Parse and results in zero heap allocations.
	var rem string
	var hasScheme bool
	if strings.HasPrefix(urlStr, "https://") {
		rem = urlStr[8:]
		hasScheme = true
	} else if strings.HasPrefix(urlStr, "http://") {
		rem = urlStr[7:]
		hasScheme = true
	}

	if hasScheme {
		// If there is userinfo (contains @) or IPv6 address (contains [), fall back to net/url.Parse.
		if !strings.Contains(rem, "@") && !strings.Contains(rem, "[") {
			// Find the first delimiter that signals the end of the host/port part.
			if idx := strings.IndexAny(rem, "/?#:"); idx != -1 {
				return rem[:idx]
			}
			return rem
		}

		// Parse full URL for complex formats (userinfo, IPv6)
		parsedURL, err := url.Parse(urlStr)
		if err != nil {
			// Fall back to string manipulation if parsing fails
			if urlsLog.Enabled() {
				urlsLog.Printf("URL parse failed, using fallback: %v", err)
			}
			return extractDomainFallback(urlStr)
		}
		return parsedURL.Hostname()
	}

	// For URLs without protocol, use string manipulation
	return extractDomainFallback(urlStr)
}

// extractDomainFallback extracts domain using string manipulation.
// This handles URLs without protocols, CONNECT requests (domain:port format),
// and plain domain names.
func extractDomainFallback(urlStr string) string {
	// Remove protocol if present (in case it wasn't http/https)
	urlStr = strings.TrimPrefix(urlStr, "https://")
	urlStr = strings.TrimPrefix(urlStr, "http://")

	// Remove port and path
	if idx := strings.Index(urlStr, ":"); idx != -1 {
		urlStr = urlStr[:idx]
	}
	if idx := strings.Index(urlStr, "/"); idx != -1 {
		urlStr = urlStr[:idx]
	}

	return strings.TrimSpace(urlStr)
}
