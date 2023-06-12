import requests
from bs4 import BeautifulSoup
import csv

def check_genre(genres):
    invalid = ['арт-хаус', 'документальный', 'пародия', 'саспенс', 'ситком', 'гей-тема']
    for x in genres:
        if x in invalid:
            return True
    return False
        

def get_title_info(title_name=''):
# Запрос к странице
    url = 'https://doramatv.live%s' % (title_name)
    headers = { 
        'User-Agent': 'Mozilla/5.0 (Macintosh; Intel Mac OS X 10.9; rv:45.0) Gecko/20100101 Firefox/45.0' 
        }
    r = requests.get(url, headers = headers)
    soup = BeautifulSoup(r.text, 'lxml')
    data = soup.find('div', class_='subject-meta')

    # Названия на двух языках
    try:
        ru_name = soup.find("span", class_="name").contents[0]
    except:
        ru_name = ''
    
    try:
        en_name = soup.find("span", class_="eng-name").contents[0]
    except:
        en_name = ''
    
    name = ru_name if ru_name != '' else en_name

    description = soup.find('div', class_='manga-description').find('div').get_text()
    
    status = False
    # количество серий
    try:
        series_cnt = data.find("p").contents[2].strip()
        status = series_cnt[series_cnt.find(",") + 2:]
        if status == 'завершено':
            status = 'finish'
        else:
            status = 'in progress'
    except:
        status = None  

    # продолжительность одной серии
    try:
        duration = data.find("span", itemprop='duration').contents[0]
        duration = int(duration)
    except:
        duration = None

    # жанры
    genres = data.find_all("span", class_='elem_genre')
    genres = [x.find('a').contents[0] for x in genres]
    if check_genre(genres):
        return []

    if (len(genres) == 0):
        return []
    
    # год
    try:
        year = int(data.find("span", class_='elem_year').find("a").contents[0])
    except:
        year = 0

    # актеры
    actors = []
    # если есть блок "в главых ролях"
    try:
        main_act = data.find_all("span", class_='elem_main_role')
        for x in main_act:
            one_actor = x.find("a").contents + [x.find("a")["href"]]
            actors.append(one_actor)
    except:
        pass
    # все актеры
    show = data.find_all("span", class_='elem_actor')
    for x in show:
        one_actor = x.find("a").contents + [x.find("a")["href"]]
        actors.append(one_actor)
    print(actors)

    # режиссеры
    directors = []
    directors_data = data.find_all("span", class_='elem_director')
    for x in directors_data:
        one_dir = x.find("a").contents + [x.find("a")["href"]]
        directors.append(one_dir)
    print('\n\n\n', directors)

    # сценаристы
    sceen = []
    sceen_data = data.find_all('span', class_='elem_screenwriter')
    for x in sceen_data:
        one_sceen = x.find("a").contents + [x.find("a")["href"]]
        sceen.append(one_sceen)
    print('\n\n\n', sceen)

    """Записывает распасенные данные в файлы для дальнейшей работы с ними"""
    with open("actors.txt", "a", encoding="utf-8") as f: 
        for actor in actors:
            s =  str(actor[1]) + '\n'
            f.write(s)

    with open("directors.txt", "a", encoding="utf-8") as f: 
        for dir in directors:
            s = str(dir[0]) + '|' + str(dir[1]) + '\n'
            f.write(s)

    with open("sceen.txt", "a", encoding="utf-8") as f: 
        for sc in sceen:
            s = str(sc[0]) + '|' + str(sc[1]) + '\n'
            f.write(s)

    return [name, description, year, status, genres[0]]


file_name = 'urls-final.txt'
content = []
with open(file_name, "r", encoding="utf-8") as f:
    lines = f.readlines()
    for line in lines:
        content.append(line.strip())

id_dorama = 0
for line in content:
    print(id_dorama+1, line)
    one = get_title_info(line)
    if len(one) == 0:
        continue
    id_dorama += 1
    # with open("urls-final.txt", "a", encoding="utf-8") as f: 
    #     s = str(line) + '\n'
    #     f.write(s)
    with open('dorama.csv', mode="a", encoding='utf-8') as f:
        file_writer = csv.writer(f, delimiter = ";", lineterminator="\r")
        str_table = [id_dorama] + one
        file_writer.writerow(str_table)
