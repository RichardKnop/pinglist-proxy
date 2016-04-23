[![Codeship Status for RichardKnop/ping](https://codeship.com/projects/84366050-eb77-0133-2fde-5680c82dbe9d/status?branch=master)](https://codeship.com/projects/148104)

# Ping List Proxy

A simple server to proxy requests through a remote server.

# Index

* [Ping List Proxy](#pinglist-proxy)
* [Index](#index)
* [Dependencies](#dependencies)
* [Setup](#setup)

# Dependencies

According to [Go 1.5 Vendor experiment](https://docs.google.com/document/d/1Bz5-UB7g2uPBdOx-rw5t9MxJwkfpx90cqG9AFL0JAYo), all dependencies are stored in the vendor directory. This approach is called `vendoring` and is the best practice for Go projects to lock versions of dependencies in order to achieve reproducible builds.

To update dependencies during development:

```
make update-deps
```

To install dependencies:

```
make install-deps
```

# Setup

Just run this command:

```
go run main.go runserver
```
