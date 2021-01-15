function randDate() {
    return new Date(Date.now() - Math.floor(Math.random() * 3600000));
}

// Seed data to go-recipes (make sure MONGO_INITDB_DATABASE is set to go-recipes).
db.users.drop();
db.users.insertMany([
    {
        username: 'admin',
        email: 'superuser@example.com',
        age: 40,
        createdAt: randDate(),
    },
    {
        username: 'chan',
        email: 'michael.chan@example.com',
        age: 24,
        createdAt: randDate(),
    },
    {
        username: 'john',
        email: 'jonny@example.com',
        age: 31,
        createdAt: randDate(),
    },
]);

// Create index on username.
db.users.createIndex({
    username: 1,
}, null);
