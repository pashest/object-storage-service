package utils

import "syscall"

// GetFreeDiskSpace returns the remaining space in bytes
func GetFreeDiskSpace(path string) (int64, error) {
	var stat syscall.Statfs_t
	err := syscall.Statfs(path, &stat)
	if err != nil {
		return 0, err
	}
	return int64(stat.Bavail) * int64(stat.Bsize), nil
}
