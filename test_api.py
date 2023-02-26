import requests


baseURL = "https://api.kinopoisk.dev/v1"
headers={
    "accept": "application/json",
    "X-API-KEY": "BP6G016-ZDTMT05-HD4M5GP-28C9CC3"
}
params = {
    "selectFields":"id name",
    "premiere.country": "Корея",
    "limit": "40",  
    "name": "!null"
}
a = requests.get(baseURL + "/movie", headers=headers, params=params)
print(a, a.text)