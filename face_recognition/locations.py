import face_recognition

human_image="data/avatars/4557/21.jpg"
panda_image="data/avatars/462/16.jpg"

def print_landmarks(image_path):
    image = face_recognition.load_image_file(image_path)
    face_locations = face_recognition.face_locations(image)
    print(face_locations)
    face_landmarks = face_recognition.face_landmarks(image)
    print(face_landmarks)

print_landmarks(human_image)

print_landmarks(panda_image)
# return empty landmarks array
