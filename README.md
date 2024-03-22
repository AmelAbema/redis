# Gedis

Redis-like storage implementation

## Description
The storage is implemented based on the map structure, which allows for the highest speed of operation. In order to implement multi-threaded access to the storage, an mutex is also added to the structure. Automatic deletion from storage is implemented using a timer.
- Stores key-value pairs
- 



## Supported operations:
- Get
- Set
- Delete

## Entity TTL
All the entries will be expired automatically after a certain period of time (10 seconds by default)



# Build info
To build the app, execute following command
```
go build -o app cmd/redis
```

# Run info
To run the app execute following command
```
./app [command]
```
| Command | Usage | Description |
| --- | --- | --- |
|`start`| `./app start` | Starts the entire application with interactive CLI interface |
|`show`| `./app show` | Starts script that show basic functionality of the app (you can also use it for testing or checking healthz) |

