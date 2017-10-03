# timestampd

This is a debugging tool that runs an HTTP server which will give you the current time.

# Usage

## Server

    $ timestampd
    2017/10/03 16:50:20 Listening on :8080

## Client

    $ curl localhost:8080
    2017-10-03 16:50:28.400799846 -0500 CDT
    $ curl localhost:8080
    2017-10-03 16:50:30.146679864 -0500 CDT
