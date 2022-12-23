package main

import (
	"fmt"
	"github.com/sony/sonyflake"
	"os"
	"time"
)

func main() {
	t, _ := time.Parse("2006-01-02", "2022-01-01")
	settings := sonyflake.Settings{
		StartTime:      t,
		MachineID:      getMachineID,
		CheckMachineID: checkMachineID,
	}

	sf := sonyflake.NewSonyflake(settings)

	for i := 0; i < 100; i++ {
		id, err := sf.NextID()
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}
		fmt.Println(id)
	}

}

func getMachineID() (uint16, error) {
	var machineID uint16 = 6
	return machineID, nil
}

func checkMachineID(machineID uint16) bool {
	existsMachines := []uint16{1, 2, 3, 4, 5}
	for _, v := range existsMachines {
		if v == machineID {
			return false
		}
	}
	return true
}
