# Music Tablature API
A (relatively) simple API for serving music tablature.
Postgres is used as the database.
Database address, port, etc. is configured in the config.yaml 
which is created on first run.

## Configuration
Default `config.yaml` generated by the program:
```yaml
server:
    port: 3000
postgres:
    address: localhost
    port: 5432
    databasename: music-tabs
    username: postgres
    password: postgres
```

## Installation & Usage
### Build from source
First clone and build:
```bash
$ git clone https://github.com/howardliam/music-tab-api && cd music-tab-api
$ go build
$ ./music-tab-api
```
Then you'll need to configure it for your Postgres instance.