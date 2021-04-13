# OAuth2 Callback (for LinkedIn)
This creates a local OAuth2 callback so you complete a 3-leg authorization flow. Resulting in an active session token which can be used for use with LinkedIn's developer API.

## Requirements
If you wish to run the application manually, you will need install golang and the associated oauth2 package:
```bash
go get golang.org/x/oauth2
```
If you wish to run the application though a container, you will need docker installed.

Before running the application, you will need to have a LinkedIn user account, enable your LinkedIn developer profile, and create/register an application.
* https://www.linkedin.com/developers/apps

Once the LinkedIn app is created, you will need to 
* Create credentials and record the following:
  * Client ID 
  * Client Secret
* Add Redirect URL for your app:
  * http://localhost:12345/oauth2/callback
* Add OAuth 2.0 scopes
  * r_emailaddress
  * r_liteprofile
  * w_member_social

Finally, it is recommended that you do NOT use the default "state" value.  Instead, create a random "state" value.  In theory, this will make the final session token more "secure'.

## Overview of OAuth2 Flow (2-leg)
This is essentially, server authentication which requires no interation with an user.  However, unless you pay for LinkedIn enterprise, you won't be able to use this option.  

## Overview of OAuth2 Flow (3-leg)
This is where the user authenticates themselves and then allows the application access to the LinkedIn resources.  This application assists in this process by creating a callback url which can be redirected to in order to finalize the flow.

More information can be found here:
* https://docs.microsoft.com/en-us/linkedin/shared/authentication/authorization-code-flow?context=linkedin%2Fcontext&tabs=https

Here are the steps (high level overview):
1. Bring up a terminal, console, or powershell (depends on which operating system you are using)
1. Set ENV variables
   * This must be done before the execution of the golang app (or passed to the container during execution)
1. Run Application
   * This can be done manually or via container  
   * The app checks for the presense of the ENV variables and sets them in memory at startup
1. Open Web Browser, Connect to the local App 
   * This will bring up a HTML form which will be already populated with the values from the ENV variables.
   * URL:  http://localhost:12345/index.html
1. Click the HTML button: "Start the LinkedIn OAUth2 Process"
   * You will automatically be redirected to LinkedIn's login url
   * The destination URL will look like this:
     * https://www.linkedin.com/oauth/v2/authorization?response_type=code&client_id={client_id}&redirect_uri=https%3A%2F%2Flocalhost:12345%2Foauth2%2Fcallback&state={client_state}&scope=r_liteprofile%20r_emailaddress%20w_member_social
1. Fill in your LinkedIn Credentials and click "Sign In"
   * If successful, this will allow linked in to authorize the app
1. The browser will redirect to local Callback URL
   * This will verify the state and exchange the OAuth2 token in order to retrieve a new AccessToken
   * NOTE: The state must match the one sent in "Step 1".
   * URL will look like this:
    * http://localhost:12345/oauth2/callback?code={code}&state={state}
1. If successful, the browser will display your Access Token.
   * If error, then it will display the error
1. You can verify that the token is valid by using LinkedIn's token inspector 
   * https://www.linkedin.com/developers/tools/token-inspector?clientId={client_id}


Once you have a valid Access Token, you can start using LinkedIn API commands.  A simple example would be to get your profile details.  If it works, it should look like this:
```bash
$ curl -X GET 'https://api.linkedin.com/v2/me' \
    -H 'Authorization: Bearer {access_token}'

{"localizedLastName":"Smith","profilePicture":{"displayImage":"urn:li:digitalmediaAsset:AAAAABBBBCCCC"},"firstName":{"localized":{"en_US":"John"},"preferredLocale":{"country":"US","language":"en"}},"lastName":{"localized":{"en_US":"Smith"},"preferredLocale":{"country":"US","language":"en"}},"id":"id1234","localizedFirstName":"John"}
``` 



## APP - Running Manually (golang)
You must 1st set the local ENV variables before you run the app. The default values are:
* LI_CALLBACK_URL = "http://localhost:12345/oauth2/callback"
* LI_CLIENT_ID = <i>N/A, you must create app and apply your own</i>
* LI_CLIENT_SECRET = <i>N/A, you must create app and apply your own</i>
* LI_CLIENT_STATE = "2BhalHArh327AHllah279"
  * NOTE: It is recommended that you change this value with a random value.
* LI_CLIENT_SCOPE = "r_liteprofile r_emailaddress w_member_social"
  * Separate each scope permission with a single space.
  * NOTE: You should set this to the same permissions which your LinkedIn app has.  If they don't match, it will cause the authentication to fail.

### Set ENV Varaibles (windows powershell)
```powershell
## Set the clientID and clientSecret
$env:LI_CALLBACK_URL = 'http://localhost:12345/oauth2/callback'
$env:LI_CLIENT_ID = 'id123'
$env:LI_CLIENT_SECRET = 'secret1234'
$env:LI_CLIENT_STATE = 'state12345'
$env:LI_CLIENT_SCOPE = 'r_liteprofile r_emailaddress w_member_social'

## Show the variables value
Get-ChildItem Env:LI_CALLBACK_URL
Get-ChildItem Env:LI_CLIENT_ID
Get-ChildItem Env:LI_CLIENT_SECRET
Get-ChildItem Env:LI_CLIENT_STATE
Get-ChildItem Env:LI_CLIENT_SCOPE

## Remove Item
Remove-Item Env:LI_CALLBACK_URL
Remove-Item Env:LI_CLIENT_ID
Remove-Item Env:LI_CLIENT_SECRET
Remove-Item Env:LI_CLIENT_STATE
Remove-Item Env:LI_CLIENT_SCOPE
```

### Set ENV Varaibles (linux bash)
```bash
## Set the clientID and clientSecret
export LI_CALLBACK_URL='http://localhost:12345/oauth2/callback'
export LI_CLIENT_ID='id123'
export LI_CLIENT_SECRET='secret1234'
export LI_CLIENT_STATE='state12345'
export LI_CLIENT_SCOPE='r_liteprofile r_emailaddress w_member_social'

## Show the variables value
echo $LI_CALLBACK_URL
echo $LI_CLIENT_ID
echo $LI_CLIENT_SECRET
echo $LI_CLIENT_STATE
echo $LI_CLIENT_SCOPE

## Remove ENV variables
unset LI_CALLBACK_URL
unset LI_CLIENT_ID
unset LI_CLIENT_SECRET
unset LI_CLIENT_STATE
unset LI_CLIENT_SCOPE
```

Once the variables are set, you can execute the app:
```bash
$ go run main.go

oauth2_callback_linkedin v2104.09a
www.2BitProgrammers.com
Copyright (C) 2020. All Rights Reserved.

LI Client Callback URL: http://localhost:12345/oauth2/callback
LI Client ID: id123
LI Client Secret: secret1234
LI Client State: state12345
LI Client Scope: r_liteprofile r_emailaddress w_member_social

2021/04/12 13:18:34 Starting App on Port 12345
2021/04/12 13:18:34 Login URL:  https://www.linkedin.com/oauth/v2/authorization?client_id=id123&redirect_uri=http%3A%2F%2Flocalhost%3A12345%2Foauth2%2Fcallback&response_type=code&scope=r_liteprofile+r_emailaddress+w_member_social&state=state12345

CTRL+C
exit status 3221225786

```

## Docker
### Build Image 
```bash
$ docker build . -t 2bitprogrammers/oauth2_callback_linkedin

[+] Building 0.3s (15/15) FINISHED
 => [internal] load build definition from Dockerfile                                                                                                                   0.0s
 => => transferring dockerfile: 1.18kB                                                                                                                                 0.0s
 => [internal] load .dockerignore                                                                                                                                      0.0s
 => => transferring context: 2B                                                                                                                                        0.0s
 => [internal] load metadata for docker.io/library/golang:alpine                                                                                                       0.0s
 => [builder 1/8] FROM docker.io/library/golang:alpine                                                                                                                 0.0s
 => [internal] load build context                                                                                                                                      0.0s
 => => transferring context: 150B                                                                                                                                      0.0s
 => CACHED [builder 2/8] RUN apk --no-cache add ca-certificates                                                                                                        0.0s
 => CACHED [builder 3/8] WORKDIR /build                                                                                                                                0.0s
 => CACHED [builder 4/8] COPY /go.mod .                                                                                                                                0.0s
 => CACHED [builder 5/8] RUN go mod download                                                                                                                           0.0s
 => CACHED [builder 6/8] COPY /main.go .                                                                                                                               0.0s
 => CACHED [builder 7/8] COPY /html/ ./html/                                                                                                                           0.0s
 => CACHED [builder 8/8] RUN go build -o oauth2_callback_linkedin .                                                                                                    0.0s
 => [stage-1 1/3] COPY --from=builder /etc/ssl/certs/ca-certificates.crt /etc/ssl/certs/                                                                               0.0s
 => [stage-1 2/3] COPY --from=builder /build/oauth2_callback_linkedin /oauth2_callback_linkedin                                                                        0.1s
 => exporting to image                                                                                                                                                 0.1s
 => => exporting layers                                                                                                                                                0.1s
 => => writing image sha256:be3a7f237761e871180d27590b5db1777f105b1539ae68925354a1360b67a622                                                                           0.0s
 => => naming to docker.io/2bitprogrammers/oauth2_callback_linkedin                                                                                                    0.0s

```
### Display Images
```bash
$ docker image ls 2bitprogrammers/oauth2_callback_linkedin

REPOSITORY                                 TAG       IMAGE ID       CREATED         SIZE
2bitprogrammers/oauth2_callback_linkedin   latest    fd18dbe46fe9   1 minutes ago   7.48MB

```

### APP - Running in a Container 
```bash
$ docker run --rm -d \
  --name oauth2_callback_linkedin \
  -e LI_CALLBACK_URL="http://localhost:12345/oauth2/callback" \
  -e LI_CLIENT_ID="id123" \
  -e LI_CLIENT_SECRET="secret1234" \
  -e LI_CLIENT_STATE="state12345" \
  -e LI_CLIENT_SCOPE="r_liteprofile r_emailaddress w_member_social" \
  -p 12345:12345 \
  2bitprogrammers/oauth2_callback_linkedin

## Verify that it is running
$ docker ps

CONTAINER ID   IMAGE                                      COMMAND                  CREATED          STATUS          PORTS                                                                                                      NAMES
546be3de23be   2bitprogrammers/oauth2_callback_linkedin   "/oauth2_callback_li…"   31 seconds ago   Up 10 seconds   0.0.0.0:12345->12345/tcp                                                                                   oauth2_callback_linkedin

## Stopping the container
$ docker stop oauth2_callback_linkedin

## Verify that it is stopped
$ docker ps

CONTAINER      ID    IMAGE    COMMAND     CREATED     STATUS      PORTS    NAMES

```
