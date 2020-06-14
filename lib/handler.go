package lib

type ProgressCallback = func(progress float32)
type FinishCallback func(err error)


type TaskHandler interface {
	Progress() float32
	Run() error
}