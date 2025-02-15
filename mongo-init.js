const username = process.env.MONGO_USERNAME;
const password = process.env.MONGO_PASSWORD;
const database = process.env.MONGO_DATABASE;


db = db.getSiblingDB(database);

db.createUser({
  user: username,
  pwd: password,
  roles: [{ role: 'readWrite', db: database }],
});
