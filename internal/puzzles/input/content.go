// Package input provides access to embedded puzzles input files.
package input

import (
	"context"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"path"
	"time"

	"github.com/obalunenko/logger"
)

// Date holds date info.
type Date struct {
	Year string
	Day  string
}

func (d Date) String() string {
	return path.Join(d.Year, d.Day)
}

// Get returns puzzle input.
func Get(ctx context.Context, d Date, session string) ([]byte, error) {
	req, err := createInputReq(ctx, d, session)
	if err != nil {
		return nil, fmt.Errorf("create input request: %w", err)
	}

	client := http.DefaultClient

	resp, err := client.Do(req)
	if err != nil {
		return nil, fmt.Errorf("send request: %w", err)
	}

	defer func() {
		if err := resp.Body.Close(); err != nil {
			logger.WithError(ctx, err).Error("Failed to close body")
		}
	}()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("read responsse body: %w", err)
	}

	switch resp.StatusCode {
	case http.StatusOK:
		return body, nil
	case http.StatusNotFound:
		return nil, fmt.Errorf("[%s] puzzle input not found", d)
	case http.StatusBadRequest:
		return nil, fmt.Errorf("unauthorised")
	default:
		return nil, fmt.Errorf("[%s] failed to get puzzle input[%s]", d, resp.Status)
	}
}

// createInputReq creates an HTTP request for retrieving the Advent of Code
// input given year/day.
func createInputReq(ctx context.Context, d Date, sessionID string) (*http.Request, error) {
	const (
		baseurl = "https://adventofcode.com"
		day     = "day"
		input   = "input"
	)

	u, err := url.Parse(baseurl)
	if err != nil {
		return nil, fmt.Errorf("parse base url: %w", err)
	}

	u.Path = path.Join(u.Path, d.Year, day, d.Day, input)

	req, err := http.NewRequestWithContext(ctx, http.MethodGet, u.String(), nil)
	if err != nil {
		return nil, err
	}

	req.AddCookie(&http.Cookie{
		Name:       "session",
		Value:      sessionID,
		Path:       "/",
		Domain:     ".adventofcode.com",
		Expires:    time.Time{},
		RawExpires: "",
		MaxAge:     0,
		Secure:     false,
		HttpOnly:   false,
		SameSite:   0,
		Raw:        "",
		Unparsed:   nil,
	})

	return req, nil
}
