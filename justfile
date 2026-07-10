mu:
    cd sql/schema && goose postgres "postgres://postgres:postgres@localhost:5432/gator" up && ..

md:
    cd sql/schema && goose postgres "postgres://postgres:postgres@localhost:5432/gator" down && ..
