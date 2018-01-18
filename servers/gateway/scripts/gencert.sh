

# Move to the above directory
cd ../

# Create a tls folder
mkdir -p tls

# Gen full chain and privkey pem files
openssl req -x509 -sha256 -nodes -days 365 -newkey rsa:2048 -subj "/CN=localhost" -keyout tls/privkey.pem -out tls/fullchain.pem