# Example Configuration

```
{
    "PORT": "80",
    "ALGORITHM": "RoundRobin",
    "SERVERS": [
        "http://5.226.143.60",
        "http://5.226.143.61",
        "http://5.226.143.62"
    ]
}
```

# Step 1 - Port Configuration:

Locate the configuration file "config.json" included
in the RapidBalance file.

The first option is the port for the balancer to run on.
You can chose a port of your choice.

# Step 2 - Algorithm Configuration:

The second option is the routing algorith. You can chose
between two algorithms, round-robin and least-connection.

These algorithms must be inputted as follows:

"RoundRobin" OR "LeastConnection"

# Step 3 - Server Configuration:

Included in the code file are three flask server files.
These servers can be used as the content nodes. 

The flask servers are preconfigured to the following
local IPs and ports.

# Flask Servers
```
Node 1: "http://127.0.0.1:5050"
Node 2: "http://127.0.0.1:5051"
Node 3: "http://127.0.0.1:5052"
```
If you would rather not use local servers for testing,
three external servers running NGINX have been set up as
usable content serving nodes. 

The IPs for those nodes are below.

# Hosted servers
```
Node 1: "http://5.226.143.60"
Node 2: "http://5.226.143.61"
Node 3: "http://5.226.143.62"
```
These nodes will remain active after the submission for use
when marking.
