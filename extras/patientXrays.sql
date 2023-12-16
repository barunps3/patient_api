create table patientXrays (
  xrayHash VARCHAR(32) NOT NULL,
  uploadDate DATE NOT NULL,
  uploadedBy VARCHAR(40) NOT NULL,
	patientUUID VARCHAR(40) NOT NULL,
  cloudLocation VARCHAR(300) NOT NULL,
  PRIMARY KEY(xRayHash)
);
insert into patientXrays (xrayHash, uploadDate, uploadedBy, patientUUID, cloudLocation) values ('685189bedf76c7bcb2bbe5b9955c0194','2023-09-10','xrayUser@1234','1c340c57-f8e9-4471-b3dc-4cfc1533a90c', 'https://baum-image-data.s3.eu-central-1.amazonaws.com/1c340c57-f8e9-4471-b3dc-4cfc1533a90c/2023-09-10/685189bedf76c7bcb2bbe5b9955c0194.png');
insert into patientXrays (xrayHash, uploadDate, uploadedBy, patientUUID, cloudLocation) values ('a53245be5f0c59127d5cb3cfde7720c5','2023-09-10','xrayUser@1234','1c340c57-f8e9-4471-b3dc-4cfc1533a90c', 'https://baum-image-data.s3.eu-central-1.amazonaws.com/1c340c57-f8e9-4471-b3dc-4cfc1533a90c/2023-09-10/a53245be5f0c59127d5cb3cfde7720c5.png');