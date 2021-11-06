// Package validator validates a given url

package validator

import "net/url"

func IsUrl(webUrl string) bool {
	u, err := url.Parse(webUrl)
	return err == nil && u.Scheme != "" && u.Host != ""
}