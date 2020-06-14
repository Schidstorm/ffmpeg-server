package lib

import (
	"bufio"
	"os/exec"
	"time"
)

type FfmpegHandler struct {
	arguments []string
	totalDuration time.Duration
	progressionTime time.Duration
}

func (h *FfmpegHandler) Progress() float32 {
	return float32(h.progressionTime) / float32(h.totalDuration)
}


func (h *FfmpegHandler) Run() error {
	cmd := exec.Command("ffmpeg", h.arguments...)

	stdout, err := cmd.StderrPipe()
	if err != nil {
		return err
	}

	err = cmd.Start()
	if err != nil {
		return err
	}

	scanner := bufio.NewScanner(stdout)
	scanner.Split(bufio.ScanLines)
	for scanner.Scan() {
		dataString := scanner.Text()
		if duration, err := FindLastDuration(dataString, "Duration: ?"); err == nil {
			h.totalDuration = duration
		}
		if t, err := FindLastDuration(dataString, "out_time=?"); err == nil {
			h.progressionTime = t
		}
	}

	return nil
}

func NewFfmpegHandler(arguments ...string) *FfmpegHandler {
	return &FfmpegHandler{
		arguments: arguments,
	}
}
