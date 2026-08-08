Gator - Blog aggregator

Requirements:
- Install Postgres
- Install Go

Installation
- Copy this in the cmd 'go install https://github.com/jpabloVega/gator'

Config file set up
- Go to your home directory
- Create a file called .gatorconfig.json
- Edit that file so it looks like this
    {"db_url":"postgres://<username>:<pasword>@localhost:5432/gator?sslmode=disable","current_user_name":""}
        By default the username is the same as your pc username
        If you didnt added a postgres password during installation you can leave it blank

How to use
- From a cmd type 'gator <command>'

Commands 
- register <username>
    Register and login into a new user, will fail if username exists
- login <username>
    Login into an existing user, will fail if no username exists
- users
    Prints the registered users, and the current login user
- addfeed <name> <url>
    Adds feed to the login user
- feeds
    Prints the feed the user has added
- follow <url>
    Follow a feed
- unfollow <url>
    Unfollow a feed
- following
    Prints the users and feeds the login user has followed
- agg <time ex: 10s, 2m, 1h> 
    Given an interval of time, it scrapes the feeds from the added feeds (Ctrl+C) to end
- browser <int>
    Given an amount(default 2), prints an amount of scraped feeds
- reset
    Delete all records, permanent
    







