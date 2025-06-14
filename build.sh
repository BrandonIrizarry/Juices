#!/usr/bin/env bash

BINS=('sitegen' 'server')

for bin in "${BINS[@]}"
do
    go build -o ./tmp/"$bin" ./cmd/"$bin"
done

cat > run.sh <<EOF
#!/usr/bin/env bash

./tmp/sitegen
./tmp/server

EOF

chmod +x run.sh
