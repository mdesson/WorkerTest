import requests
import sys

if len(sys.argv) > 1:
    user_id = sys.argv[1]
else:
    user_id = 1

user_id = sys.argv[1]

r = requests.get(
    f'https://jsonplaceholder.typicode.com/posts/{user_id}/comments')

print(r.text)
