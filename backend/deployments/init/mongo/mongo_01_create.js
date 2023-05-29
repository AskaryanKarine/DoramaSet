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
db.createCollection("_user_fav")
db.createCollection("_user_watched_episode")



db.dorama.createIndex({id:1}, {unique: true, partialFilterExpression: { id: { $exists: true} } })
db.staff.createIndex({id:1}, {unique: true, partialFilterExpression: { id: { $exists: true} } })
db.user.createIndex({username:1}, {unique: true, partialFilterExpression: { username: { $exists: true} } })
db.list.createIndex({id: 1}, {unique: true, partialFilterExpression: { id: { $exists: true} } })
db.subscription.createIndex({id:1}, {unique: true, partialFilterExpression: { id: { $exists: true} } })
