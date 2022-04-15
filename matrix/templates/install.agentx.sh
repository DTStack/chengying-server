#!/bin/bash
set -ex

############################
#Copyright (c) 2017 DTStack Inc.
#Version  0.1
############################

agent_zip='{{.AGENT_ZIP}}'

#此参数值被程序使用，修改时请@huanxiong
app_dir='{{.AGENT_DIR}}'
agent_bin='{{.AGENT_BIN}}'
run_user='{{.RUN_USER}}'
data_dir='{{.DATA_DIR}}'

#安装包下载地址
DOWNLOAD_URL='{{.AGENT_DOWNLOAD_URL}}'

trap '[ "$?" -eq 0 ] || read -p "Looks like something went wrong in step ´$STEP´"' EXIT

##install the filebeat##
install_agent() {
    mkdir -p "$app_dir"
    unzip -o "/tmp/$agent_zip" -d "$app_dir"  >/dev/null 2>&1
}

##download and installed##
install(){
    STEP='install agent'
    echo "Use the curl download and install Please Waiting..."
    cd /tmp/ && curl -L -O -s "$DOWNLOAD_URL"
    install_agent

    if [ ! -f "$agent_bin" ];then
        echo "cmd: $agent_bin not found!"
        exit 1
    fi
}

##chown##
chowns(){
    STEP='chown'
    if [ -n "$run_user" ];then
        sudo chown -R "$run_user:$run_user" "$app_dir"
        if [ -n "$data_dir" ];then
            for path in `echo $data_dir`
            do
                sudo mkdir -p $path
                sudo chown -R "$run_user:$run_user" $path
            done
        fi
    else
        if [ -n "$data_dir" ];then
            for path in `echo $data_dir`
            do
                sudo mkdir -p $path
            done
        fi
    fi
}

##delete filebeat pkg##
delete(){
    STEP='delete'
    cd /tmp/ && rm -f "$agent_zip"
}

install
chowns
delete
