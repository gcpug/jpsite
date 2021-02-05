FROM node:12.20-slim

COPY ./frontend/dist/frontend /dist/frontend

COPY ./frontend/node_modules /node_modules

CMD [ "node", "/dist/frontend/server/main.js" ]
