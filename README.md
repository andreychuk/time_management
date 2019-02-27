#TIME MANAGEMENT

Log time to Readmine issues

```
Using DOCKER

Add your credentials to .env-conf

endpoint=HOST_TO_YOUR_REDMINE
apikey=YOUR_API_KEY

docker run -i --rm --env-file .env-conf andreychuk/time-management --issueID=issue_id --hourd=3.5 --activity=activity_id

```