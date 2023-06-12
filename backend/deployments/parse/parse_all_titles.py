# парсинг названий всех дорам и ссылок на них
import requests
from bs4 import BeautifulSoup

def get_title(index):
    # Запрос к странице
    url = 'https://doramatv.live/list?sortType=POPULARITY&offset=%d' % (index)
    headers = { 
        'User-Agent': 'Mozilla/5.0 (Macintosh; Intel Mac OS X 10.9; rv:45.0) Gecko/20100101 Firefox/45.0' 
        }
    r = requests.get(url, headers = headers)
    soup = BeautifulSoup(r.text, 'lxml')

# поиск все тайтлов на одной странице
    quotes = soup.find_all('div', class_='col-sm-6')
    page = []
    for x in quotes:
        title = [x.find('h3').find('a').contents[0], x.a['href']]
        page.append(title)
    
    return page


all_titles = []
pages_cnt = 10
for i in range(pages_cnt):
    one_page = get_title(70*i)
    all_titles += one_page
    print("Страница", i + 1, "обработана")

# запись в файл
with open("urls.txt", "w", encoding="utf-8") as f: 
    for title in all_titles:
        s = str(title[1]) + '\n'
        f.write(s)