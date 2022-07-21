#!/bin/bash

VERSION=$1

latest_datasage_download_url(){
	curl -s https://api.github.com/repos/datasage-io/datasage/releases/latest \
		| grep "browser_download_url.*tar.gz" \
		| cut -d : -f 2,3 \
		| tr -d \"
}

latest_datasage_cli_download_url(){
	curl -s https://api.github.com/repos/datasage-io/datasage-cli/releases/latest \
		| grep "browser_download_url.*tar.gz" \
		| cut -d : -f 2,3 \
		| tr -d \"
}

download_datasage_latest(){
	echo "Downloading latest datasage release...$VERSION"
	curl -s -L -o datasage.tar.gz $(latest_datasage_download_url)
	tar -xvf datasage.tar.gz
	mv datasage /usr/local/bin/
}

download_datasage_cli_latest(){
	echo "Downloading latest datasage-cli..."
	curl -s -L -o datasage-cli.tar.gz $(latest_datasage_cli_download_url) 
	tar -xvf datasage-cli.tar.gz
	mv datasage-cli /usr/local/bin/
}

download_datasage_version(){
	echo "Downloading datasage version: $VERSION"
	curl -s -L -o datasage.tar.gz "https://github.com/datasage-io/datasage/releases/download/${VERSION}/datasage_${VERSION}_linux_amd64.tar.gz"
	tar -xvf datasage.tar.gz
	mv datasage /usr/local/bin/
}

download_datasage_cli_version(){
	echo "Downloading datasage-cli version: $VERSION"
	curl -s -L -o datasage-cli.tar.gz "https://github.com/datasage-io/datasage-cli/releases/download/${VERSION}/datasage-cli_${VERSION}_linux_amd64.tar.gz"
	tar -xvf datasage-cli.tar.gz
	mv datasage-cli /usr/local/bin/
}

download_csv(){
	echo "Downloading csv files..."
	curl -s -L -o class.csv https://raw.githubusercontent.com/datasage-io/datasage/main/src/storage/default/class.csv
	curl -s -L -o tags.csv https://raw.githubusercontent.com/datasage-io/datasage/main/src/storage/default/tags.csv
}

download_config(){
	echo "Downloading config files..."
	curl -s -L -o datasage.yaml https://raw.githubusercontent.com/datasage-io/datasage/main/src/conf/datasage.yaml
}

copy_csv_to_resources(){
	mkdir -p /etc/datasage/resources/
	mv class.csv /etc/datasage/resources/
	mv tags.csv /etc/datasage/resources/
}

copy_config_to_path(){
	mkdir -p /etc/datasage/conf/
	mv datasage.yaml /etc/datasage/conf/
}

cleanup(){
	rm datasage.tar.gz datasage-cli.tar.gz 
	echo "Installation complete!"
}

if [ -z "$1" ]; then
    download_datasage_latest
	download_datasage_cli_latest
else
    download_datasage_version
	download_datasage_cli_version
fi
download_csv
download_config
copy_csv_to_resources
copy_config_to_path
cleanup
