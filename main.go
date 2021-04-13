package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	hr "oauth2_callback_linkedin/html"
)

// App version info
const appName = "oauth2_callback_linkedin"
const appVersion = "2104.09a"

// Default Callback URL: http://localhost:12345/oauth2/callback
const appPort = "12345"
const appClientCallbackURL = "http://localhost:12345/oauth2/callback"
const appClientState = "2BhalHArh327AHllah279"
const appClientScope = "r_liteprofile r_emailaddress w_member_social"

// main is the application entry point
func main() {
	fmt.Printf("%s v%s\n", appName, appVersion)
	fmt.Println("www.2BitProgrammers.com\nCopyright (C) 2020. All Rights Reserved.\n")

	// Get the client callbackUrl, ID, Secret, and State from local environment variables
	var evClientCallbackUrl, evClientID, evClientSecret, evClientState, evClientScope string
	evClientCallbackUrl = os.Getenv("LI_CALLBACK_URL")
	if evClientCallbackUrl == "" {
		evClientCallbackUrl = appClientCallbackURL
	}
	evClientID = os.Getenv("LI_CLIENT_ID")
	if evClientID == "" {
		err := "[ERROR] Failed to load environmental variable:  LI_CLIENT_ID\n"
		log.Println(err)
		os.Exit(-1)
	}
	evClientSecret = os.Getenv("LI_CLIENT_SECRET")
	if evClientSecret == "" {
		err := "[ERROR] Failed to load environmental variable:  LI_CLIENT_SECRET\n"
		log.Println(err)
		os.Exit(-1)
	}
	evClientState = os.Getenv("LI_CLIENT_STATE")
	if evClientState == "" {
		evClientState = appClientState
	}
	evClientScope = os.Getenv("LI_CLIENT_SCOPE")
	if evClientScope == "" {
		evClientScope = appClientScope
	}
	hr.SetOAuthConfig(evClientCallbackUrl, evClientID, evClientSecret, evClientState, evClientScope)

	fmt.Println("")

	log.Printf("Starting App on Port %s", appPort)
	log.Println("LinkedIn Login URL: ", hr.GetLoginUrl())

	// Define endpoints, start the local webserver, and listen on default port
	http.HandleFunc("/", hr.HandleStartGet)
	http.HandleFunc("/status", hr.HandleStatusGet)
	http.HandleFunc("/oauth2/callback", hr.HandleCallbackGet)
	log.Fatal(http.ListenAndServe(":"+appPort, nil))
}
