Version:1.0 StartHTML:0000000128 EndHTML:0000010500 StartFragment:0000000128 EndFragment:0000010500 SourceURL:about:blank
product_name=wynpay_webSite
build_path=./release
build_path_dev=${build_path}_dev
current_os=`go env GOOS`

function build() {
    if [ "${current_os}" = "linux" ] || [ "${current_os}" = "darwin" ]
    then
      echo "当前平台:${current_os}，开始打包"
      rm -rf ${build_path}
      go build -o ${build_path}/${product_name}/${product_name} main.go
      mkdir ${build_path}/${product_name}/config
      cp -rf ./config/config.yaml ${build_path}/${product_name}/config
      cp -rf ./install.sh ${build_path}/${product_name}
      zip -r ${product_name}.zip ${build_path}/${product_name}
      mv -f ./${product_name}.zip ${build_path}
    else
      echo "当前平台:${current_os},仅限Linux平台下打包"
    fi
}

function build_Linux() {
    if [ "${current_os}" = "linux" ] || [ "${current_os}" = "darwin" ]
        then
          echo "当前平台:${current_os}，开始打包"
          rm -rf ${build_path}
          env GOOS=linux GOARCH=amd64 go build -o ${build_path}/${product_name}_linux/${product_name}_linux main.go
          mkdir ${build_path}/${product_name}_linux/config
          cp -rf ./config/config.yaml ${build_path}/${product_name}_linux/config
          cp -rf ./install.sh ${build_path}/${product_name}_linux
          cp -rf ./Dockerfile ${build_path}/${product_name}_linux
          zip -r ${product_name}_linux.zip ${build_path}/${product_name}_linux
          mv -f ./${product_name}_linux.zip ${build_path}
        else
          echo "当前平台:${current_os},仅限Linux平台下打包"
        fi
}



echo "============================ ${produckName} ============================"
echo "  1、打包 ${produckName}" build 包 本机环境
echo "  2、打包 ${produckName}" build 包 Linux环境
echo "======================================================================"
read -p "$(echo -e "请选择[1-2]：")" choose
case $choose in
1)
    build
    ;;
2)
    build_Linux
    ;;
*)
    echo "输入错误，请重新输入！"
    ;;
esac


