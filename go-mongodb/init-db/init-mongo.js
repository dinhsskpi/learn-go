db = db.getSiblingDB("sample_db");
db.createCollection("albums");

db.albums.insertMany([
  {
    title: "dinhpv",
    artist: "dasda",
    price: 1000,
  },
]);
