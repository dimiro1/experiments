# Go Health Check


# Running

```sh
$ go run main.go
$ curl http://localhost:8080/health/
```

```json
{
    "status": "up",
    "info": {
        "G1": {
            "status": "up",
            "info": {
                "status": 200
            }
        },
        "Websites": {
            "status": "up",
            "info": {
                "Google": {
                    "status": "up",
                    "info": {
                        "status": 200
                    }
                },
                "GuiaBolso": {
                    "status": "up",
                    "info": {
                        "status": 200
                    }
                }
            }
        }
    }
}
```