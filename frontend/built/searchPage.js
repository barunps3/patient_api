"use strict";
document.addEventListener("click", (event) => {
    const target = event.target;
    const isDropdownButton = target.matches("[data-dropdown-button]");
    let currentDropdown;
    if (isDropdownButton) {
        currentDropdown = target.closest("[data-dropdown]");
        currentDropdown.classList.toggle("active");
    }
    // Close all other dropdown
    // or when clicked somewhere other than current dropdown
    document.querySelectorAll("[data-dropdown].active").forEach(dropdown => {
        if (dropdown != currentDropdown) {
            dropdown.classList.remove('active');
        }
    });
});
const patientId = document.querySelector(".id-menu");
const searchInput = document.getElementById("patient_id");
const idSelectorButton = document.querySelector(".id-selector-btn");
patientId.addEventListener("click", event => {
    const target = event.target;
    let idType = target.textContent;
    // console.log(idType)
    switch (idType) {
        case "Aadhar Card":
            searchInput.setAttribute("placeholder", "Enter Aadhar Card Number");
            idSelectorButton.textContent = "Aadhar Card";
            break;
        case "Passport":
            searchInput.setAttribute("placeholder", "Enter Passport Number");
            idSelectorButton.textContent = "Passport";
            break;
        case "Patient ID":
            searchInput.setAttribute("placeholder", "Enter Patient ID");
            idSelectorButton.textContent = "Patient ID";
            break;
        default:
            searchInput.setAttribute("placeholder", "Select ID type");
            idSelectorButton.textContent = "ID type";
    }
    const dropdown = idSelectorButton.closest("[data-dropdown]");
    dropdown.classList.remove('active');
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
    const table = document.createElement("table");
    const headerRow = table.insertRow(0);
    headerRow.classList.add("header-row");
    let nameCell = headerRow.insertCell(0);
    let genderCell = headerRow.insertCell(1);
    let dateOfBirth = headerRow.insertCell(2);
    let idType = headerRow.insertCell(3);
    let idValue = headerRow.insertCell(4);
    nameCell.innerHTML = "Name";
    genderCell.innerHTML = "Gender";
    dateOfBirth.innerHTML = "Date of Birth";
    idType.innerHTML = "ID Type";
    idValue.innerHTML = "ID Value";
    const idInput = getIdInput();
    const patient = matchId(idInput);
    for (let i = 1; i <= 1; i++) {
        const row = table.insertRow(i);
        const nameCell = row.insertCell(0);
        const genderCell = row.insertCell(1);
        const dateOfBirth = row.insertCell(2);
        const idType = row.insertCell(3);
        const idValue = row.insertCell(4);
        nameCell.innerHTML = `${patient === null || patient === void 0 ? void 0 : patient.name.firstName} ${patient === null || patient === void 0 ? void 0 : patient.name.lastName}`;
        genderCell.innerHTML = `${patient === null || patient === void 0 ? void 0 : patient.gender}`;
        dateOfBirth.innerHTML = `${patient === null || patient === void 0 ? void 0 : patient.dateOfBirth}`;
        idType.innerHTML = `${patient === null || patient === void 0 ? void 0 : patient.idType}`;
        idValue.innerHTML = `${patient === null || patient === void 0 ? void 0 : patient.idValue}`;
    }
    resultsContainer.appendChild(table);
});
