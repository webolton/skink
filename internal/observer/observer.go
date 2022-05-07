package observer

import (
	"context"
	"fmt"
	"log"

	"github.com/radovskyb/watcher"

	"github.com/webolton/skink/internal/initializer"
)

func trackDir() error {
	w := watcher.New()

	go func() {
		for {
			select {
			case event := <-w.Event:
				fmt.Println(event) // Print the event's info.
			case err := <-w.Error:
				log.Fatalln(err)
			case <-w.Closed:
				return
			}
		}
	}()

	// Watch this folder for changes.
	if err := w.Add("."); err != nil {
		log.Fatalln(err)
	}
	return nil
}
func Watch(ctx context.Context, config *initializer.Config) error {

	select {
	case <-ctx.Done():
		return nil
	default:
		fmt.Println(config.DefaultAWSCreds)
		fmt.Println("CALLED!!!")
		// trackDir()
		// do the thing
	}
	return nil
}
