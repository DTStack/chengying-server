package limit

import (
	"context"
	"golang.org/x/time/rate"
	"io"
)

type reader struct {
	r       io.Reader
	limiter *rate.Limiter
}

func NewLimitedReader(r io.Reader, l *rate.Limiter) io.Reader {
	return &reader{
		r:       r,
		limiter: l,
	}
}

func (r *reader) Read(buf []byte) (int, error) {
	n, err := r.r.Read(buf)
	if n <= 0 {
		return n, err
	}

	r.limiter.Wait(context.Background())
	return n, err
}
