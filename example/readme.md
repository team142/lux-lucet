## Usage

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
