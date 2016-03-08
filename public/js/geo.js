var initGeo, validateGeo;

validateGeo = function(field) {
  return function(evt) {
    var fieldVal;
    fieldVal = $("div#" + field + " input").val();
    if ((fieldVal == null) || fieldVal === '' || parseInt(fieldVal) === 0) {
      alert("Marca la localización en el mapa");
      return evt.preventDefault();
    }
  };
};

initGeo = function() {
  var lat, lng, location;
  lat = $('div#lat input[type=hidden]').val();
  lng = $('div#lng input[type=hidden]').val();
  location = lat !== '' && lng !== '' ? [lat, lng] : ' ';
  return $('#geocomplete').geocomplete({
    map: '.map_canvas',
    details: 'form',
    markerOptions: {
      draggable: true
    },
    location: location
  }).on('geocode:result', function(event, res) {
    var keys;
    location = res['geometry']['location'];
    keys = Object.keys(location);
    $('div#lat input[type=hidden]').val(location[keys[0]]);
    return $('div#lng input[type=hidden]').val(location[keys[1]]);
  }).on('geocode:dragged', function(event, latLng) {
    $('div#lat input[type=hidden]').val(latLng.lat());
    return $('div#lng input[type=hidden]').val(latLng.lng());
  });
};
