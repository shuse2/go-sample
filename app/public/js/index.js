function get() {
  var xhr = new XMLHttpRequest();
  xhr.open("GET", "http://localhost:3000/twitter/login", false);
  xhr.send();
}
