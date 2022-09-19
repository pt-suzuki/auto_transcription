export
FIRESTORE_EMULATOR_HOST := localhost:8080
FIREBASE_AUTH_EMULATOR_HOST := localhost:9099
FIREBASE_STORAGE_EMULATOR_HOST := localhost:9199

emulator.start:
	firebase emulators:start --only firestore,storage,pubsub
emulator.stop:
	lsof -t -i:8080 -i:5001 -i:8085 -i:9199 -i:4000 | xargs kill -9
.PHONY: test
test:
	go test ./src/...
