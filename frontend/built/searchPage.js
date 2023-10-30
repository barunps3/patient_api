"use strict";
const patientId = document.querySelector(".id-menu");
const searchInput = document.getElementById("patient_id");
const idTypeSelector = document.querySelector("select#id-type");
idTypeSelector.addEventListener("change", () => {
    const idType = idTypeSelector.options[idTypeSelector.selectedIndex];
    // console.log(idType)
    switch (idType.value) {
        case "aadhar-card":
            searchInput.setAttribute("placeholder", "Enter Aadhar Card Number");
            break;
        case "passport":
            searchInput.setAttribute("placeholder", "Enter Passport Number");
            break;
        case "hospital-patient-id":
            searchInput.setAttribute("placeholder", "Enter Patient ID");
            break;
        default:
            searchInput.setAttribute("placeholder", "Select ID type");
    }
});
const umeDetails = {
    "name": {
        "firstName": "Ume",
        "lastName": "Hani",
    },
    "gender": "Female",
    "dateOfBirth": "11/03/1993",
    "idType": "Aadhar Card",
    "idValue": "ZKUP38U"
};
const barunDetails = {
    "name": {
        "firstName": "Barun",
        "lastName": "Mazumdar",
    },
    "gender": "Male",
    "dateOfBirth": "14/10/1992",
    "idType": "Aadhar Card",
    "idValue": "FFUP38U"
};
const searchResult = [umeDetails, barunDetails];
function getIdInput() {
    const inputText = searchInput.value;
    return inputText;
}
function matchId(idText) {
    for (const patient of searchResult) {
        if (patient.idValue == idText) {
            return patient;
        }
    }
}
const resultsContainer = document.getElementById("results-container");
const searchButton = document.getElementById("search");
searchButton.addEventListener("click", (e) => {
    const resultTable = resultsContainer.querySelector("#results-container > table");
    console.log(resultTable);
    if (resultTable != null) {
        resultsContainer.removeChild(resultTable);
    }
    const table = document.createElement("table");
    const headerRow = table.insertRow(0);
    headerRow.classList.add("header-row");
    const nameCell = headerRow.insertCell(0);
    const genderCell = headerRow.insertCell(1);
    const dateOfBirth = headerRow.insertCell(2);
    const idType = headerRow.insertCell(3);
    const idValue = headerRow.insertCell(4);
    nameCell.innerHTML = "Name";
    genderCell.innerHTML = "Gender";
    dateOfBirth.innerHTML = "Date of Birth";
    idType.innerHTML = "ID Type";
    idValue.innerHTML = "ID Value";
    const idInput = getIdInput();
    const patient = matchId(idInput);
    if (patient != null) {
        for (let i = 1; i <= 1; i++) {
            const row = table.insertRow(i);
            const nameCell = row.insertCell(0);
            const genderCell = row.insertCell(1);
            const dateOfBirth = row.insertCell(2);
            const idType = row.insertCell(3);
            const idValue = row.insertCell(4);
            nameCell.innerHTML = `${patient.name.firstName} ${patient.name.lastName}`;
            genderCell.innerHTML = `${patient.gender}`;
            dateOfBirth.innerHTML = `${patient.dateOfBirth}`;
            idType.innerHTML = `${patient.idType}`;
            idValue.innerHTML = `${patient.idValue}`;
        }
        resultsContainer.appendChild(table);
    }
});
