$(document).ready(function () {
  // Load shared templates
  loadTemplates();

  // Initialize page-specific logic if it exists
  if (typeof initializePage === "function") {
    initializePage();
  }
});

function loadTemplates() {
  $("#sidenav-main").load("partials/sidebar.html");
  $("#header").load("partials/header.html");
  $("#navbar").load("partials/navbar.html");
  $("#footer").load("partials/footer.html");
}

