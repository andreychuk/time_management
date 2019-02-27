# TIME MANAGEMENT

> 

## About
Log time to Redmine issues

## Command line parameters

    issueID - Issue ID, required
    hours   - Hours you spend on an issue (3.5), required
    spentOn - Date in format YYYY-MM-DD, if not specify will use current date
    activity - Activity ID, required
    comment - COmment to log time

## Using with Docker

    Add your credentials to .env-conf

    endpoint=HOST_TO_YOUR_REDMINE
    apikey=YOUR_API_KEY

    docker run -i --rm --env-file .env-conf andreychuk/time-management --issueID=issue_id --hourd=3.5 --activity=activity_id

