docker run -d \
  --name bootcamp-postgres \
  -p "5432:5432" \
  -e POSTGRES_USER=modanisa \
  -e POSTGRES_PASSWORD=mdns-pw \
  -e POSTGRES_DB=mdns-bootcamp \
  postgres