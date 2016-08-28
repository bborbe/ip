#!/bin/sh
# ip_server
#
# copy script to location /etc/init.d/ip_server
#

case "$1" in
	start)
		echo "Starting ip_server"
		/opt/ip/bin/ip_server -logtostderr -v=2 -port=7777 -rootdir=/rsync > /var/log/ip_server.log &
	;;
	stop)
		echo "Stopping ip_server"
		pid=`ps ax|grep ip_server | grep -v init.d |awk '{ print $1 }'`
		kill $pid  > /dev/null 2>&1
	;;
	restart)
		$0 stop
		sleep 2
		$0 start
	;;
	*)
		echo "Usage: /etc/init.d/ip_server {start|stop|restart}"
		exit 1
	;;
esac

exit 0
