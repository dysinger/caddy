package main

import (
    "github.com/caddyserver/caddy/caddy/caddymain"
    // plugins
    _ "github.com/caddyserver/dnsproviders/dnsimple"
    _ "github.com/hacdias/caddy-webdav"
)

func main() {
    caddymain.EnableTelemetry = false
    caddymain.Run()
}
