## BoltDB Web Editor

### HomePage
https://github.com/boltdbwebeditor/boltdbwebeditor/tree/develop

### DockerHub
https://hub.docker.com/r/boltdbwebeditor/boltdbwebeditor

### Usage:
```
./boltdbwebeditor:
  -db string
        Bolt Database path (default "bolt.db")
```

### Example:
`./boltdbwebeditor`

`./boltdbwebeditor -db=~/home/bolt.db`


### Docker Usage:
```
docker run \
-p 8080:8080 \
-v ~/Work/devkit/data-ee/portainer.db:/data/bolt.db \
boltdbwebeditor/boltdbwebeditor:dev \
-db /data/bolt.db
```
