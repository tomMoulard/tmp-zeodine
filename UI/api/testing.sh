banner "newuser"

curl localhost:80/newuser

banner "createws"

curl -H "Content-Type: application/json" -X POST -d \
'{"user_id":1, "ws_name":"test ws -> :)"}' \
localhost:80/createws

banner "ws"

curl -H "Content-Type: application/json" -X POST -d \
'{"user_id":1}' \
localhost:80/ws

banner "save ws"

curl -H "Content-Type: application/json" -X POST -d \
'{"user_id": 1,"ws_id": 1,"groupes": [{"groupe_id": 12,"cards": [{"card_id": 1,"card_pub": false,"card": {"card_content": "{}"}},{"card_id": 2,"card_pub": false,"card": {"card_content": "{}"}},{"card_id": 3,"card_pub": false,"card": {"card_content": "{}"}},{"card_id": 4,"card_pub": false,"card": {"card_content": "{}"}}]},{"groupe_id": 21,"cards": [{"card_pub": false,"card_id": 5,"card": {"card_content": "{\"card_pos\":12}"}}]}]}' \
localhost:80/save

banner "load ws"

curl -H "Content-Type: application/json" -X POST -d \
'{"user_id":1, "ws_id":1}' \
localhost:80/load

banner "nbcard ws"

curl -H "Content-Type: application/json" -X POST -d \
'{"user_id":1, "ws_id":1}' \
localhost:80/nbcard

banner "tag"

curl -H "Content-Type: application/json" -X POST -d \
'{"stack_id":1, "tag_val":"I am tagged"}' \
localhost:80/tag

curl -H "Content-Type: application/json" -X POST -d \
'{"stack_id":1, "tag_val":"Second tag for me !"}' \
localhost:80/tag

banner "gettag"

curl -H "Content-Type: application/json" -X POST -d \
'{"stack_id":1}' \
localhost:80/gettag
