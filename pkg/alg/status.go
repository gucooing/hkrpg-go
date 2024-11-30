package alg

import (
	"fmt"
	"time"

	"github.com/gucooing/hkrpg-go/pkg/logger"
	"github.com/shirou/gopsutil/cpu"
	"github.com/shirou/gopsutil/mem"
)

type Status struct {
	T   time.Time `json:"-"`
	Cpu float64   `json:"cpu"`
	Mem struct {
		Current uint64 `json:"current"`
		Total   uint64 `json:"total"`
	} `json:"mem"`
	Swap struct {
		Current uint64 `json:"current"`
		Total   uint64 `json:"total"`
	} `json:"swap"`
	Disk struct {
		Current uint64 `json:"current"`
		Total   uint64 `json:"total"`
	} `json:"disk"`
	Uptime     uint64    `json:"uptime"`
	Loads      []float64 `json:"loads"`
	TcpCount   int       `json:"tcpCount"`
	UdpCount   int       `json:"udpCount"`
	NetTraffic struct {
		Sent uint64 `json:"sent"`
		Recv uint64 `json:"recv"`
	} `json:"netTraffic"`
}

// func GetStatus() *Status {
// 	now := time.Now()
// 	status := &Status{
// 		T: now,
// 	}
//
// 	percents, err := cpu.Percent(0, false)
// 	if err != nil {
// 		logger.Warn("get cpu percent failed:", err)
// 	} else {
// 		status.Cpu = percents[0]
// 	}
//
// 	upTime, err := host.Uptime()
// 	if err != nil {
// 		logger.Warn("get uptime failed:", err)
// 	} else {
// 		status.Uptime = upTime
// 	}
//
// 	memInfo, err := mem.VirtualMemory()
// 	if err != nil {
// 		logger.Warn("get virtual memory failed:", err)
// 	} else {
// 		status.Mem.Current = memInfo.Used
// 		status.Mem.Total = memInfo.Total
// 	}
//
// 	swapInfo, err := mem.SwapMemory()
// 	if err != nil {
// 		logger.Warn("get swap memory failed:", err)
// 	} else {
// 		status.Swap.Current = swapInfo.Used
// 		status.Swap.Total = swapInfo.Total
// 	}
//
// 	distInfo, err := disk.Usage("/")
// 	if err != nil {
// 		logger.Warn("get dist usage failed:", err)
// 	} else {
// 		status.Disk.Current = distInfo.Used
// 		status.Disk.Total = distInfo.Total
// 	}
//
// 	avgState, err := load.Avg()
// 	if err != nil {
// 		logger.Warn("get load avg failed:", err)
// 	} else {
// 		status.Loads = []float64{avgState.Load1, avgState.Load5, avgState.Load15}
// 	}
//
// 	ioStats, err := net.IOCounters(false)
// 	if err != nil {
// 		logger.Warn("get io counters failed:", err)
// 	} else if len(ioStats) > 0 {
// 		ioStat := ioStats[0]
// 		status.NetTraffic.Sent = ioStat.BytesSent
// 		status.NetTraffic.Recv = ioStat.BytesRecv
// 	} else {
// 		logger.Warn("can not find io counters")
// 	}
//
// 	status.TcpCount, err = sys.GetTCPCount()
// 	if err != nil {
// 		logger.Warn("get tcp connections failed:", err)
// 	}
//
// 	status.UdpCount, err = sys.GetUDPCount()
// 	if err != nil {
// 		logger.Warn("get udp connections failed:", err)
// 	}
//
// 	return status
// }

func GetStatus() string {
	return fmt.Sprintf("CPU占用:%.2f%%\n内存占用%s",
		GetCpuOc(),
		MemoryOc())
}

func GetCpuOc() float64 {
	percents, err := cpu.Percent(0, false)
	if err != nil {
		logger.Warn("get cpu percent failed:", err)
		return 0
	}
	return percents[0]
}

func MemoryOc() string {
	memInfo, err := mem.VirtualMemory()
	if err != nil {
		logger.Warn("get virtual memory failed:", err)
		return "0/0"
	}
	used := float64(memInfo.Used)
	total := float64(memInfo.Total)

	if used/1024/1024 > 1024 {
		return fmt.Sprintf("%.2f/%.2fGB", used/1024/1024/1024, total/1024/1024/1024)
	}
	return fmt.Sprintf("%.2f/%.2fMB", used/1024/1024, total/1024/1024)
}
