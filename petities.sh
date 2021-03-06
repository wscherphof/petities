#!/bin/bash

# Do hand edit all named items to reflect your actual context,
# since they're all hard coded in this script.

case "$1" in

up)
	exec > >(tee -i petities.log)
	exec 2>&1
	echo -n "* start: " && date +"%H:%M"
	export DIGITALOCEAN_SIZE=2gb
	(	# petities
		essix nodes -d digitalocean -F -m 1 -w 2 create petities
		essix jmeter perfmon start petities & perfmon=$!
		essix r -n 2 create petities
		essix \
			-e DOMAIN=petities.wscherphof.nl \
			-e DB_POOL_INITIAL=100 -e DB_POOL_MAX=1000 \
			-e DB_SHARDS=1 -e DB_REPLICAS=3 \
			-e RATELIMIT=0 -e GO_ENV=test \
			run wscherphof 0.1 petities & petities=$!
		wait ${perfmon} ${petities}
		essix jmeter perfmon restart petities
	) &
	(	# slave
		essix nodes -d digitalocean -F -m 1 -w 5 create slave
		essix jmeter perfmon start slave & perfmon=$!
		essix jmeter server start slave & server=$!
		wait ${perfmon} ${server}
		essix jmeter perfmon restart slave
	) &
	(	# master
		essix nodes -d digitalocean -F -m 1 create master
		essix jmeter server start master
		essix jmeter server stop master
	) &
	wait
	echo -n "* end: " && date +"%H:%M"
;;

down)
	if [ "$2" = "force" ]; then
		answer="Y"
	else
		exec < /dev/tty
		read -p "Type Y to remove ALL nodes in swarms 'petities', 'slave', and 'master'... " answer
	fi
	if [ "$answer" = "Y" ]; then
		essix nodes -f rm petities &
		essix nodes -f rm slave &
		essix nodes -f rm master &
		wait
	fi
;;

esac
