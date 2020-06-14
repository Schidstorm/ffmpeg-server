package model

import (
	"ffmpeg-server/lib"
	"path"
)

type Task struct {
	SourceFile string
	TargetFile string
	Progress float32
	Finished bool
	Error error
	LibTask *lib.Task
}

func (t *Task) Update() {
	t.Progress = t.LibTask.Progress()
	t.Error = t.LibTask.Error()
	t.Finished = t.LibTask.Finished()
}

func NewFfmpegTask(sourceFilePath, destinationFileName string) *Task {
	destinationFilePath := path.Join(lib.GetSettings().ConversionDestinationDirectory, destinationFileName)
	task := lib.NewTask(lib.NewFfmpegHandler(
		"-y", "-re", "-hide_banner", "-progress", "pipe:2", "-i", sourceFilePath,
		"-af", "channelmap=0", "-b:a", "128k", "-map", "0:v", "-map", "0:a", destinationFilePath))
	task.Start()

	modelTask := &Task{
		SourceFile: sourceFilePath,
		TargetFile: destinationFilePath,
		Progress:   0,
		Finished:   false,
		Error:      nil,
		LibTask:    task,
	}
	RegisterTask(modelTask)
	return modelTask
}