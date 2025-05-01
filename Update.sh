rm -f Res/Game.wasm
cd Source
GOOS=js GOARCH=wasm go build -o Game.wasm -v *.go
mv Game.wasm ../Res/
cd ..
git add .
git commit -m $1 $2
git push origin main --force
