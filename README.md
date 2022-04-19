<h1 align="center" style="margin-bottom:-15px;">Final Year Project</h1>
<p align="center" style="font-weight: bold">A Performance Analysis of Modern Routing Algorithms Used in Content Delivery.</p>
<br>
<h2>Project Overview</h2>
<p>
RapidBalance is a software-defined load balancing system with two configurable algorithms, round-robin and least-connection.<br>
    <br>
    <b>PLEASE NOTE:</b> This software was created as part of a research project and as a result isnt ready for live production.
</p>

<h2>Getting Started</h2>

By default, the configuration is setup to use the RoundRobin algorithm, run on port 80 and route to the local flask servers. 

<h3>Requirements</h3>
    <ul>
        <li>GO 1.17
        <li>Python 3.10.4</li>
    </ul>

<h3>Running The Balancer</h3>

<p>The easiest way to get up an running is by using VS Code, you will be able to run the "HTTP-server" file and the software will run as intended.
Should you wish to run the software in a command terminal, you will need to build the program with <code>go build</code> and run it with 
<code>go run file_name.go</code>.</p>
<p>If you are using the flask servers, you will need to run <code>pip3 install flask</code> to install the dependencies.</p>
<p>The round-robin algorithm works as intended, but the least-connection algorithm doesnt acuratly remove idle connections.</p>

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

By default, the configuration file is set to use these three following local flask servers.

```
Node 1: "http://127.0.0.1:5050"
Node 2: "http://127.0.0.1:5051"
Node 3: "http://127.0.0.1:5052"
```
