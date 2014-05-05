package images

import "fmt"

var MaxSize = 10 * MB // todo: set this from a config file?

const (
	_       = iota // ignore first value by assigning to blank identifier
	KB size = 1 << (10 * iota)
	MB
	GB
)

type size float64

func toSize(s string) (l size) {
	fmt.Sscanf(s, "%f", &l)
	return
}

func (b size) String() string {
	switch {
	case b >= GB:
		return fmt.Sprintf("%.2fGB", b/GB)
	case b >= MB:
		return fmt.Sprintf("%.2fMB", b/MB)
	case b >= KB:
		return fmt.Sprintf("%.2fKB", b/KB)
	}
	return fmt.Sprintf("%.2fB", b)
}
