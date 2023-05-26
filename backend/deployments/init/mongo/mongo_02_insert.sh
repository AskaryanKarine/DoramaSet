#!/bin/sh

mongoimport -d DoramaSet -c staff --type csv --file csv/staff.csv --fields="_id","name","birthday","gender","type"
tr ";" "\t" < csv/dorama.csv | mongoimport -d DoramaSet -c Dorama --type csv --file csv/dorama.csv \
--fields="_id","name","description","release_year","status","genre"