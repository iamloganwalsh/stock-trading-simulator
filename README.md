# Stock Trading Simulator
Practice trading stocks and cryptocurrency using virtual funds!

## Installation Guide

### Prerequisites
##### Node.js
Ensure that Node.js is installed on your system. If not, you can download it from https://nodejs.org/en.

##### WSL2 and Redis
Ensure that WSL2 is installed on your system. It is necessary for running Redis.
It can be installed by going into Windows Powershell and executing "wsl --install".

Once installed, execute the following commands:
sudo apt update
sudo apt install redis
sudo service redis-server status

Assuming no errors pop up, this should be good to go.

To stop the redis server, execute: "sudo service redis-server stop"

### Preparing the .env file
In the backend folder stock-trading-simulator/backend, create a new file called ".env".

Copy and paste the example contents found in the ".env.example" file into the newly created ".env" file.

Then, using the information provided (either on our websites or resumes), replace "code" with the API key. NOTE: Do not put quotation marks around the key. It should read as API_KEY=12345 NOT API_KEY="12345".