Learning coding Go with Neo4j

Steps:

1. Run docker-compose 

    $ make up 

2. Run app 

    $ make run 

3. Change neo4j password by 
    
    run http://localhost:7474 

    login:  neo4j  password: neo4j 

    new password: testing

4. Add new movie, add new person, and add job
    
    $ make add_movie 

    $ make add_person 

    $ make add_job

5. To verify 

    MATCH (p:Person)-[rel]->(m:Movie)
    WHERE p.name = 'John Doe' 
    RETURN p, rel, m
    

