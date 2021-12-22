package gedis

import "sync"

type waitQueueHeadT struct {
	Lock sync.Locker
	TaskList listHead
}
