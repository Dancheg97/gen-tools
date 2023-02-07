package templates

const LegoSh = `go install github.com/go-acme/lego/v4/cmd/lego@latest
sudo lego --email="user@example.com" --domains="domain.com" --http run
sudo lego --email="user@example.com" --domains="domain.com" --http run
sudo lego --email="user@example.com" --domains="domain.com" --http run
sudo chown -R user:user .lego
`
