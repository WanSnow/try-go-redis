package gedis

import "sync"

const (
	EPOLL_CLOEXEC = 02000000
	EPOLL_NONBLOCK = 04000

	EPOLLIN = 0x001
	EPOLLPRI = 0x002
	EPOLLOUT = 0x004
	EPOLLRDNORM = 0x040
	EPOLLRDBAND = 0x080
	EPOLLWRNORM = 0x100
	EPOLLWRBAND = 0x200
	EPOLLERR = 0x008
	EPOLLHUP = 0x010
	EPOLLRDHUP = 0x2000
	EPOLLONESHOT = 1 << 30
	EPOLLET = 1 << 31

	EPOLL_CTL_ADD = 1
	EPOLL_CTL_DEL = 2
	EPOLL_CTL_MOD = 3
)

type eventPoll struct {
	Mtx sync.Mutex
	Wq waitQueueHeadT
	PollWait waitQueueHeadT
	RdlList listHead
	Lock sync.RWMutex

}

type EpollData struct {
	Ptr interface{}
	Fd int64
	U32 uint32
	U64 uint64
}

type EpollEvent struct {
	Events uint32
	Data EpollData
}

