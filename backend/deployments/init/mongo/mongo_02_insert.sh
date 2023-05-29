#!/bin/sh

mongoimport -d DoramaSet -c staff --type csv --file csv/staff.csv \
--fields="id.int32(),name.string(),birthday.date(2006-01-02),gender.string(),type.string()" \
--columnsHaveTypes


cp csv/dorama.csv csv/copy.csv
sed -i "y/\"/'/" csv/copy.csv
sed -i "y/\,/@/" csv/copy.csv
sed -i "y/\;/,/" csv/copy.csv
sed -i "y/\@/;/" csv/copy.csv
mongoimport -d DoramaSet -c dorama --type csv --file csv/copy.csv --fields="id,name,description,release_year,status,genre"
rm csv/copy.csv

cp csv/subscription.csv csv/copy.csv
sed -i "y/\;/,/" csv/copy.csv
mongoimport -d DoramaSet -c subscription --type csv --file csv/copy.csv --fields="id,name,cost,description,duration,access_lvl"
rm csv/copy.csv

mongoimport -d DoramaSet -c user --type csv --file csv/admin.csv \
--fields="username.string(),sub_id.int32(),password.string(),email.string(),registration_date.date(2006-01-02),last_active.date(2006-01-02),last_subscribe.date(2006-01-02),points.int32(),is_admin.boolean(),emoji.string(),color.string()" \
--columnsHaveTypes

# help collection
mongoimport -d DoramaSet -c _picture --type csv --file csv/picture.csv --fields="id,url"
mongoimport -d DoramaSet -c _dorama_picture --type csv --file csv/dorama-picture.csv --fields="id_dorama,id_picture"
mongoimport -d DoramaSet -c _staff_picture --type csv --file csv/staff-picture.csv --fields="id_staff,id_picture"
mongoimport -d DoramaSet -c _episode --type csv --file csv/episode.csv --fields="id,id_dorama,num_season,num_episode"
mongoimport -d DoramaSet -c _dorama_staff --type csv --file csv/dorama-staff.csv --fields="id_dorama,id_staff"
