curl \
  -X POST \
  -H "Accept: application/vnd.github.v3+json" \
  https://api.github.com/repos/youngerier/go-http-srv/dispatches \
  -d '{"event_type":"event_type"}'