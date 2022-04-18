<h1 align="center" style="margin-bottom:-15px;">Final Year Project - Rapid Balance</h1>
<p align="center" style="font-weight: bold">A Performance Analysis of Modern Routing Algorithms Used in Content Delivery.</p>

<h2>Getting Started</h2>



<h2>Example Configuration</h2>

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

<h2>Flask Servers</h2>
```
Node 1: "http://127.0.0.1:5050"
Node 2: "http://127.0.0.1:5051"
Node 3: "http://127.0.0.1:5052"
```
If you would rather not use local servers for testing,
three external servers running NGINX have been set up as
usable content serving nodes. 

The IPs for those nodes are below.

<h2>Hosted servers</h2>
```
Node 1: "http://5.226.143.60"
Node 2: "http://5.226.143.61"
Node 3: "http://5.226.143.62"
```
These nodes will remain active after the submission for use
when marking.
