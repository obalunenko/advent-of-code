package command

import (
	"context"
	"errors"
	"io"
	"net/http"
	"strings"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"

	"github.com/obalunenko/advent-of-code/internal/puzzles"
	"github.com/obalunenko/advent-of-code/internal/puzzles/input"
)

// Custom type that allows setting the func that our Mock Do func will run instead
type dofunc func(req *http.Request) (*http.Response, error)

// MockClient is the mock client
type mockHTTPClient struct {
	MockDo dofunc
}

type returnParams struct {
	status int
	body   io.ReadCloser
}

func newMockHTTPClient(p returnParams) *mockHTTPClient {
	return &mockHTTPClient{
		MockDo: func(req *http.Request) (*http.Response, error) {
			return &http.Response{
				Status:           http.StatusText(p.status),
				StatusCode:       p.status,
				Proto:            "HTTP/1.0",
				ProtoMajor:       1,
				ProtoMinor:       0,
				Header:           nil,
				Body:             p.body,
				ContentLength:    0,
				TransferEncoding: nil,
				Close:            false,
				Uncompressed:     false,
				Trailer:          nil,
				Request:          nil,
				TLS:              nil,
			}, nil
		},
	}
}

// Overriding what the Do function should "do" in our MockClient
func (m *mockHTTPClient) Do(req *http.Request) (*http.Response, error) {
	return m.MockDo(req)
}

type mockSolver struct {
	year string
	name string
}

func (m mockSolver) Year() string {
	return m.year
}

func (m mockSolver) Day() string {
	return m.name
}

func (m mockSolver) Part1(in io.Reader) (string, error) {
	split, err := read(in)
	if err != nil {
		return "", err
	}

	return split[1], nil
}

func (m mockSolver) Part2(in io.Reader) (string, error) {
	split, err := read(in)
	if err != nil {
		return "", err
	}

	return split[2], nil
}

func read(in io.Reader) ([]string, error) {
	all, err := io.ReadAll(in)
	if err != nil {
		return nil, err
	}

	split := strings.Split(string(all), ",")

	if len(split) != 3 {
		return nil, errors.New("wrong parts")
	}

	return split, nil
}

func TestRun(t *testing.T) {
	ctx := context.Background()

	year := "1992"
	day := "31"

	puzzles.Register(mockSolver{
		year: year,
		name: day,
	})

	t.Cleanup(func() {
		puzzles.UnregisterAllSolvers(t)
	})

	type expected struct {
		result  puzzles.Result
		wantErr assert.ErrorAssertionFunc
	}

	tests := []struct {
		name         string
		returnParams returnParams
		expected     expected
	}{
		{
			name: "",
			returnParams: returnParams{
				status: http.StatusOK,
				body:   io.NopCloser(strings.NewReader("1,2,3")),
			},
			expected: expected{
				result: puzzles.Result{
					Year:  "1992",
					Name:  "31",
					Part1: "2",
					Part2: "3",
				},
				wantErr: assert.NoError,
			},
		},
		{
			name: "",
			returnParams: returnParams{
				status: http.StatusNotFound,
				body:   http.NoBody,
			},
			expected: expected{
				result:  puzzles.Result{},
				wantErr: assert.Error,
			},
		},
		{
			name: "",
			returnParams: returnParams{
				status: http.StatusUnauthorized,
				body:   http.NoBody,
			},
			expected: expected{
				result:  puzzles.Result{},
				wantErr: assert.Error,
			},
		},
		{
			name: "",
			returnParams: returnParams{
				status: http.StatusOK,
				body:   http.NoBody,
			},
			expected: expected{
				result:  puzzles.Result{},
				wantErr: assert.Error,
			},
		},
	}

	for i := range tests {
		tt := tests[i]

		t.Run(tt.name, func(t *testing.T) {
			cli := input.NewFetcher(newMockHTTPClient(tt.returnParams), time.Second*5)

			got, err := run(ctx, cli, year, day)
			if !tt.expected.wantErr(t, err) {
				return
			}

			assert.Equal(t, tt.expected.result, got)
		})
	}
}
