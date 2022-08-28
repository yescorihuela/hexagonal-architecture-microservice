package common

import "strings"

func GetSlicedUrls(urls string) []string {
	slicedUrls := strings.Split(urls, ",")
	aux := make([]string, 0)
	if len(slicedUrls) > 0 {
		for _, url := range slicedUrls {
			if strings.TrimSpace(url) != "" {
				aux = append(aux, url)
			}
		}
		return aux
	}
	return nil
}

func GetStringFromSlicedUrls(slicedUrls []string) string {
	return strings.TrimSpace(strings.Join(slicedUrls, ","))
}
