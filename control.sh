#!/bin/bash
cd "$(dirname $0)"/.. || exit 1
PROC_NAME=harbour
START_COMMAND='./bin/harbour -c ./etc/harbour.json'
PROC_PORT=8080
WAIT_TIME=5

help(){
    echo "${0} <start|stop|restart|status>"
    exit 1
}

checkhealth(){
    if [[ -n "$PROC_PORT" ]] ; then
        PORT_PROC=$(/usr/sbin/ss -nlpt | grep -w :$PROC_PORT | wc -l)
        echo "[proc]$PORT_PROC"
        if [ "$PORT_PROC" = "1" ] ; then
                echo "running"
            return 0
        fi
        echo "[Port]not running"
        return 1
   else
       ps -eo comm,pid |grep -P  "^$PROC_NAME\b"
       if [ "$?" = 0 ] ; then
       echo "running"
           return 0
       fi
       echo "[Pid]not running"
       return 1
   fi
}

start(){
    checkhealth
    if [ $? = 0 ]; then
        echo "[WARN] $PROC_NAME is aleady running!"
        return 0
    fi
    mkdir -p log

    ./bin/gwc $START_COMMAND  </dev/null &>>./log/${PROC_NAME}_output.log

    for i in $(seq $WAIT_TIME) ; do
        sleep 1
        checkhealth
        if [ $? = 0 ]; then
            echo "Start $PROC_NAME success"
            return 0
        fi
    done
    echo "[ERROR] Start $PROC_NAME failed"
    return 1
}

stop(){
    if [[ -n "$PROC_PORT"  ]] ; then
        PROC_ID=$(  /usr/sbin/ss -nltp "( sport = :$PROC_PORT )" |sed 1d  | awk '{print $NF}' |  grep -oP '\,.*\,' | grep -oP "\d+" |  uniq )
    else
        PROC_ID=$(ps -eo comm,pid  | grep "^$PROC_NAME\b" |awk '{print $2}')
    fi

    GWC_ID=$(ps -f -C gwc | awk '{print $2}' | grep -v PID )

    if [[ -z "$PROC_ID" ]] ; then
        echo "[WARN] $PROC_NAME is aleady exit, skip stop"
        return 0
    fi

    checkhealth
    if [ "$?" != "0" ] ; then
        echo "[WARN] $PROC_NAME is aleady exit, skip stop"
        return 0
    fi
    kill $PROC_ID $GWC_ID
    for i in $(seq $WAIT_TIME) ; do
        sleep 1
        checkhealth
        if [ "$?" != "0" ] ; then
            echo "Stop $PROC_NAME success"
            return 0
        fi
    done

    kill -9 $PROC_ID $GWC_ID
    sleep 1
    checkhealth
    if [ "$?" != "0" ] ; then
        echo "Stop $PROC_NAME success"
        return 0
    fi

    echo "[ERROR] Stop $PROC_NAME failed"
    return 1
}

case "${1}" in
    start)
        start
        ;;
    stop)
        stop
        ;;
    status|health|checkhealth)
        checkhealth
        ;;
    restart)
        stop && start
        ;;
    *)
        help
        ;;
esac
