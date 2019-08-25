package utils

import (
	"fmt"
	"sync"

	"github.com/shirou/gopsutil/mem"
)

type VirtualSpace struct {
	Capacity int
	Used     int
}

type SysInfo struct {
	// stats *stats.Stats
	sync.Mutex
}

// var st *stats.Stats

// func init() {
// 	st = &stats.Stats{}
// }

func GetDiskSpace() (VirtualSpace, error) {
	var disk VirtualSpace
	// var stat syscall.Statfs_t

	// wd, err := os.Getwd()
	// if err != nil {
	// 	return disk, err
	// }
	// syscall.Statfs(wd, &stat)

	// // Available blocks * size per block = available space in bytes
	// fmt.Println(stat.Bavail * uint64(stat.Bsize))
	return disk, nil
}

func GetOSInfo() error {
	return nil
}

func GetCPUInfo() error {
	return nil
}

func GetMemInfo() error {
	v, _ := mem.VirtualMemory()

	// almost every return value is a struct
	fmt.Printf("Total: %v, Free:%v, UsedPercent:%f%%\n", v.Total, v.Free, v.UsedPercent)

	// convert to JSON. String() is also implemented
	fmt.Println(v)
	return nil
}

// func (sys *SysInfo) GetRamInfo() (stats.MemInfo, stats.GoMemory, error) {
// 	sys.stats.GetMemoryInfo(true, true)
// 	fmt.Println(sys.stats.GoInfo)
// 	return nil
// }

// func GetRam() (VirtualSpace, error) {
// 	var ram VirtualSpace
// 	var m runtime.MemStats
// 	runtime.ReadMemStats(&m)
// 	fmt.Printf("Alloc = %v MiB", bToMb(m.Alloc))
// 	fmt.Printf("\tTotalAlloc = %v MiB", bToMb(m.TotalAlloc))
// 	fmt.Printf("\tSys = %v MiB", bToMb(m.Sys))
// 	fmt.Printf("\tNumGC = %v\n", m.NumGC)

// 	return ram, nil
// }

// func bToMb(b uint64) uint64 {
// 	return b / 1024 / 1024
// }
