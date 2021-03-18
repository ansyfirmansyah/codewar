import pyrebase
import pprint
import time

# Sample configuration to update realtime database in Firebase
config = {
  "apiKey": "AIzaSyDAcaN6V5B7G_9BvVGtQAK1rq9-6RWNe94",
  "authDomain": "webinar-python-d64f2.firebaseapp.com",
  "databaseURL": "https://webinar-python-d64f2-default-rtdb.firebaseio.com/",
  "storageBucket": "webinar-python-d64f2.appspot.com"
}

firebase = pyrebase.initialize_app(config)
email = "ansy.firmansyah@gmail.com"
password = "p@ssword"

auth = firebase.auth()

user = auth.sign_in_with_email_and_password(email, password)

# check authentication
pprint.pprint(user)

db = firebase.database()

# Push data to database
for i in range(10):
  time.sleep(1)
  data = {
  "pesan": f"Ini pesan ke {i+1}"
  }
  result = db.child("broadcast").push(
    data,
    user['idToken']
  )
