/*
	This file contains utilities for Hecker.

	They are seperate to organize easier.
*/
package main

import (
	"fmt"
)

// Shortcuts
var spf = fmt.Sprintf

// Helper functions
func check_err(err error) {
	if err != nil {
		log.Fatal(err)
	}
}
