package skink

import (
	"context"
	"io"
	"log"
	"os"

	"github.com/webolton/skink/internal/initializer"
)

func Run(ctx context.Context, out io.Writer, config *initializer.Config) error {
	if err := config.Initialize(os.Args); err != nil {
		return err
	}

	// Set logging to standard out (whether error or info)
	log.SetOutput(out)
}
