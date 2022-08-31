.PHONY: print-version
print-version:
	echo "the new version is $(version)"
	echo $(version) > dist/latest
	cat dist/latest