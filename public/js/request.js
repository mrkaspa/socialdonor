$(function() {
  $('form#request').bind('submit', validateGeo("lat"));
  return initGeo();
});
