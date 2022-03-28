.PHONY: install
## Install all plugins/tool versions of asdf apps.
install:
	while read line; do \
		read app version <<< "$${line}"; \
		asdf plugin add $${app}; \
		asdf install $${app} $${version}; \
	done < ./.tool-versions

.PHONY: build
build:
	go build -o build/go-example
