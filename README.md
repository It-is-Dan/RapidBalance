<h1 align="center" style="margin-bottom:-15px;">Final Year Project - Rapid Balance</h1>
<p align="center" style="font-weight: bold">A Performance Analysis of Modern Routing Algorithms Used in Content Delivery.</p>

<h2>Project Overview</h2>
<p>
RapidBalance is a software-defined load balancing software with two configurable algorithms, round-robin and least-connection.<br>
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

```
Node 1: "http://127.0.0.1:5050"
Node 2: "http://127.0.0.1:5051"
Node 3: "http://127.0.0.1:5052"
```
