create table receptionistEdits (
	editDatetime DATE PRIMARY KEY,
	editType VARCHAR(6),
	editorId VARCHAR(18),
	patientUUID VARCHAR(36),
	assignedDept VARCHAR(12),
  assignedCare VARCHAR(15) CHECK (assignedCare IN ('INPATIENT', 'OUTPATIENT')),
	comments VARCHAR(41)
);
insert into receptionistEdits (editDatetime, editType, editorId, patientUUID, assignedDept, assignedCare, comments) values ('2023-04-28T13:31:07Z', 'ENTRY', 'receptionist038456', '91792716-1724-4a14-8d36-b1599be750b0', 'GenPhysician', 'OUTPATIENT', 'Fever patient');
insert into receptionistEdits (editDatetime, editType, editorId, patientUUID, assignedDept, assignedCare, comments) values ('2023-08-22T07:47:11Z', 'ENTRY', 'receptionist038456', '74d27551-87d8-4cae-8d21-c9a31b7d2571', 'Orthopaedic', 'OUTPATIENT', 'Doctor Papel recommended Orthopaedic dept');
insert into receptionistEdits (editDatetime, editType, editorId, patientUUID, assignedDept, assignedCare, comments) values ('2023-09-16T08:26:20Z', 'DELETE', 'receptionist234550', '74d27551-87d8-4cae-8d21-c9a31b7d2571', 'Orthopaedic', 'OUTPATIENT', 'Doctor Armin recommended cancer dept');
insert into receptionistEdits (editDatetime, editType, editorId, patientUUID, assignedDept, assignedCare, comments) values ('2022-11-18T13:58:42Z', 'ENTRY', 'receptionist234550', '74d27551-87d8-4cae-8d21-c9a31b7d2571', 'GenPhysician', 'OUTPATIENT','Doctor Armin recommended cancer dept');
insert into receptionistEdits (editDatetime, editType, editorId, patientUUID, assignedDept, assignedCare, comments) values ('2022-12-08T16:59:17Z', 'EDIT', 'receptionist234550', '91792716-1724-4a14-8d36-b1599be750b0', 'Orthopaedic', 'OUTPATIENT','Doctor Armin recommended cancer dept');
insert into receptionistEdits (editDatetime, editType, editorId, patientUUID, assignedDept, assignedCare, comments) values ('2023-10-08T07:00:39Z', 'ENTRY', 'receptionist234550', '91792716-1724-4a14-8d36-b1599be750b0', 'Cancer', 'INPATIENT', 'Doctor Papel recommended Orthopaedic dept');
insert into receptionistEdits (editDatetime, editType, editorId, patientUUID, assignedDept, assignedCare, comments) values ('2023-03-14T13:21:07Z', 'EDIT', 'receptionist038456', '1c340c57-f8e9-4471-b3dc-4cfc1533a90c', 'Cancer', 'INPATIENT', 'Doctor Papel recommended Orthopaedic dept');
insert into receptionistEdits (editDatetime, editType, editorId, patientUUID, assignedDept, assignedCare, comments) values ('2022-12-01T06:27:49Z', 'EDIT', 'receptionist234550', '1c340c57-f8e9-4471-b3dc-4cfc1533a90c', 'Cancer', 'INPATIENT', 'Doctor Papel recommended Orthopaedic dept');
insert into receptionistEdits (editDatetime, editType, editorId, patientUUID, assignedDept, assignedCare, comments) values ('2022-11-21T11:08:11Z', 'ENTRY', 'receptionist038456', '1c340c57-f8e9-4471-b3dc-4cfc1533a90c', 'Orthopaedic', 'OUTPATIENT', 'Doctor Papel recommended Orthopaedic dept');
insert into receptionistEdits (editDatetime, editType, editorId, patientUUID, assignedDept, assignedCare, comments) values ('2023-06-02T02:22:46Z', 'ENTRY', 'receptionist234550', '74d27551-87d8-4cae-8d21-c9a31b7d2571', 'GenPhysician', 'OUTPATIENT', 'Fever patient');
