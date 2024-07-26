#!/bin/bash
set -e

product_name=wynpay
build_path=./build
run_mode=release
OS_TYPE="Unknown"
GetOSType() {
    osName=$(uname -s | tr '[:upper:]' '[:lower:]')

    case "$osName" in
        darwin*) OS_TYPE="Darwin";;
        linux*) OS_TYPE="Linux";;
        mingw*|msys*) OS_TYPE="Windows";;
        *) OS_TYPE="Unknown";;
    esac

    echo "${OS_TYPE}"
}

GetOSType


function toBuild() {
    if [[ "$OS_TYPE" != "Linux" && "$run_mode" == "release" ]]; then
        echo "release build must to linux OS"
        exit 0
    fi

    rm -rf ${build_path}/${run_mode}
    mkdir -p ${build_path}/${run_mode}

    go_version=$(go version | awk '{print $3}')
    commit_hash=$(git show -s --format=%H)
    commit_date=$(git show -s --format="%ci")

    if [[ "$OS_TYPE" == "Darwin" ]]; then
        # macOS
        formatted_time=$(date -u -j -f "%Y-%m-%d %H:%M:%S %z" "${commit_date}" "+%Y-%m-%d_%H:%M:%S")
    else
        # Linux
        formatted_time=$(date -u -d "${commit_date}" "+%Y-%m-%d_%H:%M:%S")
    fi

    build_time=$(date +"%Y-%m-%d_%H:%M:%S")

    ld_flag_master="-X main.mGitCommitHash=${commit_hash} -X main.mGitCommitTime=${formatted_time} -X main.mGoVersion=${go_version} -X main.mPackageOS=${OS_TYPE} -X main.mPackageTime=${build_time} -X main.mRunMode=${run_mode} -X main.mPKGMode=server -s -w"
    ld_flag_salver="-X main.mGitCommitHash=${commit_hash} -X main.mGitCommitTime=${formatted_time} -X main.mGoVersion=${go_version} -X main.mPackageOS=${OS_TYPE} -X main.mPackageTime=${build_time} -X main.mRunMode=${run_mode} -X main.mPKGMode=client -s -w"

    go build -o ${build_path}/${run_mode}/${product_name}_manage/${product_name}_manage -trimpath -ldflags "${ld_flag_master}" main.go \
    && go build -o ${build_path}/${run_mode}/${product_name}_mobile/${product_name}_mobile -trimpath -ldflags "${ld_flag_salver}" main.go \
    && cp ./config/${product_name}_manage.service ${build_path}/${run_mode}/${product_name}_manage \
    && cp ./config/${product_name}_mobile.service ${build_path}/${run_mode}/${product_name}_mobile \
    && cp ./config/install_${product_name}_manage.sh ${build_path}/${run_mode}/${product_name}_manage \
    && cp ./config/install_${product_name}_mobile.sh ${build_path}/${run_mode}/${product_name}_mobile

    package_files
}

function package_files(){
    cd ${build_path}/${run_mode} \
    && if [[ "$OS_TYPE" == "Windows" ]]; then
            7z a ./${product_name}_manage_${run_mode}.zip ./${product_name}_manage >/dev/null 2>&1 \
            && 7z a ./${product_name}_mobile_${run_mode}.zip ./${product_name}_mobile >/dev/null 2>&1
        else
            zip -r ./${product_name}_manage_${run_mode}.zip ./${product_name}_manage \
            && zip -r ./${product_name}_mobile_${run_mode}.zip ./${product_name}_mobile
        fi \
    && cd ../
}


function handlerunMode() {
    if [[ "$1" == "release" || "$1" == "" ]]; then
        run_mode=release
    elif [[ "$1" == "test" ]]; then
        run_mode=test
    else
        echo "Usage: bash build.sh [release|test],default with:release"
        exit 0
    fi
}


handlerunMode "$1" && toBuild

