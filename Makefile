init-project:
	go mod init learn-neo4j-go
up:
	docker-compose up -d 
db:
	docker exec -it neo4j cypher-shell -u neo4j -p neo -d system "SHOW DATABASES;"
cli:
	docker exec -it neo4j  bin/cypher-shell -u neo4j -p neo

run:
	NEO4J_URI="bolt://127.0.0.1:7687" NEO4J_USERNAME="neo4j" NEO4J_PASSWORD="testing" NEO4J_DATABASE_NAME="neo4j" go run -race main.go  

add_movie: 
	curl -XPOST http://localhost:8080/movies -d '{"title": "My Movie", "released": 2020, "tagline": "This is not real in a real world"}'

update_movie: 
	curl -XPUT http://localhost:8080/movies -d '{"title": "My Movie", "released": 2010, "tagline": "This is not real"}'


add_person: 
	curl -XPOST http://localhost:8080/movies/person -d '{"name": "John Doe", "born": 1960 }'


update_person: 
	curl -XPUT http://localhost:8080/movies/person -d '{"name": "John Doe", "born": 1961 }'


delete_person:
	curl -XDELETE http://localhost:8080/movies/person -d '{"name": "John Doe"}'

add_job:
	curl -XPOST http://localhost:8080/movies/job -d '{"name": "John Doe", "job": "ACTED_IN", "roles": ["Mr. Doenut","Eater beater"], "movie": "My Movie" }'

delete_job:
	curl -XDELETE http://localhost:8080/movies/job -d '{"name": "John Doe", "job": "ACTED_IN", "movie": "My Movie" }'


add_job2:
	curl -XPOST http://localhost:8080/movies/job -d '{"name": "John Doe", "job": "DIRECTED", "movie": "My Movie" }'

delete_job2:
	curl -XDELETE http://localhost:8080/movies/job -d '{"name": "John Doe", "job": "DIRECTED", "movie": "My Movie" }'

add_job3:
	curl -XPOST http://localhost:8080/movies/job -d '{"name": "John Doe", "job": "PRODUCED", "movie": "My Movie" }'

delete_job3:
	curl -XDELETE http://localhost:8080/movies/job -d '{"name": "John Doe", "job": "PRODUCED", "movie": "My Movie" }'

delete_movie:
	curl -XDELETE http://localhost:8080/movies -d '{"title": "My Movie" }'
 
# list:
# 	curl http://localhost:8080/search?q=matrix

# graph:
# 	curl http://localhost:8080/graph