db = db.getSiblingDB("aygo")
db.createUser(
    {
        user: "admin",
        pwd: "admin",
        roles: ["readWrite"]
    }
);


