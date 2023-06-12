import requests
from bs4 import BeautifulSoup
import csv


def get_pic(title_name=''):
# Запрос к странице
    url = 'https://doramatv.live%s' % (title_name)
    headers = { 
        'User-Agent': 'Mozilla/5.0 (Macintosh; Intel Mac OS X 10.9; rv:45.0) Gecko/20100101 Firefox/45.0' 
        }
    r = requests.get(url, headers = headers)
    soup = BeautifulSoup(r.text, 'lxml')
    try:
        data = soup.find('div', class_='subject-cover').find_all('img')
    except:
        data = []
    print(data)

    pics = []
    for x in data:
        pics.append(x['data-full'])

    return pics



file_name = 'sceen.txt'
content = []
with open(file_name, "r", encoding="utf-8") as f:
    lines = f.readlines()
    for line in lines:
        content.append(line.strip())

id_staff = 903
id_picture = 4207
for line in content:
    print(id_staff+1, line)
    one = get_pic(line)
    id_staff += 1
    for x in one:
        id_picture += 1
        with open('picture.csv', mode="a", encoding='utf-8') as f:
            file_writer = csv.writer(f, delimiter = ",", lineterminator="\r")
            str_table = [id_picture, x]
            file_writer.writerow(str_table)

        with open('staff-picture.csv', mode="a", encoding='utf-8') as f:
            file_writer = csv.writer(f, delimiter = ",", lineterminator="\r")
            str_table = [id_staff, id_picture]
            file_writer.writerow(str_table)
