var XMLHttpRequest = require('xhr2');

var getJSON = function(url, callback) {
    var xhr = new XMLHttpRequest();
    xhr.open('GET', url, true);
    xhr.responseType = 'json';
    xhr.onload = function() {
      var status = xhr.status;
      if (status === 200) {
        callback(null, xhr.response);
      } else {
        callback(status, xhr.response);
      }
    };
    xhr.send();
};



getJSON('https://meridianbet.rs/sails/sport-with-leagues/58/date/2022-09-28T15:46:56+02:00/filter/all/filterPosition/0,0,0',
function(err, data) {
  if (err !== null) {
    console.log("error!")
  } else {
    console.log(data["events"][0]["events"][]);
  }
});