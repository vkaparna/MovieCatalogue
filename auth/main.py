from flask import request, jsonify, Flask
from passlib.hash import pbkdf2_sha256
from flask_jwt_extended import JWTManager, jwt_required, create_access_token, get_jwt_identity

import sqlite3



app = Flask(__name__)
app.config['SECRET_KEY'] = 'secret_key'
jwt = JWTManager(app)



@app.route('/register', methods=['POST'])
def register():
    username = request.json.get('username')
    password = request.json.get('password')
    print(username)
    #Sanitize username and password

    hashedPassword = pbkdf2_sha256.hash(password)

    dbConn = sqlite3.connect('Users.db')
    dbConn.execute('INSERT INTO Users (Username, Hashed_Password) VALUES (?, ?)', (username, hashedPassword))
    dbConn.commit()
    dbConn.close()

    return jsonify({'message': 'User created successfully.'})



@app.route('/listDB', methods=['GET'])
def listDB():
    dbConn = sqlite3.connect('Users.db')
    cursor = dbConn.execute('SELECT * FROM Users')
    for row in cursor:
        print ("User_ID = ", row[0], " Username = ", row[1], " Hashed_Password = ", row[2], "\n")
    dbConn.commit()
    dbConn.close()
    return jsonify({'message': 'PRINTED'})


@app.route('/login', methods=['POST'])
def login():
    username = request.json.get('username')
    password = request.json.get('password')

    dbConn = sqlite3.connect('Users.db')
    cursor = dbConn.execute('SELECT ID, Hashed_Password FROM users WHERE username = ?', (username,))
    user = cursor.fetchone()

    if user is None or not pbkdf2_sha256.verify(password, user[1]):
        return jsonify({'message': 'Bad username or password.'}), 401

    access_token = create_access_token(identity=user[0])
    return jsonify({'access_token': access_token})

@app.route('/validate')
@jwt_required()
def validate():
    userId = get_jwt_identity()
    dbConn = sqlite3.connect('Users.db')
    cursor = dbConn.execute('SELECT Username FROM Users WHERE ID = ?', (userId,))
    username = cursor.fetchone()[0]
    # return jsonify({'message': f'Protected endpoint accessed by {username}.'})
    return jsonify({'message': f'User Validated'})

if __name__ == '__main__':
    app.run()