find . -type f -name '*templ.go' -delete
templ generate
sudo docker build --tag personal-site .
