# wez - weather cli app
Using https://www.weatherapi.com API to get weather information. You need to get your own API key to use this app.

https://github.com/mshemanskyi/wez/assets/1635978/0034edf9-4548-46f2-8cca-f2a5deb7d79f



## Installation
1. Create an account at https://www.weatherapi.com and get your API key
2. Download the binary and make it executable
3. Move binary to /usr/local/bin or any other directory in your $PATH
4. Use the app
    

## Usage
    wez --apikey=<apiKey> -l "London"
    wez --apikey=<apiKey> -location "London"

    alias wezz=wez --apikey=<apiKey> -l "London"



## Build
    go build -o wez
