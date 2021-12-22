package gedis

import "time"

const (
	AE_SETSIZE = 1024*10

	AE_OK = 0
	AE_ERR = -1

	AE_NONE = 0
	AE_READABLE = 1
	AE_WRITABLE = 2

	AE_FILE_EVENTS = 1
	AE_TIME_EVENTS = 2
	AE_ALL_EVENTS = (AE_FILE_EVENTS|AE_FILE_EVENTS)
	AE_DONT_WAIT = 4

	AE_NOMORE = -1
)

type AeFileProc func(eventLoop *AeEventLoop, fd int64, clientData interface{}, mask int)
type AeTimeProc func(eventLoop *AeEventLoop, id int64, clientData interface{}) int64
type AeEventFinalizeProc func(eventLoop *AeEventLoop, clientData interface{})
type AeBeforeSleepProc func(eventLoop *AeEventLoop)

type AeFileEvent struct {
	Mask int64
	RFileProc *AeFileProc
	WFileProc *AeFileProc
	ClientData interface{}
}

type AeTimeEvent struct {
	Id int64
	WhenSec int64
	WhenMs int64
	TimeProc *AeTimeProc
	FinalizeProc *AeEventFinalizeProc
	ClientData interface{}
	Next *AeTimeEvent
}

type AeFiredEvent struct {
	Fd int64
	Mask int64
}

type AeEventLoop struct {
	MaxFd int64
	SetSize int64
	TimeEventNextId int64
	LastTime time.Time
	Events *AeFileEvent
	Fired *AeFiredEvent
	TimeEventHead *AeTimeEvent
	Stop int64
	ApiData interface{}
	BeforeSleep *AeBeforeSleepProc
}

func AeCreateEventLoop(setSize int64) *AeEventLoop{
	eventLoop := new(AeEventLoop)
	eventLoop.MaxFd = -1
	if aeapi
}
