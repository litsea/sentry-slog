# sentry-slog

## Usage

```golang
import (
	"context"
	"log"
	"log/slog"

	sentry "github.com/litsea/sentry-slog"
)

h, err := sentry.NewHandler(
	sentry.WithDSN("https://sentry-dsn"),
	sentry.WithAttachStacktrace(true),
	sentry.WithEnvironment("local"),
	sentry.WithRelease("v0.1.0")
	sentry.WithLogLevel(slog.LevelError),
	sentry.WithLogAddSource(true),
	sentry.WithLogAttrFromContext(func(ctx context.Context){ ... }),
	sentry.WithLogExtraAttrs(map[string]string{ ... }),
)
if err != nil {
	log.Fatal(err)
}
```

Default Options:

* `AttachStacktrace`: `true`
* `LogLevel`: `slog.LevelError`
* `LogAddSource`: `true`
* `LogExtraAttrs`: [defaultLogExtraAttrs](slog.go)
