#!/bin/sh
# password_server
#
# copy script to location /etc/init.d/password_server
#

case "$1" in
	start)
		echo "Starting password_server"
		/opt/password/bin/password_server -loglevel=ERROR -port=7777 -rootdir=/rsync > /var/log/password_server.log &
	;;
	stop)
		echo "Stopping password_server"
		pid=`ps ax|grep password_server | grep -v init.d |awk '{ print $1 }'`
		kill $pid  > /dev/null 2>&1
	;;
	restart)
		$0 stop
		sleep 2
		$0 start
	;;
	*)
		echo "Usage: /etc/init.d/password_server {start|stop|restart}"
		exit 1
	;;
esac

exit 0