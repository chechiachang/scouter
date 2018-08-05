from flask import Flask, request,  jsonify
from pymongo import MongoClient
import face_recognition, pickle, numpy as np

app = Flask(__name__)
client = MongoClient('mongodb://127.0.0.1:27017/scouter')

# Load encodings and index
known_faces = []
name_index = []

with open('encodings', 'rb') as fp:
    known_faces = pickle.load(fp)
with open('index', 'rb') as fp:
    name_index = pickle.load(fp)

@app.route("/", methods=['GET'])
def hello():
    return "Hello World!"

@app.route("/face_detection", methods=['POST'])
def face_detection():

    # get encoding from request
    encoding = np.asarray(request.get_json()['encoding'])
    print(encoding)

    # Get user ID by face encodings
    face_distances = face_recognition.face_distance(known_faces, encoding)
    min_index = np.argmin(face_distances)
    userid = name_index[min_index]
    print(userid)

    # Get user data by ID
    user = client.scouter.users.find_one({'_id': userid})

    return jsonify(user)

@app.route("/encoding", methods=['GET'])
def test_encoding():
    test_file = "data/avatars/4557.jpg"
    image = face_recognition.load_image_file(test_file)
    face_encoding = face_recognition.face_encodings(image)[0]
    return jsonify(face_encoding.tolist())

app.run(host='127.0.0.1', debug=True)
