curl -s https://learn.zone01oujda.ma/assets/superhero/all.json | jq -r '.[] | select(.id==($ENV.HERO_ID | tonumber)) | .connections.relatives'
