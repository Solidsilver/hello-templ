

watch:
  air -c .air.toml

install:
  go install github.com/cosmtrek/air@latest
  go install github.com/a-h/templ/cmd/templ@latest


clean:
  find . -name "*_templ.go" -delete

run:
  templ generate && go run main.go