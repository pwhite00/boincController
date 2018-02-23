package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"io/ioutil"
	"os"
	"os/exec"
	"strconv"
	"time"
)

// LaunchConfig ... does stuff with the JSON
type LaunchConfig struct {
	Containers []struct {
		BaseName      string
		RunningCount  int
		ProjectTarget string
		ProjectKey    string
	}
}

func main() {
	configFlag := flag.String("config", "", "Name the JSON config file to use.")
	noExecFlag := flag.Bool("noop", false, "Display the commands that would run but no not acutally exec any commands.")
	flag.Parse()

	var conf LaunchConfig
	configFileName := *configFlag
	noExec := *noExecFlag

	parseConfigFile(configFileName, &conf)

	fmt.Println(conf.Containers[0].BaseName)

	execCmd(&conf, noExec)

	fmt.Println("Launch commands Complete.")
}

func execCmd(l *LaunchConfig, n bool) {
	// first we figure out the commands.
	for i := 0; i < len(l.Containers); i++ {
		fmt.Printf("TypeCounter: %d\n", i)
		for x := 0; x < l.Containers[i].RunningCount; x++ {
			timeStamp := time.Now().Unix()
			time.Sleep(1 * time.Second)
			switch n {
			case true:
				fmt.Printf("NOOP: docker run -ti -d --name %s-%d -e boincurl=%s -e boinckey=%s boinc_client\n", l.Containers[i].BaseName, timeStamp, l.Containers[i].ProjectTarget, l.Containers[i].ProjectKey)

			case false:
				fmt.Printf("EXEC: instance %d, name %s\n", x+1, l.Containers[i].BaseName)
				launchCommand := "docker"
				cmdArgs := []string{"run", "-ti", "-d", "--name", l.Containers[i].BaseName + "-" + strconv.FormatInt(timeStamp, 10), "-e", "boincurl=" + l.Containers[i].ProjectTarget, "-e", "boinckey=" + l.Containers[i].ProjectKey, "boinc_client"}
				err := exec.Command(launchCommand, cmdArgs...).Run()
				if err != nil {
					fmt.Printf("%s-%d launch [Error]: [%v]\n", l.Containers[i].BaseName, timeStamp, err)
					os.Exit(10)
				}
				fmt.Printf("%s-%d launch [ok]\n", l.Containers[i].BaseName, timeStamp)
			}
		}
	}
}

func parseConfigFile(f string, l *LaunchConfig) {
	reader, rErr := os.Open(f)
	if rErr != nil {
		//fmt.Println(rErr)
		panic(rErr)
	}
	defer reader.Close()

	source, sErr := ioutil.ReadAll(reader)
	if sErr != nil {
		fmt.Println(sErr)
	}

	jsonErr := json.Unmarshal(source, &l)
	if jsonErr != nil {
		fmt.Println(jsonErr)
	}
}
