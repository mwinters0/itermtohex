package main

import (
	"encoding/json"
	"encoding/xml"
	"fmt"
	"os"
)

type plist struct {
	Dict []dict `xml:"dict"`
}

type dict struct {
	Key  []string  `xml:"key"`
	Real []float64 `xml:"real"`
	Dict []dict    `xml:"dict"`
}

type Color struct {
	Name   string
	RFloat float64
	GFloat float64
	BFloat float64
	RInt   uint8
	GInt   uint8
	BInt   uint8
	Hex    string
}

func die(y string) {
	fmt.Println("error: " + y)
	os.Exit(1)
}

func usage() {
	fmt.Println(fmt.Sprintf(
		"Usage:\n       %s myfile.itermcolors\n       %s print myfile.itermcolors",
		os.Args[0], os.Args[0],
	))
	os.Exit(1)
}

func convert(filename string) []Color {
	contents, err := os.ReadFile(filename)
	if err != nil {
		die(err.Error())
	}
	var p plist
	err = xml.Unmarshal(contents, &p)
	if err != nil {
		die(err.Error())
	}

	var colors []Color
	mainDict := p.Dict[0]
	for i, colorName := range mainDict.Key {
		c := Color{
			Name: colorName,
		}
		c.BFloat = mainDict.Dict[i].Real[0]
		c.GFloat = mainDict.Dict[i].Real[1]
		c.RFloat = mainDict.Dict[i].Real[2]
		c.RInt = uint8(c.RFloat * 255)
		c.GInt = uint8(c.GFloat * 255)
		c.BInt = uint8(c.BFloat * 255)
		c.Hex = fmt.Sprintf("#%02x%02x%02x", c.RInt, c.GInt, c.BInt)
		colors = append(colors, c)
	}
	return colors
}

func main() {
	var filename string
	var print bool
	switch len(os.Args) {
	case 2:
		filename = os.Args[1]
	case 3:
		if os.Args[1] != "print" {
			usage()
		}
		filename = os.Args[2]
		print = true
	default:
		usage()
	}

	colors := convert(filename)
	if print {
		printColor := func(c Color) {
			fmt.Printf(
				"%s (%s): \033[48;2;%d;%d;%dm          \033[0m\n",
				c.Name,
				c.Hex,
				c.RInt,
				c.GInt,
				c.BInt,
			)
		}
		for _, c := range colors {
			printColor(c)
		}
		return
	}

	j, err := json.MarshalIndent(colors, "", "  ")
	if err != nil {
		die(err.Error())
	}
	fmt.Println(string(j))
}
