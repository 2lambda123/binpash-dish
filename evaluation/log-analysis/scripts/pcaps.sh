pure_func() {
    tempfile=$(mktemp)
    cat > $tempfile 
    tcpdump -nn -r $tempfile -A 'port 53' 2> /dev/null | sort | uniq |grep -Ev '(com|net|org|gov|mil|arpa)' 2> /dev/null
    # extract URL
    tcpdump -nn -r $tempfile -s 0 -v -n -l 2> /dev/null | egrep -i "POST /|GET /|Host:" 2> /dev/null
    # extract passwords
    tcpdump -nn -r $tempfile -s 0 -A -n -l 2> /dev/null | egrep -i "POST /|pwd=|passwd=|password=|Host:" 2> /dev/null

    rm -f $tempfile
}
export -f pure_func


for item in $(hdfs dfs -ls -C $1);
do
    hdfs dfs -cat -ignoreCrc $item | pure_func
done

echo 'done';