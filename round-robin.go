package main

import "net/url"

// #-----------------------------------------------------------------------#
// #                                Round Robin                            #
// #-----------------------------------------------------------------------#

func roundRobin() *url.URL {
	// Parse the url to the reverse proxy
	url, _ := url.Parse(platformSettings.SERVERS[rrTarget])
	// If url request is greater than the index
	if (rrTarget + 1) > (len(platformSettings.SERVERS) - 1) {
		// Loop back to the start of the array
		rrTarget = 0
	} else {
		// After request, add 1 to the index
		rrTarget++
	}
	return url
}
