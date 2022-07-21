#!/bin/bash


download_binary(){
	curl -L -o datasage.tar.gz https://github.com/datasage-io/datasage/releases/download/0.1.0/datasage_0.1.0_linux_amd64.tar.gz
	tar -xvf datasage.tar.gz
	mv datasage /usr/local/bin/
}
download_csv(){
	curl -L -o class.csv https://raw.githubusercontent.com/datasage-io/datasage/main/src/storage/default/class.csv
	curl -L -o tags.csv https://raw.githubusercontent.com/datasage-io/datasage/main/src/storage/default/tags.csv
}

download_config(){
	curl -L -o datasage.yaml https://raw.githubusercontent.com/datasage-io/datasage/main/src/conf/datasage.yaml
}

copy_csv_to_resources(){
	mkdir -p /etc/datasage/resources/
	mv *.csv /etc/datasage/resources/
}

copy_config_to_path(){
	mkdir -p /etc/datasage/conf/
	mv datasage.yaml /etc/datasage/conf/
}


download_binary
download_csv
copy_csv_to_resources
download_config
copy_config_to_path
