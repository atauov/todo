FROM postgres:latest

EXPOSE 5432

ENV POSTGRES_USER=postgres
ENV POSTGRES_PASSWORD=password
ENV POSTGRES_DB=todo-db

CMD ["postgres"]
