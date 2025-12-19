package core

import "strings"

// This function filters out JavaScript endpoints from a list of URLs.
func filterJsEndpoints(urls []string) []string {
	filteredJsEndpoints := make([]string, 0, len(urls))
	for _, url := range urls {
		if !isJsEndpoint(url) {
			filteredJsEndpoints = append(filteredJsEndpoints, url)
		}
	}
	return filteredJsEndpoints
}

// isJsEndpoint checks if the given URL is a JavaScript endpoint.
func isJsEndpoint(url string) bool {
	return strings.HasSuffix(strings.Split(url, "?")[0], ".js")
}
