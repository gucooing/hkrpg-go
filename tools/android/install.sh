mkdir hkrpg
chmod 777 ./hkrpg
cd ./hkrpg
mkdir conf
chmod 777 ./conf
cd ./conf
wget https://oss.mihoyu.cn/hkrpg/2.5.0/server/hkrpg-go-pe.json
cd ../
mkdir resources
cd ./resources
wget https://oss.mihoyu.cn/hkrpg/2.5.0/server/resources.zip
unzip resources.zip
cd ../
chmod 777 ./resources
mkdir data
cd ./data
wget https://oss.mihoyu.cn/hkrpg/2.5.0/server/data.zip
unzip data.zip
cd ../
chmod 777 ./data
wget https://oss.mihoyu.cn/hkrpg/2.5.0/server/linux-arm64/hkrpg-go
chmod 777 ./hkrpg-go