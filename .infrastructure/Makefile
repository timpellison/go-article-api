.PHONY: img
img:
	docker build ../ -f ../.local/Dockerfile -t 470239745256.dkr.ecr.us-east-1.amazonaws.com/ecr-ue1-article-api-poc:latest
	aws ecr get-login-password --region us-east-1 | docker login --username AWS --password-stdin 470239745256.dkr.ecr.us-east-1.amazonaws.com
	docker push 470239745256.dkr.ecr.us-east-1.amazonaws.com/ecr-ue1-article-api-poc:latest
.PHONY: apply
apply:
	cd $(PWD)/ecs \
	&& terraform apply --auto-approve
