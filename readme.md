# Goloom 🪼

---

## Goloom - useful tool for online configuration your apps and monitoring their states.

### Services state

*Monitor services health*

![img.png](images/img.png)
### Folders and files

*Setup your complex folder structures*

![img.png](images/img1.png)
### Configurable files

*Change your `json`, `yaml` and `yml` files online and your services will be notified about the changes*

![img.png](images/img3.png)

### Configure

Configure your app files online

---

## Simple setup

### Your app should provide `GET /health` and `POST /provide` routes. 
1. The first one will let `goloom` know if app is alive and should it be 
provided with file updates or not.
2. The second one will let `goloom` make requests with data updates
Body that you will receive
```json
{
  "contentBytes": "your file bytes",
  "path": "./path_to_your_file.extension"
}
```

### To add and remove your app in `goloom` services use 
    
`Post` request    

    http://${goloom_host}:${goloom_port}/service/add?port={$app_port}

`Delete` request

    http://${goloom_host}:${goloom_port}/service/remove?port={$app_port}

```yaml
version: '3.8'

services:
  goloom:
    image: goloom
    ports:
      - '6070:6070'
    environment:
      PORT: 6070
      DIR: './data'
      CHECK_PERIOD: '3s'
```

## Goloom Configuration

### Via webpage

### Via terminal

