import sqlite3

dbConn = sqlite3.connect('Users.db')
dbConn.execute('DROP TABLE IF EXISTS Users')
dbConn.execute('''
CREATE TABLE IF NOT EXISTS Users (
    ID INTEGER PRIMARY KEY AUTOINCREMENT,
    Username TEXT NOT NULL,
    Hashed_Password TEXT NOT NULL
)
''')

dbConn.commit()
dbConn.close()