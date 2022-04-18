<h1 align="center" style="margin-bottom:-15px;">Final Year Project</h1>
<p align="center" style="font-weight: bold">A Performance Analysis of Modern Routing Algorithms Used in Content Delivery.</p>

<h2>Project Overview</h2>
<p>
RapidBalance is a software-defined load balancing system with two configurable algorithms, round-robin and least-connection.<br>
    <br>
    <b>PLEASE NOTE:</b> This software was created as part of a research project and as a result isnt ready for live production.
</p>

<h2>Getting Started</h2>



<h2>Example Configuration</h2>

```
{
    "PORT": "80",
    "ALGORITHM": "RoundRobin",
    "SERVERS": [
        "http://127.0.0.1:5050",
        "http://127.0.0.1:5051",
        "http://127.0.0.1:5052"
    ]
}
```

<h2>Flask Servers</h2>

By default, the configuration file is set to use the three following local flask servers.

```
Node 1: "http://127.0.0.1:5050"
Node 2: "http://127.0.0.1:5051"
Node 3: "http://127.0.0.1:5052"
```
