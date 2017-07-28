package main

import (
	"bufio"
	"bytes"
	"fmt"
	"net/http"
	"strings"

	"github.com/Jeffail/gabs"
	"github.com/pivotal-golang/bytefmt"
	"github.com/shirou/gopsutil/cpu"
	"github.com/shirou/gopsutil/disk"
	"github.com/shirou/gopsutil/host"
	"github.com/shirou/gopsutil/load"
	"github.com/shirou/gopsutil/mem"
	"github.com/shirou/gopsutil/net"
	"github.com/shirou/gopsutil/process"
)

type DiskUsageReport struct {
	Mountpoint string          `json:"mountpoint"`
	Usage      *disk.UsageStat `json: "usage"`
}

func saveHandler(w http.ResponseWriter, r *http.Request) {
	// allow cross domain AJAX requests
	w.Header().Set("Access-Control-Allow-Origin", "*")
}

type ProcessDetail struct {
	*process.Process
}

type ProcessState struct {
	Pid           int32   `json:"pid"`
	IsRunning     bool    `json:"is_running"`
	MemoryPercent float32 `json:"memory_percent"`
	Percent       float64 `json:"cpu_percent"`
}

func Index(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "Welcome! Nothing to see here!\n")
}

func (pdetail *ProcessDetail) Cmdinfo() (string, error) {
	cmd, err := pdetail.Cmdline()
	if err != nil {
		return "", err

	}
	return cmd, nil

}

func formatUptime(uptime uint64) string {
	buf := new(bytes.Buffer)
	w := bufio.NewWriter(buf)

	days := uptime / (60 * 60 * 24)

	if days != 0 {
		s := ""
		if days > 1 {
			s = "s"
		}
		fmt.Fprintf(w, "%d day%s, ", days, s)
	}

	minutes := uptime / 60
	hours := minutes / 60
	hours %= 24
	minutes %= 60

	fmt.Fprintf(w, "%2d:%02d", hours, minutes)

	w.Flush()
	return buf.String()
}

func Info(w http.ResponseWriter, r *http.Request) {
	// Return Host Info
	n, _ := host.Info()
	c, _ := cpu.Info()
	m, _ := mem.VirtualMemory()
	s, _ := mem.SwapMemory()
	l, _ := load.Avg()
	i, _ := net.Interfaces()
	netstats, _ := net.IOCounters(true)
	hostmap := make(map[string]interface{})
	// Host Info
	hostmap["hostname"] = n.Hostname
	hostmap["TotalMemory"] = bytefmt.ByteSize(m.Total)
	hostmap["UsedMemory"] = bytefmt.ByteSize(m.Used)
	hostmap["TotalSwap"] = bytefmt.ByteSize(s.Total)
	hostmap["UsedSwap"] = bytefmt.ByteSize(s.Used)
	hostmap["FreeSwap"] = bytefmt.ByteSize(s.Free)
	hostmap["uptime"] = formatUptime(n.Uptime)
	hostmap["load1"] = l.Load1
	hostmap["load5"] = l.Load5
	hostmap["load15"] = l.Load15
	hostmap["OS"] = n.OS
	hostmap["PlatformFamily"] = n.PlatformFamily
	hostmap["PlatformVersion"] = n.PlatformVersion
	hostmap["VirtualizationSystem"] = n.VirtualizationSystem
	hostmap["VirtualizationRole"] = n.VirtualizationRole
	// CPU Info
	hostmap["CPU"] = c[0].CPU
	hostmap["VendorID"] = c[0].VendorID
	hostmap["Cores"] = c[0].Cores
	hostmap["Mhz"] = c[0].Mhz
	hostmap["ModelName"] = c[0].ModelName
	// Network Interfaces
	hostmap["Interfaces"] = i
	hostmap["NetCounters"] = netstats
	jsonObj, _ := gabs.Consume(hostmap)
	w.Header().Set("Access-Control-Allow-Origin", "*")
	fmt.Fprint(w, jsonObj.StringIndent("", "  "))

}

func LoadAverage(w http.ResponseWriter, r *http.Request) {
	v, _ := load.Avg()
	jsonObj, _ := gabs.Consume(v)
	fmt.Fprint(w, jsonObj.StringIndent("", "  "))
}

func Memory(w http.ResponseWriter, r *http.Request) {
	v, _ := mem.VirtualMemory()
	jsonObj, _ := gabs.Consume(v)
	fmt.Fprint(w, jsonObj.StringIndent("", "  "))
}

func CPU(w http.ResponseWriter, r *http.Request) {
	v, _ := cpu.Times(false)
	cpumap := make(map[string]interface{})
	user := v[0].User
	nice := v[0].Nice
	system := v[0].System
	idle := v[0].Idle
	iowait := v[0].Iowait
	irq := v[0].Irq
	softirq := v[0].Softirq
	stolen := v[0].Stolen
	total := user + nice + system + idle + iowait + irq + softirq + stolen
	idlePercent := idle / total * 100
	cpumap["used"] = 100 - idlePercent
	cpumap["nice"] = nice / 100 * 100
	cpumap["user"] = user / total * 100
	cpumap["system"] = system / total * 100
	cpumap["idle"] = idlePercent
	cpumap["iowait"] = iowait / total * 100
	cpumap["irq"] = softirq / total * 100
	cpumap["stolen"] = stolen / total * 100
	jsonObj, _ := gabs.Consume(cpumap)
	w.Header().Set("Access-Control-Allow-Origin", "*")
	fmt.Fprint(w, jsonObj.StringIndent("", "  "))
}

func Network(w http.ResponseWriter, r *http.Request) {
	netbefore, _ := net.IOCounters(true)
	fmt.Fprint(w, netbefore)
}

func Disk(w http.ResponseWriter, r *http.Request) {
	partitions, _ := disk.Partitions(true)
	var disksUsage []string
	for i := 0; i < len(partitions); i++ {
		usage, _ := disk.Usage(partitions[i].Mountpoint)
		usageString := (*usage).String()
		if usageString != "" {
			disksUsage = append(disksUsage, usageString)
		}
	}
	w.Header().Set("Access-Control-Allow-Origin", "*")
	fmt.Fprint(w, "["+strings.Join(disksUsage, ", ")+"]")
}

func Processes(w http.ResponseWriter, r *http.Request) {
	// To be...
	w.Header().Set("Access-Control-Allow-Origin", "*")
}
