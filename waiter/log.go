package waiter

import (
	"context"
	"github.com/docker/docker/api/types/container"
	"github.com/docker/docker/client"
	"io"
	"strings"
	"time"
)

type LogWaiter struct {
	client      *client.Client
	containerID string
	log         string
}

// Wait It Waits for the specified text to appear in the log
func (w *LogWaiter) Wait(c context.Context) error {
	for {
		select {
		case <-c.Done():
			return c.Err()
		default:
			rc, err := w.client.ContainerLogs(c, w.containerID, container.LogsOptions{
				ShowStdout: true,
				ShowStderr: true,
			})
			if err != nil {
				rc.Close()
				time.Sleep(100 * time.Millisecond)
				continue
			}
			bs, err := io.ReadAll(rc)
			if err != nil {
				rc.Close()
				time.Sleep(100 * time.Millisecond)
				continue
			}
			log := string(bs)
			if strings.Count(log, w.log) >= 1 {
				rc.Close()
				return nil
			} else {
				time.Sleep(100 * time.Millisecond)
				rc.Close()
				continue
			}
		}
	}
	return nil
}

func ForLog(log string, client *client.Client, containerID string) *LogWaiter {
	return &LogWaiter{
		log:         log,
		client:      client,
		containerID: containerID,
	}
}
