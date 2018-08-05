import numpy as np, glob, face_recognition, ntpath, pickle, os
from os.path import basename
from shutil import copyfile

def copy_face_image():
    for i in glob.glob("data/avatars/*.jpg"):
        image = face_recognition.load_image_file(i)
        face_locations = face_recognition.face_locations(image)
        
        if face_locations:
            print(i)
            copyfile(i, "data/human_face/" + ntpath.basename(i))

def encoding():
    known_faces = []
    name_index = []

    for i in glob.glob("data/human_face/*.jpg"):
        print(i)
        image = face_recognition.load_image_file(i)
        face_encoding = face_recognition.face_encodings(image)[0]
        known_faces.append(face_encoding)
        filename = os.path.splitext(basename(i))[0]
        name_index.append(filename)

    with open('encodings', 'wb') as fp:
        pickle.dump(known_faces, fp)
    with open('index', 'wb') as fp:
        pickle.dump(name_index, fp)

def test_encoding():
    with open('encodings', 'rb') as fp:
        known_faces = pickle.load(fp)
    with open('index', 'rb') as fp:
        name_index = pickle.load(fp)

    test_file = "data/avatars/4557.jpg"
    image = face_recognition.load_image_file(test_file)
    face_encoding = face_recognition.face_encodings(image)[0]

    face_distances = face_recognition.face_distance(known_faces, face_encoding)

    min_index = np.argmin(face_distances)

    print(name_index[min_index])


#copy_face_image()
encoding()
test_encoding()
