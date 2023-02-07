generate-moduleone:
	buf generate --template buf.gen.moduleone.yaml buf.build/semichkin/moduleone

generate-moduletwo:
	buf generate --template buf.gen.moduletwo.yaml buf.build/semichkin/moduletwo

generate: generate-moduleone generate-moduletwo