// frontend/config.js
const API_BASE_URL = "http://localhost:8080"; // API Gateway base URL

// Helper to set up headers, especially for Authorization
function getHeaders() {
  const token = localStorage.getItem("token"); // Store the token in local storage after login
  return token ? { "Authorization": "Bearer " + token } : {};
}
