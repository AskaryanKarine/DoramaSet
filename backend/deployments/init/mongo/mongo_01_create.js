db.createCollection("dorama", {
    validator: { $and: [
            {release_year: { $gte: 1923}},
            { $or: [
                    {status: "in progress"},
                    {status: "finish"},
                    {status: "will released"},
                ]}
        ]}
})
db.createCollection("staff")
db.createCollection("user")
db.createCollection("list",
    {
        validator: { $or: [
                {type: "public"},
                {type: "private"}
            ]
        }
    })

db.createCollection("subscription")

db.user.createIndex({username:1}, {unique: true, partialFilterExpression: { login: { $exists: true} } })