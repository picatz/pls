package main

import (
	"context"
	"fmt"
	"log"
	"os"
	"runtime"
	"sync"

	"github.com/shirou/gopsutil/net"
	"github.com/shirou/gopsutil/process"
	"golang.org/x/sync/semaphore"
)

func main() {
	procs, err := process.Processes()

	if err != nil {
		log.Println(err)
		os.Exit(1)
	}

	var wg = &sync.WaitGroup{}
	var sm = semaphore.NewWeighted(int64(runtime.NumCPU()))

	wg.Add(len(procs))

	fmt.Println(fmt.Sprintf("%v\t%v\t%v\t%v\t%v\t%v\t%v\t%v\t%v\t%v\t%v", "pid", "ppid", "user", "running", "connections", "threads", "mem", "cpu", "exe", "cmdline", "children"))

	for _, proc := range procs {
		go func(p *process.Process) {
			sm.Acquire(context.TODO(), 1)
			defer sm.Release(1)
			defer wg.Done()
			var (
				user        string
				children    []int32
				cmdline     string
				exe         string
				ppid        int32
				pid         int32
				mem         float32
				cpu         float64
				threads     int32
				running     bool
				connections int
			)
			pid = p.Pid
			ppid, _ = p.Ppid()
			cmdline, _ = p.Cmdline()
			running, _ = p.IsRunning()
			user, _ = p.Username()
			exe, _ = p.Exe()
			childrenP, _ := p.Children()
			for _, child := range childrenP {
				children = append(children, child.Pid)
			}
			mem, _ = p.MemoryPercent()
			cpu, _ = p.CPUPercent()
			threads, _ = p.NumThreads()
			connectionsP, _ := net.ConnectionsPid("all", p.Pid)
			connections = len(connectionsP)
			fmt.Println(fmt.Sprintf("%v\t%v\t%v\t%v\t%v\t%v\t%v\t%v\t%v\t%v\t%v", pid, ppid, user, running, connections, threads, mem, cpu, exe, cmdline, children))
		}(proc)
	}

	wg.Wait()
}
