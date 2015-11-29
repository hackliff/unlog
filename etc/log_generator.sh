#! /bin/sh

# const (
  readonly delay=${1:-"100"}
# )

echo '{"logger": "{{lorem.words}}.{{lorem.words}}", "host":"{{domain}}", "stat": {{double}}, "timestamp": "{{unixtime}}", "method": "{{http.method}}", "level": "{{log.levels}}", "msg": "{{sentence}}"}' \
  | phony --tick ${delay}ms
