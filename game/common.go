package game

import (
)

const (
	FullScreenSetting = iota
	ResolutionSetting
)

var GameSettings = []*Setting {
    &Setting{
        content: "FullScreen",
        options: []string{"off", "on"},
        selectedOption: 0,
    },
    &Setting{
        content: "Resolution",
        options: []string{"480*360", "1024*768", "1366*768", "1440*900", "1600*900", "1920*1080"},
        selectedOption: 0,
    },
}
