package main

import (
	"fmt"
	"strconv"
)

var bcolor map[string]string = map[string]string{
	"HEADER":    "\033[95m",
	"OKBLUE":    "\033[94m",
	"OKCYAN":    "\033[96m",
	"OKGREEN":   "\033[92m",
	"WARNING":   "\033[93m",
	"FAIL":      "\033[91m",
	"ENDC":      "\033[0m",
	"BOLD":      "\033[1m",
	"UNDERLINE": "\033[4m",
}

func color(data string, col string) string {
	if col == "data" {
		return bcolor["HEADER"] + data + bcolor["ENDC"]
	} else if col == "%" {
		res, err := strconv.ParseFloat(data, 32)
		if err != nil {
			panic(err)
		}
		return bcolor["OKGREEN"] + fmt.Sprintf("%.2f", res) + bcolor["ENDC"] + " %"
	} else if col == "time" {
		return bcolor["OKCYAN"] + data + bcolor["ENDC"] + " sec"
	} else if col == "strong" {
		return bcolor["FAIL"] + data + bcolor["ENDC"]
	}
	return ""
}
