.PHONY: login create-repository

login:
	buf registry login $(REGISTRY) --username nhale --token-stdin <<< ${BSR_TOKEN}


login-local: REGISTRY=bufbuild.local
login-local: login

login-prod: REGISTRY=buf.build
login-prod: login

create-repo:
	buf beta registry repository create bufbuild.internal/nhale/buffalo --visibility public
