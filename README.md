# wez - weather cli app
Using https://www.weatherapi.com API to get weather information. You need to get your own API key to use this app.

![image](https://github.com/mshemanskyi/wez/assets/1635978/849af331-c8dd-4756-bc62-5f4271fd37f6)

## Installation
1. Create an account at https://www.weatherapi.com and get your API key
2. Download the binary and make it executable
3. Move binary to /usr/local/bin or any other directory in your $PATH
4. Use the app
    

## Usage
    -l --location <location>  Location to get weather information
    --apikey <apiKey>         API key from https://www.weatherapi.com
    -f --forecast 3            Show forecast for n days

    wez --apikey=<apiKey> -l "London"
    wez --apikey=<apiKey> -location "London"

    alias wezz=wez --apikey=<apiKey> -l "London"



## Build
    go build -o wez
