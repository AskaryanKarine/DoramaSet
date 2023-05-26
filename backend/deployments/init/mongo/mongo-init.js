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
