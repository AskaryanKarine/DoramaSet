import requests
from bs4 import BeautifulSoup
import csv


def get_person_info(person_name=''):
# Запрос к странице
    url = 'https://doramatv.live%s' % (person_name)
    headers = { 
        'User-Agent': 'Mozilla/5.0 (Macintosh; Intel Mac OS X 10.9; rv:45.0) Gecko/20100101 Firefox/45.0' 
        }
    r = requests.get(url, headers = headers)
    soup = BeautifulSoup(r.text, 'lxml')
    data = soup.find('div', class_='subject-meta')

    name = soup.find('span', class_='name').contents[0].split(" ", 1)

    if len(name) == 1:
        name = name[0]
    else:
        name = name[0] + " " + name[1]

    try:
        birthday = data.find("time")["datetime"]
    except:
        birthday = ''

    try:
        gender = data.find("span", class_='fa-lg').parent.contents[1][2]
    except:
        gender = '-'


    return [name, birthday, gender]


file_name = 'sceen.txt'
content = []
with open(file_name, "r", encoding="utf-8") as f:
    lines = f.readlines()
    for line in lines:
        content.append(line.strip().split("|"))

i = 903
for line in content:
    i += 1
    print(i, line)
    one = get_person_info(line[0])
    one = [i] + one + ['screenwriter']
    with open("staff.csv", mode="a", encoding='utf-8') as f:
        file_writer = csv.writer(f, delimiter = ",", lineterminator="\r")
        file_writer.writerow(one)

