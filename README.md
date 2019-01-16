# 🌟 Lux Lucet 
<img src="https://travis-ci.org/team142/lux-lucet.svg?branch=master" />&nbsp;
<a href="https://goreportcard.com/report/github.com/team142/lux-lucet"><img src="https://goreportcard.com/badge/github.com/team142/lux-lucet" /></a>&nbsp; 
<a href="https://codeclimate.com/github/team142/lux-lucet/maintainability"><img src="https://api.codeclimate.com/v1/badges/ee3e04d0fac7419ccae9/maintainability" /></a>&nbsp; 
[![License](http://img.shields.io/:license-mit-blue.svg?style=flat)](http://badges.mit-license.org)

System health server written in Go for systems composed of and dependant on subsystems, requiring concurrently updating health status's. 

Features
- Overall system health
- n subsystems
- Thread-safe updating of state
- Thread-safe reading of state
- Http server - listens on address and return state in json

## Usage

See [example.go](lulu/example.go)


### Initial setup

```
import "github.com/team142/lux-lucet/lulu"

...    
	healthServer := lulu.StartHealthServer()

	healthServer.UpdateOk("net/io")
	healthServer.UpdateOk("disk/io")
	healthServer.UpdateOk("queue-handler")
...
```

### Getting state

```
state := healthServer.Query()
b, _ := json.Marshal(state)
log.Println(string(b))
```

Output:
```
{
   "ok":true,
   "subsystems":[
      {
         "name":"net/io",
         "ok":true,
         "msg":""
      },
      {
         "name":"disk/io",
         "ok":true,
         "msg":""
      },
      {
         "name":"queue-handler",
         "ok":true,
         "msg":""
      }
   ]
}
```

### Updating bad health

```
...
err := someWork()
if err != nil {
    healthServer.Update("runQueueHandler", false, err.Error())
    return
}
...
```



Output:
```
{
   "ok":false,
   "subsystems":[
      {
         "name":"net/io",
         "ok":true,
         "msg":""
      },
      {
         "name":"disk/io",
         "ok":true,
         "msg":""
      },
      {
         "name":"queue-handler",
         "ok":false,
         "msg":"some error"
      }
   ]
}
```

### Starting the web server
This listens on the supplied address and returns the json of the server state on request.
```
...
lulu.StartRestServer(":9001", healthServer) //Blocking call
```
