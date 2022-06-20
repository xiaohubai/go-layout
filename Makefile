compose:
	docker-compose -f ./deploy/docker-compose.yml down
	docker-compose -f ./deploy/docker-compose.yml up -d --force-recreate

publish:
	docker build . -t "xiaohubai/go-layout:v1.0.0"
	docker push xiaohubai/go-layout:v1.0.0
