.PHONY: integration

integration:
	docker build -t pgxsqlbuilder . ;\
	set -e ;\
	test_status_code=0 ;\
	docker run --name=pgxsqlbuilder pgxsqlbuilder go test -tags integration ./... || test_status_code=$$? ;\
	docker rm -f pgxsqlbuilder ;\
	docker rmi pgxsqlbuilder ;\
