package main

import (
	"io/ioutil"
	"log"
	"os"
	"os/signal"
	"os/user"
	"runtime/debug"
	"syscall"
	"time"

	bpfwrapper2 "github.com/Zeta201/bpfwrapper"

	"github.com/Zeta201/connections"
	"github.com/Zeta201/settings"
	"github.com/iovisor/gobpf/bcc"
)

// abortIfNotRoot checks the current user permissions, if the permissions are not elevated, we abort.
func abortIfNotRoot() {
	current, err := user.Current()
	if err != nil {
		log.Panic(err)
	}

	if current.Uid != "0" {
		log.Panic("sniffer must run under superuser privileges")
	}
}

// recoverFromCrashes is a defer function that caches all panics being thrown from the application.
func recoverFromCrashes() {
	if err := recover(); err != nil {
		log.Printf("Application crashed: %v\nstack: %s\n", err, string(debug.Stack()))
	}
}

func main() {
	// if len(os.Args) != 2 {
	// 	fmt.Println("Usage: go run main.go <path to bpf source code>")
	// 	os.Exit(1)
	// }
	// bpfSourceCodeFile := os.Args[1]
	bpfSourceCodeContent, err := ioutil.ReadFile("sourcecode.c")
	if err != nil {
		log.Panic(err)
	}

	defer recoverFromCrashes()
	abortIfNotRoot()

	if err := settings.InitRealTimeOffset(); err != nil {
		log.Printf("Failed fixing BPF clock, timings will be offseted: %v", err)
	}

	// Catching all termination signals to perform a cleanup when being stopped.
	sig := make(chan os.Signal, 1)
	signal.Notify(sig, syscall.SIGHUP, syscall.SIGINT, syscall.SIGQUIT, syscall.SIGTERM)

	bpfModule := bcc.NewModule(string(bpfSourceCodeContent), nil)
	if bpfModule == nil {
		log.Panic("bpf is nil")
	}
	defer bpfModule.Close()

	connectionFactory := connections.NewFactory(time.Minute)
	go func() {
		for {
			connectionFactory.HandleReadyConnections()
			time.Sleep(10 * time.Second)
		}
	}()
	if err := bpfwrapper2.LaunchPerfBufferConsumers(bpfModule, connectionFactory); err != nil {
		log.Panic(err)
	}

	// Lastly, after everything is ready and configured, attach the kprobes and start capturing traffic.
	if err := bpfwrapper2.AttachKprobes(bpfModule); err != nil {
		log.Panic(err)
	}
	log.Println("Sniffer is ready")
	<-sig
	log.Println("Signaled to terminate")
}
