#!/bin/sh

set -e

# Get env vars in the Dockerfile to show up in the SSH session
eval $(printenv | sed -n "s/^\([^=]\+\)=\(.*\)$/export \1=\2/p" | sed 's/"/\\\"/g' | sed '/=/s//="/' | sed 's/$/"/' >> /etc/profile)

# Start SSH service
# echo "Starting SSH ..."
# /usr/sbin/sshd

# Start FileBeat
# echo "Starting Filebeat ..."
# exec filebeat -c /usr/share/filebeat/filebeat.yml &

# Start Go Epwatch
# exec ls -al app
# echo "Starting Go Epwatch ..."
# exec screen -dmS goapp /usr/share/logstash/app/main
#exec go run /usr/share/logstash/app/cmd/app/main.go &

# Start Logstash pipeline 
echo "Starting Logstash ..."
echo $GO_CONFIG_PATH 
export LS_JAVA_OPTS="-Dls.cgroup.cpuacct.path.override=/ -Dls.cgroup.cpu.path.override=/ $LS_JAVA_OPTS"
# exec logstash &
# exec /usr/share/logstash/app/main -config-path=$GO_CONFIG_PATH -profile=production_config &
exec logstash 

# if [[ -z $1 ]] || [[ ${1:0:1} == '-' ]] ; then
#   exec logstash "$@"
# else
#   exec "$@"
# fi