// Package input provides access to embedded puzzles input files.
package input

import (
	"context"
	"errors"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"path"
	"strings"
	"time"

	"github.com/obalunenko/logger"
)

var (
	// ErrNotFound returns when puzzle input is not yet unlocked or invalid date passed.
	ErrNotFound = errors.New("puzzle inout not found")
	// ErrUnauthorized returns when session is empty or invalid.
	ErrUnauthorized = errors.New("unauthorized")
)

// Date holds date info.
type Date struct {
	Year string
	Day  string
}

func (d Date) String() string {
	return path.Join(d.Year, d.Day)
}

// ClientDo provides the interface for custom HTTP client implementations.
type ClientDo interface {
	Do(*http.Request) (*http.Response, error)
}

// Client is the default Client and is used by Get, Head, and Post.
var Client ClientDo = http.DefaultClient

// Get returns puzzle input.
func Get(ctx context.Context, d Date, session string) ([]byte, error) {
	req, err := createInputReq(ctx, d, session)
	if err != nil {
		return nil, fmt.Errorf("create input request: %w", err)
	}

	const (
		timeoutSecs = 5
	)

	ctx, cancel := context.WithTimeout(ctx, time.Second*timeoutSecs)
	defer cancel()

	req = req.Clone(ctx)

	resp, err := Client.Do(req)
	if err != nil {
		return nil, fmt.Errorf("send request: %w", err)
	}

	defer func() {
		if err = resp.Body.Close(); err != nil {
			logger.WithError(ctx, err).Error("Failed to close body")
		}
	}()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("read responsse body: %w", err)
	}

	switch resp.StatusCode {
	case http.StatusOK:
		if strings.TrimSpace(string(body)) == "" {
			return nil, fmt.Errorf("empty response received")
		}

		return body, nil
	case http.StatusNotFound:
		return nil, fmt.Errorf("[%s]: %w", d, ErrNotFound)
	case http.StatusBadRequest:
		return nil, ErrUnauthorized
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

	req, err := http.NewRequestWithContext(ctx, http.MethodGet, u.String(), http.NoBody)
	if err != nil {
		return nil, fmt.Errorf("create request: %w", err)
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

	req.Header.Set("User-Agent",
		"Mozilla/5.0 (Macintosh; Intel Mac OS X 10_15_7) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/96.0.4664.93 Safari/537.36")

	return req, nil
}
