FROM node:23-alpine3.21

WORKDIR /app

COPY package*.json ./

RUN npm install --only=production

COPY . .

EXPOSE 8080

CMD ["node", "app.js"]
