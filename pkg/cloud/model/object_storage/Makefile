all: download fix_spec generate format cleanup

cleanup:
	rm -rv go.mod go.sum test spec.json README.md docs

format:
	go fmt ./...

generate:
	openapi-generator generate -i spec.json \
		-g go \
		--package-name object_storage \
		--api-name-suffix object_storage \
		--git-repo-id go-sacloud-object-storage \
		--git-user-id g1eng \
		-o .

fix_spec:
	./fix_spec.sh spec.json spec-next.json;\
	mv spec-next.json  spec.json 

download: 
	curl https://manual.sakura.ad.jp/cloud/objectstorage/api/api-json.json -o spec.json
