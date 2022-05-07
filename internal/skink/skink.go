package skink

import (
	"context"
	"fmt"
	"io"
	"log"
	"os"

	"github.com/webolton/skink/internal/initializer"
	"github.com/webolton/skink/internal/observer"
)

func Run(ctx context.Context, out io.Writer, config *initializer.Config) error {
	// Initialize struct
	if err := config.Initialize(os.Args); err != nil {
		return err
	}

	// Set logging to standard out (whether error or info)
	log.SetOutput(out)
	fmt.Println(config.SyncedDirPath)
	if err := observer.Watch(ctx, config); err != nil {
		return err
	}

	return nil
}
