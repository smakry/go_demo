package pprof

import (
	"fmt"
	"io/ioutil"
	"os"
	"os/signal"
	"strconv"
	"syscall"

	"github.com/funny/cmd"
	"github.com/funny/jsonlog/log"
	"github.com/funny/pprof"
)

func WaitCommand() {
	registerCommands()

	if pid := syscall.Getpid(); pid != 1 {
		ioutil.WriteFile("game_server.pid", []byte(strconv.Itoa(pid)), 0777)
		defer os.Remove("game_server.pid")
	}

	sigUSR1 := make(chan os.Signal, 1)
	sigTERM := make(chan os.Signal, 1)

	signal.Notify(sigUSR1, syscall.SIGUSR1)
	signal.Notify(sigTERM, syscall.SIGTERM)

	for {
		select {
		case <-sigUSR1:
			processCommand()
		case <-sigTERM:
			fmt.Println("killed")
			return
		}
	}
}

func processCommand() {
	command, err := ioutil.ReadFile("game.pprof.cmd")
	if err != nil {
		log.Error("read game.pprof.cmd", log.M{"error": err.Error()})
		return
	}
	if _, ok := cmd.Process(string(command)); !ok {
		fmt.Printf("unknow command: %s", string(command))
		cmd.Help(os.Stderr)
	}
}

func registerCommands() {
	cmd.Register(
		"cpuprof (start|stop)",
		"Start or stop cpu profile. The profile will saved to game.cpu.profile.",
		func(args []string) {
			switch args[1] {
			case "start":
				pprof.StartCPUProfile("game.cpu.profile")
			case "stop":
				pprof.StopCPUProfile()
			}
		},
	)
}
