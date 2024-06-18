package main

import (
	"fmt"
	"log"
	"math"
	"os"
	"strconv"
	"strings"
)

type Battery struct {
	Percent float64
	Charging bool
}

func ReadFileToInt32(filepath string) (int32, error) {
	data, err := os.ReadFile(filepath)
	if err != nil {
		return 0, fmt.Errorf("error reading file %s: %v", filepath, err)
	}

	strData := strings.TrimSpace(string(data))
	int64Data, err := strconv.ParseInt(strData, 10, 32)
	if err != nil {
		return 0, fmt.Errorf("error parsing int from file %s: %v", filepath, err)
	}

	return int32(int64Data), nil
}

func GetBatteryInfo()(b Battery, err error){
	b.Percent = GetBatteryState()
	b.Charging = GetChargingStatus()
	return b, nil
}

func GetBatteryState() float64 {
	energy_now, err := ReadFileToInt32("/sys/class/power_supply/BAT0/energy_now")
	if err != nil {
		log.Fatal(err)
		return float64(0)
	}
	energy_full, err := ReadFileToInt32("/sys/class/power_supply/BAT0/energy_full")
	if err != nil {
		log.Fatal(err)
		return float64(0)
	}
	return (float64(energy_now)/float64(energy_full))
}

func GetChargingStatus() bool {
	charging_info, err := os.ReadFile("/sys/class/power_supply/AC/online")

	if err != nil {
		log.Fatal(err)
	}

	strChargingData := strings.TrimSpace(string(charging_info))
	if strChargingData == "1"{
		return true
	} else {
		return false
	}
}

func (b Battery) String() string {
	return fmt.Sprintf("%d%%", int64(math.Round(b.Percent*float64(100))))
}

func (b Battery) IconColor() string {
	if b.Percent <= 0.2 {
		return "#EF3340"
	}
	if b.Percent <= 0.3 {
		return "#FF8000"
	}
	if b.Percent <= 0.4 {
		return "#EFEFEF"
	}

	return ""
}

func (b Battery) FormattedStatus() string {
	s:= fmt.Sprintf("%s %s", b.Icon(), b)
	return fmt.Sprintf("%s\n%s\n%s", s, s, b.IconColor())
}

func (b Battery) Icon() string {
	if b.Charging {
		return ""
	}
	if b.Percent <= 0.1 {
		return ""
	}
	if b.Percent <= 0.3 {
		return ""
	}
	if b.Percent <= 0.5 {
		return ""
	}
	if b.Percent <= 0.8 {
		return ""
	}
	return ""
}

func main(){
	b, _:= GetBatteryInfo()
	fmt.Println(b.FormattedStatus())	
}


