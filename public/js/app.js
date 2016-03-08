var isRetina, onSuccess, retina, visualizeMap;

isRetina = function() {
  var mediaQuery;
  mediaQuery = "(-webkit-min-device-pixel-ratio: 1.5),(min--moz-device-pixel-ratio: 1.5)(-o-min-device-pixel-ratio: 3/2),(min-resolution: 1.5dppx)";
  if (window.devicePixelRatio > 1) {
    return true;
  }
  if (window.matchMedia && window.matchMedia(mediaQuery).matches) {
    return true;
  }
  return false;
};

retina = function() {
  if (!isRetina()) {
    return;
  }
  return $('img.2x').map(function(i, image) {
    var path;
    path = $(image).attr('src');
    path = path.replace('.png', '@2x.png');
    path = path.replace('.jpg', '@2x.jpg');
    return $(image).attr('src', path);
  });
};

$(document).ready(retina);

$(document).ready(function() {
  return $(".mobile-menu").click(function() {
    return $(".menu").toggle();
  });
});

visualizeMap = function(options) {
  var urlMap;
  urlMap = 'https://bottico.cartodb.com/api/v2/viz/5b5a5396-e6e8-11e4-9020-0e853d047bba/viz.json';
  return cartodb.createVis('map', urlMap, options);
};

onSuccess = function(response) {
  var options;
  options = {
    zoom: 12,
    center: [response.coords.latitude, response.coords.longitude]
  };
  return visualizeMap(options);
};

if (navigator.geolocation) {
  navigator.geolocation.getCurrentPosition(onSuccess);
} else {
  alert("Debes habilitar la geolocalización");
  visualizeMap(null);
}
