create table patientIDs (
	patientUUID VARCHAR(40) NOT NULL,
  idType VARCHAR(20) NOT NULL,
  idValue VARCHAR(50) NOT NULL,
  PRIMARY KEY(patientUUID, idType)
);
insert into patientIDs (patientUUID, idType, idValue) values ('1c340c57-f8e9-4471-b3dc-4cfc1533a90c','Passport', 'Q6rnz6m45MAb');
insert into patientIDs (patientUUID, idType, idValue) values ('1c340c57-f8e9-4471-b3dc-4cfc1533a90c', 'AadharCard', 'a0d1I5b50ye9');
insert into patientIDs (patientUUID, idType, idValue) values ('91792716-1724-4a14-8d36-b1599be750b0', 'AadharCard', '1G2AOYsv1eqv');
insert into patientIDs (patientUUID, idType, idValue) values ('a27d0dda-67fc-449f-85cb-83b137d305fd', 'AadharCard', 'NN6qZh43PxbP');
insert into patientIDs (patientUUID, idType, idValue) values ('74d27551-87d8-4cae-8d21-c9a31b7d2571', 'Passport', 'G2QwwFndzH1A');