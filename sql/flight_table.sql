CREATE TABLE Flights (
  id SERIAL PRIMARY KEY,
  geom GEOMETRY(Point, 4326),
  latitude decimal,
  longitude decimal,
  country VARCHAR(128),
  call_sign VARCHAR(128),
  icao VARCHAR(128),
  velocity decimal
);