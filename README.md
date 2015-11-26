## Relay Controller

Do you have a relay connected to the GPIO pins of your Raspberry Pi? If so, then this application is for you. Simply run the executable and open your browser to control the relay using the intuitive web interface.

### Screenshot

A picture is worth a thousand words:

![Screenshot](http://i.stack.imgur.com/qk5yh.jpg)

### Behind the Scenes

Relay Controller consists of two components:

- the backend server providing the API written in Go
- the front-end web interface written in JavaScript using Ember

### Building the Application

To build the backend server:

- ensure that you have [Git](https://git-scm.com/) and [Go](https://golang.org/) installed
- ensure that [`GOPATH`](https://golang.org/doc/code.html#GOPATH) is properly set
- run the following command:

        go install github.com/nathan-osman/relaycontroller

The source code will then be downloaded and built. The compiled binary can be found in `$GOPATH/bin`.

### Running the Application

Relay Controller requires a single argument - a path to the JSON file that contains the configuration. A sample configuration file is provided below:

    {
        "server": {
            "addr": ":8000",
            "root": "./www"
        },
        "channels": [
            {
                "name": "gpio2",
                "title": "GPIO2",
                "number": 2
            }
        ]
    }

The first block indicates which port the server should listen on and the path to the `www` directory (located in the same directory as the source code).

The second block (`channels`) is a list of GPIO pins to control. Each entry consists of:

- the name to use internally for referring to the channel
- the text to be displayed in the web interface
- the GPIO pin number to use for the channel
