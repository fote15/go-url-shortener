package utils

import (
	"fmt"
	"net/http"
	"regexp"
	"strings"
	"time"
)

var emailRegex = regexp.MustCompile(`^[a-zA-Z0-9._%+\-]+@[a-zA-Z0-9.\-]+\.[a-zA-Z]{2,}$`)

func IsValidEmail(email string) bool {
	return emailRegex.MatchString(email)
}

func IsValidURL(raw string) (error, string) {
	// If no scheme, prepend "http://"
	if !strings.HasPrefix(raw, "http://") && !strings.HasPrefix(raw, "https://") {
		raw = "https://" + raw
	}

	client := http.Client{
		Timeout: 5 * time.Second, // avoid hanging
	}

	resp, err := client.Head(raw) // use HEAD to avoid downloading full content
	if err != nil {
		return err, ""
	}
	defer resp.Body.Close()

	if resp.StatusCode < 200 || resp.StatusCode >= 400 {
		return fmt.Errorf("URL responded with status: %s", resp.Status), ""
	}

	return nil, raw
}
