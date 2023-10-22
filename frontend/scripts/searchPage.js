document.addEventListener("click", event => {
  const isDropdownButton = event.target.matches("[data-dropdown-button]");

  let currentDropdown;
  console.log("Current dropdown: " + currentDropdown )
  if (isDropdownButton) {
    currentDropdown = event.target.closest("[data-dropdown]");
    currentDropdown.classList.toggle("active");
  }

  // Close all other dropdown
  // or when clicked somewhere other than current dropdown
  document.querySelectorAll("[data-dropdown].active").forEach(dropdown => {
    console.log("dropdown input :" + dropdown)
    if (dropdown != currentDropdown) {
      console.log("Entered if statement")
      dropdown.classList.remove('active');
    }
  })
})

const patientId = document.querySelector(".id-menu"); 
const searchPlaceholder = document.getElementById("patient_id");
const idSelectorButton = document.querySelector(".id-selector-btn");

patientId.addEventListener("click", event => {
  let idType = event.target.textContent;
  // console.log(idType)
  switch (idType) {
    case "Aadhar Card":
      searchPlaceholder.setAttribute("placeholder", "Enter Aadhar Card Number")
      idSelectorButton.textContent = "Aadhar Card";
      break;
    case "Passport":
      searchPlaceholder.setAttribute("placeholder", "Enter Passport Number")
      idSelectorButton.textContent = "Passport";
      break;
    case "Patient ID":
      searchPlaceholder.setAttribute("placeholder", "Enter Patient ID")
      idSelectorButton.textContent = "Patient ID";
      break;
    default:
      searchPlaceholder.setAttribute("placeholder", "Select ID type")
      idSelectorButton.textContent = "ID type";
  }

  idSelectorButton.closest("[data-dropdown]").classList.remove('active')
})