/*
	This file contains utilities for Hecker.

	They are seperate to organize easier.
*/
package main

import (
	"fmt"
	"net/url"
	"log"
)

// Constant data
const VERSION = "0.1"

// Shortcuts
var spf = fmt.Sprintf

// Helper functions
func check_err(err error) {
	if err != nil {
		log.Fatal(err)
	}
}

// Get the host of the URL
func parse_url_host(input string) string {
	u, err := url.Parse(input)
	check_err(err)

	return u.Hostname()
}
