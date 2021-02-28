-- 5km search around CGH airport
select * from flights 
where geom && ST_MakeEnvelope(-46.69706959626614,-23.67215378290705,-46.614768403733855,-23.582322217092948, 4326) 
