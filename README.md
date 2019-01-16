# 💀 Mortis Lux 💡
<img src="https://goreportcard.com/badge/github.com/team142/mortis-lux" />&nbsp; 

Health Monitoring in Golang.

Thread-safe system health server that can be easily added to any project.

Features
- Overall system health
- n subsystems
- Thread-safe updating of state
- Thread-safe reading of state
- Can list on address and return state in json

## Usage

### Initial setup

```
import "github.com/team142/mortis-lux"
    
func main() {
	healthServer := molu.StartHealthServer()

	healthServer.UpdateOk("net/io")
	healthServer.UpdateOk("disk/io")
	healthServer.UpdateOk("queue-handler")
}
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
molu.StartRestServer(":9001", healthServer) //Blocking call
```
