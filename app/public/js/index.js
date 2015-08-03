function get() {
  var xhr = new XMLHttpRequest();
  xhr.open("GET", "http://localhost:3000/api/user", false);
  xhr.send();
}
