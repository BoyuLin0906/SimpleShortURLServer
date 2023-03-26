# Simple Short URL Server

## Setup

1. Redis :
   - Purpose : Save a pair of short URL and original URL
   - Environment : WSL2 - Ubuntu 20.04.4 LTS
   - Install as follows:
        ```bash
        # add the repository to the apt index
        curl -fsSL https://packages.redis.io/gpg | sudo gpg --dearmor -o /usr/share/keyrings/redis-archive-keyring.gpg

        echo "deb [signed-by=/usr/share/keyrings/redis-archive-keyring.gpg] https://packages.redis.io/deb $(lsb_release -cs) main" | sudo tee /etc/apt/sources.list.d/redis.list

        # update and install redis
        sudo apt-get update
        sudo apt-get install redis
        ```
    - Ref : https://redis.io/docs/getting-started/installation/install-redis-on-windows/

2. Golang :
    - Purpose : Build a simple server using the GIN framework
    - Environment : Window 10 Professional
    - install from VSCode (go 1.17)

## Usage

1. Start redis server
    ```
    sudo service redis-server start
    ```

2. Run the API server under the root project path (Window)
    ```
    go run .\main.go
    ```

3. Call the APIs
    - Generate short URL
        ``` bash
        # request
        curl --location --request POST 'http://127.0.0.1:8888/url/shorten' \
        --header 'Content-Type: application/json' \
        --data-raw '{
            "url" : "https://www.google.com/"
        }'
        # respose
        {
            "short_url": "2R3TpRFldYE"
        }
        ```
    - Redirect with short URL
        ```bash
        curl --location --request GET 'http://127.0.0.1:8888/url/2R3TpRFldYE'
        ```

## Todo
1. 1 to 1 binding of short URL and original URL
2. Local cache
3. Web service with load balance (e.g. NGINX)
4. More hash functions