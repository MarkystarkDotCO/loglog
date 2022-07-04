
## 📄 EPwatch Logstash

This custom Logstash with default plugin for build logstash and EPwatch Python , Go code in same containers
```
 Stack
 |-------------|     |------------|     |----------|     |-----------|  
 | Source-logs | --> |  FileBeat  | --> | Logstash | --> | File /mnt |
 |_____________|     |____________|     |__________|     |___________| 
```
- Alpine : 3.15
- Logstash : 8.3.0
- FileBeat : 8.3.0
- openssh-server
- Go : 1.17.0

  
## 🐳 Build image
```
docker build -t ep-logstash .
```
    
## 🐳 Run containers on local docker
```
docker run -it -d -p 6514:6514/udp -p 6514:6514/tcp -p 9600:9600 epwatch-logstash
```

## ⚙️ Configure file for logstash
📁 config > 📁 logstash > logstash.yml

## ⚙️ Configure file for logstash pipeline

- 📁 config > 📁 logstash > pipelines.yml (for create and setting pipeline)
- 📁 config > 📁 pipeline > default.conf (store each pipeline config)

## ⚙️ Configure file for filebeat
📁 config > 📁 filebeat > filebeat.yml

## ⚙️ Configure file for startup service
You can edit in **`logstash-entrypoint.sh`** file use **`exec command`** and **`&`** for run in background

## Port

- Logstash : 0.0.0.0:9600
- pipeline : 0.0.0.0:6514/udp
- ssh : 0.0.0.0:2222 (for azure)
