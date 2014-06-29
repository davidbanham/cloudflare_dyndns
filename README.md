This script updates Cloudflare with the IP address of the machine it's running on.

This is intended to give you a thing like no-ip or dyndns, except using your cloudflare account.

Configuration is via environment variables. An example is:

    CLOUDFLARE_EMAIL=email@example.com \
    CLOUDFLARE_TOKEN=8afbe6dea02407989af4dd4c97bb6e25 \
    CLOUDFLARE_SUBDOMAIN=home \
    CLOUDFLARE_ROOT_DOMAIN=example.com \
    node index.js

There is also a version in go, with binaries for most every platform. If you don't have/want node, you can use the binary for your platform in build.

    CLOUDFLARE_EMAIL=email@example.com \
    CLOUDFLARE_TOKEN=8afbe6dea02407989af4dd4c97bb6e25 \
    CLOUDFLARE_SUBDOMAIN=home \
    CLOUDFLARE_ROOT_DOMAIN=example.com \
    ./build/app-darwin-amd64
