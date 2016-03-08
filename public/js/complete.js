$(function() {
  $('form#complete').bind('submit', validateGeo("lat"));
  return initGeo();
});
