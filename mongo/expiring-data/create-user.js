// Create a user in the user database.
db.createUser({
    user: 'nobody',
    pwd:  'secrets',
    roles: [
        {
            role: 'readWrite',
            db:   'go-recipes',
        },
        {
            role: 'dbAdmin',
            db:   'go-recipes',
        },
    ],
});
