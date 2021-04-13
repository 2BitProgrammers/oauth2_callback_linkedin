## To build this image:  docker build . -t 2bitprogrammers/oauth2_callback_linkedin
##

## Use golang image to build executable
FROM golang:alpine AS builder
## Add CA certificates to avoid error: x509: certificate signed by unknown authority
RUN apk --no-cache add ca-certificates
## Set the environmental variables
ENV GO111MODULE=on \
    CGO_ENABLED=0 \
    GOOS=linux \
    GOARCH=amd64   
## Set the working directory where built files will go         
WORKDIR /build
## Get the required golang modules
COPY $PWD/go.mod .
RUN go mod download
## Copy source and build binary
COPY $PWD/main.go .
COPY $PWD/html/ ./html/
RUN go build -o oauth2_callback_linkedin . 



## Build final image from scratch (copy executeable into empty container)
FROM scratch 
## Add CA certificates to avoid error: x509: certificate signed by unknown authority
COPY --from=builder /etc/ssl/certs/ca-certificates.crt /etc/ssl/certs/
COPY --from=builder /build/oauth2_callback_linkedin /oauth2_callback_linkedin
WORKDIR /
ENV LI_CALLBACK_URL="http://localhost:12345/oauth2/callback" \
    LI_CLIENT_STATE="2BhalHArh327AHllah279" \
    LI_CLIENT_SCOPE="r_liteprofile r_emailaddress w_member_social"
ENTRYPOINT [ "/oauth2_callback_linkedin" ]