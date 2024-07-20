package main

import (
	"fmt"
	"testing"
)

func TestGetLocationAreas(t *testing.T) {
	if err := GetLocationAreas(); err != nil {
		fmt.Println("error getting location areas")
	}
}
