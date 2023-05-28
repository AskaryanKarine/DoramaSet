#!/bin/sh

mongoimport -d DoramaSet -c staff --type csv --file csv/staff.csv --fields="id","name","birthday","gender","type"
sed -i "y/\"/'/" csv/dorama.csv
sed -i "y/\,/@/" csv/dorama.csv
sed -i "y/\;/,/" csv/dorama.csv
sed -i "y/\@/;/" csv/dorama.csv
mongoimport -d DoramaSet -c dorama --type csv --file csv/dorama.csv --fields="id","name","description","release_year","status","genre"
tr ";" "," < csv/subscription.csv | mongoimport -d DoramaSet -c subscription --type csv --file csv/subscription.csv --fields="id","name","cost","description","duration","access_lvl"
mongoimport -d DoramaSet -c user --type csv --file csv/admin.csv --fields="username","sub_id","password","email","registration_date","last_active","last_subscribe","points","is_admin","emoji","color"
# temporary collection
mongoimport -d DoramaSet -c _picture --type csv --file csv/picture.csv --fields="id","url"
mongoimport -d DoramaSet -c _dorama_picture --type csv --file csv/dorama-picture.csv --fields="id_dorama","id_picture"
mongoimport -d DoramaSet -c _staff_picture --type csv --file csv/staff-picture.csv --fields="id_staff","id_picture"
mongoimport -d DoramaSet -c _episode --type csv --file csv/episode.csv --fields="id","id_dorama","num_season","num_episode"
mongoimport -d DoramaSet -c _dorama_staff --type csv --file csv/dorama-staff.csv --fields="id_dorama","id_picture"
