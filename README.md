# sentry

## Usage

### Init Sentry

```golang
import (
	"log"

	"github.com/litsea/sentry"
)

hub, err := sentry.New(
	sentry.WithDSN("https://sentry-dsn"),
	sentry.WithAttachStacktrace(true),
	sentry.WithEnvironment("local"),
	sentry.WithRelease("v0.1.0")
)
if err != nil {
	log.Fatal(err)
}
```

Default Options:

* `AttachStacktrace`: `true`

### Logger Handler

```golang
import (
	"context"
	"log/slog"

	"github.com/litsea/sentry"
)

h := sentry.NewLogHandler(
	hub,
	sentry.LogWithLevel(slog.LevelError),
	sentry.LogWithAddSource(true),
	sentry.LogWithAttrFromContext(func(ctx context.Context){ ... }),
	sentry.LogWithExtraAttrs(map[string]string{ ... }),
)
```

Default Options:

* `Level`: `slog.LevelError`
* `AddSource`: `true`
* `ExtraAttrs`: [defaultLogExtraAttr](log.go)
