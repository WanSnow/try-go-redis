package gedis

const (
	evport_debug = 0
	MAX_EVENT_BATCHSZ = 512
)

type AeApiState struct {
	Epfd int64
	Events *EpollEvent
}

func AeApiCreate(eventLoop *AeEventLoop) int {
	state := new(AeApiState)


}
