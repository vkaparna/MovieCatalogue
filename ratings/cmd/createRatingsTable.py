import requests

import sqlite3

dbConn = sqlite3.connect('ratingsData.db')
print ("Opened database successfully")

dbConn.execute("DROP TABLE IF EXISTS Ratings")
dbConn.execute('''
CREATE TABLE Ratings(
    User_ID INTEGER,
    Movie_ID INTEGER,
    Rating_Value INTEGER,
    PRIMARY KEY (User_ID, Movie_ID)
    );
         ''')
print ("Table created successfully")

queryList = [
   "INSERT INTO Ratings(User_ID, Movie_ID, Rating_Value) VALUES (1, 1, 3);",
   "INSERT INTO Ratings(User_ID, Movie_ID, Rating_Value) VALUES (1, 2, 4);",
   "INSERT INTO Ratings(User_ID, Movie_ID, Rating_Value) VALUES (1, 3, 2);",
   "INSERT INTO Ratings(User_ID, Movie_ID, Rating_Value) VALUES (1, 4, 4);",
   "INSERT INTO Ratings(User_ID, Movie_ID, Rating_Value) VALUES (1, 5, 5);",

   "INSERT INTO Ratings(User_ID, Movie_ID, Rating_Value) VALUES (2, 1, 5);",
   "INSERT INTO Ratings(User_ID, Movie_ID, Rating_Value) VALUES (2, 2, 2);",
   "INSERT INTO Ratings(User_ID, Movie_ID, Rating_Value) VALUES (2, 3, 1);",
   "INSERT INTO Ratings(User_ID, Movie_ID, Rating_Value) VALUES (2, 4, 3);",
   "INSERT INTO Ratings(User_ID, Movie_ID, Rating_Value) VALUES (2, 5, 2);",

   "INSERT INTO Ratings(User_ID, Movie_ID, Rating_Value) VALUES (3, 1, 2);",
   "INSERT INTO Ratings(User_ID, Movie_ID, Rating_Value) VALUES (3, 2, 2);",
   "INSERT INTO Ratings(User_ID, Movie_ID, Rating_Value) VALUES (3, 3, 1);",
   "INSERT INTO Ratings(User_ID, Movie_ID, Rating_Value) VALUES (3, 4, 1);",
   "INSERT INTO Ratings(User_ID, Movie_ID, Rating_Value) VALUES (3, 5, 3);",

   "INSERT INTO Ratings(User_ID, Movie_ID, Rating_Value) VALUES (4, 1, 5);",
   "INSERT INTO Ratings(User_ID, Movie_ID, Rating_Value) VALUES (4, 2, 1);",
   "INSERT INTO Ratings(User_ID, Movie_ID, Rating_Value) VALUES (4, 3, 5);",
   "INSERT INTO Ratings(User_ID, Movie_ID, Rating_Value) VALUES (4, 4, 1);",
   "INSERT INTO Ratings(User_ID, Movie_ID, Rating_Value) VALUES (4, 5, 4);",

   "INSERT INTO Ratings(User_ID, Movie_ID, Rating_Value) VALUES (5, 1, 5);",
   "INSERT INTO Ratings(User_ID, Movie_ID, Rating_Value) VALUES (5, 2, 5);",
   "INSERT INTO Ratings(User_ID, Movie_ID, Rating_Value) VALUES (5, 3, 1);",
   "INSERT INTO Ratings(User_ID, Movie_ID, Rating_Value) VALUES (5, 4, 4);",
   "INSERT INTO Ratings(User_ID, Movie_ID, Rating_Value) VALUES (5, 5, 1);",
]

for query in queryList:
    # print(query)
    dbConn.execute(query)

dbConn.commit()

dbConn.close()

dbConn = sqlite3.connect('ratingsData.db')
cursor = dbConn.execute("SELECT * from Ratings")
for row in cursor:
   print ("User_ID = ", row[0], " Movie_ID = ", row[1], " Rating_Value = ", row[2], "\n")

dbConn.close()