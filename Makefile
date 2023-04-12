local invoke:
	sam build && sam local invoke TestFunction -e events/event.json
deploy:
	sam build && sam deploy --profile MyAWSWorker --capabilities CAPABILITY_NAMED_IAM