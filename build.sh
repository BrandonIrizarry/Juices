#!/usr/bin/env bash

BINS=('sitegen' 'server')

for bin in "${BINS[@]}"
do
    echo "Building $bin"
    go build -o ./tmp/"$bin" ./cmd/"$bin"
done

cat > run.sh <<EOF
#!/usr/bin/env bash

echo "Generating index.html from template"
./tmp/sitegen > ./app/index.html

echo "Running server"
./tmp/server

EOF

chmod +x run.sh
