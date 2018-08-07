from flask import Flask, request, jsonify
from pymongo import MongoClient
import face_recognition, pickle, numpy as np, bson, base64, io
import PIL
from PIL import Image

app = Flask(__name__)
client = MongoClient('mongodb://127.0.0.1:27017')
db = client.scouter
collection = db.users

# Load encodings and index
known_faces = []
name_index = []
with open('face_recognition/encodings', 'rb') as fp:
    known_faces = pickle.load(fp)
with open('face_recognition/index', 'rb') as fp:
    name_index = pickle.load(fp)

@app.route("/", methods=['GET'])
def hello():
    return "Hello World!"

@app.route("/count", methods=['GET'])
def count():
    return jsonify({'count':collection.count()})

@app.route("/users/<userid>", methods=['GET'])
def get_user(userid):
    user = collection.find_one({'_id': bson.Int64(userid)})
    return jsonify(user)

@app.route("/face_encoding", methods=['POST'])
def face_detection():

    # get encoding from request
    print(request.get_json())
    encoding = np.asarray(request.get_json()['encoding'])

    # Get user ID by face encodings
    face_distances = face_recognition.face_distance(known_faces, encoding)
    distance = min(face_distances)
    min_index = np.argmin(face_distances)
    userid = name_index[min_index]

    # Get user data by ID
    user = collection.find_one({'_id': bson.Int64(userid)})

    return jsonify({'user': user, 'distance': distance})

@app.route('/face_detection', methods=['POST'])
def upload_face():

    json = request.get_json()
    #print(json)

    if json :
        imageBytes = base64.b64decode(json['data'])
        
        # FIXME skip write image file
        image = open('data/decode.jpg', 'wb') 
        image.write(imageBytes)
        image.close()

        faceImage = face_recognition.load_image_file('data/decode.jpg')
        face_encodings = face_recognition.face_encodings(faceImage)

        if len(face_encodings) > 0:
            face_encoding = face_encodings[0]
            #print(face_encoding)

            # Get distances between face_encoding to all known_faces
            face_distances = face_recognition.face_distance(known_faces, face_encoding)
            # Get minimum distance
            distance = min(face_distances)
            min_index = np.argmin(face_distances)
            # Get user ID by face encodings
            userid = name_index[min_index]

            # Get user data by ID
            obj = collection.find_one({'_id': bson.Int64(userid)})
            user = obj['user']
            response = {
                    'id': obj['_id'], 
                    'contribution': obj['contribution'], 
                    'followers': user['followers'], 
                    'publicgists': user['publicgists'], 
                    'publicrepos': user['publicrepos'], 
                    'distance': distance
                    }
            #print(response)

            return jsonify(response)

        else:
            print("Failed to get encoding from image")

    else:
        print("Get image from request error")

    return jsonify({})

@app.route("/encoding", methods=['GET'])
def test_encoding():
    test_file = "data/avatars/4557.jpg"
    image = face_recognition.load_image_file(test_file)
    face_encoding = face_recognition.face_encodings(image)[0]
    return jsonify(face_encoding.tolist())

app.run(host='0.0.0.0', debug=True)
