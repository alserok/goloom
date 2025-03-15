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
      SERVICES: '127.0.0.1:6000,172.17.0.1:6001'
      CHECK_PERIOD: '3s'
```

## Goloom Configuration

### Via webpage

### Via terminal

