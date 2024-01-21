import requests
import time

def main():
  payloads = [
    {
      'stage-name': 'Index-stage',
      'url1': 'http://localhost:8001/staff/1',
      'url2': 'http://localhost:8002/staff/1'
    },
    {
      'stage-name': 'Non-index-stage',
      'url1': 'http://localhost:8001/find/staff?name=Пак',
      'url2': 'http://localhost:8002/find/staff?name=Пак'
    },
  ]
  for payload in payloads:

      headers = requests.utils.default_headers()

      for i in range(100):
        print(i)
        r = requests.get(payload["url1"],headers=headers)
        r = requests.get(payload["url2"],headers=headers)
        time.sleep(10)

  print('process ended')

if __name__ == '__main__':
  main()
