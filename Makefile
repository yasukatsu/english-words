up:
	docker-compose up -d
	docker-compose exec frontend npm run dev

down:
	docker-compose down

log:
	docker-compose logs $(V)