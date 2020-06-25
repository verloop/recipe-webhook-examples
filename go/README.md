An example of setting up a HTTP end point to serve Verloop recipes.

To know about the formats of request and response see [this](../README.md)

# How to use?

## Installing Dependencies

```
dep ensure -v
```

## Starting the http server

  * `go run main.go`
  * If you want the request and responses to be printed in `STDOUT` then enable the debug by running `export DEBUG=True` and restart the server.
  * When the webhook request is being configured from recipe make sure to add header for authorization as `CaputDraconis`

## Examples
This example is to be used with the Verloop recipe template `Recharge Bot (Using Webhooks)`