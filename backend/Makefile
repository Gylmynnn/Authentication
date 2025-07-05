start:
	docker-compose up -d
	air & echo $$! > air.pid

stop:
	docker-compose down
	-@if [ -f air.pid ]; then \
		kill `cat air.pid` && rm air.pid; \
	fi

.PHONY: start stop
