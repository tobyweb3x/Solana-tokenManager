format:
	@templ fmt .

generate:
	@templ generate 

dev:
	@templ generate --watch --proxy="http://localhost:8080" --open-browser=false -v

tailwind:
	@tailwindcss -i assets/styles/input.css  -o assets/styles/output.css --watch --minify

run:
	air

.PHONY: format generate dev run tailwind
