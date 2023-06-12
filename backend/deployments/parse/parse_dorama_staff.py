import requests
from bs4 import BeautifulSoup
import csv


"""Генерирует дочерние таблицы"""

def get_title_info(title_name=''):
# Запрос к странице
    url = 'https://doramatv.live%s' % (title_name)
    headers = { 
        'User-Agent': 'Mozilla/5.0 (Macintosh; Intel Mac OS X 10.9; rv:45.0) Gecko/20100101 Firefox/45.0' 
        }
    r = requests.get(url, headers = headers)
    soup = BeautifulSoup(r.text, 'lxml')
    data = soup.find('div', class_='subject-meta')

  
    staff = []
    # если есть блок "в главых ролях"
    try:
        main_act = data.find_all("span", class_='elem_main_role')
        for x in main_act:
            one_actor = x.find("a").contents + [x.find("a")["href"]]
            staff.append(one_actor)
    except:
        pass
    # все актеры
    show = data.find_all("span", class_='elem_actor')
    for x in show:
        one_actor = x.find("a").contents + [x.find("a")["href"]]
        staff.append(one_actor)

    directors_data = data.find_all("span", class_='elem_director')
    for x in directors_data:
        one_dir = x.find("a").contents + [x.find("a")["href"]]
        staff.append(one_dir)

    sceen_data = data.find_all('span', class_='elem_screenwriter')
    for x in sceen_data:
        one_sceen = x.find("a").contents + [x.find("a")["href"]]
        staff.append(one_sceen)

    return staff

def check_actor(all_actors, one_actor):
    for x in all_actors:
        if x[1] == one_actor[0]:
            return x[0]
    return None
    


file1 = './data/staff.csv'
actor_table = []
with open(file1, mode="r", encoding='utf-8') as w_file:
    file_reader = csv.reader(w_file, delimiter = ",")
    cnt = 0
    for row in file_reader:
        actor_table.append(row)

file_name = 'urls-final.txt'
content = []
with open(file_name, "r", encoding="utf-8") as f:
    lines = f.readlines()
    for line in lines:
        content.append(line.strip().split("|"))

id_dorama = 0
file2 = 'dorama-staff.csv'
for line in content:
    id_dorama += 1
    one = get_title_info(line[0])
    print('id dorama: ', id_dorama)
    for x in one:
        id_actor = check_actor(actor_table, x)
        if id_actor != None:
            str_table = [id_dorama, id_actor]
            with open(file2, mode="a", encoding='utf-8') as f:
                file_writer = csv.writer(f, delimiter = ",", lineterminator="\r")
                file_writer.writerow(str_table)
    