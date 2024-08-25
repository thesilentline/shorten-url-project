# URL Shortener

### Description:

URL shortener is an API service which requests URL from the user to shorten and returns shorten url with random ID or the custom ID entered by the user. Redis database is used to store the URL entered and corresponding shortened URL. The user has a quota of 10 requests to be passed in duration of 30 minutes. 

### Installation:

- clone the repository

```bash
git clone https://github.com/thesilentline/shorten-url-project.git
```

- to build and run the docker container

```bash
docker-compose up -d 
```

### Screenshot:

![image.png](Screenshot/3c8957b5-1414-41b2-b691-bf069924832a.png)

![image.png](Screenshot/b34c841f-8934-4d80-9a44-c346c905b978.png)

![image.png](Screenshot/image.png)

![image.png](Screenshot/image 1.png)

### Tech Stack:

Go | Fiber | Redis | Docker