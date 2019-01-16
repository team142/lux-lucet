# 💀 Mortis Lux 💡

Health Monitoring in Golang.

Thread-safe system health server that can be easily added to any project.


## Usage

### Initial setup

```
import "github.com/team142/mortis-lux"
    
func main() {
	healthServer := molu.StartHealthServer()

	healthServer.UpdateOk("net/io")
	healthServer.UpdateOk("disk/io")
	healthServer.UpdateOk("queue-handler")
```

### Getting state

```
state := healthServer.Query()
b, err := json.Marshal(state)
if err != nil {
    log.Fatalln(err)
}
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
