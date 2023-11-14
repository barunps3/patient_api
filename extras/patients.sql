create table patients (
	patient_id VARCHAR(40) NOT NULL PRIMARY KEY,
	first_name VARCHAR(50) NOT NULL,
	last_name VARCHAR(50) NOT NULL,
	gender VARCHAR(7) NOT NULL,
	date_of_birth DATE NOT NULL,
	insurance_id VARCHAR(50),
	phone_number VARCHAR(50),
	emergency_contact VARCHAR(50),
	address VARCHAR(200)
);
insert into patients (first_name, last_name, gender, date_of_birth, insurance_id, patient_id, phone_number, emergency_contact, address) values ('Caldwell', 'Aisman', 'Male', '2021-04-08', '8AI0F41UUG4FMPM', '1c340c57-f8e9-4471-b3dc-4cfc1533a90c', '+48 153 582 6257', '+33 950 466 8092', '9 Troy Way');
insert into patients (first_name, last_name, gender, date_of_birth, insurance_id, patient_id, phone_number, emergency_contact, address) values ('Basil', 'Fallis', 'Male', '2016-10-12', '2YHM0BTN432P040', '74d27551-87d8-4cae-8d21-c9a31b7d2571', '+351 473 330 3628', null, '2 Aberg Point');
insert into patients (first_name, last_name, gender, date_of_birth, insurance_id, patient_id, phone_number, emergency_contact, address) values ('Finlay', 'Alforde', 'Male', '1998-01-04', 'P2UCO4G1XQK30GR', 'a27d0dda-67fc-449f-85cb-83b137d305fd', '+33 694 859 9440', '+7 109 302 7656', '1603 Anthes Hill');
insert into patients (first_name, last_name, gender, date_of_birth, insurance_id, patient_id, phone_number, emergency_contact, address) values ('Katleen', 'Knoller', 'Female', '2004-11-24', null, '91792716-1724-4a14-8d36-b1599be750b0', '+86 695 248 3496', null, '0 Lawn Point');
insert into patients (first_name, last_name, gender, date_of_birth, insurance_id, patient_id, phone_number, emergency_contact, address) values ('Isaac', 'Connechie', 'Male', '2020-04-25', null, '62bc804e-6a0c-41a3-9e69-97906dcd132a', '+261 526 696 8213', null, '14543 Warbler Circle');
insert into patients (first_name, last_name, gender, date_of_birth, insurance_id, patient_id, phone_number, emergency_contact, address) values ('Yorke', 'Rubinovitch', 'Male', '2018-10-19', null, '05d024a4-470e-4c40-ade2-142c73058c1e', '+86 310 167 5175', null, '78356 Northridge Alley');
insert into patients (first_name, last_name, gender, date_of_birth, insurance_id, patient_id, phone_number, emergency_contact, address) values ('Marcellus', 'Steckings', 'Male', '2011-01-19', null, '73a1e759-7ed7-4810-bb41-c40cbd25cc48', '+86 327 122 6720', null, '62078 Clyde Gallagher Drive');
insert into patients (first_name, last_name, gender, date_of_birth, insurance_id, patient_id, phone_number, emergency_contact, address) values ('Meggy', 'Kingswood', 'Female', '2019-05-19', 'U70Y2TXYD3Y9OQ9', '4ca48a0f-2dca-4edc-a1ba-1b9444261b86', '+351 172 725 4521', null, '913 Shelley Point');
insert into patients (first_name, last_name, gender, date_of_birth, insurance_id, patient_id, phone_number, emergency_contact, address) values ('Melosa', 'Muzzall', 'Female', '1994-03-31', 'A4YU3K2CN3SEU1A', '003db288-6529-41e9-9161-c0aeba9ff619', '+86 634 351 3215', null, '85446 Novick Court');
insert into patients (first_name, last_name, gender, date_of_birth, insurance_id, patient_id, phone_number, emergency_contact, address) values ('Fonz', 'Ecles', 'Male', '2006-08-18', '52IZL8DI18AS678', 'ad37bc01-f769-4ef7-9869-a2170b6ccbd8', '+63 903 171 2993', null, '38808 Mendota Circle');
