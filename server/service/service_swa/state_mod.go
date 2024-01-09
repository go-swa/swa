package service_swa

import (
	"github.com/ServiceWeaver/weaver"
	"runtime"
	"time"

	"github.com/shirou/gopsutil/v3/cpu"
	"github.com/shirou/gopsutil/v3/disk"
	"github.com/shirou/gopsutil/v3/mem"
)

const (
	B  = 1
	KB = 1024 * B
	MB = 1024 * KB
	GB = 1024 * MB
)

var StateDict = map[string]string{
	"os":   "操作系统",
	"cpu":  "处理器",
	"ram":  "内存",
	"disk": "硬盘",

	"goos":         "适配系统",
	"numCpu":       "cpu个数",
	"compiler":     "编译器",
	"goVersion":    "go版本",
	"numGoroutine": "协程数量",
	"cores":        "内核数",
	"usedMb":       "已用(GB)",
	"totalMb":      "共计(GB)",
	"usedPercent":  "使用率",
	"usedGb":       "已用(GB)",
	"totalGb":      "容量(GB)",
}

type Server struct {
	weaver.AutoMarshal
	Os   Os   `json:"os"`
	Cpu  Cpu  `json:"cpu"`
	Ram  Ram  `json:"ram"`
	Disk Disk `json:"disk"`
}

type Os struct {
	weaver.AutoMarshal
	GOOS         string `json:"goos"`
	NumCPU       int    `json:"numCpu"`
	Compiler     string `json:"compiler"`
	GoVersion    string `json:"goVersion"`
	NumGoroutine int    `json:"numGoroutine"`
}

type Cpu struct {
	weaver.AutoMarshal
	Cpus  []float64 `json:"cpus"`
	Cores int       `json:"cores"`
}

type Ram struct {
	weaver.AutoMarshal
	UsedMB      int `json:"usedMb"`
	TotalMB     int `json:"totalMb"`
	UsedPercent int `json:"usedPercent"`
}

type Disk struct {
	weaver.AutoMarshal
	UsedGB int `json:"usedGb"`
	TotalGB     int `json:"totalGb"`
	UsedPercent int `json:"usedPercent"`
}


func InitOS() (o Os) {
	o.GOOS = runtime.GOOS

	o.NumCPU = runtime.NumCPU()

	o.Compiler = runtime.Compiler

	o.GoVersion = runtime.Version()

	o.NumGoroutine = runtime.NumGoroutine()

	return o
}


func InitCPU() (c Cpu, err error) {
	if cores, err := cpu.Counts(false); err != nil {
		return c, err
	} else {
		c.Cores = cores
	}
	if cpus, err := cpu.Percent(time.Duration(200)*time.Millisecond, true); err != nil {
		return c, err
	} else {
		c.Cpus = cpus
	}
	return c, nil
}


func InitRAM() (r Ram, err error) {
	if u, err := mem.VirtualMemory(); err != nil {
		return r, err
	} else {
		r.UsedMB = int(u.Used) / GB
		r.TotalMB = int(u.Total) / GB
		r.UsedPercent = int(u.UsedPercent)
	}
	return r, nil
}


func InitDisk() (d Disk, err error) {
	if u, err := disk.Usage("/"); err != nil {
		return d, err
	} else {
		d.UsedGB = int(u.Used) / GB
		d.TotalGB = int(u.Total) / GB
		d.UsedPercent = int(u.UsedPercent)
	}
	return d, nil
}
