package input_test

import (
	"context"
	"errors"
	"io"
	"net/http"
	"strings"
	"testing"
	"testing/iotest"

	"github.com/stretchr/testify/assert"

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

func TestGet(t *testing.T) {
	prevCli := input.Client

	t.Cleanup(func() {
		input.Client = prevCli
	})

	type client struct {
		input.ClientDo
	}

	type args struct {
		ctx     context.Context
		d       input.Date
		session string
	}

	tests := []struct {
		name    string
		client  client
		args    args
		want    []byte
		wantErr assert.ErrorAssertionFunc
	}{
		{
			client: client{
				ClientDo: newMockHTTPClient(returnParams{
					status: http.StatusOK,
					body:   io.NopCloser(strings.NewReader("test")),
				}),
			},
			args: args{
				ctx: context.Background(),
				d: input.Date{
					Year: "2021",
					Day:  "25",
				},
				session: "123",
			},
			want:    []byte("test"),
			wantErr: assert.NoError,
		},
		{
			name: "",
			client: client{
				ClientDo: newMockHTTPClient(returnParams{
					status: http.StatusNotFound,
					body:   http.NoBody,
				}),
			},
			args: args{
				ctx: context.Background(),
				d: input.Date{
					Year: "2021",
					Day:  "25",
				},
				session: "123",
			},
			want:    nil,
			wantErr: assert.Error,
		},
		{
			name: "",
			client: client{
				ClientDo: newMockHTTPClient(returnParams{
					status: http.StatusBadRequest,
					body:   io.NopCloser(strings.NewReader("no session")),
				}),
			},
			args: args{
				ctx: context.Background(),
				d: input.Date{
					Year: "2021",
					Day:  "25",
				},
				session: "123",
			},
			want:    nil,
			wantErr: assert.Error,
		},
		{
			name: "",
			client: client{
				ClientDo: newMockHTTPClient(returnParams{
					status: http.StatusInternalServerError,
					body:   io.NopCloser(strings.NewReader("no session")),
				}),
			},
			args: args{
				ctx: context.Background(),
				d: input.Date{
					Year: "2021",
					Day:  "25",
				},
				session: "123",
			},
			want:    nil,
			wantErr: assert.Error,
		},
		{
			name: "",
			client: client{
				ClientDo: &mockHTTPClient{
					MockDo: func(req *http.Request) (*http.Response, error) {
						return &http.Response{}, errors.New("error in test")
					},
				},
			},
			args: args{
				ctx: context.Background(),
				d: input.Date{
					Year: "2021",
					Day:  "25",
				},
				session: "123",
			},
			want:    nil,
			wantErr: assert.Error,
		},
		{
			name: "",
			client: client{
				ClientDo: newMockHTTPClient(returnParams{
					status: http.StatusOK,
					body:   io.NopCloser(iotest.ErrReader(errors.New("custom error"))),
				}),
			},
			args: args{
				ctx: context.Background(),
				d: input.Date{
					Year: "2021",
					Day:  "25",
				},
				session: "123",
			},
			want:    nil,
			wantErr: assert.Error,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			input.Client = tt.client

			got, err := input.Get(tt.args.ctx, tt.args.d, tt.args.session)
			if !tt.wantErr(t, err) {
				return
			}

			assert.Equal(t, tt.want, got)
		})
	}
}
