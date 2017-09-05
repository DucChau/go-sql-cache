CREATE TABLE SQL_CACHE(
    ID                  SERIAL NOT NULL,
    CREATED_At          TIMESTAMP WITH TIME ZONE NOT NULL DEFAULT NOW(),
    KEY                 TEXT NOT NULL,
    VALUE               JSONB NOT NULL,
    TTL                 INT NOT NULL,
    PRIMARY KEY(ID)
);
