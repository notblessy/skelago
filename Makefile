run:
	go run main.go httpsrv

migration:
	go run main.go migrate

model/mock/mock_welcome_repository.go:
	mockgen -destination=model/mock/mock_welcome_repository.go -package=mock github.com/notblessy/skelago/model WelcomeRepository

model/mock/mock_welcome_usecase.go:
	mockgen -destination=model/mock/mock_welcome_usecase.go -package=mock github.com/notblessy/skelago/model WelcomeUsecase

mockgen: model/mock/mock_welcome_repository.go \
	model/mock/mock_welcome_usecase.go

test: unit-test
unit-test: mockgen
	SVC_ENV=test SVC_DISABLE_CACHING=true go test ./... -v --cover