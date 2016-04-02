#!/bin/sh
# password_status_server
#
# copy script to location /etc/init.d/password_status_server
#

case "$1" in
	start)
		echo "Starting password_status_server"
		/opt/password/bin/password_status_server -loglevel=ERROR -port=7777 -rootdir=/rsync > /var/log/password_status_server.log &
	;;
	stop)
		echo "Stopping password_status_server"
		pid=`ps ax|grep password_status_server | grep -v init.d |awk '{ print $1 }'`
		kill $pid  > /dev/null 2>&1
	;;
	restart)
		$0 stop
		sleep 2
		$0 start
	;;
	*)
		echo "Usage: /etc/init.d/password_status_server {start|stop|restart}"
		exit 1
	;;
esac

exit 0