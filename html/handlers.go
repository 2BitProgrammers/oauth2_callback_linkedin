package html

import (
	"fmt"
	"log"
	"net/http"
	"net/url"
	"strings"

	// Install via:  go get golang.org/x/oauth2
	"golang.org/x/oauth2"
	"golang.org/x/oauth2/linkedin"
)

// Global variables
var liClientCallbackURL, liClientID, liClientSecret, liClientState, liClientScope string
var liOauth2Conf *oauth2.Config

// encodeScope will split the scope string (separated by spaces) and convert it to a string array
func encodeScope(scope string) []string {
	sFlds := strings.Split(scope, " ")
	//fmt.Println("[DEBUG] sflds: ", sFlds)
	return sFlds
}

// SetOAuthConfig will set the config variables: id, secret, state, scope
func SetOAuthConfig(callbackURL, id, secret, state, scope string) {
	// Set the global variables with values pulled from ENV variables
	liClientCallbackURL = callbackURL
	print(fmt.Sprintf("LI Client Callback URL: %s\n", liClientCallbackURL))
	liClientID = id
	print(fmt.Sprintf("LI Client ID: %s\n", liClientID))
	liClientSecret = secret
	print(fmt.Sprintf("LI Client Secret: %s\n", liClientSecret))
	liClientState = state
	print(fmt.Sprintf("LI Client State: %s\n", liClientState))
	liClientScope = scope
	print(fmt.Sprintf("LI Client Scope: %s\n", liClientScope))
	// Initialize the global variable for the OAuth2 config
	liOauth2Conf = &oauth2.Config{
		RedirectURL:  liClientCallbackURL,
		ClientID:     liClientID,
		ClientSecret: liClientSecret,
		Scopes:       encodeScope(liClientScope),
		Endpoint:     linkedin.Endpoint,
	}
}

// getHtmlCode_Start returns the htlm page with the params you need to start the oauth2 process
func getHtmlCode_Start(callbackURL, id, secret, state, scope string) string {
	htmlData := `<!DOCTYPE html>
	<html lang="en-US">
	
	<head>
		<meta charset="UTF-8">
		<title>LinkedIn Authorization (start)</title>
		<style>
			body {
				color: #133b77;
				background-color: #b8dee0;
				font: 16px/1 sans-serif;
				margin: 18px;
				line-height: 108%;
			}
			#divErrors {
				display: block;
				color: #BC2325;
				padding: 5px;
				margin-left: 28px;
				margin-right: 28px;
				font-weight: bold;
			}
			.cHeader {
				font-weight: bold;			
			}
			.cAlignCenter {
				text-align: center;
			}
			.cAlignRight {
				text-align: right;
			}
			.cAlignCenter {
				text-align: center;
			}
			.cAlignRight {
				text-align: right;
			}
			.cAlignLeft {
				text-align: left;
			}
			.btn {
				width: 444px;
				font-weight: bold;
				background-color: darkgrey;
				color: black;
			}
		</style>
	</head>
	
	<body>
		<div id="divErrors"></div>

		This is the starting point for the LinkedIn OAuth2 process.
    <br><br>
    <u>NOTE</u>: Change values at your own peril.
    <br><br>
	
		<form action="`
	htmlData += liOauth2Conf.Endpoint.AuthURL
	htmlData += `" method="get" target="li_oauth_login">
			<table>
				<tr class="cAlignRight">
					<td><span class="cHeader">LinkedIn Auth URL: &nbsp;</span></td>
					<td><input name="linkedin_auth_url" type="text" value="`
	htmlData += liOauth2Conf.Endpoint.AuthURL
	htmlData += `" size="75" disabled /></td>
			    </tr>
				<tr class="cAlignRight">
					<td><span class="cHeader">Client Redirect URL: &nbsp;</span></td>
					<td><input name="redirect_uri" type="text" value="`
	htmlData += callbackURL
	htmlData += `" size="75" /></td>
				</tr>
				<tr class="cAlignRight">
					<td><span class="cHeader">Client Response Type: &nbsp;</span></td>
					<td><input name="response_type" type="text" value="code" size="75" /></td>
				</tr>				
				<tr class="cAlignRight">
					<td><span class="cHeader">Client ID: &nbsp;</span></td>
					<td><input name="client_id" type="text" value="`
	htmlData += id
	htmlData += `" size="75" /></td>
				</tr>
				<tr class="cAlignRight">
					<td><span class="cHeader">Client Secret: &nbsp;</span></td>
					<!-- <td class="cAlignLeft"><span id="spanClientSecret">secret123 </span></td> -->
					<td><input name="client_secret" type="text" value="`
	htmlData += secret
	htmlData += `" size="75" disabled /></td>
				</tr>
				<tr class="cAlignRight">
					<td><span class="cHeader">Client State: &nbsp;</span></td>
					<td><input name="state" type="text" value="`
	htmlData += state
	htmlData += `" size="75" /></td>
				</tr>
				<tr class="cAlignRight">
					<td><span class="cHeader">Client Scope: &nbsp;</span></td>
					<td><input name="scope" type="text" value="`
	htmlData += scope
	htmlData += `" size="75" /></td>
				</tr>
				<tr>
					<td></td>
					<td class="cAlignCenter">
						<br><br>
						<input class="btn" type="submit" value="Start the LinkedIn OAuth2 Process" />
						</td>
				</tr>
			</table>
		</form>
		
	</body>
	</html>
`
	return htmlData
}

// getHtmlCode_AccessToken returns the html code for the page after the access token is succesfully retrieved
func getHtmlCode_AccessToken(accessToken string) string {
	htmlData := `<!DOCTYPE html>
	<html lang="en-US">
	
	<head>
		<meta charset="UTF-8">
		<title>LinkedIn AccessToken</title>
		<style>
			body {
				color: #133b77;
				background-color: #b8dee0;
				font: 16px/1 sans-serif;
				margin: 18px;
				line-height: 108%;
			}
			#divErrors {
				display: block;
				color: #BC2325;
				padding: 5px;
				margin-left: 28px;
				margin-right: 28px;
				font-weight: bold;
			}
			.cHeader {
				font-weight: bold;			
			}
			.cAlignCenter {
				text-align: center;
			}
			.cAlignRight {
				text-align: right;
			}
			.btn {
				width: 444px;
				font-weight: bold;
				background-color: darkgrey;
				color: black;
				//color: red;
			}
		</style>
	</head>
	
	<body>
		<div id="divErrors"></div>
	`
	htmlData += "<b>LinkedIn Access Token:</b>  \n<br>\n"
	htmlData += accessToken
	htmlData += `
	</body>
	</html>
	`
	return htmlData
}

// getHtmlCode_Error return html code of an webpage which contains a description of an error that happened.
func getHtmlCode_Error(error, errorDescription string) string {
	htmlData := `<!DOCTYPE html>
	<html lang="en-US">
	
	<head>
		<meta charset="UTF-8">
		<title>LinkedIn Error</title>
		<style>
			body {
				color: #133b77;
				background-color: #b8dee0;
				font: 16px/1 sans-serif;
				margin: 18px;
				line-height: 108%;
			}
			#divErrors {
				display: block;
				color: #BC2325;
				padding: 5px;
				margin-left: 28px;
				margin-right: 28px;
			}
			.cHeader {
				font-weight: bold;			
			}
			.cAlignCenter {
				text-align: center;
			}
			.cAlignRight {
				text-align: right;
			}
		</style>
	</head>
	
	<body>
		<div id="divErrors">
			Failed to retrieve LinkedIn Access token.
			<br><br>

			<span class="cHeader">Error: &nbsp;</span>

	`
	htmlData += error
	htmlData += `
			<br><br>
			<span class="cHeader">Description: &nbsp;</span> 
			`
	htmlData += errorDescription
	htmlData += `
	</div>

	</body>
	</html>
	`
	return htmlData
}

// GetLoginUrl returns the LinkedIn login url.  This relies on the global variables to complete successfully.
func GetLoginUrl() string {
	loginUrl := liOauth2Conf.AuthCodeURL(liClientState)
	return loginUrl
}

// HandleStatusGet will handle all incoming status requests. It returns a simple json value.
func HandleStatusGet(w http.ResponseWriter, r *http.Request) {
	method := r.Method
	uri := r.RequestURI
	responsePayload := `{ "healthy": true}`
	returnResponseJson(w, method, uri, "", http.StatusOK, "OK", responsePayload, "")
}

// HandleLoginGet will handle the authentication start page
func HandleStartGet(w http.ResponseWriter, r *http.Request) {
	htmlData := getHtmlCode_Start(liClientCallbackURL, liClientID, liClientSecret, liClientState, liClientScope)
	returnResponseHtml(w, r, htmlData)
}

// HandleCallbackGet will handle the oauth2 callback request
func HandleCallbackGet(w http.ResponseWriter, r *http.Request) {
	state := r.FormValue("state")
	code := r.FormValue("code")
	error, _ := url.QueryUnescape(r.FormValue("error"))
	if error != "" {
		errorDescription, _ := url.QueryUnescape(r.FormValue("error_description"))
		msg := fmt.Sprintf("[ERROR] HandleCallbackGet() - Failed to retrieve access token (%s) - %s", error, errorDescription)
		log.Println(msg)
		htmlData := getHtmlCode_Error(error, errorDescription)
		returnResponseHtml(w, r, htmlData)
		return
	}
	if state != liClientState {
		msg := fmt.Sprintf("[ERROR] HandleCallbackGet() - invalid oauth state (%s)", state)
		log.Println(msg)
		htmlData := getHtmlCode_Error("State Mismatch", fmt.Sprintf("Invalid OAuth State (%s)", state))
		returnResponseHtml(w, r, htmlData)
		return
	}
	token, err := liOauth2Conf.Exchange(oauth2.NoContext, code)
	if err != nil {
		msg := fmt.Sprintf("[ERROR] HandleCallbackGet() - OAuth Code Exchange Failed - %s", err.Error())
		log.Println(msg)
		htmlData := getHtmlCode_Error("OAuth Code Exchange Failed", err.Error())
		returnResponseHtml(w, r, htmlData)
		return
	}

	log.Println("LinkedIn AccessToken: ", token.AccessToken)
	returnResponseHtml(w, r, getHtmlCode_AccessToken(token.AccessToken))
}
