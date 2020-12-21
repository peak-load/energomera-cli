package main

import (
	"fmt"
	"log"
	"strings"
	"time"

	"github.com/peak-load/energomera"
	"github.com/tarm/serial"
	"github.com/tkanos/gonfig"
)

type Configuration struct {
	Port          string
	SleepInterval time.Duration
	Counters      []string
}

func main() {

	configuration := Configuration{}
	err := gonfig.GetConf("config.json", &configuration)
	if err != nil {
		panic(err)
	}

	buf := make([]byte, 128)

	port := "/dev/ttyS0"
	counters := []string{""}
	SleepInterval := time.Millisecond * 500

	if len(configuration.Port) > 0 {
		port = configuration.Port
	}

	if len(configuration.Counters) > 0 {
		counters = configuration.Counters
	}

	SleepInterval = configuration.SleepInterval

	for counter := range counters {
		c := &serial.Config{Name: port, Baud: 9600, Size: 7, StopBits: 1, Parity: 'E', ReadTimeout: time.Millisecond * SleepInterval}
		time.Sleep(time.Second * 1)
		s, err := serial.OpenPort(c)
		if err != nil {
			log.Fatal(err)
		}
		// initialize commands
		fmt.Println("========== COUNTER " + counters[counter] + " ==========")
		n, _ := s.Write([]byte("/?" + counters[counter] + "!\r\n"))
		time.Sleep(time.Millisecond * SleepInterval)
		n, _ = s.Read(buf)

		ident := strings.Split(string(buf[:n]), "")
		time.Sleep(time.Millisecond * SleepInterval)

		n, _ = s.Write([]byte("\x060" + ident[4] + "1\r\n"))
		time.Sleep(time.Millisecond * SleepInterval)
		n, _ = s.Read(buf)

		// send commands
		commands := []string{"VOLTA", "CURRE", "POWEP", "POWPP", "FREQU", "ET0PE"}
		commandline := map[string]string{"head": "R1", "body": ""}
		for i := range commands {
			commandline["body"] = commands[i] + "()"
			command := energomera.DataEncode(commandline)
			n, _ = s.Write(command)
			time.Sleep(time.Millisecond * SleepInterval)
			n, _ = s.Read(buf)
			switch commands[i] {
			case "VOLTA":
				phasev := strings.Split(string(buf[1:n]), "\r\n")
				fmt.Printf("phase1v: %q\n", strings.Trim(phasev[0], "VOLTA()"))
				fmt.Printf("phase2v: %q\n", strings.Trim(phasev[1], "VOLTA()"))
				fmt.Printf("phase3v: %q\n", strings.Trim(phasev[2], "VOLTA()"))

			case "CURRE":
				phasea := strings.Split(string(buf[1:n]), "\r\n")
				fmt.Printf("phase1a: %q\n", strings.Trim(phasea[0], "CURRE()"))
				fmt.Printf("phase2a: %q\n", strings.Trim(phasea[1], "CURRE()"))
				fmt.Printf("phase3a: %q\n", strings.Trim(phasea[2], "CURRE()"))

			case "POWPP":
				phasep := strings.Split(string(buf[1:n]), "\r\n")
				fmt.Printf("phase1p: %q\n", strings.Trim(phasep[0], "POWPP()"))
				fmt.Printf("phase2p: %q\n", strings.Trim(phasep[1], "POWPP()"))
				fmt.Printf("phase3p: %q\n", strings.Trim(phasep[2], "POWPP()"))

			case "ET0PE":
				tarif := strings.Split(string(buf[1:n]), "\r\n")
				fmt.Printf("tarif1: %q\n", strings.Trim(tarif[0], "ET0PE()"))
				fmt.Printf("tarif2: %q\n", strings.Trim(tarif[1], "ET0PE()"))
				fmt.Printf("tarif3: %q\n", strings.Trim(tarif[2], "ET0PE()"))

			case "POWEP":
				power := strings.Split(string(buf[1:n]), "\r\n")
				fmt.Printf("power: %q\n", strings.Trim(power[0], "POWEP()"))

			case "FREQU":
				freq := strings.Split(string(buf[1:n]), "\r\n")
				fmt.Printf("freq: %q\n", strings.Trim(freq[0], "FREQU()"))
			}
		}
		end := []byte("\x01\x42\x30\x03\x75")
		n, _ = s.Write(end)
		s.Close()
	}
}
