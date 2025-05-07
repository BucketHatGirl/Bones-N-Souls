rm -f Application/Assets/Game.wasm
cd Source
GOOS=js GOARCH=wasm go build -o Game.wasm -v *.go
mv Game.wasm ../Application/Assets
cd ..
if [ -f ./Application/Assets/Game.wasm ]; then
  git add .
  git commit -m "$1" $2
  git push origin main --force
else
  echo "Cannot push WebAssembly does not exist"
fi
