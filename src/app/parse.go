package app

import (
	"net"
	"net/url"
	"strings"

	"github.com/pkg/errors"
)

func parseTelegramUsername(rawURL string) (string, error) {
	if !strings.HasPrefix(rawURL, "https://") && !strings.HasPrefix(rawURL, "http://") {
		rawURL = "https://" + rawURL
	}

	u, err := url.Parse(rawURL)
	if err != nil {
		return "", errors.Wrap(err, "can not parse url")
	}

	host := u.Host
	if strings.Contains(host, ":") {
		host, _, err = net.SplitHostPort(u.Host)
		if err != nil {
			return "", errors.Wrap(err, "can not split host")
		}
	}

	if host != "t.me" {
		return "", errors.New("this is not a telegram url")
	}

	return strings.TrimPrefix(u.Path, "/"), nil
}
