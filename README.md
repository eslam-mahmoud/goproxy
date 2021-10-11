
# goproxy

Simple proxy app created using Golang to create secure tunnel

## How to launch the application
### Run the commited build
run `./goproxy`

Using CURL you can send request to the app example `curl https://api.giphy.com/v1/gifs/search?q=morning\&api_key=APIKEY\&limit=1 -x localhost:8081`

### With Go env
If `go` environment installed use this command
`go run *.go` to run from the code 

or You can build the app using `go build -o goproxy` then run it using `./goproxy`

Using CURL you can send request to the app example `curl https://api.giphy.com/v1/gifs/search?q=morning\&api_key=APIKEY\&limit=1 -x localhost:8081`

### With docker compose
Using this comand `docker compose up --build`

Using CURL you can send request to the app as `curl https://api.giphy.com/v1/gifs/search?q=morning\&api_key=APIKEY\&limit=1 -x localhost:9501`

## Notes
### Port number
Notice the difference in port number between running the app using go vs docker compose, this is because the app run locally on port 8081 which is configrable using env value `PORT` which take port as string ex. `8081`

### Allowed destinations
Application have pre defined list of Allowed destinations which are giphy API & google.com which can be configered using env var `PROVIDERS` as string comma sprated strings ex `api.giphy.com:443,google.com:443`

## Enhancments and known issues
### Using api clien apps like postman and insomnia
We get `requestURI` as `"/api.giphy.com/v1/gifs/search?q=morning&api_key=123"` from api client apps

We get `requestURI` as `"api.giphy.com:443"` from `CURL`

So prefared to use `CURL` for now

### Deploying to production
For now I have trupple running this app to prod/staging env, all the requests fails with message like as nginx does not support http connect out of the box https://github.com/chobits/ngx_http_proxy_connect_module

Using nginx

```
$ curl https://api.giphy.com/v1/gifs/search?q=morning\&api_key=231\&limit=1 -v -x https://goproxy.eslam.me
*   Trying 157.245.246.176:443...
* TCP_NODELAY set
* Connected to goproxy.eslam.me (157.245.246.176) port 443 (#0)
* ALPN, offering http/1.1
* successfully set certificate verify locations:
*   CAfile: /etc/ssl/certs/ca-certificates.crt
  CApath: /etc/ssl/certs
* TLSv1.3 (OUT), TLS handshake, Client hello (1):
* TLSv1.3 (IN), TLS handshake, Server hello (2):
* TLSv1.3 (IN), TLS handshake, Encrypted Extensions (8):
* TLSv1.3 (IN), TLS handshake, Certificate (11):
* TLSv1.3 (IN), TLS handshake, CERT verify (15):
* TLSv1.3 (IN), TLS handshake, Finished (20):
* TLSv1.3 (OUT), TLS change cipher, Change cipher spec (1):
* TLSv1.3 (OUT), TLS handshake, Finished (20):
* SSL connection using TLSv1.3 / TLS_AES_256_GCM_SHA384
* ALPN, server accepted to use http/1.1
* Proxy certificate:
*  subject: CN=goproxy.eslam.me
*  start date: Oct 10 21:33:32 2021 GMT
*  expire date: Jan  8 21:33:31 2022 GMT
*  subjectAltName: host "goproxy.eslam.me" matched cert's "goproxy.eslam.me"
*  issuer: C=US; O=Let's Encrypt; CN=R3
*  SSL certificate verify ok.
* allocate connect buffer!
* Establish HTTP proxy tunnel to api.giphy.com:443
> CONNECT api.giphy.com:443 HTTP/1.1
> Host: api.giphy.com:443
> User-Agent: curl/7.68.0
> Proxy-Connection: Keep-Alive
> 
* TLSv1.3 (IN), TLS handshake, Newsession Ticket (4):
* TLSv1.3 (IN), TLS handshake, Newsession Ticket (4):
< HTTP/1.1 400 Bad Request
< Server: nginx
< Date: Mon, 11 Oct 2021 01:03:26 GMT
< Content-Type: text/html
< Content-Length: 150
< Connection: close
< 
* Received HTTP code 400 from proxy after CONNECT
* CONNECT phase completed!
* Closing connection 0
curl: (56) Received HTTP code 400 from proxy after CONNECT
```

Using heroku

```
$ curl https://api.giphy.com/v1/gifs/search?q=morning\&api_key=1fYKUz7KHLRa88bqR3CJeGIEXuPqNvCI\&limit=1 -v -x https://golangproxy.herokuapp.com
*   Trying 54.83.6.65:443...
* TCP_NODELAY set
* Connected to golangproxy.herokuapp.com (54.83.6.65) port 443 (#0)
* ALPN, offering http/1.1
* successfully set certificate verify locations:
*   CAfile: /etc/ssl/certs/ca-certificates.crt
  CApath: /etc/ssl/certs
* TLSv1.3 (OUT), TLS handshake, Client hello (1):
* TLSv1.3 (IN), TLS handshake, Server hello (2):
* TLSv1.2 (IN), TLS handshake, Certificate (11):
* TLSv1.2 (IN), TLS handshake, Server key exchange (12):
* TLSv1.2 (IN), TLS handshake, Server finished (14):
* TLSv1.2 (OUT), TLS handshake, Client key exchange (16):
* TLSv1.2 (OUT), TLS change cipher, Change cipher spec (1):
* TLSv1.2 (OUT), TLS handshake, Finished (20):
* TLSv1.2 (IN), TLS handshake, Finished (20):
* SSL connection using TLSv1.2 / ECDHE-RSA-AES128-GCM-SHA256
* ALPN, server did not agree to a protocol
* Proxy certificate:
*  subject: CN=*.herokuapp.com
*  start date: Jun  1 00:00:00 2021 GMT
*  expire date: Jun 30 23:59:59 2022 GMT
*  subjectAltName: host "golangproxy.herokuapp.com" matched cert's "*.herokuapp.com"
*  issuer: C=US; O=Amazon; OU=Server CA 1B; CN=Amazon
*  SSL certificate verify ok.
* allocate connect buffer!
* Establish HTTP proxy tunnel to api.giphy.com:443
> CONNECT api.giphy.com:443 HTTP/1.1
> Host: api.giphy.com:443
> User-Agent: curl/7.68.0
> Proxy-Connection: Keep-Alive
> 
* TLSv1.2 (IN), TLS alert, close notify (256):
* Proxy CONNECT aborted
* CONNECT phase completed!
* Closing connection 0
* TLSv1.2 (OUT), TLS alert, close notify (256):
curl: (56) Proxy CONNECT aborted
```

### Handle redirect
`$ curl -p --proxy 127.0.0.1:8081 https://google.com/search?q=hello`
```
<HTML><HEAD><meta http-equiv="content-type" content="text/html;charset=utf-8">
<TITLE>301 Moved</TITLE></HEAD><BODY>
<H1>301 Moved</H1>
The document has moved
<A HREF="https://www.google.com/search?q=hello">here</A>.
</BODY></HTML>
```