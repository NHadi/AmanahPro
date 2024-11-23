// 1. Base API URL
const API_BASE_URL = "http://localhost:8080";

// 2. Utility to Get JWT Token
function getAuthToken() {
  return localStorage.getItem("authToken");
}

// 3. AJAX Setup for Authenticated Requests
$.ajaxSetup({
  beforeSend: (xhr) => {
    const token = getAuthToken();
    if (token) {
      xhr.setRequestHeader("Authorization", `Bearer ${token}`);
    }
  },
  error: (xhr) => {
    if (xhr.status === 401) {
      alert("Session expired. Please log in again.");
      window.location.href = "login.html";
    } else if (xhr.status === 403) {
      alert("Access denied.");
    } else if (xhr.status >= 500) {
      alert("Server error. Please try again later.");
    }
  },
});

// 4. Check if User is Authenticated
function isAuthenticated() {
  return !!getAuthToken();
}

// 5. Redirect to Login if Not Authenticated
function requireAuth() {
  if (!isAuthenticated()) {
    alert("Unauthorized access. Redirecting to login.");
    window.location.href = "login.html";
  }
}

// 6. Utility Function: Format Dates
function formatDate(dateString) {
  const options = { year: "numeric", month: "long", day: "numeric" };
  return new Date(dateString).toLocaleDateString(undefined, options);
}

// 7. Utility Function: Format Currency
function formatCurrency(amount) {
  return new Intl.NumberFormat("en-US", {
    style: "currency",
    currency: "USD",
  }).format(amount);
}

// 8. Initialize on Page Load
$(document).ready(() => {
  // Ensure the user is authenticated
  requireAuth();


});
