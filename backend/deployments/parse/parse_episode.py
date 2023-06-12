import requests
from bs4 import BeautifulSoup
import csv


def get_episode(title_name=''):
# Запрос к странице
    url = 'https://doramatv.live%s' % (title_name)
    headers = { 
        'User-Agent': 'Mozilla/5.0 (Macintosh; Intel Mac OS X 10.9; rv:45.0) Gecko/20100101 Firefox/45.0' 
        }
    r = requests.get(url, headers = headers)
    soup = BeautifulSoup(r.text, 'lxml')
    try:
        data = soup.find('div', class_='chapters-link chapters').find('table').find_all('tr')
    except:
        data = []
    # print(data)

    eps = []
    for x in data:
        eps.append([int(x['data-vol'])+1, int(x['data-num'])//10])
        # eps.append(x['data-full'])

    return eps



file_name = 'urls-final.txt'
content = []
with open(file_name, "r", encoding="utf-8") as f:
    lines = f.readlines()
    for line in lines:
        content.append(line.strip())

id_dorama = 0
id_episode = 0
for line in content:
    print(id_dorama+1, line)
    one = get_episode(line)
    id_dorama += 1
    for x in one:
        id_episode += 1
        with open('episode.csv', mode="a", encoding='utf-8') as f:
            file_writer = csv.writer(f, delimiter = ",", lineterminator="\r")
            str_table = [id_episode, id_dorama] + x
            file_writer.writerow(str_table)


