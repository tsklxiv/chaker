/*This file contains utilities for Chaker.
	They are seperate to organize easier.
*/
package main

import (
	"fmt"
	"log"
	"net/url"
)

// Constant data
const VERSION = "0.1"

// Shortcuts
var spf = fmt.Sprintf

// Helper functions
func checkErr(err error) {
	if err != nil {
		log.Fatal(err)
	}
}

// Get the host of the URL
func parseURLHost(input string) string {
	u, err := url.Parse(input)
	checkErr(err)

	return u.Hostname()
}
