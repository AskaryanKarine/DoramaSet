db.createUser(
    {
        user: "karine",
        pwd: "12346",
        roles: [
            {
                role: "readWrite",
                db: "DoramaSet"
            }
        ]
    }
);

db.createCollection("Dorama")
db.createCollection("Staff")
db.createCollection("User")
db.createCollection("List")
db.createCollection("Subscription")

