compile:
	oapi-codegen -generate server,spec -package=api ./api/docs/release/client.yaml > ./internal/interfaces/rest/api/services.gen.go
	oapi-codegen -generate types -package=api ./api/docs/release/client.yaml > ./internal/interfaces/rest/api/types.gen.go

local:
	docker-compose -f ./docker/docker-compose.yml up --build