package helpers

import (
	"context"
	"time"

	"github.com/go-kit/log"
)

type LogginMiddleware struct {
	Logger log.Logger
	Next   EncryptService
}

func (mw LogginMiddleware) Encrypt(ctx context.Context, key string, text string) (output string, err error) {
	defer func(begin time.Time) {
		_ = mw.Logger.Log(
			"method", "encrypt",
			"key", key,
			"text", text,
			"output", output,
			"err", err,
			"took", time.Since(begin),
		)
	}(time.Now())
	output, err = mw.Next.Encrypt(ctx, key, text)
	return
}

func (mw LogginMiddleware) Decrypt(ctx context.Context, key string, text string) (output string, err error) {
	defer func(begin time.Time) {
		_ = mw.Logger.Log(
			"method", "decrypt",
			"key", key,
			"text", text,
			"output", output,
			"err", err,
			"took", time.Since(begin),
		)
	}(time.Now())
	output, err = mw.Next.Decrypt(ctx, key, text)
	return
}
