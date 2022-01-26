import requests


print(requests.post('https://httpbin.org/post', json={'key': 'value'}).text)
