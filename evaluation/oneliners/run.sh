#!/bin/bash

export DISH_TOP=$(realpath $(dirname "$0")/../../..)
export PASH_TOP=$(realpath $DISH_TOP/pash)
export TIMEFORMAT=%R
cd "$(dirname "$0")"

if [[ "$@" == *"--small"* ]]; then
    scripts_inputs=(
        "nfa-regex;1M"
        "sort;1M"
        "top-n;1M"
        "wf;1M"
        "spell;1M"
        "diff;1M"
        "bi-grams;1M"
        "set-diff;1M"
        "sort-sort;1M"
        "shortest-scripts;all_cmds"
    )
else
    scripts_inputs=(
        "nfa-regex;1G"
        "sort;3G"
        "top-n;3G"
        "wf;3G"
        "spell;3G"
        "diff;3G"
        "bi-grams;3G"
        "set-diff;3G"
        "sort-sort;3G"
        "shortest-scripts;all_cmdsx100"
    )
fi

mkdir -p "outputs"

oneliners_bash() {
    echo executing oneliners $(date)

    mkdir -p "outputs/bash"

    for script_input in ${scripts_inputs[@]}
    do
        IFS=";" read -r -a parsed <<< "${script_input}"
        script_file="./scripts/${parsed[0]}.sh"
        input_file="/oneliners/${parsed[1]}.txt"
        output_file="./outputs/bash/${parsed[0]}.out"
        time_file="./outputs/bash/${parsed[0]}.time"
        log_file="./outputs/bash/${parsed[0]}.log"

        (time $script_file $input_file > $output_file) 2> $time_file

        echo "$script_file $(cat "$time_file")" 
    done
}

PASH_FLAGS='--width 8 --r_split'

oneliners_pash(){
    flags=${1:-$PASH_FLAGS}
    prefix=${2:-par}
    prefix=$prefix

    times_file="$prefix.res"
    outputs_suffix="$prefix.out"
    time_suffix="$prefix.time"
    outputs_dir="outputs"
    pash_logs_dir="pash_logs_$prefix"

    mkdir -p "$outputs_dir"
    mkdir -p "$pash_logs_dir"

    touch "$times_file"
    cat $times_file >> $times_file.d
    echo executing one-liners with $prefix pash with data $(date) | tee "$times_file"
    echo '' >> "$times_file"

    for script_input in ${scripts_inputs[@]}
    do
        IFS=";" read -r -a script_input_parsed <<< "${script_input}"
        script="${script_input_parsed[0]}"
        input="${script_input_parsed[1]}"

        export IN="/oneliners/$input"
        export dict=

        printf -v pad %30s
        padded_script="${script}.sh:${pad}"
        padded_script=${padded_script:0:30}

        outputs_file="${outputs_dir}/${script}.${outputs_suffix}"
        pash_log="${pash_logs_dir}/${script}.pash.log"
        single_time_file="${outputs_dir}/${script}.${time_suffix}"

        echo -n "${padded_script}" | tee -a "$times_file"
        { time "$PASH_TOP/pa.sh" $flags --log_file "${pash_log}" ${script}.sh > "$outputs_file"; } 2> "${single_time_file}"
        cat "${single_time_file}" | tee -a "$times_file"
    done
}

# For testing purposes
# hdfs dfs -rm -r "/outputs/hadoop-streaming/oneliners"
# hadoop jar "/opt/hadoop-3.4.0/share/hadoop/tools/lib/hadoop-streaming-3.4.0.jar" -files nfa-regex.sh -D mapred.reduce.tasks=0 -D dfs.checksum.type=NULL -input "/oneliners/1G.txt" -output "/outputs/hadoop-streaming/oneliners/nfa-regex" -mapper nfa-regex.sh # nfa-regex
oneliners_hadoopstreaming(){
    jarpath="/opt/hadoop-3.4.0/share/hadoop/tools/lib/hadoop-streaming-3.4.0.jar" # Adjust as required
    basepath="/oneliners" # Adjust as required
    times_file="hadoopstreaming.res"
    outputs_suffix="hadoopstreaming.out"
    outputs_dir="/outputs/hadoop-streaming/oneliners"
    . bi-gram.aux.sh

    cd "hadoop-streaming/"

    hdfs dfs -rm -r "$outputs_dir"
    hdfs dfs -mkdir -p "$outputs_dir"

    touch "$times_file"
    cat "$times_file" >> "$times_file".d
    echo executing oneliners $(date) | tee "$times_file"
    echo '' >> "$times_file"

    while IFS= read -r line; do
        printf -v pad %20s
        name=$(cut -d "#" -f2- <<< "$line")
        name=$(sed "s/ //g" <<< $name)
        padded_script="${name}.sh:${pad}"
        padded_script=${padded_script:0:20} 
        echo "${padded_script}" $({ time { eval $line &> /dev/null; } } 2>&1) | tee -a "$times_file"
    done <"run_all.sh"
    cd ".."
    mv "hadoop-streaming/$times_file" .
}

oneliners_bash

# oneliners_pash "$PASH_FLAGS" "par"

# oneliners_pash "$PASH_FLAGS --distributed_exec" "distr"

# oneliners_hadoopstreaming
