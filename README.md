<h1 align="center">GoChat</h1>

## Showcase
![alt text](https://github.com/KerbsOD/GoChat/blob/master/LoginImage.png?raw=true)
![alt text](https://github.com/KerbsOD/GoChat/blob/master/ChatImage.png?raw=true)

## Run Locally
To run the app locally you'll need to have installed Go >= 1.19 and NodeJS >= 18.19
1. Download the repo
```
git clone https://github.com/KerbsOD/GoChat.git
cd GoChat
```
2. Run the frontend
```
cd Frontend/
npm install
npm run build
npm start
```
3. Run the backend
```
cd Backend/
go mod tidy
go run main.go
```
4. Open the page \
Normally the frontend will automatically open the page. If not, you can access it by going to http://localhost:3000 

## Run on Docker
```
sudo docker-compose up --build --force-recreate --no-deps
```
