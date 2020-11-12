.PHONY: clean
API_URL := "https://api.sandbox.hierarchy.enlight.skf.com"

RM     ?= rm
WGET   ?= wget
MKDIR  ?= mkdir
DOCKER ?= docker

rest/models/: rest/swagger.json
	$(RM) -rf "$@" && $(MKDIR) -p "$@"
	$(DOCKER) run --rm \
		--volume "$(shell pwd):/src" \
		--user "$(shell id -u):$(shell id -g)" \
		quay.io/goswagger/swagger:v0.25.0 \
			generate model --spec="/src/$<" --target="/src/$@.."

rest/swagger.json:
	$(WGET) "$(API_URL)/swagger/doc.json" -O "$@"
	./scripts/patch-skf-uuids.sh "$@"
	./scripts/patch-x-nullable.sh "$@"

clean:
	$(RM) -rf rest/models
