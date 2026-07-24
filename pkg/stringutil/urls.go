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

	// Fast path for standard http:// and https:// URLs to avoid net/url.Parse overhead
	var rest string
	hasProtocol := false
	if strings.HasPrefix(urlStr, "https://") {
		rest = urlStr[8:]
		hasProtocol = true
	} else if strings.HasPrefix(urlStr, "http://") {
		rest = urlStr[7:]
		hasProtocol = true
	}

	if hasProtocol {
		// Find end of host part (before path, query, or fragment separator)
		hostEnd := len(rest)
		for i := range len(rest) {
			c := rest[i]
			if c == '/' || c == '?' || c == '#' {
				hostEnd = i
				break
			}
		}
		hostPart := rest[:hostEnd]

		// Fallback to url.Parse for complex hosts (e.g. userinfo or IPv6)
		if strings.ContainsAny(hostPart, "@[") {
			parsedURL, err := url.Parse(urlStr)
			if err != nil {
				if urlsLog.Enabled() {
					urlsLog.Printf("URL parse failed, using fallback: %v", err)
				}
				return extractDomainFallback(urlStr)
			}
			return parsedURL.Hostname()
		}

		// Strip port if present
		if idx := strings.IndexByte(hostPart, ':'); idx != -1 {
			hostPart = hostPart[:idx]
		}
		return hostPart
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
