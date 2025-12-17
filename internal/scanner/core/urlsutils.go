package core

import "strings"

func filterJsEndpoints(urls []string) []string {
	jsEndpoints := make([]string, 0, len(urls))
	for _, url := range urls {
		if isJsEndpoint(url) {
			jsEndpoints = append(jsEndpoints, url)
		}
	}
	return jsEndpoints
}

func isJsEndpoint(url string) bool {
	return !strings.HasSuffix(strings.Split(url, "?")[0], ".js")
}
