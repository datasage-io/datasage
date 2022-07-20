#!/bin/bash

download_csv(){
	curl -L -o class.csv https://raw.githubusercontent.com/datasage-io/datasage/main/src/storage/default/class.csv
	curl -L -o tags.csv https://raw.githubusercontent.com/datasage-io/datasage/main/src/storage/default/tags.csv
}

copy_csv_to_resources(){
	mkdir -p /etc/datasage/resources/
	mv *.csv /etc/datasage/resources/
}

download_csv
copy_csv_to_resources
