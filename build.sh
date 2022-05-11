#!/usr/bin/env bash

tag=$2
if [[ -z "tag" ]]; then
  echo "usage: $0 <package-name> <tag>"
  exit 1
fi

package=$1
if [[ -z "$package" ]]; then
  echo "usage: $0 <package-name> <tag>"
  exit 1
fi
package_split=(${package//\// })
package_name=${package_split[-1]}
	
platforms=(
    "windows/amd64" 
    "windows/arm64" 
    "darwin/amd64"
    "darwin/arm64"
    "linux/amd64"
    "linux/arm64"
)

for platform in "${platforms[@]}"
do
	platform_split=(${platform//\// })
	GOOS=${platform_split[0]}
	GOARCH=${platform_split[1]}
	output_name=$package_name'-'$GOOS'-'$GOARCH
	if [ $GOOS = "windows" ]; then
		output_name+='.exe'
	fi	

	env GOOS=$GOOS GOARCH=$GOARCH go build -ldflags="-s -w" -o release/builds/$output_name $package
	if [ $? -ne 0 ]; then
   		echo 'An error has occurred. Aborting the script execution...'
		exit 1
	fi
done

for file in ./release/builds/*
do 
	mkdir -p ./release/$tag
	fname="$(basename ${file})"

	if [[ $fname == *".exe"* ]]; then
		fname=${fname/%.exe}
	fi

	zip -9jpr ./release/$tag/$fname.zip $file ./LICENSE ./.env
done 